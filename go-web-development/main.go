package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var serverPort string
var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}
func handleRequests() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", home)
	http.HandleFunc("/about-us", aboutus)
	http.HandleFunc("/services", services)
	http.HandleFunc("/projects", projects)
	http.HandleFunc("/blogs", blogs)
	http.HandleFunc("/shop", shop)
	http.HandleFunc("/contact-us", contact)
	log.Fatal(http.ListenAndServe(serverPort, nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello world")
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
func aboutus(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello world")
	tmpl.ExecuteTemplate(w, "about.html", nil)
}
func services(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello world")
	tmpl.ExecuteTemplate(w, "services-dark.html", nil)
}
func projects(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello world")
	tmpl.ExecuteTemplate(w, "projects-classic.html", nil)
}
func blogs(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello world")
	tmpl.ExecuteTemplate(w, "blog.html", nil)
}
func shop(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello world")
	tmpl.ExecuteTemplate(w, "shop.html", nil)
}
func contact(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello world")
	tmpl.ExecuteTemplate(w, "contact.html", nil)
}

func main() {

	serverPort = ":9002"
	fmt.Printf("Website Started at http://localhost%v", serverPort)
	handleRequests()
}
