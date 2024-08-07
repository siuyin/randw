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

func Select(n int) []string {
	perm := rand.Perm(len(Words))
	s := []string{}
	for i := 0; i < n; i++ {
		s = append(s, Words[perm[i]])
	}
	return s
}
