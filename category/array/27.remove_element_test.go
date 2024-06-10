package array_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/chensylz/leetcode/category/array"
)

type RemoveElementData struct {
	Nums   []int
	Target int
	K      int
}

type RemoveElementTestSuite struct {
	suite.Suite

	data []RemoveElementData
}

func (s *RemoveElementTestSuite) SetupSuite() {
	s.data = []RemoveElementData{
		{
			Nums:   []int{3, 2, 2, 3},
			Target: 3,
			K:      2,
		},
		{
			Nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			Target: 2,
			K:      5,
		},
	}
}

func TestRemoveElementTestSuite(t *testing.T) {
	suite.Run(t, new(RemoveElementTestSuite))
}

func (s *RemoveElementTestSuite) TestNormal() {
	for _, d := range s.data {
		nums := slices.Clone(d.Nums)
		k := array.RemoveElement(nums, d.Target)
		s.Equal(d.K, k)
	}
}

func (s *RemoveElementTestSuite) TestDoublePointer() {
	for _, d := range s.data {
		nums := slices.Clone(d.Nums)
		k := array.RemoveElementDoublePointer(nums, d.Target)
		s.Equal(d.K, k)
	}
}
