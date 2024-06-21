package hash_test

import (
	"github.com/chensylz/leetcode/category/hash"
	"github.com/stretchr/testify/suite"
	"testing"
)

type IsHappyData struct {
	n      int
	result bool
}

type IsHappyTestSuite struct {
	suite.Suite

	data []IsHappyData
}

func (s *IsHappyTestSuite) SetupSuite() {
	s.data = []IsHappyData{
		{
			n:      19,
			result: true,
		},
		{
			n:      2,
			result: false,
		},
	}
}

func TestIsHappyTestSuite(t *testing.T) {
	suite.Run(t, new(IsHappyTestSuite))
}

func (s *IsHappyTestSuite) TestNormal() {
	for _, data := range s.data {
		s.Equal(data.result, hash.IsHappy(data.n))
	}
}
