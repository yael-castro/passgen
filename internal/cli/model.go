package cli

import (
	"encoding/json"
	"flag"
	"reflect"
	"strconv"
)

// Code used as error for exit codes
type Code uint8

// Error returns the uint value formated
func (c Code) Error() string {
	return strconv.FormatInt(int64(c), 10)
}

// Flags contains all cli flags (parameters) used to make a generation password
type Flags struct {
	// Length password length
	Length uint
	// String contains all characters used to generate passwords
	String string
	// Output specify the file name where the password dictionary will save
	Output string
	// Verbose specify if print the program execution step by step (show more details about the password generation)
	Verbose bool
	// State contains the state of password generation
	State *FlagValue
}

// _ flag.Value implementation constraint for *FlagValue
var _ flag.Value = (*FlagValue)(nil)

// NewValue constructs a flag.Value
func NewValue(value interface{}) *FlagValue {
	return &FlagValue{value: value}
}

// FlagValue implementation of flag.Value used to decode a json data in any
type FlagValue struct {
	value interface{}
}

// Value returns saved value
func (s *FlagValue) Value() interface{} {
	return s.value
}

// Set takes the string passed as parameter, and it assumes
// that the string is encoded in json format and decode
// the serialized data to save it
func (s *FlagValue) Set(value string) error {
	return json.Unmarshal([]byte(value), s.value)
}

// String returns the type of saved data
func (s *FlagValue) String() string {
	return reflect.TypeOf(s).String()
}
