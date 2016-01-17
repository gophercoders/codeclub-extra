package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Print out some information about the program
	fmt.Println("rune-counter counts the number or runes (or characters)")
	fmt.Println("in your message.")

	// Ask the user to enter a message
	fmt.Println("Enter your message:")
	var message string
	// Your Turn
	// Now use the input pattern to set the value of the variable
	// called messge to the users input.

	var numberOfRunesInMessage int
	// The next line in new.
	// To count the number of runes in the string we use the
	// RuneCountInString function, like this.
	numberOfRunesInMessage = utf8.RuneCountInString(message)

	// Next we print out the message
	fmt.Print("Your message ")
	fmt.Print(message)
	fmt.Print(" has ")

	// But the next line in the output depends on how many
	// runes there are in the string. The number of runes in
	// the message is in the variable called numberOfRunesInMessage
	//
	// Your Turn
	// You need to use some if-else statements to do this.
	// If there are no runes then print
	//		"no runes at all!"
	// if there is exactly one rune then print
	//		"just one rune."
	// if there are less than 11 runes then print
	//		"between 2 and 10 runes"
	// if there are less than 21 runes print
	//		"between 2 and 20 runes"
	// else print
	//		"more than 20 runes. Wow, What a long message!"
}
