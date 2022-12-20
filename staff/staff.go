package staff

import (
	"database/sql"
	"errors"
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

func (sm *StaffMenu) Login(nama string, password string) (int, error) {
	selectid := sm.DB.QueryRow("SELECT id_staff FROM staffs WHERE name = ? AND password = ?", nama, password)
	var IDada int
	err := selectid.Scan(&IDada)
	if err != nil {
		log.Println("error saat select dari database", err.Error())
		return 0, errors.New("password atau username salah")
	}
	return IDada, nil
}
