package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type URL struct {
	ID            string    `json:"Id"`
	Original_Url  string    `json:"original_url"`
	Shorten_Url   string    `json:"shorten_url"`
	Creation_Time time.Time `json:"created_time"`
}

type PageData struct {
	ShortenedURL string
}

var PORT int = 3000
var serverurl string = "http://localhost:" + strconv.Itoa(PORT)
var tmpl = template.Must(template.ParseFiles("index.html"))
var Url_DB = make(map[string]URL)

func generateShortUrl(url string) string {
	hasher := md5.New()
	hasher.Write([]byte(url))

	data := hasher.Sum(nil)
	finalstr := hex.EncodeToString(data)

	return finalstr[:8] // Returns first 8 chars
}

func createUrl(originalUrl string) string {
	shortedUrl := generateShortUrl(originalUrl)

	id := shortedUrl
	Url_DB[id] = URL{
		ID:            id,
		Original_Url:  originalUrl,
		Shorten_Url:   shortedUrl,
		Creation_Time: time.Now(),
	}

	return shortedUrl
}

func getURL(id string) (URL, error) {
	url, ok := Url_DB[id]
	if !ok {
		return URL{}, fmt.Errorf("url not found")
	}
	return url, nil
}

func shortUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		originalUrl := r.FormValue("url")
		shortUrl := createUrl(originalUrl)
		// Render the form with the shortened URL
		tmpl.Execute(w, PageData{ShortenedURL: serverurl + "/redirect/" + shortUrl})
	} else {
		// Render the empty form
		tmpl.Execute(w, PageData{})
	}
}

func redirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]

	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusNotFound)
		return
	}

	// Redirect to the original URL
	http.Redirect(w, r, url.Original_Url, http.StatusFound)
}

func main() {
	fmt.Println("Making URL Shortner...")

	http.HandleFunc("/shorturl", shortUrlHandler)
	http.HandleFunc("/redirect/", redirectUrlHandler)

	// Start the server
	fmt.Println("Server running on port", PORT)
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
	if err != nil {
		fmt.Println("Error in starting server:", err)
	}
}
