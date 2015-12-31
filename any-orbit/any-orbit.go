// The any-orbits program calculates the orbital
// speed of any satellite in orbit around Earth.
//
// This program is incomplete. You need to finish it before it will run.
//
// You'll need to know the radius of the Earth is 6378km to complete the program.
//
// Try the program out with these numbers
//
// The average altitude of the ISS is 412.5 km
// The average time of the ISS orbit is 1.525 hours
// The average altitude of the Hubble Space Telescope is 595 km
// The average time of the Hubble orbit is 1.5933 hours
// The average altitude of the Astra 1KR satellite is 35786 km
// The Astra 1KR orbit time is 24 hours

package main

import (
	"fmt"
	// The "math" package is new. We need to use it because it contains a very
	// accurate value for Pi
	"math"

	// Your Turn
	// You need to use the import part of the input pattern
)

func main() {
	// These are the variables
	var earthsRadiusInKm float64 // Earths radius
	// Your Turn
	// What type should timeOfOrbitInHrs be?
	var timeOfOrbitInHrs        // the time of the satellites orbit in hours
	var altitudeOfOrbit float64 // the altitude of the satellite above Earth
	// Your turn
	// What type should nameOfSatellite be?
	var nameOfSatellite

	// Start by telling the user what the program does
	fmt.Println("The any-orbit program computes the orbital speed of")
	fmt.Println("any satellite, using just the satellites altitude in km")
	fmt.Println("and the satellites time of orbit in hours.")

	// Now we need to ask the user for the first bit of information
	// the satellites name
	fmt.Println("Please enter the name of the satellite: ")
	// Your turn
	// What do you need to do here to read the users answer and store it
	// in the variable nameOfSatellite?
	nameOfSatellite = simpleio.

		// Now we need to ask the user for the altitude of the satellite
		fmt.Println("Please enter the altitude of the satellite in km: ")
	// Your turn
	// What do you need to do here to read in the users answer
	// and store it in the variable altitudeOfOrbit?
	altitudeOfOrbit =

		// Your turn!
		// What do you need to do to ask the user for the time of the satellites
		// orbit in hours
		fmt.Println("Please enter the time of the satellites orbit in hours: ")
	// Your turn
	// What do you need to do to read in the users answer and store it
	// in the varible timeOfOrbitInHrs

	// Set the value for the Earths radius
	earthsRadiusInKm = 6378.0

	// this is the variable for the distance of the orbitof the satellite
	var orbitalDistanceInKm float64
	// work out the distance of the orbit using the formula
	// distance = 2 * pi * radius
	orbitalDistanceInKm = 2 * math.Pi * (earthsRadiusInKm + altitudeOfOrbit)

	// ths is the variable for the speed of the satellite
	var speedInKmPerHr float64

	// Now we can work out the speed in km per hour. The speed is just the
	// distance of the orbit divided by the time the orbit takes.
	speedInKmPerHr = orbitalDistanceInKm / timeOfOrbitInHrs

	// Now we can print the results!
	fmt.Print("The distance of one orbit of ")
	// Your Turn
	// What do you need to do here to print out the satellites name?

	fmt.Print(" is ")
	// Your Turn
	// What do you need to do here to print out the distance of the
	// satellites orbit?

	fmt.Println("km")
	fmt.Print("The speed of ")
	// Your Turn
	// What do you need to do here to print out the distance of the
	// satellites orbit? (Yes, you need this line twice)

	fmt.Print(" in orbit is ")
	// Your Turn
	// What do you need to do here to print out the speed of the
	// satellite in orbit?

	fmt.Println("km/h")

	// Your Turn
	// Change the questions that the program asks for both the time of
	// the oribit and the altitude of the satellite so that the questions
	// use the name of the satellite that the user types in instead of
	// saying "satellite"

	// Bonus Points
	// Change the program again so that it asks the user for the
	// orbital time in minutes. Then convert the users answer into hours and
	// use that in the rest of the program.
}
