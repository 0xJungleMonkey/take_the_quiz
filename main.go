package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// struct connect between csv [][]string and the quiz program, separate ask and answer for future use.
type Question struct {
	ask    string
	answer string
}

// get the questions.
func readCSV(rs io.ReadSeeker) ([][]string, error) {
	//skip first row
	row1, err := bufio.NewReader(rs).ReadSlice('\n')
	if err != nil {
		return nil, err
	}
	_, err = rs.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		return nil, err
	}
	// Read remaining rows
	r := csv.NewReader(rs)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// the main quiz process
func quiz(records [][]string) {
	score := 0
	for _, record := range records {
		ques := Question{
			ask:    record[0],
			answer: record[1],
		}
		fmt.Println(ques.ask)
		// var then variable name then variable type
		var first string

		// Taking input from user
		fmt.Scanln(&first)
		if first == ques.answer {
			fmt.Println("correct!")
			score += 1
		}

	}
	fmt.Printf("You get %d.", score)
}

func main() {
	fi, err := os.Open("math.csv")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	records, err := readCSV(fi)
	if err != nil {
		panic(err)
	}
	quiz(records)
}
