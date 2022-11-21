package web

import (
	"net/http"
	"time"

	"garuda.com/m/web/cmd/activity"
	"garuda.com/m/web/cmd/auth"
)

func Register(username, password string) error {
	return auth.Register(username, password)
}

func Login(username, password string) (string, time.Time, error) {
	return auth.Login(username, password)
}

func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request, claims *auth.Claims)) http.HandlerFunc {
	return auth.VerifyJWT(endpointHandler)
}

func CreatePost(username, title, content string) error {
	return activity.CreatePost(username, title, content)
}

func GetPosts(username string) ([]map[string]string, error) {
	return activity.GetPosts(username)
}

func AddFollowing(username, following string) error {
	return activity.AddFollowing(username, following)
}

func GetFollowings(username string) ([]string, error) {
	return activity.GetFollowings(username)
}

func DeleteFollowing(username, following string) error {
	return activity.DeleteFollowing(username, following)
}
