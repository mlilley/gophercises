package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

var (
	argFile    = flag.String("f", "problems.csv", "CSV questions file")
	argTimeout = flag.Int("t", 10, "Seconds within which to compete quiz")
)

func loadQuiz(filename string) (pz []problem, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return []problem{}, err
	}
	defer f.Close()

	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return []problem{}, err
	}

	for _, p := range data {
		pz = append(pz, problem{question: p[0], answer: p[1]})
	}

	return pz, nil
}

func poseQuiz(problems []problem) {
	stdIn := bufio.NewReader(os.Stdin)

	fmt.Printf("You have %d seconds to complete this quiz. Press Enter to begin.", *argTimeout)
	stdIn.ReadLine()

	done := make(chan bool)
	ticker := time.NewTicker(time.Second * time.Duration(*argTimeout))
	n := 0

	go func() {
		for _, p := range problems {
			fmt.Println(p.question)
			answer, _ := stdIn.ReadString('\n')
			answer = strings.TrimSpace(answer)
			if answer == p.answer {
				n++
			}
		}
		done <- true
	}()

	select {
	case <-done:
	case <-ticker.C:
		fmt.Printf("You got %d out of %d in the allotted time.\n", n, len(problems))
	}
}

func main() {
	flag.Parse()

	problems, err := loadQuiz(*argFile)
	if err != nil {
		panic(err)
	}

	poseQuiz(problems)
}
