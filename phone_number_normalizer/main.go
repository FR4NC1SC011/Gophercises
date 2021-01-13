package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	valid "github.com/asaskevich/govalidator"
	_ "github.com/mattn/go-sqlite3"
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
	insertStudent(sqliteDatabase, "Martin", "614 1561325")
	insertStudent(sqliteDatabase, "Javier", "614-123-4567")
	insertStudent(sqliteDatabase, "Francisco", "(614) 163 4881")
	insertStudent(sqliteDatabase, "Hector", "614 1561325")
	insertStudent(sqliteDatabase, "Ousmane", "6141561325")

	displayStudents(sqliteDatabase)
	fmt.Println("\n\n")
	normalizeNumbers(sqliteDatabase)
	fmt.Println("\n\n")
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

func normalizeNumbers(db *sql.DB) {
	numbers := getNumbers(db)

	for i, number := range numbers {
		normalizeNumberSQL := fmt.Sprintf(`UPDATE client SET number = %s WHERE idClient = %d`, number, i+1)
		statement, err := db.Prepare(normalizeNumberSQL)
		check(err)
		statement.Exec()
	}

	//fmt.Println("DUPLICATE:", unique(numbers))
}

func displayStudents(db *sql.DB) {
	deleteDuplicateNumbersSQL := `DELETE FROM client WHERE rowid NOT IN (SELECT min(rowid) FROM client GROUP BY number)`
	statement, err := db.Prepare(deleteDuplicateNumbersSQL)
	check(err)
	statement.Exec()

	row, err := db.Query("SELECT * FROM client ORDER BY idClient")
	check(err)
	defer row.Close()

	for row.Next() {
		var id int
		var name string
		var number string
		row.Scan(&id, &name, &number)
		log.Println("Student: ", id, " ", name, " ", number)
	}
}

func getNumbers(db *sql.DB) []string {
	var numbers []string
	var normalized_numbers []string

	row, err := db.Query("SELECT number FROM client ORDER BY idClient")
	check(err)
	defer row.Close()

	for row.Next() {
		var number string
		row.Scan(&number)
		numbers = append(numbers, number)
	}

	for _, number := range numbers {
		normalized_numbers = append(normalized_numbers, normalize(number))
	}

	return normalized_numbers
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

func unique(arr []string) []int {
	occured := map[string]bool{}
	result := []int{}
	for e := range arr {
		if occured[arr[e]] != true {
			occured[arr[e]] = true

			fmt.Println(arr[e])
			result = append(result, e+1)
		}
	}

	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
