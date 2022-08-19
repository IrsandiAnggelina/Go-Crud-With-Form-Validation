package pasiencontroller

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/IrsandiAnggelina/go-crud/entities"
	"github.com/IrsandiAnggelina/go-crud/libraries"
	"github.com/IrsandiAnggelina/go-crud/models"
)

var validation = libraries.NewValidation()
var pasienModel = models.NewPasienModel()

func Index(rw http.ResponseWriter, r *http.Request){
	pasien, _ := pasienModel.FindAll()

	data := map[string]interface{}{
		"pasien":pasien,
	}

//manggil view
	temp, err := template.ParseFiles("views/pasien/index.html")
	if err != nil{
		panic(err)
	}
	temp.Execute(rw, data)
}

func Add(rw http.ResponseWriter, r *http.Request){
	//atribut request
	if r.Method == http.MethodGet{
	//response
	temp, err := template.ParseFiles("views/pasien/add.html")
	if err != nil{
		panic(err)
	}
	temp.Execute(rw, nil)
	} else if r.Method == http.MethodPost{

		r.ParseForm()

		var pasien entities.Pasien
		//namalengkap yang di buat di create.html berupa nama_lengkap, akan disimpan di entities.Pasien,NamaLengkap
		pasien.NamaLengkap = r.Form.Get("nama_lengkap")
		pasien.Nik = r.Form.Get("nik")
		pasien.JenisKelamin = r.Form.Get("jenis_kelamin")
		pasien.TempatLahir = r.Form.Get("tempat_lahir")
		pasien.TanggalLahir = r.Form.Get("tanggal_lahir")
		pasien.Alamat = r.Form.Get("alamat")
		pasien.NoHp = r.Form.Get("no_hp")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pasien)

		if vErrors != nil{
			data["pasien"] = pasien
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data Pasien Berhasil Disimpan" 
			pasienModel.Create(pasien) //insert data ke database
		}
		
		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(rw, data)
	}

}

func Update(rw http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		queryString := r.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64) //parameter 10 dan side 64. convert ke int 64

		var pasien entities.Pasien
		pasienModel.Find(id, &pasien)

		data := map[string]interface{}{
			"pasien" : pasien,
		}

		temp, err := template.ParseFiles("views/pasien/udate.html")
		if err != nil{
			panic(err)
		}
		temp.Execute(rw, data)
		} else if r.Method == http.MethodPost{
	
			r.ParseForm()
	
			var pasien entities.Pasien
			pasien.Id, _ = strconv.ParseInt(r.Form.Get("id"), 10,64)
			//namalengkap yang di buat di create.html berupa nama_lengkap, akan disimpan di entities.Pasien,NamaLengkap
			pasien.NamaLengkap = r.Form.Get("nama_lengkap")
			pasien.Nik = r.Form.Get("nik")
			pasien.JenisKelamin = r.Form.Get("jenis_kelamin")
			pasien.TempatLahir = r.Form.Get("tempat_lahir")
			pasien.TanggalLahir = r.Form.Get("tanggal_lahir")
			pasien.Alamat = r.Form.Get("alamat")
			pasien.NoHp = r.Form.Get("no_hp")
	
			var data = make(map[string]interface{})
	
			vErrors := validation.Struct(pasien)
	
			if vErrors != nil{
				data["pasien"] = pasien
				data["validation"] = vErrors
			} else {
				data["pesan"] = "Data Pasien Berhasil Diperbarui" 
				pasienModel.Update(pasien) //insert data ke database
			}
			
			temp, _ := template.ParseFiles("views/pasien/update.html")
			temp.Execute(rw, data)
		}
	
}

func Delete(rw http.ResponseWriter, r *http.Request){
	
	queryString  := r.URL.Query()
	id, _ :=  strconv.ParseInt(queryString.Get("id"), 10, 64)

	pasienModel.Delete(id)

	http.Redirect(rw, r, "/pasien", http.StatusSeeOther)
}