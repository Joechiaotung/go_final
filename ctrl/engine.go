package ctrl

import (
	"../model"
	"../view"
	// "image"
	// "image/draw"
	// "math"
	"math/rand"
	"time"
	"fmt"
)

// LoopDelay is the delay between the iterations of the main loop of the game engine, in milliseconds.
var LoopDelay = 20 // ~20 FPS

// InitNew initializes a new game.
func initNew() {
	// Initialize random number generator
	rand.Seed(time.Now().Unix())

	fmt.Println("Trying to init new")
	model.InitNew()
	view.InitNew()
}

// StartEngine starts the game engine in a new goroutine and returns as soon as possible.
func StartEngine() {
	model.NewGameCh <- 1 // Cannot block as application was just started, no incoming requests processed yet

	// initNew()
	model.Mutex.Lock()

	go simulate()
}

// Delta time since our last iteration
var dt float64

// simulate implements the game cycle
func simulate() {


	for {
		// Check if we have to start a new game
		select {
		case <-model.NewGameCh:
			initNew()
		default:
		}

		// Process mouse clicks
	clickLoop:
		for {
			select {
			case click := <-model.ClickCh:
				handleClick(click)
			default:
				break clickLoop
			}
		}

		// now := time.Now().UnixNano()
		// dt = float64(now-t) / 1e9
		// t = now

		// Sleep some time.
		// Iterations might not be exact, but we don't rely on it:
		// We calculate delta time and calculate moving and next positions
		// based on the delta time.

		model.Mutex.Unlock() // While sleeping, clients can request view images
		if model.Won {
			// If won, nothing has to be done, just wait for a new game signal
			<-model.NewGameCh // Blocking receive
			// Send back value to detect it at the proper place
			model.NewGameCh <- 1
		}
		time.Sleep(time.Millisecond * time.Duration(LoopDelay))
		model.Mutex.Lock() // We will modify model now, labyrinth image might change so lock.
	}
}

// handleClick handles a mouse click
func handleClick(c model.Click) {

	col, row := c.X/model.BlockSize, c.Y/model.BlockSize

	// AI
	model.PlayerTurn = true
	if model.PlayerTurn == true {
		if model.Board[row][col] != model.StoneWhite &&  model.Board[row][col] != model.StoneBlack{
			model.Board[row][col] = model.StoneWhite

			aiMove := NextMove()
			fmt.Println(aiMove)
			model.Board[aiMove.X][aiMove.Y] = model.StoneBlack

			model.PlayerTurn = !model.PlayerTurn

			model.DrawColaRow(col, row)
		}
	} else {
		// aiMove := NextMove()
		// fmt.Println(aiMove)
		// model.Board[aiMove.X][aiMove.Y] = model.StoneBlack
		// model.Board[row][col] = model.StoneBlack
	}


	// model.Board[aiMove.X][aiMove.Y] = model.StoneBlack

	// model.PlayerTurn = false

}

// handleWinning handles the winning of game event.
func handleWinning() {
	
}

func init() {

}