package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Post represents a blog post
type Post struct {
	Title  string `json:"title"`
	Routes string `json:"function"`
}

// Sample data for demonstration purposes
var homesite = []Post{
	{Title: "Learn Goland Basic", Routes: "/learn"},
	{Title: "Blogs", Routes: "/blog"},
}
var posts = []Post{
	{Title: "Hello World", Routes: "/hello-world"},
	{Title: "Variables in Go", Routes: "/variables"},
	{Title: "Data Types in Go", Routes: "/datatype"},
	{Title: "Arrays in Go", Routes: "/arrays"},
}

func blog(c *gin.Context) {
	c.HTML(http.StatusOK, "blog.html", nil)
}

func learn(c *gin.Context) {
	c.HTML(http.StatusOK, "learn.html", nil)
}

func helloworld(c *gin.Context) {
	c.HTML(http.StatusOK, "hello-world.html", nil)
}

func variables(c *gin.Context) {
	c.HTML(http.StatusOK, "variables.html", nil)
}

func datatype(c *gin.Context) {
	c.HTML(http.StatusOK, "datatype.html", nil)
}

func arrays(c *gin.Context) {
	c.HTML(http.StatusOK, "arrays.html", nil)
}

func search(c *gin.Context) {
	query := c.Query("q")
	var results []Post

	if query != "" {
		for _, post := range posts {
			if strings.Contains(strings.ToLower(post.Title), strings.ToLower(query)) {
				results = append(results, post)
			}
		}
	}

	c.HTML(http.StatusOK, "search.html", gin.H{
		"query":   query,
		"results": results,
	})
}
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1200"
		log.Println("No PORT environment variable set. Using default port:", port)
	} else {
		log.Println("Using port from environment variable:", port)
	}

	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	dir, _ := os.Getwd()
	fmt.Println("Current working directory:", dir)

	files, err := os.ReadDir("template/*")
	if err != nil {
		log.Fatal("Error reading template directory:", err)
	}
	for _, file := range files {
		fmt.Println("Template file:", file.Name())
	}
	//
	r.GET("/", func(c *gin.Context) {
		fmt.Println("Posts:", posts)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"posts": posts})
	})

	r.GET("/hello-world", helloworld)
	r.GET("/blog", blog)
	r.GET("/learn", learn)
	r.GET("/variables", variables)
	r.GET("/datatype", datatype)
	r.GET("/arrays", arrays)
	r.GET("/search", search)

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
