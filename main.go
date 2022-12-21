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
	"tokoku/transaction"
)

func main() {
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var staffMenu = staff.StaffMenu{DB: conn}
	var custMenu = customer.CustMenu{DB: conn}
	var productMenu = product.ProductMenu{DB: conn}
	var transactionMenu = transaction.TransMenu{DB: conn}

	var inputMenu = 1

	for inputMenu != 0 {
		fmt.Println("WELCOME TO TOKOKU")
		fmt.Println("------------------")
		fmt.Print("1. Login\n0. Exit\n------------------\nPlease choose a menu [1, 0] : ")
		fmt.Scanln(&inputMenu)
		fmt.Println("=======================")
		if inputMenu == 1 {
			var inputName, inputPassword string
			fmt.Println()
			fmt.Println("LOGIN MENU")
			fmt.Println("------------------")
			fmt.Print("Please insert your username : ")
			fmt.Scanln(&inputName)
			fmt.Print("Please insert your password : ")
			fmt.Scanln(&inputPassword)
			res, err := staffMenu.Login(inputName, inputPassword)
			if err != nil {
				fmt.Println("------------------")
				fmt.Println(err.Error())
				fmt.Println("------------------")

			}
			if res.ID > 0 {
				fmt.Println("=======================")
				fmt.Println("Logged in succesfully!")
				fmt.Println("=======================")
				if res.ID > 1 {
					islogin := true
					for islogin {
						fmt.Println()
						fmt.Printf("WELCOME STAFF, %s!\n", inputName)
						fmt.Println("------------------")
						fmt.Println("transaction menu")
						fmt.Println("------------------")
						fmt.Println("1. New transaction")
						fmt.Println("2. Transactions history")
						fmt.Println("------------------")
						fmt.Println("product menu")
						fmt.Println("------------------")
						fmt.Println("3. Insert a new product")
						fmt.Println("4. Show products")
						fmt.Println("5. Add a product stock")
						fmt.Println("6. Update a product name")
						fmt.Println("------------------")
						fmt.Println("customer menu")
						fmt.Println("------------------")
						fmt.Println("7. Insert a new customer")
						fmt.Println("8. Show customers")
						fmt.Println("------------------")
						fmt.Println("9. Logout")
						fmt.Println("------------------")
						fmt.Print("Please choose a menu [1, 2, 3, 4, 5, 6, 7, 8, 9] : ")

						var choice int
						fmt.Scanln(&choice)
						fmt.Println("=======================")

						switch choice {
						case 1:

							// NEW TRANSACTION
							fmt.Println(transactionMenu.ShowTransaction(1))
							fmt.Println(transactionMenu.ShowTransaction(2))
							fmt.Println(transactionMenu.ShowTransaction(3))
							fmt.Println(transactionMenu.ShowTransaction(4))
						case 2:

							// TRANSACTIONS HISTORY

						case 3:

							// INSERT A NEW PRODUCT
							inputProduct := product.Product{}
							inputProduct.IDStaff = res.ID
							fmt.Println()
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

						case 4:

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

							var showPrd = true
							for showPrd {
								var prodMenu string
								fmt.Print("Back to main menu? [Y / N] : ")
								fmt.Scanln(&prodMenu)

								if prodMenu == "Y" || prodMenu == "y" {
									showPrd = false
								}
							}

						case 5:

							// UPDATE A PRODUCT STOCK
							prodMenu := 1
							fmt.Println()
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

									res, err := productMenu.InsertStock(addQty, prodID)

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

						case 6:

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

						case 7:

							// INSERT A NEW CUSTOMER
							fmt.Println()
							fmt.Println("INSERT A NEW CUSTOMER")
							var CusName string
							fmt.Print("Please insert customer name : ")
							fmt.Scanln(&CusName)
							ifada, err := custMenu.AddCustomer(CusName, res.ID)
							if ifada {
								fmt.Println("------------------")
								fmt.Println("Added a new customer successfully!")
							} else {
								fmt.Println("------------------")
								fmt.Println("Sorry unable to add a new customer, please insert correctly")
							}
							if err != nil {
								fmt.Println(err.Error())
							}

						case 8:

							// SHOW ALL CUSTOMERS
							fmt.Println()
							fmt.Println("LIST OF CUSTOMERS")
							fmt.Println("------------------")
							customers, _ := custMenu.ShowCustomer()
							for i := 0; i < len(customers); i++ {
								fmt.Println("Customer ID   	: ", customers[i].ID)
								fmt.Println("Customer Name   : ", customers[i].Name)
								fmt.Println("Inserted by  	: ", customers[i].StaffName)
								fmt.Println("------------------")
							}

							var showCust = true
							for showCust {
								var menu string
								fmt.Print("Back to main menu? [Y / N] : ")
								fmt.Scanln(&menu)

								if menu == "Y" || menu == "y" {
									showCust = false
								}
							}

						case 9:

							// LOGOUT
							islogin = false
							fmt.Println("Logged out succesfully!")
							fmt.Println("=======================")
							fmt.Println()

						}
					}
				} else if res.ID == 1 {
					islogin := true
					for islogin {
						fmt.Println()
						fmt.Println("WELCOME ADMIN")
						fmt.Println("------------------")
						fmt.Println("transaction menu")
						fmt.Println("------------------")
						fmt.Println("1. Delete a transaction")
						fmt.Println("2. Delete a product")
						fmt.Println("3. Delete a customer")
						fmt.Println("------------------")
						fmt.Println("staff menu")
						fmt.Println("------------------")
						fmt.Println("4. Insert a new staff")
						fmt.Println("5. Edit a staff")
						fmt.Println("6. Delete a staff")
						fmt.Println("------------------")
						fmt.Println("9. Logout")
						fmt.Println("------------------")
						fmt.Print("Please Insert Menu [1, 2, 3, 4, 5, 6, 9] : ")
						var choice int
						fmt.Scanln(&choice)

						switch choice {
						case 1:

							// DELETE A TRANSACTION

						case 2:

							// DELETE A PRODUCT
							prodMenu := 1

							for prodMenu != 9 {
								fmt.Println("LIST OF PRODUCTS")
								fmt.Println("------------------")
								products, _ := productMenu.Show()
								if len(products) == 0 {
									fmt.Println("No product available.")
								} else {
									for i := 0; i < len(products); i++ {
									fmt.Println("Product Code   : ", products[i].ID)
									fmt.Println("Product Name   : ", products[i].Name)
									fmt.Println("QTY            : ", products[i].Qty)
									fmt.Println("Staff Name     : ", products[i].StaffName)
									fmt.Println("------------------")
									}
								}

								fmt.Println("=======================")
								fmt.Println("1. Delete a product")
								fmt.Println("2. Delete all products")
								fmt.Println("9. Back to main menu")
								fmt.Println("------------------")
								fmt.Print("Please choose a menu [1, 2, 9] : ")
								fmt.Scanln(&prodMenu)

								if prodMenu == 1 {

									// DELETE A PRODUCT
									var productID int
									fmt.Println("=======================")
									fmt.Println("DELETE A PRODUCT")
									fmt.Println("------------------")
									fmt.Print("Please insert a product code : ")
									fmt.Scanln(&productID)

									res, err := productMenu.Delete(productName)
									
									if err != nil {
										fmt.Println("------------------")
										fmt.Println(err.Error())
										fmt.Println("=======================")
									}

									if res {
										fmt.Println("------------------")
										fmt.Printf("Product `%s` has been deleted successfully.\n", products[productID].Name)
										fmt.Println("=======================")
									}

								} else if prodMenu == 2 {

									// DELETE ALL PRODUCTS
									var deleteAll string
									fmt.Println("------------------")
									fmt.Print("Are you sure to delete all the products [Y, N] : ")
									fmt.Scanln(&deleteAll)

									if deleteAll == "Y" || deleteAll == "y" {
										res, err := productMenu.DeleteAll()

										if err != nil {
											fmt.Println("------------------")
											fmt.Println(err.Error())
											fmt.Println("=======================")
										}

										if res {
											fmt.Println("=======================")
											fmt.Println("All the products has been deleted successfully!")
											fmt.Println("=======================")
										}

									} else {
										fmt.Println("=======================")
									}

								} else {
									fmt.Println("=======================")
								}
							}

						case 3:

							// DELETE A CUSTOMER
							var custName string
							fmt.Println()
							fmt.Println("LIST OF CUSTOMERS")
							fmt.Println("------------------")
							customers, _ := custMenu.ShowCustomer()
							for i := 0; i < len(customers); i++ {
								fmt.Println("Customer ID   	: ", customers[i].ID)
								fmt.Println("Customer Name   : ", customers[i].Name)
								fmt.Println("Inserted by  	: ", customers[i].StaffName)
								fmt.Println("------------------")
							}

							fmt.Println("DELETE A CUSTOMER")
							fmt.Println("------------------")
							fmt.Print("Insert Customer Name : ")
							fmt.Scanln(&custName)
							fmt.Println("------------------")

							ifberhasil, err := custMenu.RemoveCustomer(custName)
							if err != nil {
								fmt.Println(err.Error())
							}
							if ifberhasil {
								fmt.Println("Data Customer ", custName, " has been deleted successfully!")
							} else {
								fmt.Println("Sorry can't delete customer, please input correctly")
							}

							fmt.Println("=======================")

						case 4:

							// INSERT A NEW STAFF
							var newStaff staff.Staff
							fmt.Println("INSERT A NEW STAFF")
							fmt.Println("------------------")
							fmt.Print("Insert new username : ")
							fmt.Scanln(&newStaff.Name)
							fmt.Print("Insert new password : ")
							fmt.Scanln(&newStaff.Password)
							res, err := staffMenu.Register(newStaff)
							if err != nil {
								fmt.Println("------------------")
								fmt.Println(err.Error())
							}
							if res {
								fmt.Println("------------------")
								fmt.Println("New Staff :", newStaff.Name, "has been inserted succesfully!")
								fmt.Println("=======================")
							} else {
								fmt.Println("------------------")
								fmt.Println("Sorry the username has been used, unable to insert new staff.")
								fmt.Println("=======================")
							}

						case 5:

							// EDIT A STAFF
							staffs, _ := staffMenu.Show()
							for i := 0; i < len(staffs); i++ {
							fmt.Println("Staff Id           : ", staffs[i].ID)
							fmt.Println("Staff Name         : ", staffs[i].Name)
							fmt.Println("------------------")
							}

							// var staffName, updateName, updatePass string
							// fmt.Print("Masukkan password baru: ")
							// fmt.Scanln(&inputPass)
							// isChanged, err := authMenu.GantiPassword(inputPass, res.ID)
							// if err != nil {
							// 	fmt.Println(err.Error())
							// }
							// if isChanged {
							// 	fmt.Println("Berhasil ganti password")
							// 	isLogin = false
							// }

							// UPDATE STAFF ACCOUNT
							staffEdit := 1

							for staffEdit != 9 {
								fmt.Println("LIST OF STAFFS")
								fmt.Println("------------------")
								staffs, _ := staffMenu.Show()
								if len(staffs) == 0 {
									fmt.Println("Staffs is empty.")
								} else {
									for i := 0; i < len(staffs); i++ {
										fmt.Println("Staff ID     : ", staffs[i].ID)
										fmt.Println("Staff Name   : ", staffs[i].Name)
										fmt.Println("------------------")
									}
								}
					
								fmt.Println("1. Update a staff accout")
								fmt.Println("9. Back to main menu")
								fmt.Print("Please choose a menu [1, 9] : ")
								fmt.Scanln(&staffEdit)

								if staffEdit == 1 {
									var staffID int
									var newName string
									var newPass string
									fmt.Println("=======================")
									fmt.Println("UPDATE STAFF ACCOUNT")
									fmt.Println("------------------")
									fmt.Print("Please insert id staff : ")
									fmt.Scanln(&staffID)
									fmt.Print("Please insert new name : ")
									fmt.Scanln(&newName)
									fmt.Print("Please insert new password : ")
									fmt.Scanln(&newPass)

									res, err := staffMenu.UpdateStaff(newName, newPass, staffID)

									if err != nil {
										fmt.Println("------------------")
										fmt.Println(err.Error())
										fmt.Println("=======================")
									}

									if res {
										fmt.Println("------------------")
										fmt.Println("Updated a staff account succesfully!")
										fmt.Println("=======================")
									}

								} else {
									fmt.Println("=======================")
								}
							}

						case 6:

							// DELETE A STAFF
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
							fmt.Println("=======================")

						case 9:

							// LOGOUT
							islogin = false
							fmt.Println("Logged out successfully!")
							fmt.Println("=======================")
							fmt.Println()

						}
					}
				}
			}
		}
	}
}
