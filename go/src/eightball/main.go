package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
)

// enter `go get` from cmd line to install the colors mod

var answers = []string{
	"Yes.",
	"No.",
	"Maybe.",
}
var responces = []easterEgg{
	{"die|death", "Death is quite the grim topic."},
	{"love", "Do I look like cupid?"},
}

type easterEgg struct {
	key      string
	responce string
}

func main() {
	fmt.Println(c.Clear + c.Multi("Welcome to the Magic Eight Ball!"))
}
