package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Abhichede/blog/posts"
)

var templ = template.Must(template.New("posts").Parse(readPosts()))
var _posts = posts.AllPosts()

func main() {
	fmt.Println(_posts)
	http.Handle("/", http.HandlerFunc(LoadPosts))
	fmt.Println("Server is listening on :8085 ")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

//LoadPosts ...
func LoadPosts(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, _posts)
}

func readPosts() string {
	postsTemplate, _ := ioutil.ReadFile("./templates/posts.html")
	return string(postsTemplate)
}
