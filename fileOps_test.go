package standard

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/dscottboggs/attest"
)

func TestBuildPath(t *testing.T) {
	t.Run("Absolute path", func(t *testing.T) {
		test := attest.Test{t}
		test.Equals(
			"/test/path/number/one",
			BuildPath(ROOT, "test", "path", "number", "one"))
		test.Equals("/var/log/www", BuildPath(ROOT, "var", "log", "www"))
	})
	t.Run("Relative path", func(t *testing.T) {
		test := attest.Test{t}
		test.Equals("test/path/two", BuildPath("test", "path", "two"))
		test.Equals(
			".config/i3/config",
			BuildPath(".config", "i3", "config"))
	})
}

func TestReadFileAsync(t *testing.T) {
	// setup
	test := attest.Test{t}
	tFileContents :=
		"A test string which should be expected to be in the file.\n"
	test.Handle(
		ioutil.WriteFile(
			"/tmp/fileOps_test.go.tmp.txt",
			[]byte(tFileContents),
			0644))
	ch := make(chan string)
	// Test actually starts here
	go ReadFileAsync(ch, ROOT, "tmp", "fileOps_test.go.tmp.txt")
	test.Equals(tFileContents, <-ch)
	// teardown -- delete the temp file
	test.Handle(os.Remove("/tmp/fileOps_test.go.tmp.txt"))
	go ReadFileAsync(ch, ROOT, "invalid", "path", "causes", "error")
	test.Equals(
		"Error reading file '/invalid/path/causes/error'.", <-ch)
	// Output: Error reading file '/invalid/path/causes/error'.
}
