package helper

import (
	"flag"
)

type (

	// Parser struct to hold parser data
	Parser struct {
		Owner  string
		Cakes  uint64
		Apples uint64
	}
)

// ParseFlag function to parse flag data
func ParseFlag(defaultParser *Parser) (parser Parser) {

	// set up parser data
	parser = Parser{}

	// get flag data
	owner := flag.String("owner", defaultParser.Owner, "who is the name of owner")
	apples := flag.Uint64("apples", defaultParser.Apples, "how much apples to give")
	cakes := flag.Uint64("cakes", defaultParser.Cakes, "how much cakes to give")

	// parse data
	flag.Parse()

	if owner != nil {
		parser.Owner = *owner
	}

	if apples != nil {
		parser.Apples = *apples
	}

	if cakes != nil {
		parser.Cakes = *cakes
	}

	return
}
