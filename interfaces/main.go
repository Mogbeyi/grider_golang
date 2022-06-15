package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}
type frenchBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	fb := frenchBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(fb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func (frenchBot) getGreeting() string {
	return "Bonjour"
}