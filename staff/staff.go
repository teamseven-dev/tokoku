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

func (sm *StaffMenu) Login(name string, password string) (Staff, error) {
	// selectid := sm.DB.QueryRow("SELECT id_staff FROM staffs WHERE name = ? AND password = ?", nama, password)
	// var IDada int
	// err := selectid.Scan(&IDada)
	// if err != nil {
	// 	log.Println("error saat select dari database", err.Error())
	// 	return 0, errors.New("password atau username salah")
	// }
	// return IDada, nil

	loginQry, err := sm.DB.Prepare("SELECT id_staff FROM staffs WHERE name = ? AND password = ?")
	if err != nil {
		log.Println("prepare insert staff ", err.Error())
		return Staff{}, errors.New("prepare statement insert staff error")
	}

	row := loginQry.QueryRow(name, password)

	if row.Err() != nil {
		log.Println("login query ", row.Err().Error())
		return Staff{}, errors.New("tidak bisa login, data tidak ditemukan")
	}
	res := Staff{}
	err = row.Scan(&res.ID)

	if err != nil {
		log.Println("after login query ", err.Error())
		return Staff{}, errors.New("tidak bisa login, kesalahan setelah error")
	}

	res.Name = name

	return res, nil
}

func (sm *StaffMenu) Duplicate(name string) bool {
	res := sm.DB.QueryRow("SELECT id_staff FROM staffs where name = ?", name)
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

func (sm *StaffMenu) Register(newStaff Staff) (bool, error) {
	// menyiapakn query untuk insert
	registerQry, err := sm.DB.Prepare("INSERT INTO staffs (name, password, created_date, updated_date) values (?, ?, now(), null)")
	if err != nil {
		log.Println("prepare insert staff ", err.Error())
		return false, errors.New("prepare statement insert staff error")
	}
	
	if sm.Duplicate(newStaff.Name) {
		log.Println("duplicated information")
		return false, errors.New("nama sudah digunakan")
	}
	
	// menjalankan query dengan parameter tertentu
	res, err := registerQry.Exec(newStaff.Name, newStaff.Password)
	if err != nil {
		log.Println("insert staff ", err.Error())
		return false, errors.New("insert staff error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()
	
	if err != nil {
		log.Println("after insert staff ", err.Error())
		return false, errors.New("error setelah insert")
	}
	
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}
	
	return true, nil
}

func (sm *StaffMenu) Show() ([]Staff, error) {
	rows, err := sm.DB.Query("SELECT id_staff, name FROM staffs")
	if err != nil {
    	log.Println(err)
	}

	res := []Staff{} // creating empty slice
	defer rows.Close()

	for rows.Next() {
		staffs := Staff{} // creating new struct for every row
		err = rows.Scan(&staffs.ID, &staffs.Name)
		if err != nil {
			log.Println(err)
		}
		res = append(res, staffs)
	}

	return res, nil
}

func (sm *StaffMenu) Remove(name string) (bool, error) {
	removeQry, err := sm.DB.Prepare("DELETE FROM staffs WHERE name = ?")
	if err != nil {
		log.Println("prepare remove staff ", err.Error())
		return false, errors.New("prepare statement remove staff error")
	}

	if name == "admin" {
		log.Println("admin information")
		return false, errors.New("admin tidak dapat dihapus")
	}

	res, err := removeQry.Exec(name)
	if err != nil {
		log.Println("remove staff ", err.Error())
		return false, errors.New("remove staff error")
	}

	affRows, err := res.RowsAffected()
	
	if err != nil {
		log.Println("after remove staff ", err.Error())
		return false, errors.New("error setelah remove")
	}
	
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}
	
	return true, nil
}

func (sm *StaffMenu) UpdateStaff(updateName string, updatePassword string, id int) (bool, error) {
	updateStaffQry, err := sm.DB.Prepare("UPDATE staffs SET name = ? AND password = ? WHERE id_staff = ?")
	if err != nil {
		log.Println("prepare update name and password ", err.Error())
		return false, errors.New("prepare statement update name and password error")
	}

	res, err := updateStaffQry.Exec(updateName, updatePassword)
	if err != nil {
		log.Println("update name and password ", err.Error())
		return false, errors.New("update name and password error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after update name and password ", err.Error())
		return false, errors.New("error setelah update name and password")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}