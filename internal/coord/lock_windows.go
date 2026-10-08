package coord

import (
	"errors"
	"os"
)

// Coordination is not supported on Windows yet. These stubs let the module
// build there; taking the state lock fails with a clear error rather than
// running without one, because two sessions writing state.json unlocked
// would lose each other's records.

var errNoFileLock = errors.New("session coordination is not supported on Windows yet")

func tryLockExclusive(f *os.File) error { return errNoFileLock }

func lockBusy(err error) bool { return false }

func unlock(f *os.File) error { return nil }
