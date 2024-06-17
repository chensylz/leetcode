package linkedlist_test

import (
	"github.com/chensylz/leetcode/category/linkedlist"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RemoveNthFromEndData struct {
	Input  *linkedlist.ListNode
	N      int
	Result []int
}

type RemoveNthFromEndTestSuite struct {
	suite.Suite

	data []RemoveNthFromEndData
}

func (s *RemoveNthFromEndTestSuite) SetupSuite() {
	s.data = []RemoveNthFromEndData{
		{
			Input:  linkedlist.NewListWithArray([]int{1, 2, 3, 4, 5}),
			Result: []int{1, 2, 3, 5},
			N:      2,
		},
		{
			Input:  linkedlist.NewListWithArray([]int{1}),
			Result: []int{},
			N:      1,
		},
		{
			Input:  linkedlist.NewListWithArray([]int{1, 2}),
			Result: []int{1},
			N:      1,
		},
	}
}

func TestRemoveNthFromEndTestSuite(t *testing.T) {
	suite.Run(t, new(RemoveNthFromEndTestSuite))
}

func (s *RemoveNthFromEndTestSuite) TestNormal() {
	for _, data := range s.data {
		s.Equal(linkedlist.ExtendArrayToStr(data.Result),
			linkedlist.ExtendLinkedToArrayStr(linkedlist.RemoveNthFromEnd(data.Input, data.N)))
	}
}
