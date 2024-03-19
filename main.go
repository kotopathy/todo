package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type task struct {
	Name string
	Done bool
}

var t []task
var tpl *template.Template

func main() {
	t = []task{}

	tpl, _ = template.ParseGlob("*.html")
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/postform", postformHandler)
	http.ListenAndServe(":8080", nil)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcomeHandler is running")
	tpl.ExecuteTemplate(w, "index.html", t)
}

func postformHandler(w http.ResponseWriter, r *http.Request) {
	newtask := r.FormValue("newtask")
	postedTask := &task{newtask, false}
	jpostedTask, err := json.Marshal(postedTask)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jpostedTask))
	fmt.Fprintf(w, "Имя: %s Возраст: %s", name, age)
}
