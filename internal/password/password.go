// Package password contains everything required to generate passwords
package password

import (
	"math"
)

// Generator implemented by password generators
type Generator interface {
	// Generate use to generate a new password random, based on time or some state
	Generate() []rune
}

// _ Generator implementation constraint for Password
var _ Generator = (*Password)(nil)

// Password struct used to make passwords based on a state and characters
type Password struct {
	// characters are the chars that contains the generated passwords
	characters []rune
	// state slice that represents the state of password generation
	state []int
}

// SetState sets the state used to make a passwords
func (p *Password) SetState(state []int) {
	p.state = state
}

// State returns a slice that represents the state of password generation
// more specifically defines the length of passwords that will generate, and
// it is used during the password generation.
//
// How it works:
//
// For an empty state with length 3 the initial state will be: [0, 0, 0]
// but the state can be init with any value always it drive for the following points...
//
// For each slot the maximum value is: len(Characters())
//
// When a slot exceeds the maximum allowed value the slot value resets to 0
// and the previous slot increments its value by one.
//
// When a password is generated using the method Generate the last slot increases by one.
//
// When the first slot exceeds the maximum allowed value the password generation ends.
func (p *Password) State() []int {
	return p.state
}

// SetCharacters use to set characters without repeat
func (p *Password) SetCharacters(characters []rune) {
	set := make(map[rune]struct{})

	// Extracting unique characters
	for _, char := range characters {
		set[char] = struct{}{}
	}

	// Setting unique characters
	p.characters = make([]rune, 0, len(set))

	for char := range set {
		p.characters = append(p.characters, char)
	}
}

// Characters returns the characters used to generate passwords
func (p *Password) Characters() []rune {
	return p.characters
}

// Generate takes the current state (for example [1, 0, 1]) to get your corresponding char for each slot and with it make a password.
// Following the previous example the password generation should be:
//
// Password = Character()[1] + Character()[0] + Character()[1]
//
// Note: panics if the password use Generate when Next method returns false
func (p *Password) Generate() (word []rune) {
	if !p.Next() {
		panic("can not create more passwords")
	}

	for i := range p.state {
		word = append(word, p.characters[p.state[i]])
	}

	index := len(p.state) - 1
	p.state[index]++

	for index != 0 && p.state[index] == len(p.characters) {
		p.state[index] = 0
		index--
		p.state[index]++
	}

	return
}

// Next check if a password can be generated
func (p *Password) Next() bool {
	return len(p.characters) > 0 && len(p.state) > 0 && p.state[0] != len(p.characters)
}

// N returns the number of passwords that be created with an initial state.
// The initial state refers an empty slice of length "n": [0, 0, ...]
func (p *Password) N() int {
	return int(math.Pow(float64(len(p.characters)), float64(len(p.state))))
}
