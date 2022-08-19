package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/IrsandiAnggelina/go-crud/config"
	"github.com/IrsandiAnggelina/go-crud/entities"
)

type PasienModel struct {
	conn *sql.DB
}

func NewPasienModel() *PasienModel{
	conn, err := config.DBConnection()
	if err != nil{
		panic(err)
	}

	return &PasienModel{
		conn : conn, //conn NewPasienModel masuk ke Sturct PasienModel
	}
}
//menampilkan data yang disimpan ke browser
func (p *PasienModel) FindAll() ([]entities.Pasien, error){

	rows, err := p.conn.Query("select * from pasien")
	if err != nil {
		return []entities.Pasien{}, err
	}
	defer rows.Close()

	var dataPasien []entities.Pasien
	for rows.Next() {
		var pasien entities.Pasien
		rows.Scan(&pasien.Id, 
			&pasien.NamaLengkap,
			&pasien.Nik, 
			&pasien.JenisKelamin, 
			&pasien.TempatLahir, 
			&pasien.TanggalLahir, 
			&pasien.Alamat, 
			&pasien.NoHp,
		)
		if pasien.JenisKelamin == "1"{
			pasien.JenisKelamin = "Laki-Laki"
		} else {
			pasien.JenisKelamin = "Perempuan"
		}

		//2022-08-18 => yyyy-mm-dd
		tgl_lahir, _ := time.Parse("2022-08-18", pasien.TanggalLahir)
		//18-08-2022 => dd-mm-yyyy
		pasien.TanggalLahir = tgl_lahir.Format("18-08-2022")

		dataPasien = append(dataPasien, pasien)
	}
	return dataPasien, nil	
}

//proses menyimpan data ke db
func (p *PasienModel) Create(pasien entities.Pasien) bool{
	result, err := p.conn.Exec("insert into pasien (nama_lengkap, nik, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat, no_hp) values(?,?,?,?,?,?,?)", 
	pasien.NamaLengkap, pasien.Nik, pasien.JenisKelamin, pasien.TempatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHp)

	if err != nil {
		fmt.Println(err)
		return false 
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *PasienModel) Find(id int64, pasien *entities.Pasien) error{
	return p.conn.QueryRow("select * from pasien where id = ?", id).Scan(&pasien.Id, 
		&pasien.NamaLengkap,
		&pasien.Nik, 
		&pasien.JenisKelamin, 
		&pasien.TempatLahir, 
		&pasien.TanggalLahir, 
		&pasien.Alamat, 
		&pasien.NoHp,
	)
}

func (p *PasienModel) Update(pasien entities.Pasien) error{

	_, err := p.conn.Exec("update pasien set nama_lengkap = ?, nik = ?, jenis_kelamin = ?, tempat_lahir = ?, tanggal_lahir = ?, alamat = ?, no_hp = ? where id = ?",
					pasien.NamaLengkap, pasien.Nik, pasien.JenisKelamin, pasien.TempatLahir, pasien.TanggalLahir, pasien.Alamat, pasien.NoHp, pasien.Id)
	
	if err != nil{
		return err
	}

	return nil
}

func (p *PasienModel) Delete(id int64){
	p.conn.Exec("delete from pasien where id = ?", id)
}