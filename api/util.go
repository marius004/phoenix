package api

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/marius004/phoenix/models"
	"github.com/marius004/phoenix/util"
)

const (
	HttpOnlyCookieAuthKey = "auth-token-http-only"
	CookieAuthKey         = "auth-token"
	CookieUserKey         = "user"
)

func (s *API) canManageProblemResources(r *http.Request, problem *models.Problem) bool {
	user := util.UserFromRequestContext(r)

	if user == nil || problem == nil {
		return false
	}

	if util.IsRAdmin(r) {
		return true
	}

	if util.IsRProposer(r) && problem.AuthorId == user.Id {
		return true
	}

	return false
}

func (s *API) canSeeSubmissionSourceCode(r *http.Request, submission *models.Submission, problem *models.Problem) bool {
	user := util.UserFromRequestContext(r)
	return util.IsRAdmin(r) ||
		s.canProposerSeeSourceCode(r, problem) ||
		(uint64(submission.UserId) == user.Id)
}

func (s *API) canProposerSeeSourceCode(r *http.Request, problem *models.Problem) bool {
	user := util.UserFromRequestContext(r)
	return util.IsRProposer(r) && user != nil && problem.AuthorId == user.Id
}

type serverAuthCookie struct {
	Token string `json:"token"`
}

func newServerAuthCookie(token string) *serverAuthCookie {
	return &serverAuthCookie{
		Token: token,
	}
}

func (c *serverAuthCookie) Cookie(exp time.Duration) *http.Cookie {
	return &http.Cookie{
		Name:     HttpOnlyCookieAuthKey,
		Value:    c.Token,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24 * time.Duration(exp)),
	}
}

type clientAuthCookie struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

func newClientAuthCookie(token string, user *models.User) *clientAuthCookie {
	return &clientAuthCookie{
		Token: token,
		User:  user,
	}
}

func (c *clientAuthCookie) Cookie(exp time.Duration) *http.Cookie {
	return &http.Cookie{
		Name:     CookieAuthKey,
		Value:    encodeToBase64(c),
		Path:     "/",
		HttpOnly: false,
		Expires:  time.Now().Add(time.Hour * 24 * time.Duration(exp)),
	}
}

func encodeToBase64(data interface{}) string {
	json, err := json.Marshal(data)

	if err != nil {
		log.Println(err)
		return ""
	}

	return base64.StdEncoding.EncodeToString(json)
}
