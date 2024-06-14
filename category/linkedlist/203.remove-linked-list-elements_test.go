package linkedlist_test

import (
	"github.com/chensylz/leetcode/category/linkedlist"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RemoveElementsData struct {
	Input  *linkedlist.ListNode
	Val    int
	Result []int
}

type RemoveElementsTestSuite struct {
	suite.Suite

	data []RemoveElementsData
}

func (s *RemoveElementsTestSuite) SetupSuite() {
	s.data = []RemoveElementsData{
		{
			Input:  linkedlist.NewListWithArray([]int{1, 2, 6, 3, 4, 5, 6}),
			Val:    6,
			Result: []int{1, 2, 3, 4, 5},
		},
		{
			Input:  linkedlist.NewListWithArray([]int{}),
			Val:    1,
			Result: []int{},
		},
		{
			Input:  linkedlist.NewListWithArray([]int{7, 7, 7, 7}),
			Val:    7,
			Result: []int{},
		},
	}
}

func TestRemoveElementsTestSuite(t *testing.T) {
	suite.Run(t, new(RemoveElementsTestSuite))
}

func (s *RemoveElementsTestSuite) TestNormal() {
	for _, d := range s.data {
		s.Equal(
			linkedlist.ExtendArrayToStr(d.Result),
			linkedlist.ExtendLinkedToArrayStr(linkedlist.RemoveElements(d.Input, d.Val)),
		)
	}
}
