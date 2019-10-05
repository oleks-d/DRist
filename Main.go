package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.gohtml", "templates/header.gohtml", "templates/footer.gohtml")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/login.gohtml", "templates/header.gohtml", "templates/footer.gohtml")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "login", nil)

}

func newThemeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/new.gohtml", "templates/header.gohtml", "templates/footer.gohtml")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "newtheme", nil)
}

func addThemeHandler(w http.ResponseWriter, r *http.Request) {
	var name string = r.FormValue("name")
	fmt.Println("Added " + name)

	theme := models.newTheme{"", "", "", "", ""}
	fmt.Println("Added " + theme.name)
}

func main() {
	fmt.Println("Welcome!")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/new", newThemeHandler)

	http.HandleFunc("/add_theme", addThemeHandler)

	http.ListenAndServe(":8080", nil)
}
