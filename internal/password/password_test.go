package password

import (
	"reflect"
	"strconv"
	"testing"
)

// TestPassword test password generation using the Password struct
func TestPassword_Generate(t *testing.T) {
	cases := [...]struct {
		password *Password
		expected [][]rune
	}{
		{
			expected: [][]rune{
				{'A', 'A', 'A'},
				{'A', 'A', 'B'},
				{'A', 'B', 'A'},
				{'A', 'B', 'B'},
				{'B', 'A', 'A'},
				{'B', 'A', 'B'},
				{'B', 'B', 'A'},
				{'B', 'B', 'B'},
			},
			password: &Password{
				characters: []rune{'A', 'B'},
				state:      []int{0, 0, 0},
			},
		},
		{
			expected: [][]rune{
				{'🐥', '🐣'},
				{'🐥', '🐥'},
			},
			password: &Password{
				characters: []rune{'🐣', '🐥'},
				state:      []int{1, 0},
			},
		},
		{
			password: &Password{},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			password := c.password

			for n, expected := range c.expected {
				if !password.Next() {
					t.Fatal("could not generated more passwords ", n)
				}

				generated := password.Generate()

				if !reflect.DeepEqual(expected, generated) {
					t.Errorf(`expected "%v" got "%v"`, expected, generated)
					continue
				}

				t.Log(string(generated))
			}

			if password.Next() {
				t.Fatal("the Password creates more passwords than it should")
			}

			t.Log(password.State())
		})
	}
}
