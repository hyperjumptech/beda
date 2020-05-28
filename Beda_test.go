package beda

import "testing"

type TestLehvenstein struct {
	S1 string
	S2 string
	D  int
}

func TestLevenshteinDistance(t *testing.T) {
	testData := make([]*TestLehvenstein, 0)
	testData = append(testData, &TestLehvenstein{
		S1: "abc",
		S2: "abd",
		D:  1,
	}, &TestLehvenstein{
		S1: "abc",
		S2: "abc",
		D:  0,
	}, &TestLehvenstein{
		S1: "abc",
		S2: "ade",
		D:  2,
	}, &TestLehvenstein{
		S1: "abc",
		S2: "def",
		D:  3,
	}, &TestLehvenstein{
		S1: "abc",
		S2: "abca",
		D:  1,
	}, &TestLehvenstein{
		S1: "abc",
		S2: "abcabc",
		D:  3,
	}, &TestLehvenstein{
		S1: "abc",
		S2: "ab",
		D:  1,
	}, &TestLehvenstein{
		S1: "abc",
		S2: "",
		D:  3,
	})

	for _, td := range testData {
		sd := NewStringDiff(td.S1, td.S2)
		if sd.LevenshteinDistance() != td.D {
			t.Error("Distance between", td.S1, "and", td.S2, "expected to", td.D, "but", sd.LevenshteinDistance())
		}
	}
}

type TestTrigram struct {
	S1 string
	S2 string
	D  float32
}

func TestTrigramCompare(t *testing.T) {
	testData := make([]*TestTrigram, 0)
	testData = append(testData, &TestTrigram{
		S1: "Twitter v1",
		S2: "Twitter v2",
		D:  0.6666667,
	}, &TestTrigram{
		S1: "Twitter v1",
		S2: "Twitter v1",
		D:  1,
	})
	for _, td := range testData {
		sd := NewStringDiff(td.S1, td.S2)
		if sd.TrigramCompare() != td.D {
			t.Error("trigram Compare between", td.S1, "and", td.S2, "expected to", td.D, "but", sd.TrigramCompare())
		}
	}
}

type TestJaroDistancce struct {
	S1 string
	S2 string
	DJ float32
}

func TestJaroDistance(t *testing.T) {
	testData := make([]*TestJaroDistancce, 0)
	testData = append(testData, &TestJaroDistancce{
		S1: "martha",
		S2: "marhta",
		DJ: 0.9444444,
	}, &TestJaroDistancce{
		S1: "martha",
		S2: "martha",
		DJ: 1,
	})
	for _, td := range testData {
		sd := NewStringDiff(td.S1, td.S2)
		if sd.JaroDistance() != td.DJ {
			t.Error("Jaro Distance between", td.S1, "and", td.S2, "expected to", td.DJ, "but", sd.JaroDistance())
		}
	}
}

func TestJaroWinklerDistance(t *testing.T) {
	testData := make([]*TestJaroDistancce, 0)
	testData = append(testData, &TestJaroDistancce{
		S1: "martha",
		S2: "marhta",
		DJ: 0.96111107,
	}, &TestJaroDistancce{
		S1: "martha",
		S2: "martha",
		DJ: 1,
	})
	for _, td := range testData {
		sd := NewStringDiff(td.S1, td.S2)
		if sd.JaroWinklerDistance(0.1) != td.DJ {
			t.Error("Jaro Distance between", td.S1, "and", td.S2, "expected to", td.DJ, "but", sd.JaroWinklerDistance(0.1))
		}
	}
}
