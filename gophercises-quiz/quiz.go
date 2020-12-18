package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

type quiz []problem

func readQuestions(filename *string) quiz {
	file, err := os.Open(*filename)
	if err != nil {
		exit(fmt.Sprintf("Cannot open specified file %s\n", *filename))
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Cannot open CSV file"))
	}

	retQuiz := make(quiz, len(lines))
	for i, line := range lines {
		var p problem
		p.q = strings.TrimSpace(line[0])
		p.a = strings.TrimSpace(line[1])
		retQuiz[i] = p
	}
	return retQuiz
}
