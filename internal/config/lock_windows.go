package config

import (
	"errors"
	"os"
)

// Burst does not run on Windows yet, and Windows has no flock. This stub
// lets the module build there; an Update fails with a clear error rather than
// saving without the lock, which is the lost-update race Update exists to
// prevent.
func lockExclusive(f *os.File) error {
	return errors.New("changing config.json is not supported on Windows yet: it needs a file lock")
}

func unlock(f *os.File) {}
