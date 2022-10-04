package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Users struct {
	Name string
	Age int
	Email string
	Password string
}

type UserDetail struct {
	Name string
	Age int
}

var idx int

var port = ":8000"

var user = []Users{
	{Name: "Budi", Age: 19, Email: "budi@gmail.com", Password: "123456"},
	{Name: "Amir", Age: 20, Email: "amir@gmail.com", Password: "123457"},
	{Name: "Indra", Age: 21, Email: "indra@gmail.com", Password: "123458"},
}

var userDetail = []UserDetail{}

func main () {
	http.HandleFunc("/submit", submit)
	http.HandleFunc("/", root)
	http.HandleFunc("/detail", detail)
	http.HandleFunc("/login", login)
	http.ListenAndServe(port, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func submit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	email := r.FormValue("email")
	if r.Method == "POST" {
		for index := range user{
			if(user[index].Email == email) {
				idx = index
				userDetail = []UserDetail{
					{Name: user[index].Name, Age: user[index].Age},
				}
				http.Redirect(w, r, "/detail", http.StatusSeeOther)
				return
			}
		}
	}

	return
}

func login(w http.ResponseWriter, r *http.Request) {
	var tpl, err = template.ParseFiles("login.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, nil)
}

func detail(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	fmt.Fprint(w, email)
	var tpl, err = template.ParseFiles("index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, userDetail)
	return
}
