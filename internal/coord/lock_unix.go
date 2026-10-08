//go:build unix

package coord

import (
	"errors"
	"os"
	"syscall"
)

// The state lock is an flock: released when the descriptor closes, so a hook
// that crashes cannot leave it held. Windows has no flock; see
// lock_windows.go.

// tryLockExclusive takes an exclusive lock on f without waiting.
func tryLockExclusive(f *os.File) error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
}

// lockBusy reports whether err from tryLockExclusive means another process
// holds the lock, as opposed to a real failure.
func lockBusy(err error) bool {
	return errors.Is(err, syscall.EWOULDBLOCK)
}

func unlock(f *os.File) error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
}
