package main

import "net/http"
import "html/template"
import "fmt"
import "database/sql"
import "os"
import _ "github.com/go-sql-driver/mysql"

var username, pass, db string

func connect() *sql.DB {
	var db, err = sql.Open("mysql", username+":"+pass+"@/"+db)
	err = db.Ping()
	if err != nil {
		fmt.Println("database tidak bisa dihubungi")
		os.Exit(0)

	}
	return db

}

func input_db_akses() {
	fmt.Print("masukkan username mysql : ")
	fmt.Scanln(&username)
	fmt.Print("masukkan password mysql : ")
	fmt.Scanln(&pass)
	fmt.Print("masukkan database mysql : ")
	fmt.Scanln(&db)

	db := connect()
	defer db.Close()

}

const html = `<html>

<form action="/mau_input" method="post">
masukan id mhs : <br>
<input type="number" name="id"><br>
masukan nama mhs : <br>
<input type="text" name="nama"><br>
<input type="submit" value="insert"><br>
</form>

</html>`

func utama(res http.ResponseWriter, req *http.Request) {
	data := map[string]string{

		"nama": "eko",
	}
	halaman, _ := template.New("tmp").Parse(html)
	halaman.Execute(res, data)

}

func input(res http.ResponseWriter, req *http.Request) {
	db := connect()
	defer db.Close()

	id := req.FormValue("id")
	name := req.FormValue("nama")
	db.Exec("insert into mhs value (?,?)", id, name)

	http.Redirect(res, req, "/", 301)
}

func main() {
	input_db_akses()
	http.HandleFunc("/", utama)
	http.HandleFunc("/mau_input", input)
	http.ListenAndServe(":8080", nil)
}
