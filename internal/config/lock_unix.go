//go:build unix

package config

import (
	"os"
	"syscall"
)

// lockExclusive blocks until it holds an exclusive flock on f, retrying a
// wait that a signal interrupted.
func lockExclusive(f *os.File) error {
	for {
		err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
		if err != syscall.EINTR {
			return err
		}
	}
}

func unlock(f *os.File) {
	_ = syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
}
