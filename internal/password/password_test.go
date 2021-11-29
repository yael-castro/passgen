package password

import (
	"reflect"
	"strconv"
	"testing"
)

// tdtPassword is table driven testing for Password struct
type tdtPassword struct {
	expectedPasswords [][]rune
	*Password
}

var (
	tests = []tdtPassword{
		{
			expectedPasswords: [][]rune{
				{'A', 'A', 'A'},
				{'A', 'A', 'B'},
				{'A', 'B', 'A'},
				{'A', 'B', 'B'},
				{'B', 'A', 'A'},
				{'B', 'A', 'B'},
				{'B', 'B', 'A'},
				{'B', 'B', 'B'},
			},
			Password: &Password{
				characters: []rune{'A', 'B'},
				state:      []int{0, 0, 0},
			},
		},
		{
			expectedPasswords: [][]rune{
				{'🐥', '🐣'},
				{'🐥', '🐥'},
			},
			Password: &Password{
				characters: []rune{'🐣', '🐥'},
				state:      []int{1, 0},
			},
		},
		{
			Password: &Password{},
		},
	}

	benchmark = tests[0]
)

// TestPassword test password generation using the Password struct
func TestPassword(t *testing.T) {
	for i, v := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			for n, expected := range v.expectedPasswords {
				if !v.Next() {
					t.Fatal("could not generated more passwords ", n)
				}

				got := v.Generate()

				if !reflect.DeepEqual(expected, got) {
					t.Errorf(`expected "%v" got "%v"`, expected, got)
					continue
				}

				t.Log(string(got))
			}

			if v.Next() {
				t.Fatal("the Password creates more passwords than it should")
			}

			t.Log(v.State())
		})
	}
}
