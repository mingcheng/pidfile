// Package pidfile provides structure and helper functions to create and remove
// PID file. A PID file is usually a file used to store the process ID of a
// running process.
//
// @ref https://github.com/moby/moby/tree/master/pkg/pidfile
package pidfile

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Common error for pidfile package
var (
	ErrProcessRunning = errors.New("process is running")
	ErrFileStale      = errors.New("pidfile exists but process is not running")
	ErrFileInvalid    = errors.New("pidfile has invalid contents")
)

// PIDFile is a file used to store the process ID of a running process.
type PIDFile struct {
	path string
	pid  int
}

// New creates a PIDfile using the specified path.
func New(path string) (*PIDFile, error) {
	file := PIDFile{
		path: path,
		pid:  os.Getpid(),
	}

	if pid, err := file.Content(); err == nil || processExists(pid) {
		return nil, ErrProcessRunning
	}

	if err := file.Write(); err != nil {
		return nil, err
	}

	return &file, nil
}

// Remove the PIDFile.
func (file PIDFile) Remove() error {
	return os.Remove(file.path)
}

// Content reads the PIDFile content.
func (file PIDFile) Content() (int, error) {
	contents, err := ioutil.ReadFile(file.path)
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(contents)))
	if err != nil || file.pid != pid {
		return 0, ErrFileInvalid
	}

	return pid, nil
}

// Write writes a pidfile, returning an error
// if the process is already running or pidfile is orphaned
func (file PIDFile) Write() error {
	return file.WriteControl(os.Getpid(), false)
}

func (file PIDFile) WriteControl(pid int, overwrite bool) error {
	// Check for existing pid
	if oldPid, err := file.Content(); err != nil && !os.IsNotExist(err) {
		return err
	} else if err == nil {
		// We have a pid
		if processExists(oldPid) {
			return ErrProcessRunning
		}
		if !overwrite {
			return ErrFileStale
		}
	}

	// Note MkdirAll returns nil if a directory already exists
	if err := os.MkdirAll(filepath.Dir(file.path), os.FileMode(0700)); err != nil {
		return err
	}

	// We're clear to (over)write the file
	return ioutil.WriteFile(file.path, []byte(fmt.Sprintf("%d\n", pid)), 0600)
}

// Running returns true if the process is running.
func (file PIDFile) Running() bool {
	return processExists(file.pid)
}
