package main

import (
	// This is the graphics library we are going to use. It is called the
	// Simple Direct Media Library. SDL for short. We need this to create the
	// window and to provide the drawing functions we need.
	"github.com/veandco/go-sdl2/sdl"
)

// These are the variables for the graphics library
// They have to be outside of the main function because the functions at the
// end of the file need them.
// This the window we are going to draw into
var window *sdl.Window

// This the abstraction to the graphics hardware inside the computer
// that actually does the drawing
var renderer *sdl.Renderer

// ---- Game State variables ----
// The quit flag this is used to control the main game loop.
// If quit is true then the user wants to finish the game. This will
// break the main game loop.
var quit bool

// The programs main function
func main() {
	// ---- This is the start of Owen's graphics setup code ----

	// First we have to initalise the SDL library, before we can use it
	sdl.Init(sdl.INIT_EVERYTHING)
	// defer is a go keyword and a special feature.
	// This means that go will automatically call the function sdl.Quit() before
	// the program exits for us. We don't have to remember to put this at the end!
	defer sdl.Quit()

	// These variabels are important. They are the width and height of the window
	// If you change these you will change the size of the image
	var windowWidth int
	var windowHeight int
	// if you want to change these try 800 for the width and 600 for the height
	windowWidth = 1024
	windowHeight = 768

	// Now we have to create the window we want to use.
	// We need to tell the SDL library how big to make the window of the correct
	// size - that's what the bit in the brackets does
	window = createWindow(windowWidth, windowHeight)
	// automatically destroy the window when the program finishes
	defer window.Destroy()
	// Now we have a window we need to create a renderer so we can draw into
	// it. In this case we want to use the first graphics card that supports faster
	// drawing
	renderer = createRenderer(window)
	// automatically destroy the renderer when the program exits.
	defer renderer.Destroy()

	// Set a black i.e. RGBA (0,0,0,0) background colour and clear the window
	renderer.SetDrawColor(0, 0, 0, 0)
	renderer.Clear()
	// ---- This is the end of Owen's graphics setup code ----

	// initialise the games variables.
	initialise()

	// Your Turn
	// 		You need to call the  main game loop function of the game here.
	// Put your answer on the next line

}

// Initialise sets the inital values of the game state variables.
// Initialise must be called before the games main loop starts.
func initialise() {
	// initially set the quit flag to false.
	quit = false
}

// GameMainLoop controls the game. It performs three main tasks. The first task
// is to get the users input. The second task is to update the games state based
// on the user input and the rules of the game. The final task is to update, or
// draw or render, the changes to the screen. Each of these tasks is also
// a separate function.
// The function also has to repeatedly do these three tasks while
// the value of the quit flag is false.
//
// Your Turn
// 		Write a main game loop funtion according to the instructions above
// Start your answer on the next line

// GetInput gets the users input and updates the game state variables that realte
// to the users input, for example, the direction that the user wants to move their
// bat in.
//
// Your Turn
//		See if you can write the getInput function
//		getInput has to check if the user has clicked the close button.
//		You can do this be calling the checkQuit function. Don't worry about how
//		chckQuit works. We will look at this later!
//		For now all you need to do is call checkQuit from getInput()
//	Start your answer on the next line

// UpdateGameState updates the game state variables based on the user input and
// the rules of the game.
//
// Your Turn
//		See if you can write the updateStateFunction
//		For now, lets just get updateState to print the value of the quit variable.
// Start your answer on the next line

// Render updates the screen, based on the new positions of the bats and the ball.
//
// Your Turn
//		See if you can write the render function
//		render needs to call the Present function. You can do this like this:
//			renderer.Present()
//		This will display the window, with a black background
// Start your answer on the next line

// CheckQuit checks if the user has clicked the window's close button.
// If the user has then the quit variable is set it true.
func checkQuit() {
	var event sdl.Event
	event = sdl.PollEvent()

	if event != nil {
		switch event.(type) {
		case *sdl.QuitEvent:
			quit = true
		}
	}
}

// Create the graphics window using the SDl library or crash trying
func createWindow(w, h int) *sdl.Window {
	var window *sdl.Window
	var err error

	window, err = sdl.CreateWindow("Pong Game", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	return window
}

// Create the graphics renderer or crash trying
func createRenderer(w *sdl.Window) *sdl.Renderer {
	var r *sdl.Renderer
	var err error
	r, err = sdl.CreateRenderer(w, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	return r
}
