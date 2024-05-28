package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type Question struct{
	question string
	options []string
	answer string
}
var questions= []Question{

	{question:"What is the capital of France?",
	options:[]string{"a) Paris,b) London, c)Berlin, d)Madrid"},
	answer:"a"},
	{
		question:"What is 2 +2?",
		options:[]string{"a)4,b)3,c)5,d)11"},
		answer:"a",

	},{
		question: "Who wrote 'To Kill a Mockingbird'?",
        options:  []string{"a) Harper Lee", "b) J.K. Rowling", "c) Ernest Hemingway", "d) Mark Twain"},
        answer:   "a",
	},
}
func main()  {
	reader:=bufio.NewReader(os.Stdin)
	score:=0
	for i,q:=range questions {
		fmt.Printf("Question %d:%s\n",i+1,q.question)
		for _,option:=range q.options{
			fmt.Println(option)
		}
		fmt.Print("Enter your answer:")
		answer,_:=reader.ReadString('\n')
		answer=strings.TrimSpace(answer)
		
		if answer==q.answer{
       fmt.Println("Correct")
	   score++
	
	} else{
		fmt.Println("Wrong! The correct answer was ",q.answer)
	}
	fmt.Println()
	}
fmt.Printf("You scored %d out of  %d.\n",score ,len(questions))
}
