package util

import "net/http"

// IsRAuthed returns true if the user creating the request is authenticated
func IsRAuthed(r *http.Request) bool {
	user := UserFromRequestContext(r.Context())
	return user != nil
}

// IsRAdmin returns true if the user creating the request is an admin
func IsRAdmin(r *http.Request) bool {
	user := UserFromRequestContext(r.Context())
	return user != nil && user.IsAdmin
}

// IsRProposer returns true if the user creating the request is a proposer
func IsRProposer(r *http.Request) bool {
	user := UserFromRequestContext(r.Context())
	return user != nil && (user.IsProposer || user.IsAdmin)
}
