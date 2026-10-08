//go:build !darwin || !cgo

package touchid

import "errors"

// errUnsupported is what Authenticate returns where there is no Touch ID
// sheet to show: any platform but macOS, or a macOS build without cgo. A
// refusal rather than a pass, because the actions behind this gate hand a
// stored secret back in the clear, and a build that cannot ask the owner
// must not reveal it.
var errUnsupported = errors.New("device-owner authentication needs Touch ID, which this build does not have (macOS with cgo only)")

// Authenticate always refuses on this build; see errUnsupported.
func Authenticate(reason string) error {
	return errUnsupported
}
