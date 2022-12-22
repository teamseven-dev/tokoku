package transaction

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Transaction struct {
	ID           int
	IDStaff      int
	IDCustomer   int
	StaffName    string
	CustomerName string
	CreatedDate  string
}

type Items struct {
	IDTransaction int
	IDProduct     int
	Qty           int
	ProductName   string
}

type TransMenu struct {
	DB *sql.DB
}

func (tm *TransMenu) AddTransaction(idstaff, IdCustomer int) (int, error) {
	inserttrans, err := tm.DB.Prepare("INSERT INTO transactions (id_staff,id_customer) VALUE (?,?)")
	if err != nil {
		log.Println("prepare insert transaction error.")
		return 0, errors.New("prepare insert transaction error.")
	}
	res, err := inserttrans.Exec(idstaff, IdCustomer)
	if err != nil {
		log.Println("insert transaction error")
		return 0, errors.New("insert transaction error")
	}
	rowaffect, err := res.RowsAffected()
	if err != nil {
		log.Println("error row affected")
		return 0, errors.New("no row affected")
	}
	if rowaffect <= 0 {
		log.Println("error rowaffect")
		return 0, errors.New("no row record")
	}
	id, _ := res.LastInsertId()
	return int(id), nil
}

// TRANSACTION SHOW
func (tm *TransMenu) ShowTransaction(id int) ([]Transaction, error) {
	showTransQry, err := tm.DB.Query("SELECT t.id_transaction, s.name, c.name, t.created_date FROM transactions t JOIN staffs s ON t.id_staff  = s.id_staff JOIN customers c ON t.id_customer = c.id_customer WHERE t.id_transaction = ?", id)
	if err != nil {
		log.Println("Prepare show transaction table", err.Error())
		return []Transaction{}, errors.New("prepare statement show transaction table error")
	}

	res := []Transaction{} // creating empty slice
	defer showTransQry.Close()

	for showTransQry.Next() {
		transaction := Transaction{} // creating new struct for every row
		err = showTransQry.Scan(&transaction.ID, &transaction.StaffName, &transaction.CustomerName, &transaction.CreatedDate)
		if err != nil {
			log.Println("------------------")
			log.Println(err)
		}
		res = append(res, transaction)
	}

	return res, nil
}

// SHOW ALL TRANSACTIONS
func (tm *TransMenu) ShowAllTransaction() ([]Transaction, error) {
	showTransQry, err := tm.DB.Query("SELECT t.id_transaction, s.name, c.name, t.created_date FROM transactions t JOIN staffs s ON t.id_staff  = s.id_staff JOIN customers c ON t.id_customer = c.id_customer")
	if err != nil {
		log.Println("Prepare show transaction table", err.Error())
		return []Transaction{}, errors.New("prepare statement show transaction table error")
	}

	res := []Transaction{} // creating empty slice
	defer showTransQry.Close()

	for showTransQry.Next() {
		transaction := Transaction{} // creating new struct for every row
		err = showTransQry.Scan(&transaction.ID, &transaction.StaffName, &transaction.CustomerName, &transaction.CreatedDate)
		if err != nil {
			log.Println("------------------")
			log.Println(err)
		}
		res = append(res, transaction)
	}

	return res, nil
}
// INSERT ITEMS
func (tm *TransMenu) InsertItem(idTrx, idProd, qty int) (bool, error) {
	insertItem, err := tm.DB.Prepare("INSERT INTO items(id_transaction,id_product,qty) VALUE (?,?,?)")
	if err != nil {
		log.Println("error query insert")
		return false, errors.New("error prepare")
	}
	res, err := insertItem.Exec(idTrx, idProd, qty)
	if err != nil {
		log.Println("error insert item")
		return false, errors.New("error insert item")
	}
	rowaffec, err := res.RowsAffected()
	if err != nil {
		log.Println("error di row affected")
		return false, errors.New("no row record")
	}
	if rowaffec <= 0 {
		log.Println("no row record")
		return false, errors.New("no row record")
	}
	return true, nil
}

// ITEMS SHOW
func (tm *TransMenu) ShowItems(id int) ([]Items, error) {
	showItemsQry, err := tm.DB.Query("SELECT p.id_product, p.product_name, i.qty FROM items i, products p WHERE i.id_product  = p.id_product AND i.id_transaction = ?", id)
	if err != nil {
		log.Println("Prepare show transaction items : ", err.Error())
		return []Items{}, errors.New("Prepare statement show transaction items error")
	}

	res := []Items{} // creating empty slice
	defer showItemsQry.Close()

	for showItemsQry.Next() {
		item := Items{} // creating new struct for every row
		err = showItemsQry.Scan(&item.IDProduct, &item.ProductName, &item.Qty)
		if err != nil {
			log.Println("------------------")
			log.Println(err)
		}
		res = append(res, item)
	}

	return res, nil
}

// DELETE TRANSACTION
func (tm *TransMenu) Delete(id int) (bool, error) {
	deleteQry, err := tm.DB.Prepare("DELETE FROM transactions WHERE id_transaction = ?;")
	if err != nil {
		fmt.Println("------------------")
		log.Println("Prepare delete transaction : ", err.Error())
		return false, errors.New("Prepare statement delete transaction error.")
	}

	res, err := deleteQry.Exec(id)
	if err != nil {
		fmt.Println("------------------")
		log.Println("Delete transaction : ", err.Error())
		return false, errors.New("Delete transaction error.")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		fmt.Println("------------------")
		log.Println("Afer delete transaction : ", err.Error())
		return false, errors.New("Error after delete transaction.")
	}

	if affRows <= 0 {
		fmt.Println("------------------")
		log.Println("No rows affected.")
		return false, errors.New("No record affected.")
	}

	return true, nil
}
