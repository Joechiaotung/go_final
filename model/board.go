package model

import (
	"image"
	"image/color"
	"fmt"
	"image/draw"
	// "math/rand"
	"sync"
)

// Mutex to be used to synchronize model modifications
var Mutex sync.Mutex

// The model/data of the board
var Board [][]Stone

var PlayerTurn bool = true

type PiecePos struct {
	X, Y int
}

// Tells if we won
var Won bool

// Channel to signal new game
var NewGameCh = make(chan int, 1)

// Constant for the right Mouse button value in the Click struct.
// Button value for left and middle may not be the same for older browsers, but right button always has this value.
const MouseBtnRight = 2

// Click describes a mouse click.
type Click struct {
	// X, Y are the mouse coordinates in pixel, in the coordinate system of the Labyrinth
	X, Y int
	// Btn is the mouse button
	Btn int
}

// Channel to receive mouse clicks on (view package sends, ctrl package (engine) processes them)
var ClickCh = make(chan Click, 10)

// InitNew initializes a new game.
func InitNew() {
	fmt.Println("BOARD INIT NEW........");

	// backfill entire surface with grey
	BoardImg = image.NewRGBA(image.Rect(0, 0, BoardWidth+100, BoardHeight+100))
    bg := color.RGBA{50, 50, 50, 255}
    draw.Draw(BoardImg, BoardImg.Bounds(), &image.Uniform{bg}, image.ZP, draw.Src)

	initBoard()
	initBoardImg()
}

// initBoard initializes and generates a new Board.
func initBoard() {

	fmt.Println("INIT BOARD........ %v, %v", Rows, Cols);
	Board = make([][]Stone, Rows)

	for i := range Board {
		Board[i] = make([]Stone, Cols)
	}
	genBoard()
}


type circle struct {
    p image.Point
    r int
    c color.Color
}


func (c *circle) ColorModel() color.Model {
    return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
    return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
    xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
    if xx*xx+yy*yy < rr*rr {
        return c.c
    }
    return color.Alpha{0}
}


func (c *circle) Color(circleColor color.Color) color.Color {
	return color.Alpha{0}
}



// initBoardImg initializes and draws the image of the Labyrinth.
func initBoardImg() {

	emptycolor := color.RGBA{233, 169, 94, 255}
	for ri, row := range Board {
		for ci, block := range row {
			if (block == StoneEmpty) && ci != Cols-1 && ri != Rows-1 {
				x, y := ci*BlockSize+BlockSize/2+BlockSize, ri*BlockSize+BlockSize/2+BlockSize
				myimage := image.NewRGBA(image.Rect(x, y, x+BlockSize-1, y+BlockSize-1))

			    // backfill entire surface with beige
			    draw.Draw(BoardImg, myimage.Bounds(), &image.Uniform{emptycolor}, image.ZP, draw.Src)
			} 
		}
	}
}

func DrawColaRow(col int, row int) {
	blackColor := color.RGBA{0, 0, 0, 255}
	whiteColor := color.RGBA{255, 255, 255, 255}
	for ri, row := range Board {
		for ci, block := range row {
			x, y := ci*BlockSize+BlockSize, ri*BlockSize+BlockSize
			myimage := image.NewRGBA(image.Rect(x, y, x+BlockSize-3, y+BlockSize-3))

			pieceSize := BlockSize/2 - 2
			if block == StoneBlack {
			    cr := &circle{image.Point{pieceSize, pieceSize}, pieceSize, blackColor}
			    draw.DrawMask(BoardImg, myimage.Bounds(), cr, image.ZP, cr, image.ZP, draw.Over)
			} else if block == StoneWhite {
			    
			    cr := &circle{image.Point{pieceSize, pieceSize}, pieceSize, whiteColor}
				draw.DrawMask(BoardImg, myimage.Bounds(), cr, image.ZP, cr, image.ZP, draw.Over)
			} 

			// draw.Draw(BoardImg, myimage.Bounds(), &image.Uniform{pieceColor}, image.ZP, draw.Src)
		}
	}
}

// genLab generates a random labyrinth.
func genBoard() {
	// Create a "frame":
	for ri := range Board {
		Board[ri][0] = StoneEmpty
	}
	for ci := range Board[0] {
		Board[0][ci] = StoneEmpty
	}
}
