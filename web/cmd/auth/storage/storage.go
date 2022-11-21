package storage

import (
	"errors"

	"garuda.com/m/web/cmd/auth/storage/memory"
	"garuda.com/m/web/cmd/auth/storage/structs"
)

type Storage interface {
	AddUser(string, string) error
	GetUser(string) (structs.User, error)
	UpdateUser(string, string) error
	DeleteUser(username string) error
	CreatePost(username string, title string, content string) error
	GetPosts(username string) ([]structs.Post, error)
	UpdatePost(username string, title string, content string) error
	DeletePost(username string, title string) error
	AddFollowing(follower, following string) error
	GetFollowings(username string) ([]string, error)
	DeleteFollowing(follower, following string) error
}

func CreateNewStorage(storageType string) (Storage, error) {
	if storageType == "memory" {
		return memory.CreateNewMemory(), nil
	}
	if storageType == "raft" {
		return nil, errors.New("Not implemented")
	}
	return nil, errors.New("Unidentified storage type")
}

var OurStore, _ = CreateNewStorage("memory")
