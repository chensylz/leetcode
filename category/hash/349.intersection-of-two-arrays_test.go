package hash_test

import (
	"testing"

	"github.com/chensylz/leetcode/category/hash"
	"github.com/stretchr/testify/suite"
)

type IntersectionData struct {
	nums1  []int
	nums2  []int
	result []int
}

type IntersectionTestSuite struct {
	suite.Suite

	data []IntersectionData
}

func (s *IntersectionTestSuite) SetupSuite() {
	s.data = []IntersectionData{
		{
			nums1:  []int{1, 2, 2, 1},
			nums2:  []int{2, 2},
			result: []int{2},
		},
		{
			nums1:  []int{4, 9, 5},
			nums2:  []int{9, 4, 9, 8, 4},
			result: []int{9, 4},
		},
	}
}

func TestIntersectionTestSuite(t *testing.T) {
	suite.Run(t, new(IntersectionTestSuite))
}

func (s *IntersectionTestSuite) TestNormal() {
	for _, data := range s.data {
		result := hash.Intersection(data.nums1, data.nums2)
		resultMap := make(map[int]bool)
		for _, v := range result {
			resultMap[v] = true
		}
		for _, v := range data.result {
			s.Equal(resultMap[v], true)
		}
	}
}
