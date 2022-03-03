package models

type Board struct {
	NumColor int
	NumLen   int
	Matrix   [20][20]byte
}

var Letters = []byte{'B', 'G', 'P', 'T', 'Y', 'L', 'M', 'R', 'A'}

type Vector struct {
	Vec []int
}
