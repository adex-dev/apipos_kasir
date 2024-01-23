package Master

import (
	"apipos/database"
	"apipos/models/master"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func StoreResources(c *fiber.Ctx) ([]*master.Storelist, error) {
	query := `SELECT prm_outlet.id_outlet,prm_outlet.nomor,prm_outlet.nama_outlet,prm_outlet.add_ons,prm_outlet.create_user,prm_outlet.status,prm_outlet.token,prm_outlet.brands,prm_outlet.prefix,prm_outlet.service_charge,prm_outlet.tax,prm_outlet.alamat,prm_outlet.telp_1,prm_outlet.telp_2,prm_outlet.longitude,prm_outlet.latitude,prm_outlet.thumbnail,users.nama_user FROM prm_outlet INNER JOIN users ON users.nik = prm_outlet.create_user order by  prm_outlet.nama_outlet ASC`
	rows, err := database.DBMASTER.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var stores []*master.Storelist
	var thumbnail, longitude, latitude, telp2 sql.NullString
	for rows.Next() {
		var st master.Storelist
		err := rows.Scan(&st.Idoutlet, &st.Nomor, &st.Namaoutlet, &st.Addons, &st.Createuser, &st.Status, &st.Token, &st.Brands, &st.Prefix, &st.ServiceCharge, &st.Tax, &st.Alamat, &st.Telp1, &telp2, &longitude, &latitude, &thumbnail, &st.Nama_user)
		if err != nil {
			return nil, err
		}
		if thumbnail.Valid {
			st.Thumbnail = thumbnail.String
		} else {
			st.Thumbnail = ""
		}
		if longitude.Valid {
			st.Longitude = longitude.String
		} else {
			st.Longitude = ""
		}
		if latitude.Valid {
			st.Latitude = latitude.String
		} else {
			st.Latitude = ""
		}
		if telp2.Valid {
			st.Telp2 = telp2.String
		} else {
			st.Telp2 = ""
		}
		stores = append(stores, &st)
	}
	return stores, nil
}
