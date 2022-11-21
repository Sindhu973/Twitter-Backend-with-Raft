package memory

import (
	"errors"
	"sync"

	"garuda.com/m/web/cmd/auth/storage/structs"
)

type Memory struct {
	users sync.Map
}

func CreateNewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) AddUser(username string, hashedPassword string) error {
	if _, ok := m.users.Load(username); ok {
		return errors.New("User already exists")
	}
	m.users.Store(username, structs.User{Username: username, Hash_password: hashedPassword})
	return nil
}

func (m *Memory) GetUser(username string) (structs.User, error) {
	if _, ok := m.users.Load(username); !ok {
		return structs.User{}, errors.New("User not found")
	}
	user, _ := m.users.Load(username)
	return user.(structs.User), nil
}

func (m *Memory) UpdateUser(username, hash_password string) error {
	if _, ok := m.users.Load(username); !ok {
		return errors.New("User does not exist")
	}
	m.users.Store(username, structs.User{Username: username, Hash_password: hash_password})
	return nil
}

func (m *Memory) DeleteUser(username string) error {
	if _, ok := m.users.Load(username); !ok {
		return errors.New("User does not exist")
	}
	m.users.Delete(username)
	return nil
}

func (m *Memory) CreatePost(username string, title string, content string) error {
	if _, ok := m.users.Load(username); !ok {
		return errors.New("User does not exist")
	}
	user, _ := m.GetUser(username)
	user.Posts = append(user.Posts, structs.Post{Title: title, Content: content})
	m.users.Store(username, user)
	return nil
}

func (m *Memory) GetPosts(username string) ([]structs.Post, error) {
	if _, ok := m.users.Load(username); !ok {
		return nil, errors.New("User does not exist")
	}
	user, _ := m.GetUser(username)
	return user.Posts, nil
}

func (m *Memory) DeletePost(username string, title string) error {
	if _, ok := m.users.Load(username); !ok {
		return errors.New("User does not exist")
	}
	user, _ := m.GetUser(username)
	for i, post := range user.Posts {
		if post.Title == title {
			user.Posts = append(user.Posts[:i], user.Posts[i+1:]...)
			m.users.Store(username, user)
			return nil
		}
	}
	return errors.New("Post not found")
}

func (m *Memory) UpdatePost(username string, title string, content string) error {
	if _, ok := m.users.Load(username); !ok {
		return errors.New("User does not exist")
	}
	user, _ := m.GetUser(username)
	for i, post := range user.Posts {
		if post.Title == title {
			user.Posts[i].Content = content
			m.users.Store(username, user)
			return nil
		}
	}
	return errors.New("Post not found")
}

func (m *Memory) AddFollowing(follower, following string) error {
	if _, ok := m.users.Load(follower); !ok {
		return errors.New("User does not exist")
	}
	if _, ok := m.users.Load(following); !ok {
		return errors.New("User does not exist")
	}
	if follower == following {
		return errors.New("Cannot follow yourself")
	}
	user, _ := m.GetUser(follower)
	user.Following = append(user.Following, following)
	m.users.Store(follower, user)
	return nil
}

func (m *Memory) GetFollowings(username string) ([]string, error) {
	if _, ok := m.users.Load(username); !ok {
		return nil, errors.New("User does not exist")
	}
	user, _ := m.GetUser(username)
	return user.Following, nil
}

func (m *Memory) DeleteFollowing(follower, following string) error {
	if _, ok := m.users.Load(follower); !ok {
		return errors.New("User does not exist")
	}
	if _, ok := m.users.Load(following); !ok {
		return errors.New("User does not exist")
	}
	if follower == following {
		return errors.New("Cannot unfollow yourself")
	}
	user, _ := m.GetUser(follower)
	for i, follow := range user.Following {
		if follow == following {
			user.Following = append(user.Following[:i], user.Following[i+1:]...)
			m.users.Store(follower, user)
			return nil
		}
	}
	return errors.New("User not found")
}
