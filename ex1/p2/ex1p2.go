package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type args struct {
	file    string
	timeout int
}

type problem struct {
	question string
	answer   string
}

type quiz struct {
	problems []problem
	correct  int
	timeout  int
}

func parseArgs() args {
	file := flag.String("f", "problems.csv", "CSV questions file")
	timeout := flag.Int("t", 10, "Seconds within which to compete quiz")
	flag.Parse()
	return args{
		file:    *file,
		timeout: *timeout,
	}
}

func parseProblem(line string) (problem, error) {
	terms := strings.Split(line, ",")
	if len(terms) != 2 {
		return problem{}, errors.New("unable to parse problem")
	}
	return problem{question: terms[0], answer: terms[1]}, nil
}

func loadQuiz(a args) quiz {
	data, err := ioutil.ReadFile(a.file)
	if err != nil {
		panic(err)
	}

	var q = quiz{timeout: a.timeout}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		p, err := parseProblem(line)
		if err == nil {
			q.problems = append(q.problems, p)
		}
	}

	return q
}

func promptToBegin(q quiz) {
	fmt.Printf("You have %d seconds to complete this quiz. Press enter to begin.", q.timeout)
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func poseProblem(p problem) bool {
	fmt.Println(p.question)
	txt, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	txt = strings.Trim(txt, "\n")
	return txt == p.answer
}

func poseQuiz(q *quiz) {
	for _, p := range q.problems {
		if poseProblem(p) {
			q.correct++
		}
	}
}

func printResult(q quiz) {
	fmt.Printf("You got %d out of %d in the allotted time.\n", q.correct, len(q.problems))
}

func main() {
	args := parseArgs()
	quiz := loadQuiz(args)

	promptToBegin(quiz)
	poseQuiz(&quiz)
	printResult(quiz)
}
