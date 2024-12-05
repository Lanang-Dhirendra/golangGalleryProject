package model

import (
	"database/sql"
	"math"
)

type Images struct {
	ID      string
	Name    string
	FileExt string
	Owner   string
	Score   int
	Tags    string

	CreatedAt int64
	UpdatedAt int64
	DeletedAt sql.NullInt64
}

const DbAddr string = "root:@tcp(127.0.0.1:3306)"
const Database string = "gallery"
const Table string = "images"
const ConnTimeout = 5 // seconds
const MaxGet = 200

var AcceptedExt []string = []string{
	".jpg", ".jpeg", ".png", ".svg", ".webp", ".gif",
}

var symbols []string = []string{
	"abcdefghijklmnopqrstuvwxyz",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"0123456789",
	" ",
	"_",
}

func Syms(val int) string {
	syms := ""
	if val < 1 || val > int(math.Pow(2, float64(len(symbols)))-1) {
		return ""
	}
	for i := range symbols {
		revIdx := len(symbols) - i
		bin := int(math.Pow(2, float64(revIdx-1)))
		if val-bin >= 0 {
			val = val - bin
			syms += symbols[revIdx-1]
		}
	}
	return syms
}
