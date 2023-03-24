package controllers

import (
	"database/sql"
	"fmt"
	"tokokuserviceapp/config"
	"tokokuserviceapp/entities"

	_ "github.com/go-sql-driver/mysql"
)

var LoggedInUser entities.User

type Auth struct {
	DB   *sql.DB
	Data []entities.User
}

func InitAuthen(DB *sql.DB) *Auth {
	return &Auth{DB: DB}

}

func (Auth *Auth) Login() bool {
	user, username, password := entities.User{}, "", ""
	fmt.Println("=============== Login ==============")
	fmt.Printf("Enter Username \t: ")
	fmt.Scanln(&username)
	fmt.Printf("Enter Password \t: ")
	fmt.Scanln(&password)
	fmt.Println()

	querySelectUser := "SELECT u.id, u.role, u.name, u.username,u.password FROM users u WHERE u.username = ?"
	err := config.DB.QueryRow(querySelectUser, username).Scan(&user.ID, &user.Role, &user.Nama, &user.Username, &user.Password)
	Auth.Data = append(Auth.Data, user)
	if Auth.Login() {
		if LoggedInUser.Role == "admin" {
			fmt.Println("User is admin")
		} else if LoggedInUser.Role == "pegawai" {
			fmt.Println("User is pegawai")
		}
	} else {
		fmt.Println("Login failed")
	}
	if err != nil {
		return false
	}
	return true
}

func (Auth *Auth) AddPegawai() bool {
	if LoggedInUser.Role != "admin" {
		fmt.Println("Error: Only admin can add pegawai")
		return false
	}

	var user entities.User
	fmt.Println("=============== Tambah Pegawai ==============")
	fmt.Printf("Enter Nama \t\t: ")
	fmt.Scanln(&user.Nama)
	fmt.Printf("Enter Username \t\t: ")
	fmt.Scanln(&user.Username)
	fmt.Printf("Enter Password \t\t: ")
	fmt.Scanln(&user.Password)
	fmt.Printf("Enter Role \t\t: ")
	fmt.Scanln(&user.Role)

	fmt.Println()

	queryInsertUser := "INSERT INTO users (name, username, password, role) VALUES (?, ?, ?, ?)"
	result, err := Auth.DB.Exec(queryInsertUser, user.Nama, user.Username, user.Password, user.Role)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected > 0 {
		fmt.Println("Pegawai berhasil ditambahkan")
		return true
	} else {
		fmt.Println("Pegawai gagal ditambahkan")
		return false
	}
}
