//go:build unix

package admin

import (
	"os"
	"syscall"
)

// terminateSelf sends this process SIGTERM, which the drain handler answers.
func terminateSelf() error {
	return syscall.Kill(os.Getpid(), syscall.SIGTERM)
}
