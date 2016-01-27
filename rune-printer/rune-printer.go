// The rune printer. rune-printer prints the the last rune (or character)
// and up to the first three runes of a message that the user
// enters to the screen.
package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/gophercoders/simpleio"
)

func main() {
	// tell the user what the program does.
	fmt.Println("rune-printer prints the the last rune (or character)")
	fmt.Println("and up to the first three runes in a message")
	fmt.Println("that the user enters.")
	// Ask the user to enter their message
	fmt.Println("Enter a message:")
	var message string
	//
	// Your turn
	//
	// Use the input pattern to set the value of the variable message
	// to the users input.
	// Put your answer on the next line


	// Before you can print the individual runes, you need to know
	// how many runes are in the string. Otherwise how do you know
	// which one is the last one?
	var numberOfRunesInMessage int
	//
	// Your turn
	//
	// Now you have to use the utf8.RuneCountInString function to
	// count the number of runes in the message.
	// Complete the next line
	numberOfRunesInMessage =

	// Now you are ready to create the "slice of runes"
	// so you can access each individual rune (or charcter)
	//
	// Your Turn
	//
	// You need to create a new variable called runes that is "the
	// slice of runes" type
	// Complete the next line
	var

	// Your Turn
	//
	// Now you need to convert the message, which is a string, to a
	// slice of runes.
	// Complete the next line
	runes =


	// Now print out the ansers to the user
	fmt.Print("Your message has ")
	fmt.Print(numberOfRunesInMessage)
	fmt.Println(" runes in it.")

	// Your turn
	//
	// You'll need to if-else statement to complete this.
	// If there is exactly one rune in the message the program should
	// 		print out "The first rune in your message is: "
	//		followed by the first rune
	// if there are exactly two runes in the message the program should
	//		print out "The first two runes in your message are:"
	//		followed by the first two runes
	// otherwise print out the message
	// 		"the first three runes in your message are"
	//		followed by the first three runes
	// Start you answer on the next line

	// Your turn
	//
	// Before the program ends it always prints
	// 		"The last rune in your message is: "
	//		followed by the last rune.
	// start your answer on the next line
}
