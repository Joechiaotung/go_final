package model

const (
	// BlockSize is the size of the board unit in pixels.
	BlockSize = 30
)

var (
	// Rows is the number of rows in the Labyrinth
	Rows int
	// Cols is the number of columns in the Labyrinth
	Cols int

	// LabWidth is the width of the labyrinth's image in pixels.
	BoardWidth int

	// LabHeight is the height of the labyrinth's image in pixels.
	BoardHeight int
)

// V is the moving speed of Gopher and the Buddlogs in pixel/sec.
var V float64

// Type of the unit of the labyrinth
type Stone int

// Pieces
const (
	StoneEmpty Stone = iota
	StoneWhite
	StoneBlack
)
