package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Abhichede/posts"
)
var _posts = posts.AllPosts()

func main() {
	http.Handle("/", http.HandlerFunc(LoadPosts))
	http.Handle("/post", http.HandlerFunc(LoadPost))
	fmt.Println("Server is listening on :8085 ")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

//LoadPosts ...
func LoadPosts(w http.ResponseWriter, req *http.Request) {
	var templ = template.Must(template.New("posts").Parse(readPosts()))
	templ.Execute(w, _posts)
}

//LoadPost ...
func LoadPost(w http.ResponseWriter, req *http.Request) {
	var templ = template.Must(template.New("post").Parse(readPost()))
	_ = templ
	id, _ := strconv.Atoi(req.FormValue("id"))
	templ.Execute(w, _posts.Posts[(id - 1)])
}

func readPosts() string {
	postsTemplate, _ := ioutil.ReadFile("./templates/posts.html")
	return string(postsTemplate)
}
func readPost() string {
	postsTemplate, _ := ioutil.ReadFile("./templates/post.html")
	return string(postsTemplate)
}
