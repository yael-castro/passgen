// Package cli there is everything related to the actions as obtain and validate the received parameters
package cli

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/yael-castro/passgen/internal/password"
)

// binaryName is the application name
var binaryName string

const (
	logo = `
     ____                      _____   _____  _     __
    /    \ ____   ___ ___     //  //  / ___/ / \   / /
   /  ___//   /| //_ //_  == //  __  / __/  / / \ / /  Password Generator ${version} (https://github.com/yael-castro) 
  /__/    \___\|__//__//    //___// /____/ /_/   \_/
  `
	version  = `v0.1.0`
	howToUse = "%s\nHow to use:\n\n  %s [options]\n\nOptions:\n\n"
)

func init() {
	_, binaryName = filepath.Split(os.Args[0])
}

// usage default handler for cli misuse
func usage() {
	fmt.Printf(howToUse, strings.ReplaceAll(logo, "${version}", version), binaryName)

	flag.CommandLine.SetOutput(os.Stdout)
	flag.CommandLine.PrintDefaults()
	flag.CommandLine.SetOutput(io.Discard) // hide flag errors

	fmt.Println()
}

// InitFlags initialize all fields of the model.Flags passed as parameter
func InitFlags(flags *Flags) {
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + ".txt"

	flag.CommandLine.SetOutput(io.Discard) // hide flag errors
	flag.CommandLine.Init(os.Args[0], flag.ExitOnError)

	flag.CommandLine.UintVar(&flags.Length, "length", 0, "use to set length for words (this flag will ignore if the flag 'state' is present)")
	flag.CommandLine.StringVar(&flags.String, "string", "", "use to set characters used to generate values")
	flag.CommandLine.StringVar(&flags.Output, "out", fileName, "use to set name to the output file")
	flag.CommandLine.BoolVar(&flags.Verbose, "v", false, "verbose")

	flags.State = NewValue(&[]int{})
	flag.CommandLine.Var(flags.State, "state", "password generation state, must be an array")

	flag.CommandLine.Usage = usage

	flag.CommandLine.Parse(os.Args[1:])
}

// ValidateFlags validate the received flags and if it founds an error exit the program with a exit code = 2
func ValidateFlags(flags Flags) {
	length := len(*(flags.State.Value().(*[]int))) // length of state

	if flags.String == "" || (flags.Length < 1 && length < 1) {
		flag.CommandLine.Usage()
		os.Exit(2)
	}
}

// Run executes the function RunContext with the context.Background() by default
func Run(w io.Writer, flags Flags) error {
	return RunContext(context.Background(), w, flags)
}

// RunContext concurrently executes password generation based on the received Flags.
// When a password is generated immediately it is saved using the io.Writer
//
// Note: If the context.Context passed as parameter is canceled the passwords
// is canceled generation
func RunContext(ctx context.Context, w io.Writer, flags Flags) error {
	doneCh := make(chan struct{})

	_, err := os.Stat(flags.Output)
	if err == nil {
		fmt.Fprintf(w, "[ERROR] the file '%s' already exists\n", flags.Output)
		return Code(2)
	}

	// saving information in file
	file, err := os.Create(flags.Output)
	if err != nil {
		fmt.Fprintln(w, "[ERROR] failed creating file")
		return Code(2)
	}

	// defer works a LIFO
	defer file.Close()
	defer file.Sync()

	p := password.Password{}

	state := *(flags.State.Value().(*[]int))

	if len(state) < 1 {
		state = make([]int, flags.Length)
	}

	p.SetState(state)
	p.SetCharacters([]rune(flags.String))

	go func() {
		fmt.Fprintf(w, "Generating '%d' passwords...\n", p.N())

		for p.Next() {
			word := p.Generate()
			file.Write([]byte(string(word) + "\n"))
		}

		doneCh <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		fmt.Fprintln(w, "\nCanceled process with state:", p.State())
		fmt.Fprintf(w, "Saved successfully in file '%v'\n", flags.Output)
		return ctx.Err()
	case <-doneCh:
		fmt.Fprintf(w, "[SUCCESS] password generation\n")
		fmt.Fprintf(w, "Saved successfully in file '%v'\n", flags.Output)
		return nil
	}
}
