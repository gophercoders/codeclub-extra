package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/gophercoders/simpleio"
)

func main() {
	fmt.Println("The reverse-message program prints out the message")
	fmt.Println("that you type in, but backwards!")
	fmt.Println("Please enter your message:")
	// Your turn
	// message is the variable that stores what the user types in.
	// What type does it need to be?
	// You need to complete the next line.
	var message
	// Your turn
	// How do you store what the user typed in in the variable called message?
	// Write you answer on the next line.


	// Your turn
	// runes is the variable we are going to store the "characters" of the
	// message in. Runes has to be use the new "slice of runes type"
	// Can you remember what the type name is?
	// You need to complete the next line
	var runes
	// Your turn
	// Now you need to convert the variable called message, which is a string,
	// to a slice of runes, and store the answer in the variable called "runes"
	// You need to complete the next line.
	runes =

	// The program needs to know the number of runes in the message.
	// You've seen how to do this in the rune-printer program.
	// The next two lines do this for your.
	var numberOfRunesInMessage int
	numberOfRunesInMessage = utf8.RuneCountInString(message)
	// Print out the number of runes to the screen
	fmt.Print("numberOfRunesInMessage = ")
	fmt.Println(numberOfRunesInMessage)
	// We are going to use the variable i as the loop counter.
	// The name is traditional.
	var i int
	// Your Turn
	// What value does the variable i need to start with?
	// Remember you are going to go backwards though the charcters in the message
	// Hint: What's the last valid position in the message?
	i =
	// Your Turn
	// Now you need to use the new loop pattern.
	// You want to create a loop that counts down from the
	// last character in the slice of runes, "runes", to the beginning.
	// The loop variable i has to start at the last valid position in the
	// string.
	// Each time through the loop you want to print out the the
	// character at position i
	// Each time through the loop you want to make i one smaller.
	// You have to keep the loop going until you have printed
	// all of the characters in the rune slice. So the last character
	// you print should be the first character in the slice of runes
	// Write your answer starting on the line below.
	for
	// Your Turn
	// Remember to print a new line to the screen at the end.
	// If you don't the reversed message and the terminal prompt will be
	// on the same line.

}
