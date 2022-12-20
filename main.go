package main

import (
	"fmt"
	"tokoku/config"
	"tokoku/staff"
)

func main() {
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var staffMenu = staff.StaffMenu{DB: conn}

	fmt.Println(conn)
	var inputMenu = 1
	for inputMenu != 0 {
		fmt.Println("Welcome To Tokoku")
		fmt.Print("1. Login\n0. Exit\nInsert Your Menu :")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			var inputName, inputPassword string
			fmt.Print("Username :")
			fmt.Scanln(&inputName)
			fmt.Print("Password :")
			fmt.Scanln(&inputPassword)
			res, err := staffMenu.Login(inputName, inputPassword)
			if err != nil {
				fmt.Println(err.Error())
			}
			if res > 0 {
				fmt.Println("Success Login ")
				if res > 1 {
					islogin := true
					for islogin {
						fmt.Println("Staff Menu")
						fmt.Println("1. Transaction")
						fmt.Println("2. Add a new product")
						fmt.Println("3. Update product stock")
						fmt.Println("4. Update product information")
						fmt.Println("5. Add a new customer")
						fmt.Println("9. Logout")
						fmt.Print("Please Insert Menu :")
						var choice int
						fmt.Scanln(&choice)
						switch choice {
						case 1:
							{
								fmt.Println("ini transaction")

							}
						case 2:
							{
								fmt.Println("add new product")
							}
						case 3:
							{
								fmt.Println("update product stock")
							}
						case 4:
							{
								fmt.Println("updated product information")
							}
						case 5:
							{
								fmt.Println("add new customer")
							}
						case 9:
							{
								fmt.Println("bye")
								islogin = false
							}
						}
					}
				} else if res == 1 {
					islogin := true
					for islogin {
						fmt.Println("Admin Menu")
						fmt.Println("1. Add a new staff")
						fmt.Println("2. Remove a staff")
						fmt.Println("3. Remove a product")
						fmt.Println("4. Remove a transaction")
						fmt.Println("5. Remove a customer")
						fmt.Println("9. Logout")
						fmt.Print("Please Insert Menu :")
						var choice int
						fmt.Scanln(&choice)
						switch choice {
						case 1:
							{
								fmt.Println("add a new staff")

							}
						case 2:
							{
								fmt.Println("remove a staff")
							}
						case 3:
							{
								fmt.Println("remove product")
							}
						case 4:
							{
								fmt.Println("remove transaction")
							}
						case 5:
							{
								fmt.Println("remove customer")
							}
						case 9:
							{
								fmt.Println("bye")
								islogin = false
							}
						}

					}

				}
			}
		}
	}
}