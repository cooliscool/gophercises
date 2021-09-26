// Ajmal Moochingal
// Gophercises - #1 Quiz Game
// https://courses.calhoun.io/lessons/les_goph_01

// A Quiz program that:
// 1. Accepts questions as a CSV
// 2. Takes in commmandline args to configure certain parameters
// 3. Has a Timer that starts ticking and concludes the quiz if time runs out

// Tasks:
// [✓] commandline args to binary
// [✓] open file and parse questions as csv
// [✓] timer, scoring

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type QuizProblem struct {
	Question string
	Answer   string
}

var correctAnswers int
var totalQuestions int

func createQuestions(data [][]string) []QuizProblem {
	var problems []QuizProblem
	for _, line := range data {
		var p QuizProblem

		p.Question = line[0]
		p.Answer = line[1]
		problems = append(problems, p)
	}

	return problems
}

func printResults(x int, y int) {
	fmt.Println("Quiz ended!\nResults : \n", x, "correct answers out of", y, "\nThankyou!")
	// TODO: print which anwsers went wrong
}

func timerElapse(s int) {
	time.Sleep(time.Duration(int(time.Second) * s))
	fmt.Println("Quiz Timed out!")
	printResults(correctAnswers, totalQuestions)
	os.Exit(0)
}

func main() {

	totalQuestions = 0
	correctAnswers = 0
	print := fmt.Println

	// Define flags
	csvSourceFile := flag.String("infile", "problems.csv", "input csv that contains quiz problems")
	timeLimit := flag.Int("time", 5, "time limit")

	flag.Parse()
	// print(*csvSourceFile)

	f, err := os.Open(*csvSourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // close file at the end of program

	// read problems
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// print(data)
	problems := createQuestions(data)
	// fmt.Printf("%+v\n", problems)
	totalQuestions = len(problems)

	////// start timer
	// there should be a non-blocking thread running timer, and taking action if timer lapses.
	//
	// while questionsCompleted <= totalQuestions {
	//     show question
	//     keep track of score
	// }
	//////

	print("You have", *timeLimit, "seconds to complete the Quiz. Press Enter to start")
	var start int
	fmt.Scanf("%d", &start)

	// Set the timer
	// concurrency in go is cool
	go timerElapse(*timeLimit)

	// Start Quiz
	for i, problem := range problems {
		var input string
		print("Question #", i+1, "\n", problem.Question, " : ")
		fmt.Scanf("%s", &input)

		if input == problem.Answer {
			correctAnswers += 1
			print("Correct!")
		} else {
			print("Wrong!")
		}
	}

	printResults(correctAnswers, totalQuestions)
}
