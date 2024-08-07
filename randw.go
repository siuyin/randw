package randw

import (
	_ "embed"
	"math/rand"
	"strings"
)

//go:embed words.txt
var b []byte

var Words []string

func init() {
	w := strings.Split(string(b), "\n")
	Words = w[0 : len(w)-1]
}

func Word() string {
	return Words[rand.Intn(len(Words))]
}
