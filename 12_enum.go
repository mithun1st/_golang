package main

import (
	"fmt"
)

// * Enum with single data
type LevelE int

const (
	Low LevelE = iota + 101
	Mid
	High
	Extreme LevelE = 4
	Max     LevelE = 5
)

// * Enum with multiple data
type RgbE string

const (
	R RgbE = "#FF0000"
	G RgbE = "#00FF00"
	B RgbE = "#0000FF"
)

func (rgb RgbE) title() string {
	switch rgb {
	case R:
		return "Red"
	case G:
		return "Green"
	case B:
		return "Blue"
	}
	return ""
}

func (rgb RgbE) chh() rune {
	switch rgb {
	case R:
		return 'r'
	case G:
		return 'g'
	case B:
		return 'b'
	}
	return ' '
}

func main() {

	fmt.Println(Low, Mid, High, Extreme, Max)

	fmt.Println(R, G, B)
	fmt.Println(R.title(), G.title(), B.title())
	fmt.Println(R.chh(), G.chh(), B.chh())

	/*
		101 102 103 4 5
		#FF0000 #00FF00 #0000FF
		Red Green Blue
		114 103 98
	*/
}
