package api

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/marius004/phoenix/internal/models"
	"github.com/marius004/phoenix/internal/util"
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

func convertUrlValuesToInt(values []string) []int {
	ret := make([]int, 0)

	for _, str := range values {
		if val, err := strconv.Atoi(str); err != nil {
			ret = append(ret, val)
		}
	}

	return ret
}

func convertUrlValuesToLangArr(values []string) []models.SubmissionLang {
	langs := make([]models.SubmissionLang, len(values))

	for ind, val := range values {
		langs[ind] = models.SubmissionLang(val)
	}

	return langs
}

func convertUrlValuesToStatusArr(values []string) []models.SubmissionStatus {
	statuses := make([]models.SubmissionStatus, len(values))

	for ind, val := range values {
		statuses[ind] = models.SubmissionStatus(val)
	}

	return statuses
}

// https://en.gravatar.com/site/implement/hash/
func (s *API) calculateEmailHash(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	hash := md5.Sum([]byte(email))
	str := hex.EncodeToString(hash[:])

	return str
}
