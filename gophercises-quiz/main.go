package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func exit(message string) {
	fmt.Println(message)
	os.Exit(-1)
}

func main() {
	csvFileName := flag.String("csv", "questions.csv", "A CSV file containing questions and answers")
	timeLimit := flag.Int("limit", 15, "Time in seconds for entire quiz")
	flag.Parse()

	quiz := readQuestions(csvFileName)

	fmt.Println("Welcome to the quiz !!!")

	correct := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for i, q := range quiz {
		fmt.Printf("Question %d. %s = ", i+1, q.q)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You answered %d correctly out of %d (Timeout)", correct, len(quiz))
			return

		case answer := <-answerChannel:
			if answer == q.a {
				correct++
			}
		}
	}

	fmt.Printf("You answered %d correctly out of %d", correct, len(quiz))
}
