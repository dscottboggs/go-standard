package fileops

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/dscottboggs/attest"
)

const (
	tFileContents          = "A test string which should be expected to be in the file.\n"
	testReadFilePath       = "/tmp/fileOps_read_test.txt"
	testWriteFilePath      = "/tmp/fileOps_write_test.txt"
	testPermissionFilePath = "/tmp/fileOps_perm_test.txt"
	readOnlyPermissions    = 0444
)

func TestMain(m *testing.M) {
	// setup
	err := ioutil.WriteFile(testReadFilePath, []byte(tFileContents), 0644)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	// run tests
	exitCode := m.Run()
	// teardown
	err = os.RemoveAll(path.Join(os.TempDir(), testReadFilePath))
	if err != nil && !os.IsExist(err) && !os.IsNotExist(err) {
		log.Fatal(err)
	}
	err = os.RemoveAll(path.Join(os.TempDir(), testWriteFilePath))
	if err != nil && !os.IsExist(err) && !os.IsNotExist(err) {
		log.Fatal(err)
	}
	os.Exit(exitCode)
}

func TestReadFileAsync(t *testing.T) {
	t.Run("successfully read from a real file", func(t *testing.T) {
		test := attest.New(t)
		ch, err := ReadFileAsync(ROOT, "tmp", "fileOps_test.go.tmp.txt")
		test.Handle(err)
		test.Equals(tFileContents, strings.Trim(string(<-ch), "\x00"))
	})
	t.Run("throws an error when the file doesn't exist", func(t *testing.T) {
		test := attest.New(t)
		ch, err := ReadFileAsync(ROOT, "invalid", "path", "causes", "error")
		if err == nil {
			test.NotNil(err, "got nil error from ReadFileAsync on invalid path.")
			time.Sleep(1 * time.Second)
			select {
			case bytes := <-ch:
				test.Error("received %v bytes from expected nonexistent file", bytes)
			default:
				test.Log("nothing received on the channel")
			}
		}
	})
}

func TestWrite(t *testing.T) {
	t.Run("successfully writes to a real file", func(t *testing.T) {
		test := attest.New(t)
		dataChan := make(chan []byte)
		errorChan := WriteFileAsync(dataChan, testWriteFilePath)
		time.Sleep(1 * time.Second)
		select {
		case err := <-errorChan:
			test.Error("got error %#+v when trying to write file", err)
		default:
		}
		dataChan <- []byte(tFileContents)
		close(dataChan)
		time.Sleep(1 * time.Second)
		select {
		case err := <-errorChan:
			test.Error("got error %#+v when trying to write contents to file", err)
		default:
		}
	})
	t.Run("sends an error when permission denied", func(t *testing.T) {
		test := attest.New(t)
		dataChan := make(chan []byte)
		if f, err := os.Create(testPermissionFilePath); err != nil {
			if !os.IsPermission(err) {
				t.Fatalf(
					"got error %#v (that wasn't a permission error) trying to create %s",
					err,
					testPermissionFilePath,
				)
			}
		} else {
			test.Handle(os.Chmod(testPermissionFilePath, readOnlyPermissions))
			f.Close()
		}
		errorChan := WriteFileAsync(dataChan, testPermissionFilePath)
		time.Sleep(1 * time.Second) // to remove necesity for blocking `select`
		select {
		case err := <-errorChan:
			if !os.IsPermission(err) {
				test.Error(
					"got unexpected error %#+v while trying to invoke permission error",
				)
			} else if err == nil {
				test.Error(
					"didn't get expected permission error trying to write to read " +
						"only file",
				)
			}
		default:
			test.Error(
				"didn't get an error on the error channel when trying to invoke " +
					"permission error",
			)
		}
	})
}
