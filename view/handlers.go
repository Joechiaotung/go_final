package view

import (
	"../model"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var _board_id = 0

/*
var Params = struct {
	Title         string
	Width, Height *int
	RunId         int64
	ShowFreezeBtn bool
	Board_id      int
}{AppTitle, &ViewWidth, &ViewHeight, time.Now().Unix(), false, 1}
*/
// Template of the play html page
var playTempl = template.Must(template.New("t").Parse(play_html))

// The client's (browser's) view position
var Pos image.Point

// init registers the http handlers.
func init() {
	model.InitAll()
	http.HandleFunc("/", playHtmlHandle)
	// handle_new_id()
	// handle_new_id(1)
	http.HandleFunc("/runid", runIdHandle)
	http.HandleFunc("/img", imgHandle)
	http.HandleFunc("/clicked", clickedHandle)
	http.HandleFunc("/new", newGameHandle)
	// http.HandleFunc("/help", helpHtmlHandle)
}
func handle_new_id(board_id int) {
	http.HandleFunc("/table/"+strconv.Itoa(board_id), playHtmlHandle)
}

// InitNew initializes a new game.
func InitNew() {

}

type Param struct {
	Title         string
	Width, Height *int
	RunId         int64
	ShowFreezeBtn bool
	Board_id      int
}

var Params [1000]Param

func init_params() {
	for i := range Params {
		Params[i] = Param{AppTitle, &ViewWidth, &ViewHeight, time.Now().Unix(), false, 1}
	}

}

// playHtmlHandle serves the html page where the user can play.
func playHtmlHandle(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Path, "/")
	fmt.Println(r.URL.Path)
	if len(s) == 2 {
		port := 3000
		url := fmt.Sprintf("http://localhost:%d/table/%d", port, _board_id)
		/*
			if err := open(url); err != nil {
				fmt.Println("Auto-open failed:", err)
				fmt.Printf("Open %s in your browser.\n", url)
			}
		*/
		http.Redirect(w, r, url, http.StatusSeeOther)
		_board_id += 1
	} else {
		board_id, _ := strconv.Atoi(s[len(s)-1])
		Params[board_id].Board_id = board_id

		playTempl.Execute(w, Params[board_id])
	}

}

// runidHandle serves the running app id which changes if app is restarted
// (so browser clients can detect if app was restarted).
func runIdHandle(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "runid: %d", Params.RunId)
}

// imgHandle serves images of the player's view.

var quality int // is this right?
func imgHandle(w http.ResponseWriter, r *http.Request) {
	board_id, _ := strconv.Atoi(r.FormValue("board_id"))
	fmt.Println("Board id:", board_id)
	quality = 100

	rect := image.Rect(0, 0, ViewWidth, ViewHeight).Add(image.Pt(10, 10))
	model.Mutex.Lock()
	jpeg.Encode(w, model.BoardImgs[board_id].SubImage(rect), &jpeg.Options{quality})
	model.Mutex.Unlock()

}

// clickedHandle receives mouse click (mouse button pressed) events with mouse coordinates.
func clickedHandle(w http.ResponseWriter, r *http.Request) {
	board_id, err := strconv.Atoi(r.FormValue("board_id"))
	fmt.Println("Board id:", board_id)
	x, err := strconv.Atoi(r.FormValue("x"))
	if err != nil {
		return
	}

	y, err := strconv.Atoi(r.FormValue("y"))
	if err != nil {
		return
	}

	btn, err := strconv.Atoi(r.FormValue("b"))
	if err != nil {
		return
	}

	// x, y are in the coordinate system of the client's view.
	// Translate them to the Labyrinth's coordinate system:
	select {
	case model.ClickCh <- model.Click{Pos.X + x, Pos.Y + y, btn, board_id}:
	default:
	}
}

// // newGameHandle signals to start a newgame.
func newGameHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new game handle")
	// Use non-blocking send
	select {
	case model.NewGameCh <- 1:
	default:
	}
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)

	return exec.Command(cmd, args...).Start()
}
