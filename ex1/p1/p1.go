package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type args struct {
	filename string
}

func parseArgs() args {
	f := flag.String("f", "problems.csv", "Name of quiz CSV file")
	flag.Parse()
	return args{filename: *f}
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

func giveQuiz(quiz [][]string) {
	r := bufio.NewReader(os.Stdin)
	numCorrect := 0

	for x := range quiz {
		fmt.Printf("%s ", quiz[x][0])
		answer, _ := r.ReadString('\n')
		if normalizeAnswer(answer) == normalizeAnswer(quiz[x][1]) {
			numCorrect++
		}
	}

	fmt.Printf("\nYou got %d of %d correct!", numCorrect, len(quiz))
}

func main() {
	args := parseArgs()
	quiz := loadQuiz(args.filename)
	giveQuiz(quiz)
}
