package main

import (
	"errors"
	"syscall"
)

// Burst does not run on Windows yet. These stubs let the module build there.

// signalProcess cannot signal another process on Windows: there are no POSIX
// signals. Failing for signal 0 too means a caller never mistakes "cannot
// tell" for "still running".
func signalProcess(pid int, sig syscall.Signal) error {
	return errors.New("signalling a process is not supported on Windows yet")
}

// newSessionAttr returns nil: Windows has no sessions in the POSIX sense, and
// a child there does not get the terminal's hangup anyway.
func newSessionAttr() *syscall.SysProcAttr {
	return nil
}
