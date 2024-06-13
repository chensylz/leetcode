package array_test

import (
	"slices"
	"testing"

	"github.com/chensylz/leetcode/category/array"
	"github.com/stretchr/testify/suite"
)

type MinimumSubArrayData struct {
	Nums   []int
	Target int
	Result int
}

type MinimumSubArrayTestSuite struct {
	suite.Suite

	data []MinimumSubArrayData
}

func (s *MinimumSubArrayTestSuite) SetupSuite() {
	s.data = []MinimumSubArrayData{
		{
			Nums:   []int{2, 3, 1, 2, 4, 3},
			Target: 7,
			Result: 2,
		},
		{
			Nums:   []int{1, 4, 4},
			Target: 4,
			Result: 1,
		},
		{
			Nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			Target: 11,
			Result: 0,
		},
		{
			Nums:   []int{1, 2, 3, 4, 5},
			Target: 11,
			Result: 3,
		},
	}
}

func TestMinimumSubArrayTestSuite(t *testing.T) {
	suite.Run(t, new(MinimumSubArrayTestSuite))
}

func (s *MinimumSubArrayTestSuite) TestNormal() {
	for _, d := range s.data {
		nums := slices.Clone(d.Nums)
		result := array.MinSubArrayLen(d.Target, nums)
		s.Equal(d.Result, result)
	}
}

func (s *MinimumSubArrayTestSuite) TestWindow() {
	for _, d := range s.data {
		nums := slices.Clone(d.Nums)
		result := array.MinSubArrayLenWindow(d.Target, nums)
		s.Equal(d.Result, result)
	}
}
