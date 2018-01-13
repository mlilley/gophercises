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
	file string
}

type problem struct {
	question string
	answer   string
}

type quiz struct {
	problems []problem
	correct  int
}

func parseArgs() args {
	file := flag.String("f", "problems.csv", "CSV questions file")
	flag.Parse()
	return args{file: *file}
}

func parseProblem(line string) (problem, error) {
	terms := strings.Split(line, ",")
	if len(terms) != 2 {
		return problem{}, errors.New("unable to parse problem")
	}
	return problem{question: terms[0], answer: terms[1]}, nil
}

func loadQuiz(filename string) quiz {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var q = quiz{}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		p, err := parseProblem(line)
		if err == nil {
			q.problems = append(q.problems, p)
		}
	}

	return q
}

func poseProblem(p problem) bool {
	fmt.Println(p.question)
	txt, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	txt = strings.Trim(txt, "\n")
	return txt == p.answer
}

func main() {
	args := parseArgs()
	quiz := loadQuiz(args.file)

	for _, p := range quiz.problems {
		if poseProblem(p) {
			quiz.correct++
		}
	}

	fmt.Println("You got", quiz.correct, "out of", len(quiz.problems))
}
