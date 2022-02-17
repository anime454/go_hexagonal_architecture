package repository

import (
	"database/sql"
)

type userRepositoryDB struct {
	db *sql.DB
}

func NewUserRepositoryDB(db *sql.DB) userRepositoryDB {
	return userRepositoryDB{db: db}
}

func (u userRepositoryDB) Create(user User) (string, error) {
	stmt, err := u.db.Prepare(`INSERT INTO user (id, username, password, fullname, email, role, auto_datetime) VALUES( ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return "", err
	}

	_, err = stmt.Exec(user.Id, user.Username, user.Password, user.FullName, user.Email, user.Role, user.AutoDatetime)
	if err != nil {
		return "", err
	}

	return user.Id, nil
}

func (u userRepositoryDB) GetAll() ([]User, error) {
	res := []User{}
	stmt, err := u.db.Prepare(`select id, username, password, fullname, email, role, auto_datetime from user`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		utmp := User{}
		err = rows.Scan(&utmp.Id, &utmp.Username, &utmp.Password, &utmp.FullName, &utmp.Email, &utmp.Role, &utmp.AutoDatetime)
		if err != nil {
			return nil, err
		}
		res = append(res, utmp)
	}
	rows.Scan(&res)

	return res, nil
}

func (u userRepositoryDB) GetById(id string) (*User, error) {
	res := User{}
	stmt, err := u.db.Prepare(`select id, username, password, fullname, email, role, auto_datetime from user where id = ?`)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&res.Id, &res.Username, &res.Password, &res.FullName, &res.Email, &res.Role, &res.AutoDatetime)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (u userRepositoryDB) Update(user UserDetail) (string, error) {

	stmt, err := u.db.Prepare(`UPDATE user SET username = ?, fullname = ?, email = ?, role = ?, auto_datetime = ? WHERE id = ?`)
	if err != nil {
		return "", err
	}

	result, err := stmt.Exec(user.Username, user.FullName, user.Email, user.Role, user.AutoDatetime, user.Id)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err

	}
	if rowsAffected <= 0 {
		return "", nil
	}

	return user.Id, nil
}

func (u userRepositoryDB) Delete(id string) error {
	stmt, err := u.db.Prepare(`DELETE FROM user WHERE id = ?`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
