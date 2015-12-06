/*
This pattern for handling command line flags with a
Config struct, DefineFlags() and ValidateConfig()
allows for great reuse/combination of libraries
and easy testing of code.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

var ProgramName string = path.Base(os.Args[0])

// For true ease of reuse and testing, put
// ProjectConfig, DefineFlags() and ValidateConfig()
// in their own library.
//
// Replace 'Project' with your library's name,
// define all variables to be parsed from command
// line flags in your Config struct.
type ProjectConfig struct {
	InputPath string
	Int64Flag int64
}

// call DefineFlags before myflags.Parse()
func (c *ProjectConfig) DefineFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.InputPath, "inputPath", "", "describe this flag")
	fs.Int64Var(&c.Int64Flag, "intFlag", -1, "describe this flag")
}

// call c.ValidateConfig() after myflags.Parse()
func (c *ProjectConfig) ValidateConfig() error {
	if c.InputPath == "" {
		return fmt.Errorf("-inputPath required and missing")
	}
	if c.Int64Flag == -1 {
		return fmt.Errorf("-intFlag required and missing")
	}
	return nil
}

// demonstrate the sequence of calls to DefineFlags() and ValidateConfig()
func main() {

	myflags := flag.NewFlagSet("myflags", flag.ExitOnError)
	cfg := &ProjectConfig{}
	cfg.DefineFlags(myflags)

	err := myflags.Parse(os.Args[1:])
	err = cfg.ValidateConfig()
	if err != nil {
		log.Fatalf("%s command line flag error: '%s'", ProgramName, err)
	}

	fmt.Printf("flag parsing done, the rest of program goes here...\n")
}
