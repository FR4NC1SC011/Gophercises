package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	var puntuacion int

	file := flag.String("csvfile", "problems.csv", "CSV file with the questions")
	seconds := flag.Int("time", 10, "Time in seconds to use as a timer")
	flag.Parse()

	sec := time.Duration(*seconds) * time.Second
	reader := bufio.NewReader(os.Stdin)
	csvfile, err := os.Open(*file)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(csvfile)

	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question: %s Answer: ", record[0])
		answer, _ := reader.ReadString('\n')
		// Convert CRLF to LF
		answer = strings.Replace(answer, "\n", "", -1)

		if strings.Compare(answer, record[1]) == 0 {
			fmt.Println("Respuesta Correcta")
			puntuacion += 1
		} else {
			fmt.Println("Respuesta Incorrecta Sigue Jugando")
		}
		timer := time.AfterFunc(sec, func() {
			fmt.Println()
			fmt.Println("Tiempo Agotado")
			fmt.Printf("Puntos Totales: %d\n", puntuacion)
			os.Exit(1)
		})

		defer timer.Stop()
	}

	fmt.Printf("Puntos Totales: %d", puntuacion)

}
