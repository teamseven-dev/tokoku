package transaction

import (
	"database/sql"
	"errors"
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

// SHOW ALL TRANSACTIONS

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

// DELETE TRANSACTION
