package models

import (
	"crud/config"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id int
	Name string
	Email string
	Active bool
}


func HashPassword(pPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pPassword), 14)
	return string(bytes), err
}

func CheckHash(pPassword string, pHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pHash), []byte(pPassword))
	return err == nil 
}

func IndexUser() []User {
	connection := config.Connection()

	rows, err := connection.Query("SELECT id, name, email, active  FROM users WHERE active = 1")

	if err != nil {
		fmt.Println("Error al listar usuarios : ")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Users Retrieved:")
	}

	connection.Close()

	//
	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.Id, &u.Name, &u.Email, &u.Active)
		users = append(users, u)
	}
	return users
}

func ShowUser(pId string) User {
	conection := config.Connection()

	query := "SELECT id, name, email, active FROM users WHERE id = ?"

	var u User
	err := conection.QueryRow(query, pId).Scan(&u.Id, &u.Name, &u.Email, &u.Active)

	if err != nil {
		fmt.Println("Error al buscar usuario : ")
		fmt.Println(err.Error())
	} else {
		fmt.Println("User Retrieved:")
	}

	return u
}

func StoreUser(pNombre string, pCorreo string, pPassword string ) {
	// DB connection
	connection := config.Connection()

	// prepare the query
	query, err := connection.Prepare("INSERT INTO users (name, email, password, active) VALUES (?, ?, ?, ?)")
	// handle errs
	if err != nil {
		fmt.Println("Error al preparar consulta de creación : ")
		fmt.Println(err.Error())
	}

	// encrypt password
	securePassword, err := HashPassword(pPassword)
	if err != nil {
		fmt.Println("Error al cifrar password!")
	} else {
		fmt.Println("Password cifrada!")
	} 

	// execute query
	result, qErr := query.Exec(pNombre, pCorreo, securePassword, 1)
	// handle errs
	if qErr != nil {
		fmt.Println("Error al crear usuario : ")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Usuario Creado : ", result)
	}

	// close connection
	connection.Close()
}

func DeleteUser(pId string) {
	// DB connection
	connection := config.Connection()

	// prepare the query
	query, err := connection.Prepare("UPDATE users SET active = 0 WHERE id = ?")
	// handle errs
	if err != nil {
		fmt.Println("Error al preparar consulta de eliminación : ")
		fmt.Println(err.Error())
	}

	// execute query
	result, qErr := query.Exec(pId)
	// handle errs
	if qErr != nil {
		fmt.Println("Error al eliminar usuario : ")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Usuario Eliminado : ", result)
	}

	// close connection
	connection.Close()
}

func UpdateUser(pId string, pName string, pEmail string) {
	// DB connection
	connection := config.Connection()

	// prepare the query
	query, err := connection.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	// handle errs
	if err != nil {
		fmt.Println("Error al preparar consulta de actualización : ")
		fmt.Println(err.Error())
	}

	// execute query
	result, qErr := query.Exec(pName, pEmail, pId)
	// handle errs
	if qErr != nil {
		fmt.Println("Error al actualizar usuario : ")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Usuario actulizado : ", result)
	}

	// close connection
	connection.Close()
}