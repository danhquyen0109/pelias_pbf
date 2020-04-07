package command

import (
	"log"
	"os"

	"github.com/danhquyen0109/pelias_pbf/handler"
	"github.com/danhquyen0109/pelias_pbf/lib"
	"github.com/danhquyen0109/pelias_pbf/parser"

	"github.com/codegangsta/cli"
)

// BitmaskSuperRelations cli command
func BitmaskSuperRelations(c *cli.Context) error {

	// create parser
	parser := parser.NewParser(c.Args()[0])

	// don't clobber existing bitmask file
	if _, err := os.Stat(c.Args()[1]); err == nil {
		log.Println("bitmask file already exists; don't want to override it")
		os.Exit(1)
	}

	// open database for writing
	handle := &handler.BitmaskSuperRelations{
		Masks: lib.NewBitmaskMap(),
	}

	// write to disk
	defer handle.Masks.WriteToFile(c.Args()[1])

	// Parse will block until it is done or an error occurs.
	parser.Parse(handle)

	return nil
}
