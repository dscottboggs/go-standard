package fileops

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
	"strings"

	"github.com/dscottboggs/attest"
)


func TestReadFileAsync(t *testing.T) {
	// setup
	test := attest.New(t)
	tFileContents :=
		"A test string which should be expected to be in the file.\n"
	test.Handle(
		ioutil.WriteFile(
			"/tmp/fileOps_test.go.tmp.txt",
			[]byte(tFileContents),
			0644))
	// Test actually starts here
	ch, err := ReadFileAsync(ROOT, "tmp", "fileOps_test.go.tmp.txt")
	test.Handle(err)
	test.Equals(tFileContents, strings.Trim(string(<-ch), "\x00"))
	// teardown -- delete the temp file
	test.Handle(os.Remove("/tmp/fileOps_test.go.tmp.txt"))
	ch, err = ReadFileAsync(ROOT, "invalid", "path", "causes", "error")
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
}
