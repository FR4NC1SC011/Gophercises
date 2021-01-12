package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	valid "github.com/asaskevich/govalidator"
	_ "github.com/mattn/go-sqlite3"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "fco"
	password = "pass"
)

func main() {
	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db")
	check(err)
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
	defer sqliteDatabase.Close()

	createTable(sqliteDatabase)
	insertStudent(sqliteDatabase, "Francisco", "(614) 1634 881")
	insertStudent(sqliteDatabase, "Martin", "6141561325")
	insertStudent(sqliteDatabase, "Javier", "614-123-4567")

	displayStudents(sqliteDatabase)

}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE client (
		"idClient" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"number" TEXT
	  );`

	log.Println("Create phone number table...")
	statement, err := db.Prepare(createStudentTableSQL)
	check(err)
	statement.Exec()
	log.Println("phone number table created")
}

func insertStudent(db *sql.DB, name string, number string) {
	log.Println("Inserting client record")
	insertStudentSQL := `INSERT INTO client(name, number) VALUES (?, ?)`
	statement, err := db.Prepare(insertStudentSQL)
	check(err)
	_, err = statement.Exec(name, number)
	check(err)
}

func displayStudents(db *sql.DB) {
	row, err := db.Query("SELECT * FROM client ORDER BY name")
	check(err)
	defer row.Close()

	for row.Next() {
		var id int
		var name string
		var number string
		row.Scan(&id, &name, &number)
		log.Println("Student: ", name, " ", number)
	}

}

func normalize(number string) string {
	var normalized_number string

	noSpaceNumber := strings.ReplaceAll(number, " ", "")
	phone_number := strings.Split(noSpaceNumber, "")

	for _, digit := range phone_number {
		if valid.IsInt(digit) {
			normalized_number += digit
		}
	}

	return normalized_number
}

// re := regexp.MustCompile("\\D")
// return re.ReplaceAllString(number, "")

func check(e error) {
	if e != nil {
		panic(e)
	}
}
