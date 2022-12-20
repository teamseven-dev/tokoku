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
			}
		}
	}
}
