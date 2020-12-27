package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Superhero struct {
	Name    string
	Alias   string
	Friends []string
}

func (s Superhero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

// Muncul Otomatis

// Test 1
// func handlerIndex(w http.ResponseWriter, r *http.Request) {
// 	var message = "Welcome"
// 	w.Write([]byte(message))
// }
// func handlerHello(w http.ResponseWriter, r *http.Request) {
// 	var message = "Hello world!"
// 	w.Write([]byte(message))
// }

// Test 5
// type M map[string]interface{}

// Test 6
// type Info struct {
// 	Affiliation string
// 	Address     string
// }
// type Person struct {
// 	Name    string
// 	Gender  string
// 	Hobbies []string
// 	Info    Info
// }

func main() {
	// Test 1
	// http.HandleFunc("/", handlerIndex)
	// http.HandleFunc("/index", handlerIndex)
	// http.HandleFunc("/hello", handlerHello)
	// var address = "localhost:9000"
	// fmt.Printf("server started at %s\n", address)
	// err := http.ListenAndServe(address, nil)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// Test 2
	// handlerIndex := func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello"))
	// }
	// http.HandleFunc("/", handlerIndex)
	// http.HandleFunc("/index", handlerIndex)
	// http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello again"))
	// })
	// fmt.Println("server started at localhost:9000")
	// http.ListenAndServe(":9000", nil)

	// Test 3
	// http.Handle("/static/",
	// 	http.StripPrefix("/static/",
	// 		http.FileServer(http.Dir("assets"))))
	// fmt.Println("server started at localhost:9000")
	// http.ListenAndServe(":9000", nil)

	// Test 4
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// not yet implemented
	// 	var filepath = path.Join("views", "index.html")
	// 	var tmpl, err = template.ParseFiles(filepath)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	var data = map[string]interface{}{
	// 		"title": "Learning Golang Web",
	// 		"name":  "Batman",
	// 	}
	// 	err = tmpl.Execute(w, data)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })
	// fmt.Println("server started at localhost:9000")
	// http.ListenAndServe(":9000", nil)

	// Test 5
	// var tmpl, err = template.ParseGlob("views/*")
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }
	// http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
	// 	var data = M{"name": "Batman"}
	// 	err = tmpl.ExecuteTemplate(w, "index", data)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })
	// http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	// 	var data = M{"name": "Batman"}
	// 	err = tmpl.ExecuteTemplate(w, "about", data)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })
	// fmt.Println("server started at localhost:9000")
	// http.ListenAndServe(":9000", nil)

	// Test 6
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	var person = Person{
	// 		Name:    "Bruce Wayne",
	// 		Gender:  "male",
	// 		Hobbies: []string{"Reading Books", "Traveling", "Buying things"},
	// 		Info:    Info{"Wayne Enterprises", "Gotham City"},
	// 	}
	// 	var tmpl = template.Must(template.ParseFiles("view.html"))
	// 	if err := tmpl.Execute(w, person); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })
	// fmt.Println("server started at localhost:9000")
	// http.ListenAndServe(":9000", nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Superhero{
			Name:    "Bruce Wayne",
			Alias:   "Batman",
			Friends: []string{"Superman", "Flash", "Green Lantern"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
