package main

var winningMoves = map[string]bool{
	"a1:a2:a3": true,
	"b1:b2:b3": true,
	"c1:c2:c3": true,

	"a1:b1:c1": true,
	"a2:b2:c2": true,
	"a3:b3:c3": true,

	"a1:b2:c3": true,
	"a3:b2:c1": true,
}
