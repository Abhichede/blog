package main

import (
	"fmt"

	posts "./posts"
)

// Create a blog which displayes the blog posts from the json

// to do that first we read the json file which contains the blogs

func main() {
	posts := posts.AllPosts()

	for i := 0; i < len(posts.Posts); i++ {
		fmt.Println("Title:", posts.Posts[i].Title)
		fmt.Println("Description:", posts.Posts[i].Description)
	}
}
