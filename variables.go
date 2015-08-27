package main

type Player struct {
	Name   string
	Field  Field
	Points int
	Combo  int
}

type Cell struct {
	X int8
	Y int8
}

type Strategy struct {
	BurnK   int
	StepK   int
	DamageK int
	PostyK  int
}
