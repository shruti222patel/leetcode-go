package problem0001

import(
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums           []int
	target         int
	expectedResult []int
}

var cases = []Case{
	Case{nums: []int{1, 2, 3, 4, 5},
		target:         8,
		expectedResult: []int{2, 4},
	},
	Case{nums: []int{1, 2, 3, 4, 5},
		target:         10,
		expectedResult: []int{},
	},
	Case{nums: []int{3, 2, 4},
		target:         11,
		expectedResult: []int{},
	},
	Case{nums: []int{3, 2, 4},
	target:         6,
	expectedResult: []int{1,2},
	},
	Case{nums: []int{3, 3},
	target:         6,
	expectedResult: []int{0, 1},
	},
}

func TestTwoSumBruteForceOptimzed(t *testing.T) {
	for _, cas := range cases {
		result := twoSumBruteForceOptimzed(cas.nums, cas.target)
		check(t, result, cas)
	}
}


func TestTwoSumHashMap(t *testing.T) {
	for _, cas := range cases {
		result := twoSumHashMap(cas.nums, cas.target)
		check(t, result, cas)
	}
}


func check(t *testing.T, result []int, cas Case) {
	assert.Equal(t, len(result), len(cas.expectedResult))
	if len(result) > 0 {
		assert.Equal(t, result[0], cas.expectedResult[0])
		assert.Equal(t, result[1], cas.expectedResult[1])
	}
}