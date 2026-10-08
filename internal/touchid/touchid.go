//go:build darwin && cgo

package touchid

// The real prompt: LocalAuthentication through cgo. Other builds get
// touchid_other.go, which refuses.

/*
#cgo CFLAGS: -x objective-c -fobjc-arc
#cgo LDFLAGS: -framework LocalAuthentication -framework Foundation
#include <stdlib.h>
#include "tid.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

// Authenticate presents the system sheet with reason as its explanation and
// returns nil only if the device owner authenticated.
func Authenticate(reason string) error {
	cReason := C.CString(reason)
	defer C.free(unsafe.Pointer(cReason))

	var cErr *C.char
	ok := C.tid_auth(cReason, &cErr, C.int(timeoutSeconds))
	if cErr != nil {
		msg := C.GoString(cErr)
		C.free(unsafe.Pointer(cErr))
		if ok != 1 {
			return errors.New(msg)
		}
	}
	if ok != 1 {
		return errors.New("authentication failed")
	}
	return nil
}
