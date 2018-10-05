package fileops

import (
	"os"
	"path"
)

// ROOT - a literal forward slash
const ROOT = "/"

// Read -- read the contents of the file built by filepath into the
// contents channel.
func Read(filepath ...string) (chan []byte, chan error) {
	var (
		buffer    = make([]byte, 1024)
		bufchan   = make(chan []byte)
		fp        = path.Join(filepath...)
		file, err = os.Open(fp)
		errch     = make(chan error)
	)
	go func() {
		if err != nil && !os.IsExist(err) {
			close(bufchan)
			errch <- err
			close(errch)
		}
		for {
			_, err := file.Read(buffer)
			bufchan <- buffer
			if err != nil {
				close(bufchan)
				return
			}
		}
	}()
	return bufchan, errch
}

// Write -- Write data received on the given channel to the given file
// until the channel is closed
func Write(filepath ...string) (chan []byte, chan error) {
	var (
		buf    []byte
		ok     bool
		err    error
		file   *os.File
		errch  = make(chan error)
		datach = make(chan []byte)
	)
	go func() {
		file, err = os.Create(path.Join(filepath...))
		if err != nil {
			errch <- err
			close(errch)
			return
		}
		for {
			if buf, ok = <-datach; ok {
				if _, err := file.Write(buf); err != nil {
					file.Close()
					errch <- err
					break
				}
			} else {
				file.Close()
				break
			}
		}
	}()
	return datach, errch
}

// Copy the given file, trying several methods to make sure it gets done as fast
// as possible.
func Copy(fromPath, toPath string, overwrite bool) error {
	var (
		err              error
		fromFile, toFile os.FileInfo
	)
	if f, statErr := os.Stat(fromPath); statErr != nil {
		return statErr
	} else {
		fromFile = f
	}
	if f, statErr := os.Stat(toPath); statErr != nil {
		if os.IsNotExist(statErr) {
			toFile = f
		} else {
			return statErr
		}
	} else if overwrite {
		toFile = f
	} else {
		return err
	}
	if os.SameFile(fromFile, toFile) {
		return nil
	}
	if err := os.Link(fromPath, toPath); err == nil {
		return nil
	}
	// try copying one kilobyte at at time
	var (
		writer, wErr = Write(toPath)
		reader, rErr = Read(fromPath)
		e            error
		ok           bool
		kilobyte     []byte
	)
	for {
		select {
		case e = <-wErr:
			return e
		case e = <-rErr:
			return e
		case kilobyte, ok = <-reader:
			if ok {
				writer <- kilobyte
			} else {
				close(writer)
				return nil
			}
		}
	}
}
