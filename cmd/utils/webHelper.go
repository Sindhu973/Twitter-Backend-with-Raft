package utils

func AddAuthorToPosts(author string, posts *[]map[string]string) {
	for _, post := range *posts {
		post["author"] = author
	}
	return
}
