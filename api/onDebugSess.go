package api

import "context"

// onDebugSess handles a debug client authentication request
func (srv *server) onDebugSess(
	ctx context.Context,
	username, password string,
) []byte {
	// Check debug credentials
	if username != srv.opts.DebugUser.Username ||
		password != srv.opts.DebugUser.Password {
		return nil
	}

	// Return session key
	return srv.debugSessionKey
}
