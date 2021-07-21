package password

import (
	"reflect"
	"strconv"
	"testing"
)

type tableDrivenTesting struct {
	Length      int
	Letters     []rune
	ExpectedSet [][]rune
}

var (
	tests = []tableDrivenTesting{
		{
			Length:  2,
			Letters: []rune{'A', 'B'},
			ExpectedSet: [][]rune{
				{'B', 'B'},
				{'B', 'A'},
				{'A', 'B'},
				{'A', 'A'},
			},
		},
	}

	benchmark = tests[0]
)

func TestGenerateWordSet(t *testing.T) {
	for i, test := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			gotSet := GenerateWordSet(test.Length, test.Letters)

			if !reflect.DeepEqual(gotSet, test.ExpectedSet) {
				t.Errorf("expected '%v' got '%v'", test.ExpectedSet, gotSet)
			}
		})
	}
}

func BenchmarkGenerateWordSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateWordSet(benchmark.Length, benchmark.Letters)
	}
}
