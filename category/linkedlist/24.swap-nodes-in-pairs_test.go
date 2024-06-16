package linkedlist_test

import (
	"github.com/chensylz/leetcode/category/linkedlist"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SwapPairsData struct {
	Input  *linkedlist.ListNode
	Result []int
}

type SwapPairsTestSuite struct {
	suite.Suite

	data []SwapPairsData
}

func (s *SwapPairsTestSuite) SetupSuite() {
	s.data = []SwapPairsData{
		{
			Input:  linkedlist.NewListWithArray([]int{1, 2, 3, 4}),
			Result: []int{2, 1, 4, 3},
		},
		{
			Input:  linkedlist.NewListWithArray([]int{}),
			Result: []int{},
		},
		{
			Input:  linkedlist.NewListWithArray([]int{1}),
			Result: []int{1},
		},
	}
}

func TestSwapPairsTestSuite(t *testing.T) {
	suite.Run(t, new(SwapPairsTestSuite))
}

func (s *SwapPairsTestSuite) TestNormal() {
	for _, d := range s.data {
		s.Equal(
			linkedlist.ExtendArrayToStr(d.Result),
			linkedlist.ExtendLinkedToArrayStr(linkedlist.SwapPairs(d.Input)),
		)
	}
}
