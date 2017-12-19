package tasks

import (
	"encoding/csv"
	"log"
	"os"
	"bufio"
)

func GetTasks(file string) (records [][]string) {

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	reader := csv.NewReader(r)

	records, err = reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func GetQuestion(task []string) string {
	return task[0]
}

func GetAnswer(task []string) string {
	return task[1]
}