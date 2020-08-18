package main

import "fmt"

func main() {
	printMessage("Hello")
}

// TODO
func printMessage(msg string) string {
	// FIXME
	return msg + " World"
}

func run() {
	var target, num = -5, 3
	target =- num // Noncompliant; target = -3. Is that really what's meant?
	target =+ num // Noncompliant; target = 3
	fmt.Println(target)
	printMessage("This should be a constant") // Noncompliant; 'This should ...' is duplicated 3 times
	printMessage("This should be a constant")
	printMessage("This should be a constant")
}
