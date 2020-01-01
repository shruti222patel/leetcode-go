package scrapper

import(
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	HTMLFile string
}

var cases = []Case{
	Case{
		HTMLFile: "test.html",
		LeetCodeProblem: LeetCodeProblem{
			Name:          "Add Two Numbers",
			Number: "0002",
			Description:        `You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

			You may assume the two numbers do not contain any leading zero, except the number 0 itself.`,
			Example:       `Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
			Output: 7 -> 0 -> 8
			Explanation: 342 + 465 = 807.`,
			Difficulty:    "Medium",
			Url:           "https://leetcode.com/problems/add-two-numbers/",
			RelatedTopics: "LinkedList, Math",
		},
	},
}

func Test_LeetCodeScraper(t *testing.T) {
	for _, cas := range cases {
		result := FuncName(?)
		check(t, result, cas)
	}
}

