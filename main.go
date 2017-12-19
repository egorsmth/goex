package main

import (
	"fmt"
	"time"
	"flag"
	"github.com/egorsmth/gopher_1/tasks"
	"github.com/egorsmth/gopher_1/input"
)

func main() {
	file := flag.String("file", "tasks/tasks.csv", "tasks file path")
	limit := flag.Int("limit", 30, "time limit to answer quizes")
	flag.Parse()

	timer := time.NewTimer(time.Second * time.Duration(*limit))
	problems := tasks.GetTasks(*file)
	total := len(problems)
	correct := 0
	timerCh := make(chan bool)
	answerCh := make(chan bool)
	go func () {
		<-timer.C
		timerCh <- true
	}()

	
	
	for _, task := range problems {
		go func () {
			answer := input.Ask(tasks.GetQuestion(task))
			if answer == tasks.GetAnswer(task) {
				answerCh <- true
			} else {
				answerCh <- false
			}
		}()

		select {
		case <- timerCh:
			fmt.Print("Time out!\n")
			fmt.Print("Your result: ", correct, "/", total, "\n")
			return
		case ans := <- answerCh:
			if ans {
				correct += 1
			}
		}
	}
	fmt.Print("Your result: ", correct, "/", total, "\n")
}