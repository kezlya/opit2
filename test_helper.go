package main

import (
	"fmt"
)

func PrintVisual(f Field) {
	fmt.Println()
	for i := range f.Grid {
		fmt.Print(f.Height-i, "	")
		for _, c := range f.Grid[f.Height-i] {
			if c {
				fmt.Print("⬛ ")
			} else {
				fmt.Print("⬜ ")
			}
		}
		fmt.Println()
	}
	fmt.Println("	 0 1 2 3 4 5 6 7 8 9")
}

func PrintVisuals(a, b Field) {
	fmt.Println()
	for i := range a.Grid {
		fmt.Print(a.Height-i, "	")
		for _, c := range a.Grid[a.Height-i] {
			if c {
				fmt.Print("⬛ ")
			} else {
				fmt.Print("⬜ ")
			}
		}
		fmt.Print("   ")
		for _, c := range b.Grid[b.Height-i] {
			if c {
				fmt.Print("⬛ ")
			} else {
				fmt.Print("⬜ ")
			}
		}
		fmt.Println()
	}
	fmt.Println("	 0 1 2 3 4 5 6 7 8 9    0 1 2 3 4 5 6 7 8 9")
}

func (a Picks) isEqual(b Picks) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
