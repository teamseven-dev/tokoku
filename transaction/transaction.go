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
		log.Println("error prepare")
		return 0, errors.New("query insert error")
	}
	res, err := inserttrans.Exec(idstaff, IdCustomer)
	if err != nil {
		log.Println("error perintah exec")
		return 0, errors.New("error exec")
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

// INSERT ITEMS

// ITEMS SHOW
