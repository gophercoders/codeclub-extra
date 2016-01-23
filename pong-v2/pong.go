package main

import
// This is the graphics library we are going to use. It is called the
// Simple Direct Media Library. SDL for short. We need this to create the
// window and to provide the drawing functions we need.

"github.com/gophercoders/toolbox"

// These are the variables for the graphics library
// They have to be outside of the main function because the functions at the
// end of the file need them.
// This the window we are going to draw into
var window toolbox.Window

// These variabels are important. They are the width and height of the window
// If you change these you will change the size of the image
var windowWidth int
var windowHeight int

// The programs main function
func main() {
	// ---- This is the start of Owen's graphics setup code ----

	// First we have to initalise the toolbox, before we can use it
	toolbox.Initialise()

	// defer is a go keyword and a special feature.
	// This means that go will automatically call the function toolbox.Close() before
	// the program exits for us. We don't have to remember to put this at the end!
	defer toolbox.Close()

	// if you want to change these try 800 for the width and 600 for the height
	windowWidth = 1024
	windowHeight = 768

	// Now we have to create the window we want to use.
	// We need to tell the SDL library how big to make the window of the correct
	// size - that's what the bit in the brackets does
	window = toolbox.CreateWindow("Pong Game", windowWidth, windowHeight)
	// automatically destroy the window when the program finishes
	defer toolbox.DestroyWindow(window)

	// Set a black i.e. RGBA (0,0,0,0) background colour and clear the window
	toolbox.SetBackgroundColour(0, 0, 0)
	toolbox.ClearBackground()

	// ---- This is the end of Owen's graphics setup code ----

	// defer any cleanup actions
	defer cleanup()
	// initialise the games variables.
	initialise()
	// render everything initially so that we can see the game before it starts
	render()
	// now start the main game loop of the game.
	gameMainLoop()
}

// Initialise sets the inital values of the game state variables.
// Initialise must be called before the games main loop starts.
func initialise() {
}

// GameMainLoop controls the game. It performs three manin tasks. The first task
// is to get the users input. The second task is to update the games state based
// on the user input and the rules of the game. The final task is to update, or
// render, the changes to the screen.
func gameMainLoop() {
	for {
		getInput()
		updateState()
		render()
	}
}

// Cleanup is used to ensure that we free all memory before the program
// exists.
func cleanup() {
}

// GetInput gets the users input and updates the game state variables that realte
// to the users input, for example, the direction that the user wants to move their
// bat in.
func getInput() {
}

// UpdateGameState updates the game state variables based on the user input and
// the rules of the game.
func updateState() {
}

// Render updates the screen, based on the new positions of the bats and the ball.
func render() {
	// Set a black i.e. RGBA (0,0,0,0) background colour and clear the window
	toolbox.ShowWindow()
}
