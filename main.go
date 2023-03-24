package main

import (
	"tokokuserviceapp/config"
	"tokokuserviceapp/menu"
)

func init() {
	// config.InitEnvironment()
	config.ConnectToDB()
}

func main() {
	menu.AuthMenu()

}
