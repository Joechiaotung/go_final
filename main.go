//https://golang-lab5.herokuapp.com/operation/num1/num2

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w,"hello world!")

	pathParts := strings.SplitN(r.URL.Path, "/", -1)

	operation := pathParts[1]
	var a, b int
	if len(pathParts) == 4 || len(pathParts) == 5 {
		a, _ = strconv.Atoi(pathParts[2])
		b, _ = strconv.Atoi(pathParts[3])
		if len(pathParts) == 5 {
			c := pathParts[4]
			if c != "" {
				fmt.Fprintf(w, "hello world!")
				return
			}
		}
	} else {
		fmt.Fprintf(w, "hello world!")
		return
	}

	if operation == "add" {
		//fmt.Printf("%d%d",a,b)
		fmt.Fprintf(w, "%v + %v = %v", a, b, a+b)
	} else if operation == "sub" {
		fmt.Fprintf(w, "%v - %v = %v", a, b, a-b)
	} else if operation == "mul" {
		fmt.Fprintf(w, "%v * %v = %v", a, b, a*b)
	} else if operation == "div" {
		fmt.Fprintf(w, "%v / %v = %v, reminder = %v", a, b, a/b, a%b)
	} else {
		fmt.Fprintf(w, "hello world!")
	}
}

/*func add(a, b int) int {

}

func div(a, b int) int int {

}*/

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
