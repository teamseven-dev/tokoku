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
		fmt.Println("------------------")
		log.Println("prepare insert staff ", err.Error())
		return false, errors.New("prepare statement insert staff error")
	}
	
	if sm.Duplicate(newStaff.Name) {
		fmt.Println("------------------")
		log.Println("duplicated information")
		return false, errors.New("nama sudah digunakan")
	}
	
	// menjalankan query dengan parameter tertentu
	res, err := registerQry.Exec(newStaff.Name, newStaff.Password)
	if err != nil {
		fmt.Println("------------------")
		log.Println("insert staff ", err.Error())
		return false, errors.New("insert staff error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()
	
	if err != nil {
		fmt.Println("------------------")
		log.Println("after insert staff ", err.Error())
		return false, errors.New("error setelah insert")
	}
	
	if affRows <= 0 {
		fmt.Println("------------------")
		log.Println("no record affected")
		return false, errors.New("no record")
	}
	
	return true, nil
}

func (sm *StaffMenu) Show() ([]Staff, error) {
	rows, err := sm.DB.Query("SELECT id_staff, name FROM staffs")
	if err != nil {
		fmt.Println("------------------")
    	log.Println(err)
	}
	
	res := []Staff{} // creating empty slice
	defer rows.Close()
	
	for rows.Next() {
		staffs := Staff{} // creating new struct for every row
		err = rows.Scan(&staffs.ID, &staffs.Name)
		if err != nil {
			fmt.Println("------------------")
			log.Println(err)
		}
		res = append(res, staffs)
	}
	
	return res, nil
}

func (sm *StaffMenu) Remove(name string) (bool, error) {
	removeQry, err := sm.DB.Prepare("DELETE FROM staffs WHERE name = ?")
	if err != nil {
		fmt.Println("------------------")
		log.Println("prepare remove staff ", err.Error())
		return false, errors.New("prepare statement remove staff error")
	}
	
	if name == "admin" {
		fmt.Println("------------------")
		log.Println("admin information")
		return false, errors.New("admin can't be deleted")
	}
	
	res, err := removeQry.Exec(name)
	if err != nil {
		fmt.Println("------------------")
		log.Println("remove staff ", err.Error())
		return false, errors.New("remove staff error")
	}
	
	affRows, err := res.RowsAffected()
	
	if err != nil {
		fmt.Println("------------------")
		log.Println("after remove staff ", err.Error())
		return false, errors.New("error setelah remove")
	}
	
	if affRows <= 0 {
		fmt.Println("------------------")
		log.Println("no record affected")
		return false, errors.New("no record")
	}
	
	return true, nil
}

func (sm *StaffMenu) DeleteAll() (bool, error) {
	deleteAllQry, err := sm.DB.Prepare("DELETE FROM staffs WHERE id_staff != 1")
	if err != nil {
		fmt.Println("------------------")
		log.Println("Prepare delete staffs : ", err.Error())
		return false, errors.New("prepare statement delete staffs error")
	}

	res, err := deleteAllQry.Exec()
	if err != nil {
		fmt.Println("------------------")
		log.Println("Delete staffs : ", err.Error())
		return false, errors.New("delete product staffs")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		fmt.Println("------------------")
		log.Println("Afer delete staffs : ", err.Error())
		return false, errors.New("error after delete staffs")
	}

	if affRows <= 0 {
		fmt.Println("------------------")
		log.Println("No rows affected")
		return false, errors.New("no record affected")
	}
	
	return true, nil
}

func (sm *StaffMenu) UpdateStaff(newName string, newPassword string, id int) (bool, error) {
	updateStaffQry, err := sm.DB.Prepare("UPDATE staffs SET name = ?, password = ?, updated_date = now() WHERE id_staff = ?")
	if err != nil {
		fmt.Println("------------------")
		log.Println("prepare update name and password ", err.Error())
		return false, errors.New("prepare statement update name and password error")
	}
	
	if id == 1 {
		fmt.Println("------------------")
		log.Println("admin information")
		return false, errors.New("you don't have permission to edit admin account")
	}
	
	res, err := updateStaffQry.Exec(newName, newPassword, id)
	if err != nil {
		fmt.Println("------------------")
		log.Println("update name and password ", err.Error())
		return false, errors.New("update name and password error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()
	
	if err != nil {
		fmt.Println("------------------")
		log.Println("after update name and password ", err.Error())
		return false, errors.New("error setelah update name and password")
	}
	
	if affRows <= 0 {
		fmt.Println("------------------")
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}