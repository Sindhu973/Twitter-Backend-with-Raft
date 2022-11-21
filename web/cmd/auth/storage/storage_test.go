package storage

import (
	"testing"
)

// Test to check if the storage of memory type is working
func TestCreateNewStorage(t *testing.T) {
	_, err := CreateNewStorage("memory")
	if err != nil {
		t.Error(err)
	}
}

// Test to check if adding a single user to the storage of memory type is working
func TestCreateUsers(t *testing.T) {
	storage, err := CreateNewStorage("memory")
	testUser := "testuser"
	testPassword := "password"
	if err != nil {
		t.Error(err)
	}
	if storage == nil {
		t.Error("Storage is nil")
	}
	err = storage.AddUser(testUser, testPassword)
	if err != nil {
		t.Error(err)
	}
	user, err := storage.GetUser(testUser)
	if err != nil {
		t.Error(err)
	}
	if user.Username != testUser {
		t.Error("User with username " + testUser + " not found in memory storage")
	}
}

// Test to UpdateUser to the storage of memory type
func TestUpdateUser(t *testing.T) {
	storage, err := CreateNewStorage("memory")
	testUser := "testuser"
	testPassword := "password"
	if err != nil {
		t.Error(err)
	}
	if storage == nil {
		t.Error("Storage is nil")
	}
	err = storage.AddUser(testUser, testPassword)
	if err != nil {
		t.Error(err)
	}
	prevUser, err := storage.GetUser(testUser)
	err = storage.UpdateUser(testUser, "newpassword")
	if err != nil {
		t.Error(err)
	}
	currUser, err := storage.GetUser(testUser)
	if prevUser.Hash_password == currUser.Hash_password {
		t.Error("User with username " + testUser + " not updated in memory storage")
	}
}

// Test to DeleteUser to the storage of memory type
func TestDeleteUser(t *testing.T) {
	storage, err := CreateNewStorage("memory")
	testUser := "testuser"
	testPassword := "password"
	if err != nil {
		t.Error(err)
	}
	if storage == nil {
		t.Error("Storage is nil")
	}
	err = storage.AddUser(testUser, testPassword)
	if err != nil {
		t.Error(err)
	}
	err = storage.DeleteUser(testUser)
	if err != nil {
		t.Error(err)
	}
	_, err = storage.GetUser(testUser)
	if err == nil {
		t.Error("User with username " + testUser + " not deleted in memory storage")
	}
}

// Test to CreatePost to the storage of memory type
func TestCreatePost(t *testing.T) {
	storage, err := CreateNewStorage("memory")
	testUser := "testuser"
	testPassword := "password"
	if err != nil {
		t.Error(err)
	}
	if storage == nil {
		t.Error("Storage is nil")
	}
	err = storage.AddUser(testUser, testPassword)
	if err != nil {
		t.Error(err)
	}
	err = storage.CreatePost(testUser, "title", "content")
	if err != nil {
		t.Error(err)
	}
	posts, err := storage.GetPosts(testUser)
	if err != nil {
		t.Error(err)
	}
	if len(posts) != 1 {
		t.Error("Post not created")
	}
	if len(posts) > 0 {
		if posts[0].Title != "title" {
			t.Error("Post title not created")
		}
		if posts[0].Content != "content" {
			t.Error("Post content not created")
		}
	}
}

// Test to UpdatePost to the storage of memory type
func TestUpdatePost(t *testing.T) {
	user := "testuser"
	password := "password"
	title := "title"
	content := "content"
	storage, err := CreateNewStorage("memory")

	if err != nil {
		t.Error(err)
	}
	if storage == nil {
		t.Error("Storage is nil")
	}

	err = storage.AddUser(user, password)
	if err != nil {
		t.Error(err)
	}

	err = storage.CreatePost(user, title, content)
	if err != nil {
		t.Error(err)
	}

	err = storage.UpdatePost(user, title, "newcontent")
	if err != nil {
		t.Error(err)
	}

	posts, err := storage.GetPosts(user)
	if err != nil {
		t.Error(err)
	}

	if len(posts) != 1 {
		t.Error("Post not created")
	}

	if len(posts) > 0 {
		if posts[0].Title != title {
			t.Error("Post title not created")
		}
		if posts[0].Content != "newcontent" {
			t.Error("Post content not created")
		}
	}
}
