package menu

import (
	"fmt"
	"tokokuserviceapp/config"
	"tokokuserviceapp/controllers"
	"tokokuserviceapp/helpers"
)

func AuthMenu() {
	var authent = controllers.InitAuthen(config.DB)
	auth := -2
	for auth != -1 {
		fmt.Println("========== AUTH =========")
		fmt.Println("\t1. Login")
		fmt.Println("\t0. Exit")
		fmt.Println("=========================")
		fmt.Printf("Enter your choice : ")
		fmt.Scanln(&auth)
		if auth <= -1 || auth >= 2 {
			continue
		}
		if auth == 1 {
			if authent.Login() {
				fmt.Println("Masuk")
				MainMenu()
			}
		} else if auth == 0 {
			helpers.CloseCmd()
			break
		}
	}
}

func MainMenu() {
	var choice int
	for choice != 99 {
		helpers.StartCmd()
		fmt.Println("Welcome Admin!")
		fmt.Println("========== MENU =========")
		fmt.Println("1. Add User")
		fmt.Println("2. ")
		fmt.Println("3. ")
		fmt.Println("4. ")
		fmt.Println("5. ")
		fmt.Println("6. ")
		fmt.Println("7. ")
		fmt.Println("8. ")
		fmt.Println("99. Logout")
		fmt.Println("=========================")
		fmt.Printf("Enter your choice : ")
		fmt.Scanln(&choice)
		fmt.Println()
		// if choice <= -1 || choice >= 9 {
		// 	continue
		// }
		if choice == 1 {
			var Call controllers.Auth
			Call.AddPegawai()
		}
		// } else if choice == 2 {
		// } else if choice == 3 {
		// } else if choice == 4 {
		// } else if choice == 5 {
		// } else if choice == 6 {
		// } else if choice == 7 {
		// } else if choice == 8 {
		// } else if choice == 99 {
		// 	break
		// }
	}

}
