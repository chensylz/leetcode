package array_test

import (
	"github.com/chensylz/leetcode/category/array"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SquaresOfASortedArrayData struct {
	Nums   []int
	NewArr []int
}

type SquaresOfASortedArrayTestSuite struct {
	suite.Suite

	data []SquaresOfASortedArrayData
}

func (s *SquaresOfASortedArrayTestSuite) SetupSuite() {
	s.data = []SquaresOfASortedArrayData{
		{
			Nums:   []int{-4, -1, 0, 3, 10},
			NewArr: []int{0, 1, 9, 16, 100},
		},
		{
			Nums:   []int{-7, -3, 2, 3, 11},
			NewArr: []int{4, 9, 9, 49, 121},
		},
	}
}

func TestSquaresOfASortedArrayTestSuite(t *testing.T) {
	suite.Run(t, new(SquaresOfASortedArrayTestSuite))
}

func (s *SquaresOfASortedArrayTestSuite) TestDoublePointer() {
	for _, data := range s.data {
		result := array.SortedSquares(data.Nums)
		for i := 0; i < len(result); i++ {
			s.Equal(data.NewArr[i], result[i])
		}
	}
}
