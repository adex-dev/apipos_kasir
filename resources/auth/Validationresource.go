package auth

import (
	"apipos/database"
	"apipos/models"
	"database/sql"
	"errors"
)

func ValidasiLogin(nik, password string) (*models.Users, error) {
	row := database.DBMASTER.QueryRow("SELECT * FROM users WHERE nik=? AND  password=?", nik, password)

	var user models.Users
	var state sql.NullString

	if err := row.Scan(&user.Nik, &user.Nama_user, &user.Add_ons, &user.Status, &user.Password, &user.Level, &state); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	if state.Valid {
		user.State = state.String
	} else {
		user.State = ""
	}

	return &user, nil
}
