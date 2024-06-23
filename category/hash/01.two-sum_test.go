package hash_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/chensylz/leetcode/category/common"
	"github.com/chensylz/leetcode/category/hash"
)

type TwoSumData struct {
	nums   []int
	target int
	result []int
}

type TwoSumTestSuite struct {
	suite.Suite

	data []TwoSumData
}

func (s *TwoSumTestSuite) SetupSuite() {
	s.data = []TwoSumData{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			result: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			result: []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			result: []int{0, 1},
		},
	}
}

func TestTwoSumTestSuite(t *testing.T) {
	suite.Run(t, new(TwoSumTestSuite))
}

func (s *TwoSumTestSuite) TestNormal() {
	for _, data := range s.data {
		result := hash.TwoSum(data.nums, data.target)
		s.True(common.EqualSliceWithoutIndex(data.result, result))
	}
}
