package hash_test

import (
	"testing"

	"github.com/chensylz/leetcode/category/hash"

	"github.com/stretchr/testify/suite"
)

type IsAnagramData struct {
	s  string
	t  string
	ok bool
}

type IsAnagramTestSuite struct {
	suite.Suite

	data []IsAnagramData
}

func (s *IsAnagramTestSuite) SetupSuite() {
	s.data = []IsAnagramData{
		{
			s:  "anagram",
			t:  "nagaram",
			ok: true,
		},
		{
			s:  "rat",
			t:  "car",
			ok: false,
		},
	}
}

func TestIsAnagramTestSuite(t *testing.T) {
	suite.Run(t, new(IsAnagramTestSuite))
}

func (s *IsAnagramTestSuite) TestNormal() {
	for _, d := range s.data {
		s.Equal(d.ok, hash.IsAnagram(d.s, d.t))
	}
}

func (s *IsAnagramTestSuite) TestByArray() {
	for _, d := range s.data {
		s.Equal(d.ok, hash.IsAnagramByArray(d.s, d.t))
	}
}
