package main

type Player struct {
	Name   string
	Field  Field
	Points int
	Combo  int
}

type Hole struct {
	X int
	Y int
}
