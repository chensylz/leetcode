package array_test

import (
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/chensylz/leetcode/category/array"
)

type Data struct {
	Nums   []int
	Target int
	Result int
}

type BinarySearchTestSuite struct {
	suite.Suite

	data []Data
}

func (s *BinarySearchTestSuite) SetupSuite() {
	s.data = []Data{
		{
			Nums:   []int{-1, 0, 3, 5, 9, 12},
			Target: 9,
			Result: 4,
		},
		{
			Nums:   []int{-1, 0, 3, 5, 9, 12},
			Target: 2,
			Result: -1,
		},
	}
}

func TestBinarySearchTestSuite(t *testing.T) {
	suite.Run(t, new(BinarySearchTestSuite))
}

func (s *BinarySearchTestSuite) TestNormalSearch() {
	for _, ts := range s.data {
		res := array.Search(ts.Nums, ts.Target)
		s.Equal(ts.Result, res)
	}
}

func (s *BinarySearchTestSuite) TestSearchFirst() {
	data := []Data{
		{
			Nums:   []int{-1, 0, 0, 0, 9, 12},
			Target: 0,
			Result: 1,
		},
		{
			Nums:   []int{-1, 1, 9, 9, 9, 12},
			Target: 9,
			Result: 2,
		},
	}
	for _, ts := range data {
		res := array.SearchFirst(ts.Nums, ts.Target)
		s.Equal(ts.Result, res)
	}
}

func (s *BinarySearchTestSuite) TestSearchLast() {
	data := []Data{
		{
			Nums:   []int{-1, 0, 0, 0, 9, 12},
			Target: 0,
			Result: 3,
		},
		{
			Nums:   []int{-1, 1, 9, 9, 9, 12},
			Target: 9,
			Result: 4,
		},
	}
	for _, ts := range data {
		res := array.SearchLast(ts.Nums, ts.Target)
		s.Equal(ts.Result, res)
	}
}
