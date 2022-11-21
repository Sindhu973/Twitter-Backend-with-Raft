package structs

type User struct {
	Username      string
	Hash_password string
	Posts         []Post
	Following     []string
}

type Post struct {
	Title   string
	Content string
}
