package posts

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

// Posts contains array of all posts
type Posts struct {
	Posts []Post `json:"posts"`
}

// Post is single
type Post struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// JSONParser which parses the json file provided
func JSONParser(filename string) Posts {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error while opening the file: ", err)
	}
	defer file.Close()

	jsonBytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal("Error while Parsing json: ", err)
	}

	var posts Posts

	json.Unmarshal(jsonBytes, &posts)

	return posts
}

func AllPosts() Posts {
	blogPosts := flag.String("src", "blogs.json", "A JSON file which contains the blog posts.")
	flag.Parse()
	return JSONParser(*blogPosts)
}
