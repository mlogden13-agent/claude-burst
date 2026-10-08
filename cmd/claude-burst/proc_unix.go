//go:build unix

package main

import "syscall"

// signalProcess sends sig to pid; signal 0 only checks the process exists.
func signalProcess(pid int, sig syscall.Signal) error {
	return syscall.Kill(pid, sig)
}

// newSessionAttr starts a child in its own session: no SIGHUP when the
// terminal closes, and outside the gateway's process group.
func newSessionAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{Setsid: true}
}
