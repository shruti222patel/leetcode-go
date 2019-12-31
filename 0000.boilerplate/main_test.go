package problem

import(
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
}

var cases = []Case{}

func Test_FuncName(t *testing.T) {
	for _, cas := range cases {
		result := FuncName(?)
		check(t, result, cas)
	}
}

func check(t *testing.T, result ?, cas Case) {
	// Assert
}