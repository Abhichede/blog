package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Abhichede/posts"
	_ "github.com/lib/pq"
)
var _posts = posts.AllPosts()
type Post struct {
	Id int
	Title, Description string
}

const (
	host = "localhost"
	port     = 5432
  user     = "techverito_abhijit"
  password = "postgres"
  dbname   = "go_practice"
)

func main() {
	http.Handle("/", http.HandlerFunc(LoadPosts))
	http.Handle("/post", http.HandlerFunc(LoadPost))

	// inserPost("Now second post", "This is now second post!!!")

	fmt.Println("Server is listening on :8085 ")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

//LoadPosts ...
func LoadPosts(w http.ResponseWriter, req *http.Request) {
	var templ = template.Must(template.New("posts").Parse(readPosts()))
	templ.Execute(w, getAllPosts())
}

//LoadPost ...
func LoadPost(w http.ResponseWriter, req *http.Request) {
	var templ = template.Must(template.New("post").Parse(readPost()))
	_ = templ
	id, _ := strconv.Atoi(req.FormValue("id"))
	templ.Execute(w, getAllPosts()[id])
}

func readPosts() string {
	postsTemplate, _ := ioutil.ReadFile("./templates/posts.html")
	return string(postsTemplate)
}
func readPost() string {
	postsTemplate, _ := ioutil.ReadFile("./templates/post.html")
	return string(postsTemplate)
}

func getConnection() (db *sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to postgres!")
	return db
}

func inserPost(title, description string)  {
	db := getConnection();
	defer db.Close()
	sqlStatement := `INSERT INTO posts(title, description) VALUES($1, $2)`;

	_, err := db.Exec(sqlStatement, title, description)
	checkErr(err)

	fmt.Println("Post added!!!")
}

func getAllPosts() map[int]Post {
	db := getConnection();
	defer db.Close()
	sqlStatement := `SELECT * FROM posts`;
	posts__, err := db.Query(sqlStatement)
	checkErr(err)
	defer posts__.Close()
	posts := make(map[int]Post)
	i := 1
	for posts__.Next() {
		var post Post
		err := posts__.Scan(&post.Id, &post.Title, &post.Description)
		checkErr(err)
		posts[i] = post
		i++
	}
	return posts
}

func checkCount(rows *sql.Rows) (count int) {
	count = 0
 	for rows.Next() {
		count++
  }
  return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}