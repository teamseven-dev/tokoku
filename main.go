package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tokoku/config"
	"tokoku/customer"
	"tokoku/product"
	"tokoku/staff"
)

func main() {
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var staffMenu = staff.StaffMenu{DB: conn}
	var custmenu = customer.CustMenu{DB: conn}
	var productMenu = product.ProductMenu{DB: conn}

	var inputMenu = 1

	for inputMenu != 0 {
		fmt.Println("Welcome To Tokoku")
		fmt.Print("1. Login\n0. Exit\nInsert Your Menu : ")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			var inputName, inputPassword string
			fmt.Print("Username : ")
			fmt.Scanln(&inputName)
			fmt.Print("Password : ")
			fmt.Scanln(&inputPassword)
			res, err := staffMenu.Login(inputName, inputPassword)
			if err != nil {
				fmt.Println(err.Error())
			}
			if res.ID > 0 {
				fmt.Println("------------------")
				fmt.Println("Logged in succesfully!")
				fmt.Println("=======================")
				if res.ID > 1 {
					islogin := true
					for islogin {
						fmt.Printf("Welcome staff, %s!\n", inputName)
						fmt.Println("------------------")
						fmt.Println("1. Transaction")
						fmt.Println("2. Insert a new product")
						fmt.Println("3. Show products")
						fmt.Println("4. Add a product stock")
						fmt.Println("5. Update a product name")
						fmt.Println("6. Insert a new customer")
						fmt.Println("9. Logout")
						fmt.Print("Please choose a menu [1,2,3,4,5,6,9] : ")

						var choice int
						fmt.Scanln(&choice)
						fmt.Println("=======================")
						
						switch choice {
						case 1:
							
							// TRANSACTION

						case 2:

							// INSERT A NEW PRODUCT
							inputProduct := product.Product{}
							inputProduct.IDStaff = res.ID
							fmt.Println("INSERT A NEW PRODUCT")
							fmt.Println("------------------")
							fmt.Print("Insert product name : ")
							consoleReader := bufio.NewReader(os.Stdin)
							newProd, _ := consoleReader.ReadString('\n')
							newProd = strings.TrimSuffix(newProd, "\n")
							inputProduct.Name = newProd
							fmt.Print("Insert product quantity : ")
							fmt.Scanln(&inputProduct.Qty)

							prodRes, err := productMenu.Insert(inputProduct)
							if err != nil {
								fmt.Println("------------------")
								fmt.Println(err.Error())
							} else {
								fmt.Println("------------------")
								fmt.Println("Inserted a new product successfully!")
								fmt.Println("=======================")
							}
							inputProduct.ID = prodRes

						case 3:
							
							// SHOW PRODUCTS
							prodMenu := 1

							for prodMenu != 9 {
								fmt.Println("LIST OF PRODUCTS")
								fmt.Println("------------------")
								products, _ := productMenu.Show()
								for i := 0; i < len(products); i++ {
								fmt.Println("Product Code   : ", products[i].ID)
								fmt.Println("Product Name   : ", products[i].Name)
								fmt.Println("QTY            : ", products[i].Qty)
								fmt.Println("Staff Name     : ", products[i].StaffName)
								fmt.Println("------------------")
								}

								fmt.Println("1. Delete a product")
								fmt.Println("9. Back to main menu")
								fmt.Print("Please choose a menu [1, 9] : ")
								fmt.Scanln(&prodMenu)
								
								if prodMenu == 1 {
									var productName string
									fmt.Println("=======================")
									fmt.Println("DELETE A PRODUCT")
									fmt.Println("------------------")
									fmt.Print("Please insert a product name : ")
									consoleReader := bufio.NewReader(os.Stdin)
									productName, _ = consoleReader.ReadString('\n')
									productName = strings.TrimSuffix(productName, "\n")

									res, err := productMenu.Delete(productName)

									if err != nil {
										fmt.Println("------------------")
										fmt.Println(err.Error())
										fmt.Println("=======================")
									}

									if res {
										fmt.Println("------------------")
										fmt.Printf("Product `%s` has been deleted successfully.\n", productName)
										fmt.Println("=======================")
									}

								} else {
									fmt.Println("=======================")
								}
							}
							
						case 4:
							
							// UPDATE A PRODUCT STOCK
							prodMenu := 1

							for prodMenu != 9 {
								fmt.Println("LIST OF PRODUCTS")
								fmt.Println("------------------")
								products, _ := productMenu.Show()
								for i := 0; i < len(products); i++ {
								fmt.Println("Product Code   : ", products[i].ID)
								fmt.Println("Product Name   : ", products[i].Name)
								fmt.Println("QTY            : ", products[i].Qty)
								fmt.Println("Staff Name     : ", products[i].StaffName)
								fmt.Println("------------------")
								}

								fmt.Println("1. add stock")
								fmt.Println("9. Back to main menu")
								fmt.Print("Please choose a menu [1, 9] : ")
								fmt.Scanln(&prodMenu)
								
								if prodMenu == 1 {
									var prodID, addQty int
									fmt.Println("=======================")
									fmt.Println("ADD A PRODUCT STOCK")
									fmt.Println("------------------")
									fmt.Print("Please insert product id : ")
									fmt.Scanln(&prodID)
									fmt.Print("Please insert additional Qty : ")
									fmt.Scanln(&addQty)

									res, err := productMenu.UpdateStock(addQty, prodID)

									if err != nil {
										fmt.Println("------------------")
										fmt.Println(err.Error())
										fmt.Println("=======================")
									}

									if res {
										fmt.Println("------------------")
										fmt.Println("Added a product stock succesfully!")
										fmt.Println("=======================")
									}

								} else {
									fmt.Println("=======================")
								}
							}

						case 5:
							
							// UPDATE A PRODUCT NAME
							prodMenu := 1

							for prodMenu != 9 {
								fmt.Println("LIST OF PRODUCTS")
								fmt.Println("------------------")
								products, _ := productMenu.Show()
								for i := 0; i < len(products); i++ {
								fmt.Println("Product Code   : ", products[i].ID)
								fmt.Println("Product Name   : ", products[i].Name)
								fmt.Println("QTY            : ", products[i].Qty)
								fmt.Println("Staff Name     : ", products[i].StaffName)
								fmt.Println("------------------")
								}

								fmt.Println("1. update a product name")
								fmt.Println("9. Back to main menu")
								fmt.Print("Please choose a menu [1, 9] : ")
								fmt.Scanln(&prodMenu)
								
								if prodMenu == 1 {
									var prodID int 
									var newName string
									fmt.Println("=======================")
									fmt.Println("UPDATE A PRODUCT NAME")
									fmt.Println("------------------")
									fmt.Print("Please insert product id : ")
									fmt.Scanln(&prodID)
									fmt.Print("Please insert new name : ")
									fmt.Scanln(&newName)

									res, err := productMenu.UpdateName(newName, prodID)

									if err != nil {
										fmt.Println("------------------")
										fmt.Println(err.Error())
										fmt.Println("=======================")
									}

									if res {
										fmt.Println("------------------")
										fmt.Println("Updated a product name succesfully!")
										fmt.Println("=======================")
									}

								} else {
									fmt.Println("=======================")
								}
							}

						case 6:

							// INSERT A NEW CUSTOMER
							fmt.Println("INSERT A NEW CUSTOMER")
							var CusName string
							fmt.Println("Please Insert Data Customer")
							fmt.Print("New Customer Name :")
							fmt.Scanln(&CusName)
							ifada, err := custmenu.AddCustomer(CusName, res.ID)
							if ifada == true {
								fmt.Println("Success Add Customer")
							} else {
								fmt.Println("Sorry Can't Add Customer")
							}
							if err != nil {
								fmt.Println(err.Error())
							}

						case 9:

							// LOGOUT
							fmt.Println("Logged out succesfully!")
							fmt.Println("=======================")
							islogin = false

						}
					}
				} else if res.ID == 1 {
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
							var newStaff staff.Staff
							fmt.Println("\n", "== Insert New Staff ==")
							fmt.Print("Masukkan nama: ")
							fmt.Scanln(&newStaff.Name)
							fmt.Print("Masukkan password: ")
							fmt.Scanln(&newStaff.Password)
							res, err := staffMenu.Register(newStaff)
							if err != nil {
								fmt.Println(err.Error())
							}
							if res {
								fmt.Println("Sukses mendaftarkan data")
							} else {
								fmt.Println("Gagal mendaftarn data")
							}
						case 2:
							var removeStaff staff.Staff
							fmt.Println("\n", "== Remove Staff ==")
							fmt.Print("Staff name: ")
							fmt.Scanln(&removeStaff.Name)
							res, err := staffMenu.Remove(removeStaff.Name)
							if err != nil {
								fmt.Println(err.Error())
							}
							if res {
								fmt.Println("Berhasil menghapus data")
							} else {
								fmt.Println("Gagal menghapus data")
							}
						case 3:
							fmt.Println("remove product")
						case 4:
							fmt.Println("remove transaction")
						case 5:
							{
								var namacus string
								fmt.Println("Delete Customer Menu")
								fmt.Print("Insert Customer Name :")
								fmt.Scanln(&namacus)
								ifberhasil, err := custmenu.RemoveCustomer(namacus)
								if err != nil {
									fmt.Println(err.Error())
								}
								if ifberhasil {
									fmt.Println("Data Customer ", namacus, " Success To Delete")
								} else {
									fmt.Println("Sorry Can't Delete Customer")
								}
							}
							fmt.Println("remove customer")
						case 9:
							fmt.Println("bye")
							islogin = false
						}
					}
				}
			} 
		}
	}
}