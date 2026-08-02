// Command atlasloader prints the schema described by the bot's gorm models.
// Atlas runs it to work out what the database should look like, and compares
// that against the migrations directory to generate the next migration.
//
// It deliberately imports only the db package: importing env would try to open
// a database connection, which this has no need of.
package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"phoenixbot/bot/db"
)

func main() {
	statements, err := gormschema.New("mysql").Load(db.Models()...)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	_, _ = io.WriteString(os.Stdout, statements)
}
