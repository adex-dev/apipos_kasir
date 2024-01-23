package Master

import (
	"apipos/database"
	"apipos/models/master"
)

func StoreResources(brands, status string) ([]*master.Storelist, error) {
	query := `SELECT prm_outlet.id_outlet,prm_outlet.nomor,prm_outlet.nama_outlet,prm_outlet.add_ons,prm_outlet.create_user,prm_outlet.status,prm_outlet.token,prm_outlet.brands,prm_outlet.prefix,prm_outlet.service_charge,prm_outlet.tax,prm_outlet.alamat,prm_outlet.telp_1,prm_outlet.telp_2,users.nama_user FROM prm_outlet INNER JOIN users ON users.nik = prm_outlet.create_user WHERE prm_outlet.brands=? AND prm_outlet.status=?`
	rows, err := database.DBMASTER.Query(query, brands, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var stores []*master.Storelist
	for rows.Next() {
		var st master.Storelist
		err := rows.Scan(&st.Idoutlet, &st.Nomor, &st.Namaoutlet, &st.Addons, &st.Createuser, &st.Status, &st.Token, &st.Brands, &st.Prefix, &st.ServiceCharge, &st.Tax, &st.Alamat, &st.Telp1, &st.Telp2, &st.Nama_user)
		if err != nil {
			return nil, err
		}
		stores = append(stores, &st)
	}
	return stores, nil
}
