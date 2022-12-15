package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"
    "html/template"
    "net/http"
	"strconv"

)

func main() {
	http.HandleFunc("/", blog)
	http.HandleFunc("/admin", login)

	err := http.ListenAndServe(":8080", nil) 
    if err != nil {
        panic(err.Error())
    }
}




func blog(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")

	if err != nil {
		panic(err.Error())
	}else if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}



func login(w http.ResponseWriter, r *http.Request) {

	Author := ""
	Post := 0
	Name := ""
	Intro := ""
	Body := ""

	if r.Method == "POST" {
		r.ParseForm()
		
		author := [][]string{r.Form["Author"]}
		Author = author[0][0]
		
		postarray := [][]string{r.Form["post"]}
		Post, _ = strconv.Atoi(postarray[0][0])

		name := [][]string{r.Form["name"]}
		Name = name[0][0]

		intro := [][]string{r.Form["intro"]}
		Intro = intro[0][0]

		body := [][]string{r.Form["body"]}
		Body = body[0][0]

		fmt.Println(Author,Post,Name,Intro,Body)


		db, err := sqlConnect()
		if err != nil {
			panic(err.Error())
		}
		post := Posts{
			Author: Author,
			Post:   Post,
			Name:   Name,
			Intro:  Intro,
			Body:   Body,
		}
	
		db.Create(&post)

		t, err := template.ParseFiles("complete.html")

		if err != nil {
			panic(err.Error())
		}else if err := t.Execute(w, nil); err != nil {
			panic(err.Error())
		}
	}else {
		t, err := template.ParseFiles("admin.html")

		if err != nil {
			panic(err.Error())
		}else if err := t.Execute(w, nil); err != nil {
			panic(err.Error())
		}
	}
}








func sqlConnect() (database *gorm.DB, err error) {
	dsn := "root:chacha0503@tcp(127.0.0.1:3306)/goProject?charset=utf8&parseTime=True&loc=Local"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}





type Posts struct {
	Author string
	Post   int    
	Name   string 
	Intro  string 
	Body   string 
}