package ci_server

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	addr string
	dbuser string
	dbpasswd string
	database string
}

const usage string = `Usage %s [OPTIONS]
Options:
`

func ParseArgs() (opts *Options) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flag.PrintDefaults()
	}

	addr := flag.String("addr", ":8000", "Listened http address.")
	dbuser := flag.String("dbuser", "DEBUG", "SQL database user.")
	dbpasswd := flag.String("dbpasswd", "DEBUG", "SQL database password.")
	database := flag.String("database", "ci", "SQL database name")

	flag.Parse()

	opts = &Options{
		addr: *addr,
		dbuser: *dbuser,
		dbpasswd: *dbpasswd,
		database: *database,
	}
	return
}