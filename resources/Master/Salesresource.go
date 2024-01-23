package Master

import (
	"apipos/database"
	"apipos/models/master"
	"database/sql"
	"fmt"
)

func SalesResource(start, end string) (*master.Result, error) {
	rows, err := database.DBMASTER.Query("SELECT id_outlet, nama_outlet, brands FROM prm_outlet ORDER BY nama_outlet ASC")
	if err != nil {
		return nil, fmt.Errorf("database query error: %v", err)
	}
	defer rows.Close()

	var result master.Result
	result.DataIsoide = make([]master.StoreSales, 0)
	result.DataNahm = make([]master.StoreSales, 0)

	for rows.Next() {
		var db *sql.DB
		var outlet master.Outlet
		err := rows.Scan(
			&outlet.Idoutlet,
			&outlet.Name,
			&outlet.Brands,
		)

		if err != nil {
			return nil, err
		}

		// Add debugging prints to check values

		var where map[string]interface{}
		if outlet.Brands == "ISOIDE" {
			where = map[string]interface{}{
				"id_lokasi":        outlet.Idoutlet,
				"start_date":       start + " 00:01",
				"end_date":         end + " 23:59",
				"status_transaksi": "CLOSED",
			}
			db = database.DB
		} else if outlet.Brands == "NAHM" {
			where = map[string]interface{}{
				"id_lokasi":        outlet.Idoutlet,
				"start_date":       start + " 00:01",
				"end_date":         end + " 23:59",
				"status_transaksi": "CLOSED",
			}
			db = database.DBNAHM
		}

		intNetsales, err := salesSum(db, where)
		storeSales := master.StoreSales{
			NamaStore: outlet.Name,
			Jumlah:    intNetsales,
		}
		if outlet.Brands == "ISOIDE" {
			result.DataIsoide = append(result.DataIsoide, storeSales)
			result.TotalIsoide += intNetsales
		} else if outlet.Brands == "NAHM" {
			result.DataNahm = append(result.DataNahm, storeSales)
			result.TotalNahm += intNetsales
		}
	}

	return &result, nil
}

func salesSum(db *sql.DB, where map[string]interface{}) (int, error) {
	if db == nil {
		return 0, fmt.Errorf("database connection is nil")
	}

	var netsales sql.NullFloat64
	query := fmt.Sprintf(`
        SELECT SUM(fk.net_sales) as netsales
        FROM transaction_parameter
        INNER JOIN trm_payment fk ON fk.no_trans = transaction_parameter.no_trans
        WHERE transaction_parameter.id_lokasi = ? AND fk.close_payment >= ? AND fk.close_payment <= ? AND transaction_parameter.status_transaksi = ?`)

	// Extract values from the 'where' map
	idLokasi, startDate, endDate, statusTransaksi :=
		where["id_lokasi"], where["start_date"], where["end_date"], where["status_transaksi"]

	err := db.QueryRow(query, idLokasi, startDate, endDate, statusTransaksi).Scan(&netsales)
	if err != nil {
		return 0, err
	}

	// Check if the value is valid before accessing it
	var intNetsales int
	if netsales.Valid {
		intNetsales = int(netsales.Float64)
	} else {
		intNetsales = 0
	}
	return intNetsales, nil
}
