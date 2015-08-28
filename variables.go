package main

type Player struct {
	Name   string
	Field  Field
	Points int
	Combo  int
}

type Cell struct {
	X int
	Y int
}

type Strategy struct {
	BurnK   int
	StepK   int
	DamageK int
	PostyK  int
}
