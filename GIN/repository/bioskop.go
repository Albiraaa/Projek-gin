package repository

import (
	"database/sql"
	"projek/models"
)

func GetAllbioskop(db *sql.DB) (result []models.Bioskop, err error) {
	sql := "SELECT * FROM bioskop"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var bioskop models.Bioskop

		err = rows.Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)
		if err != nil {
			return
		}

		result = append(result, bioskop)
	}

	return
}

func InsertBioskop(db *sql.DB, bioskop models.Bioskop) (err error) {
	sql := "INSERT INTO bioskop(id, nama, lokasi, rating) VALUES ($1, $2,$3,$4)"

	errs := db.QueryRow(sql, bioskop.ID, bioskop.Nama, bioskop.Lokasi, bioskop.Rating)

	return errs.Err()
}

func UpdateBioskop(db *sql.DB, bioskop models.Bioskop) (err error) {
	sql := "UPDATE bioskop SET nama = $1, lokasi = $2, rating = $3 WHERE id = $4"

	errs := db.QueryRow(sql, bioskop.Nama, bioskop.Lokasi, bioskop.Rating, bioskop.ID)

	return errs.Err()
}

func DeleteBioskop(db *sql.DB, bioskop models.Bioskop) (err error) {
	sql := "DELETE FROM bioskop WHERE id = $1"

	errs := db.QueryRow(sql, bioskop.ID)
	return errs.Err()
}

func GetBioskopByID(db *sql.DB, id int) (bioskop models.Bioskop, err error) {
	sql := "SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1"

	err = db.QueryRow(sql, id).Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)

	return
}
