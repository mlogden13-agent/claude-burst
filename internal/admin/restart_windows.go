package admin

import "errors"

// terminateSelf has no Windows equivalent yet: there is no SIGTERM to send
// and no launchd to start the gateway again, so the restart button says so
// instead of stopping a gateway that nothing would bring back.
func terminateSelf() error {
	return errors.New("restarting the gateway is not supported on Windows yet")
}
