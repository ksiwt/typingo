package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/tjarratt/babble"
)

const timeLimit = 30

var score int = 0

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func generateRandomWord() string {
	babbler := babble.NewBabbler()
	babbler.Count = 1
	return babbler.Babble()
}

func isCorrect(answer, question string) bool {
	return answer == question
}

func main() {
	// start.
	fmt.Println("Complete a Typing Test in 30 Seconds!")

	ctx := context.Background()
	timeout := timeLimit * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// get answer.
	strCh := input(os.Stdin)

	for {
		// set question.
		question := generateRandomWord()
		fmt.Println(question)

		select  {
		case input := <-strCh:
			if isCorrect(input, question) {
				score++
				fmt.Println("Good!")
			} else {
				fmt.Println("Wops!")
			}

		case <-ctx.Done():
			fmt.Println("Time up!")
			fmt.Printf("Your score is %d.", score)
			return
		}
	}
}
