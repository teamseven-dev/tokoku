package customer

import (
	"database/sql"
	"errors"
	"log"
)

type CustMenu struct {
	DB *sql.DB
}

type Customer struct {
	ID      int
	Name    string
	IDStaff int
}

func (cm *CustMenu) AddCustomer(nama string, StaffID int) (bool, error) {
	addcost, err := cm.DB.Prepare("INSERT INTO customers(name,id_staff) VALUE (?,?)")
	if err != nil {
		log.Println("error query insert", err.Error())
		return false, errors.New("can't add data")
	}
	res, err := addcost.Exec(nama, StaffID)
	if err != nil {
		log.Println("error insert to database", err.Error())
		return false, errors.New("can't add data")
	}
	affctrow, err := res.RowsAffected()
	if err != nil {
		log.Println("no row affected", err.Error())
		return false, errors.New("no row affected")
	}
	if affctrow <= 0 {
		log.Println("erorr at rowaffected", err.Error())
		return false, errors.New("no row affected")
	}
	return true, nil
}

func (cm *CustMenu) RemoveCustomer(nama string) (bool, error) {
	delcust, err := cm.DB.Prepare("DELETE FROM customers WHERE name = ?")
	if err != nil {
		log.Println("error di query delete")
		return false, errors.New("error query delete")
	}
	res, err := delcust.Exec(nama)
	if err != nil {
		log.Println("error exec")
		return false, errors.New("error exec")
	}
	affctrow, err := res.RowsAffected()
	if err != nil {
		log.Println("no row affected")
		return false, errors.New("no row affected")
	}
	if affctrow <= 0 {
		log.Println("erorr at rowaffected")
		return false, errors.New("no row affected")
	}

	return true, nil
}

func (cm *CustMenu) ShowCustomer() ([]Customer, error) {
	res, err := cm.DB.Query("SELECT FROM customers")
	if err != nil {
		log.Println("error query", err.Error())
		return []Customer{}, errors.New("error select database")
	}
	cus := []Customer{}
	defer res.Close()

	for res.Next() {
		customer := Customer{} // creating new struct for every row
		err = res.Scan(&customer.ID, &customer.Name, &customer.IDStaff)
		if err != nil {
			log.Println("------------------")
			log.Println(err)
		}
		cus = append(cus, customer)
	}
	return cus, nil
}
