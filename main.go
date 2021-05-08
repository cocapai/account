package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, v := range r.URL.Query() {
		fmt.Println("key:", k, ",value:", v[0])
	}
	for k, v := range r.PostForm {
		fmt.Println("k:", k, "value:", v[0])
	}
	fmt.Println(r.Method)
	if r.Method == "POST" {
		po := r.PostForm
		fmt.Println("hello world", po["asdf"][0])
	}
	// t := template.New("Index")
	// t.Parse("<div id='templateTextDiv'>Hi,{{.name}},{{.someStr}}</div>")
	t, error := template.ParseFiles("./templates/index.html")
	data := map[string]string{
		"name":    "zeta",
		"someStr": "kaishila",
	}
	t.Execute(w, data)
	if error != nil {
		fmt.Println("cuowu", error)
	}
}

func main() {
	http.HandleFunc("/", myWeb)

	staticHandle := http.FileServer(http.Dir("./static"))
	http.Handle("/js/", staticHandle)

	fmt.Println("9988")

	err := http.ListenAndServe(":9988", nil)
	if err != nil {
		fmt.Println("有错误: ", err)
	}
}
