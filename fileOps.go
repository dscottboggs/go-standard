package standard

import (
	"fmt"
	"io/ioutil"
	"log"
)

// ROOT - a literal forward slash
const ROOT = "/"

// BuildPath -- Build an absolute or relative path a la python's os.path.join
func BuildPath(filepath ...string) string {
	var abspath string
	if filepath[0] == ROOT {
		for _, fp := range filepath[1:] {
			abspath += "/" + fp
		}
	} else {
		abspath = filepath[0]
		for _, fp := range filepath[1:] {
			abspath += "/" + fp
		}
	}
	return abspath
}

// ReadFileAsync -- read the contents of the file built by filepath into the
// contents channel.
func ReadFileAsync(contents chan string, filepath ...string) {
	fp := BuildPath(filepath...)
	b, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Printf("Error reading file '%s'.", fp)
		contents <- fmt.Sprintf("Error reading file '%s'.", fp)
		return
	}
	contents <- string(b)
	return
}
