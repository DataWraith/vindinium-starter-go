package vindinium

// Direction represents one of the possible movement directions in Vindinium
type Direction string

const (
	North Direction = "North"
	East  Direction = "East"
	South Direction = "South"
	West  Direction = "West"
	Stay  Direction = "Stay"
)

// NESW is an array containing the cardinal directions for convenient use with
// a range statement.
var NESW = [4]Direction{North, East, South, West}
