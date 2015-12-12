// The orbits program calculates the orbital speed of the International Space
// Station, the Hubble Space Telescope, the Astra 1KR TV broadcast satellite and
// a GPS satellite.
//
// This program is incomplete. You need to finish it before it will run.
//
// Important numbers you will need in the program
//
// The radius of the Earth is 6378km
// The average altitude of the ISS is 412.5 km
// The average time of the ISS orbit is 1.525 hours
// The average altitude of the Hubble Space Telescope is 595 km
// The average time of the Hubble orbit is 1.5933 hours
// The average altitude of the Astra 1KR satellite is 35786 km
// The Astra 1KR orbit time is 24 hours
// Bonus Task
// GPS satellites orbit at 20350 km and orbit every 12 hours.

package main

import (
	"fmt"
    // The "math" package is new. We need to use it because it contains a very
    // accurate value for Pi
	"math"
)

func main() {
    // First the variables for the Earths radius and time of the ISS orbit.
    // These are floating point types to keeps things accruate - we want to
    // know what the numbers are after the decimal point. This also avoids
    // having to do any conversions from integer to floating point.
	var earthsRadiusInKm float64
	var timeOfIssOrbitInHrs float64

    // Your turn
    // The ISS altitude is 412.5 km. What type of variable do you need?
	var issAverageAltitudeInKm

    // Set the value for the Earths radius
	earthsRadiusInKm = 6378.0
    // Set the value for the ISS orbit time in hours
	timeOfIssOrbitInHrs = 1.545
    // Your Turn
    // Now set the value of the ISS's average altitude - see the important
    // numbers at the top of the rpogram
	issAverageAltitudeInKm =

    // Before you can work out the speed you need to work out the total
    // distance of the orbit of the ISS. We need a variable to store the
    // answer is.
	var issOrbitalDistanceInKm float64
    // If we assume the orbit is a circle then the circumfrence of the circle is
    // 2 * pi * the circle radius
    // Remember that the radius of the orbit is the earths radius plus the altitude,
    // so we have to work this out.
    // The "math.Pi" bit is just how you access the value of Pi in the math package
	issOrbitalDistanceInKm = 2 * math.Pi * (earthsRadiusInKm + issAverageAltitudeInKm)

    // We need another variable to store the speed of the ISS orbit in km/h
	var issSpeedInKmPerHr float64
    // Now we can work out the speed in km per hour. The speed is just the
    // distance of the orbit divided by the time the orbit takes.
	issSpeedInKmPerHr = issOrbitalDistanceInKm / timeOfIssOrbitInHrs

    // Now we can print the results!
	fmt.Print("The distance of one orbit of the ISS is ")
	fmt.Print(issOrbitalDistanceInKm)
	fmt.Println("km")
	fmt.Print("The speed of the ISS in orbit is ")
	fmt.Print(issSpeedInKmPerHr)
	fmt.Println("km/h")

    // Your Turn
    // Now you know how to calcualte the orbital speed of the ISS
    // see if you can calculate the oribital speed of the Hubble Space
    // Telescope.
    // Hint: You need new variables for the altitude, orbital distance and
    // orbital speed of the Hubble Space Telescope. But the maths is the same,
    // you just need to change the variable names.
    // The numbers you need are at the top of the program.
    // ....Add your new code below this line....

    // Your Turn Again
    // Now do the same for the Astra 1KR satellite to work out its speed.
    // ....Add your new code below this line.....

    // Bonus Points
    // See if you can work out the orbital speed of a GPS satellite.
    // ....Add your new code below this line....
}
