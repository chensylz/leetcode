package linkedlist_test

import (
	"github.com/chensylz/leetcode/category/linkedlist"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ReverseLinkedListData struct {
	Input  *linkedlist.ListNode
	Result []int
}

type ReverseLinkedListTestSuite struct {
	suite.Suite

	data []ReverseLinkedListData
}

func (s *ReverseLinkedListTestSuite) SetupSuite() {
	s.data = []ReverseLinkedListData{
		{
			Input:  linkedlist.NewListWithArray([]int{1, 2, 3, 4, 5}),
			Result: []int{5, 4, 3, 2, 1},
		},
		{
			Input:  linkedlist.NewListWithArray([]int{}),
			Result: []int{},
		},
		{
			Input:  linkedlist.NewListWithArray([]int{1, 2}),
			Result: []int{2, 1},
		},
	}
}

func TestReverseLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(ReverseLinkedListTestSuite))
}

func (s *ReverseLinkedListTestSuite) TestNormal() {
	for _, data := range s.data {
		s.Equal(linkedlist.ExtendArrayToStr(data.Result),
			linkedlist.ExtendLinkedToArrayStr(linkedlist.ReverseList(data.Input)))
	}
}
