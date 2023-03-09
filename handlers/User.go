package handlers

import (
	"crud/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("views/principal.html", "views/layout.html")
	
	if err != nil {
		log.Fatal("Index view error")
	}
	view.Execute(w, nil)
}

func IndexUsers(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("views/listUsers.html", "views/layout.html")
	
	if err != nil {
		log.Fatal("List User view error")
	}

	users := models.IndexUser()

	view.Execute(w, users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("views/createUsers.html", "views/layout.html")

	if err != nil {
		fmt.Println("Create User view error")
		fmt.Println(err.Error())
	}

	view.Execute(w, nil)
}

func StoreUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		contra := r.FormValue("contra")

		models.StoreUser(nombre, correo, contra)

		http.Redirect(w, r, "/users", http.StatusFound)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteUser(id)
	http.Redirect(w, r, "/users", http.StatusFound)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("views/editUsers.html", "views/layout.html")
	if err != nil {
		fmt.Println("Edit User view error")
		fmt.Println(err.Error())
	}

	id := r.URL.Query().Get("id")
	user := models.ShowUser(id)

	view.Execute(w, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entro al hanlde update user")

	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("nombre")
		email := r.FormValue("correo")

		models.UpdateUser(id, name, email)

		http.Redirect(w, r, "/users", http.StatusFound)
	} else {
		fmt.Println("los metodos no cinciden")
	}
}