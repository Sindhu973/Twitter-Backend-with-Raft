package activity

import (
	"testing"

	"garuda.com/m/web/cmd/auth"
)

// Test to check if creating a post is working
func TestCreatePost(t *testing.T) {
	user := "testuser"
	password := "password"

	err := auth.Register(user, password)
	if err != nil {
		t.Error(err)
	}

	token, _, err := auth.Login(user, password)
	if err != nil {
		t.Error(err)
	}
	if token == "" {
		t.Error("Token is empty")
	}

	title := "testtitle"
	content := "testcontent"

	err = CreatePost(user, title, content)
	if err != nil {
		t.Error(err)
	}
}

// Test to check if getting posts is working
func TestGetPosts(t *testing.T) {
	user := "testuser"
	password := "password"

	err := auth.Register(user, password)
	if err != nil {
		t.Error(err)
	}

	token, _, err := auth.Login(user, password)
	if err != nil {
		t.Error(err)
	}
	if token == "" {
		t.Error("Token is empty")
	}

	title := "testtitle"
	content := "testcontent"

	err = CreatePost(user, title, content)
	if err != nil {
		t.Error(err)
	}

	posts, err := GetPosts(user)
	if err != nil {
		t.Error(err)
	}
	if len(posts) != 1 {
		t.Error("Expected 1 post, got", len(posts))
	}
	if posts[0]["title"] != title {
		t.Error("Expected title", title, "got", posts[0]["title"])
	}
	if posts[0]["content"] != content {
		t.Error("Expected content", content, "got", posts[0]["content"])
	}
}

// Test to check if adding to following is working
func TestAddFollowings(t *testing.T) {
	user := "testuser"
	password := "password"
	following := "following"

	err := auth.Register(user, password)
	if err != nil {
		t.Error(err)
	}

	err = auth.Register(following, password)
	if err != nil {
		t.Error(err)
	}

	token, _, err := auth.Login(user, password)
	if err != nil {
		t.Error(err)
	}
	if token == "" {
		t.Error("Token is empty")
	}

	err = AddFollowing(user, following)
	if err != nil {
		t.Error(err)
	}

	followings, err := GetFollowings(user)
	if err != nil {
		t.Error(err)
	}
	if len(followings) != 1 {
		t.Error("Expected 1 following, got", len(followings))
	}
	if followings[0] != following {
		t.Error("Expected following", following, "got", followings[0])
	}
}

// Test to check if getting followings is working
func TestGetFollowings(t *testing.T) {
	user := "testuser"
	password := "password"
	following := "following"

	err := auth.Register(user, password)
	if err != nil {
		t.Error(err)
	}

	err = auth.Register(following, password)
	if err != nil {
		t.Error(err)
	}

	token, _, err := auth.Login(user, password)
	if err != nil {
		t.Error(err)
	}
	if token == "" {
		t.Error("Token is empty")
	}

	err = AddFollowing(user, following)
	if err != nil {
		t.Error(err)
	}

	followings, err := GetFollowings(user)
	if err != nil {
		t.Error(err)
	}
	if len(followings) != 1 {
		t.Error("Expected 1 following, got", len(followings))
	}
	if followings[0] != following {
		t.Error("Expected following", following, "got", followings[0])
	}
}

// Test to check if deleting a following is working
func TestDeleteFollowing(t *testing.T) {
	user := "testuser"
	password := "password"
	following := "following"

	err := auth.Register(user, password)
	if err != nil {
		t.Error(err)
	}

	err = auth.Register(following, password)
	if err != nil {
		t.Error(err)
	}

	token, _, err := auth.Login(user, password)
	if err != nil {
		t.Error(err)
	}
	if token == "" {
		t.Error("Token is empty")
	}

	err = AddFollowing(user, following)
	if err != nil {
		t.Error(err)
	}

	err = DeleteFollowing(user, following)
	if err != nil {
		t.Error(err)
	}

	followings, err := GetFollowings(user)
	if err != nil {
		t.Error(err)
	}
	if len(followings) != 0 {
		t.Error("Expected 0 following, got", len(followings))
	}
}
