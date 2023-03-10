package product

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Product struct {
	ID          int
	Name        string
	Qty         int
	IDStaff     int
	StaffName   string
	CreatedDate string
	UpdatedDate string
}

type ProductMenu struct {
	DB *sql.DB
}

func (pm *ProductMenu) Duplicate(productName string) bool {
	res := pm.DB.QueryRow("SELECT id_product FROM products where product_name = ?", productName)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			log.Println("Result scan error", err.Error())
		}
		return false
	}
	return true
}

func (pm *ProductMenu) Insert(newProduct Product) (bool, error) {
	insertQry, err := pm.DB.Prepare("INSERT INTO products (product_name, qty, id_staff, created_date, updated_date) VALUES (?,?,?,now(),now())")
	if err != nil {
		fmt.Println("------------------")
		log.Println("Prepare insert newProduct : ", err.Error())
		return false, errors.New("Prepare statement insert new product error.")
	}

	if pm.Duplicate(newProduct.Name) {
		fmt.Println("------------------")
		log.Println("duplicated information")
		return false, errors.New("product name already exist")
	}

	res, err := insertQry.Exec(newProduct.Name, newProduct.Qty, newProduct.IDStaff)
	if err != nil {
		fmt.Println("------------------")
		log.Println("Insert new product : ", err.Error())
		return false, errors.New("Insert new product error.")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		fmt.Println("------------------")
		log.Println("Afer inser new product : ", err.Error())
		return false, errors.New("Error after insert new product.")
	}

	if affRows <= 0 {
		fmt.Println("------------------")
		log.Println("No rows affected.")
		return false, errors.New("No record affected.")
	}

	return true, nil
}

func (pm *ProductMenu) Delete(id int) (bool, error) {
	turnoffFK, _ := pm.DB.Prepare("SET FOREIGN_KEY_CHECKS=0")
	res1, _ := turnoffFK.Exec()
	ToffFK, _ := res1.RowsAffected()
	if ToffFK > 0 {
		fmt.Println("error Turn off FK")
	}
	deleteQry, err := pm.DB.Prepare("DELETE FROM products WHERE id_product = ?;")
	if err != nil {
		fmt.Println("------------------")
		log.Println("Prepare delete product : ", err.Error())
		return false, errors.New("Prepare statement delete product error.")
	}

	res, err := deleteQry.Exec(id)
	if err != nil {
		fmt.Println("------------------")
		log.Println("Delete product : ", err.Error())
		return false, errors.New("Delete product error.")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		fmt.Println("------------------")
		log.Println("Afer delete product : ", err.Error())
		return false, errors.New("Error after delete product.")
	}

	if affRows <= 0 {
		fmt.Println("------------------")
		log.Println("No rows affected.")
		return false, errors.New("No record affected.")
	}
	turnonFK, _ := pm.DB.Prepare("SET FOREIGN_KEY_CHECKS=1")
	res2, _ := turnonFK.Exec()
	TonFK, _ := res2.RowsAffected()
	if TonFK > 0 {
		fmt.Println("error Turn on FK")
	}
	return true, nil
}

func (pm *ProductMenu) DeleteAll() (bool, error) {
	deleteAllQry, err := pm.DB.Prepare("DELETE FROM products;")
	if err != nil {
		fmt.Println("------------------")
		log.Println("Prepare delete product : ", err.Error())
		return false, errors.New("Prepare statement delete product error.")
	}

	res, err := deleteAllQry.Exec()
	if err != nil {
		fmt.Println("------------------")
		log.Println("Delete product : ", err.Error())
		return false, errors.New("Delete product error.")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		fmt.Println("------------------")
		log.Println("Afer delete product : ", err.Error())
		return false, errors.New("Error after delete product.")
	}

	if affRows <= 0 {
		fmt.Println("------------------")
		log.Println("No rows affected.")
		return false, errors.New("No record affected.")
	}

	return true, nil
}

func (pm *ProductMenu) Show() ([]Product, error) {
	rows, err := pm.DB.Query("SELECT p.id_product, p.product_name, p.qty, s.name FROM staffs s, products p WHERE s.id_staff = p.id_staff")
	if err != nil {
		log.Println("------------------")
		log.Println(err)
	}

	res := []Product{} // creating empty slice
	defer rows.Close()

	for rows.Next() {
		product := Product{} // creating new struct for every row
		err = rows.Scan(&product.ID, &product.Name, &product.Qty, &product.StaffName)
		if err != nil {
			log.Println("------------------")
			log.Println(err)
		}
		res = append(res, product)
	}

	return res, nil
}

func (pm *ProductMenu) UpdateName(newName string, id int) (bool, error) {
	updateNameQry, err := pm.DB.Prepare("UPDATE products SET product_name = ? WHERE id_product = ?")
	if err != nil {
		log.Println("Prepare update product name : ", err.Error())
		return false, errors.New("Prepare statement update product name error.")
	}

	res, err := updateNameQry.Exec(newName, id)
	if err != nil {
		log.Println("Update product name : ", err.Error())
		return false, errors.New("Update product name error.")
	}

	affRows, err := res.RowsAffected()
	if err != nil {
		log.Println("After update product name : ", err.Error())
		return false, errors.New("After update product name error.")
	}

	if affRows <= 0 {
		log.Println("No rows affected.")
		return false, errors.New("No record affected.")
	}

	return true, nil
}

func (pm *ProductMenu) InsertStock(addQty, id int) (bool, error) {
	insertStockQry, err := pm.DB.Prepare("UPDATE products SET qty = qty + ? WHERE id_product = ?")
	if err != nil {
		log.Println("Prepare update product stock : ", err.Error())
		return false, errors.New("Prepare statement update product stock error.")
	}

	res, err := insertStockQry.Exec(addQty, id)
	if err != nil {
		log.Println("Update product stock : ", err.Error())
		return false, errors.New("Update product stock error.")
	}

	affRows, err := res.RowsAffected()
	if err != nil {
		log.Println("After update product stock : ", err.Error())
		return false, errors.New("After update product stock error.")
	}

	if affRows <= 0 {
		log.Println("No rows affected.")
		return false, errors.New("No record affected.")
	}

	return true, nil
}

func (pm *ProductMenu) UpdateStock(addQty, id int) (bool, error) {
	updateStockQry, err := pm.DB.Prepare("UPDATE products SET qty = qty - ? WHERE id_product = ?")
	if err != nil {
		log.Println("Prepare update product stock : ", err.Error())
		return false, errors.New("Prepare statement update product stock error.")
	}

	res, err := updateStockQry.Exec(addQty, id)
	if err != nil {
		log.Println("Update product stock : ", err.Error())
		return false, errors.New("Update product stock error.")
	}

	affRows, err := res.RowsAffected()
	if err != nil {
		log.Println("After update product stock : ", err.Error())
		return false, errors.New("After update product stock error.")
	}

	if affRows <= 0 {
		log.Println("No rows affected.")
		return false, errors.New("No record affected.")
	}

	return true, nil
}

func (pm *ProductMenu) GetQty(id int) (int, error) {
	getQty := pm.DB.QueryRow("SELECT qty FROM products WHERE id_product = ?", id)
	var qty int
	err := getQty.Scan(&qty)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			log.Println("Result scan error", err.Error())
		}
		return 0, errors.New("Error when getting qty.")
	}
	return qty, nil
}
