package fileops

import (
	"os"
	"path"
)

// ROOT - a literal forward slash
const ROOT = "/"

// ReadFileAsync -- read the contents of the file built by filepath into the
// contents channel.
func ReadFileAsync(filepath ...string) (chan []byte, error) {
	var (
		buffer    = make([]byte, 1024)
		bufchan   = make(chan []byte)
		fp        = path.Join(filepath...)
		file, err = os.Open(fp)
	)
	if err != nil && !os.IsExist(err) {
		close(bufchan)
		return bufchan, err
	}
	go func() {
		_, err := file.Read(buffer)
		bufchan <- buffer
		if err != nil {
			close(bufchan)
			return
		}
	}()
	return bufchan, nil
}

// WriteFileAsync -- Write data received on the given channel to the given file
// until the channel is closed
func WriteFileAsync(contents chan []byte, permissions os.FileMode, filepath ...string) chan error {
	var (
		buf   []byte
		ok    bool
		err   error
		errch chan error
		file  *os.File
	)
	go func() {
		file, err = os.Create(path.Join(filepath...))
		if err != nil {
			errch <- err
		}
		for {
			if buf, ok = <-contents; ok {
				_, err := file.Write(buf)
				if err != nil {
					errch <- err
					return
				}
			} else {
				return
			}
		}
	}()
	return errch
}
