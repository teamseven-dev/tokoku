package staff

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type StaffMenu struct {
	DB *sql.DB
}

type Staff struct {
	ID          int
	Name        string
	Password    string
	CreatedDate string
	UpdatedDate string
}

func (sm *StaffMenu) Login(name string, password string) (Staff, error) {
	loginQry, err := sm.DB.Prepare("SELECT id_staff FROM staffs WHERE name = ? AND password = ?")
	if err != nil {
		fmt.Println("------------------")
		log.Println("prepare insert staff ", err.Error())
		return Staff{}, errors.New("prepare statement insert staff error")
	}

	row := loginQry.QueryRow(name, password)

	if row.Err() != nil {
		fmt.Println("------------------")
		log.Println("login query ", row.Err().Error())
		return Staff{}, errors.New("tidak bisa login, data tidak ditemukan")
	}
	res := Staff{}
	err = row.Scan(&res.ID)

	if err != nil {
		fmt.Println("------------------")
		log.Println("after login query ", err.Error())
		return Staff{}, errors.New("tidak bisa login, kesalahan setelah error")
	}

	res.Name = name

	return res, nil
}
