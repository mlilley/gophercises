package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type args struct {
	filename string
	timeout  int
}

func parseArgs() args {
	f := flag.String("f", "problems.csv", "Name of quiz CSV file")
	t := flag.Int("t", 30, "Quiz time limit (0 for none)")
	flag.Parse()
	return args{filename: *f, timeout: *t}
}

func loadQuiz(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

func normalizeAnswer(v string) string {
	return strings.ToUpper(strings.Trim(v, " \t\n"))
}

func giveQuiz(quiz [][]string, timeout int) {
	c := make(chan bool)
	r := bufio.NewReader(os.Stdin)
	numCorrect := 0

	go func() {
		for x := range quiz {
			fmt.Printf("%s ", quiz[x][0])
			answer, _ := r.ReadString('\n')
			if normalizeAnswer(answer) == normalizeAnswer(quiz[x][1]) {
				numCorrect++
			}
		}
		c <- true
	}()

	select {
	case <-c:
		fmt.Printf("\nYou got %d of %d correct!", numCorrect, len(quiz))
	case <-time.After(time.Duration(timeout) * time.Second):
		fmt.Printf("\nYou got %d of %d correct!", numCorrect, len(quiz))
	}
}

func main() {
	args := parseArgs()
	quiz := loadQuiz(args.filename)
	giveQuiz(quiz, args.timeout)
}
