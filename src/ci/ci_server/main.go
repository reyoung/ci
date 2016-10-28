package ci_server

import (
	"net/http"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Main() {
	opts := ParseArgs()
	db, err := sql.Open("postgres", fmt.Sprintf("%s:%s@/%s",
		opts.dbuser, opts.dbpasswd, opts.database))
	CheckNoErr(err)
	defer func() { CheckNoErr(db.Close()) }()
	buildChan := make(chan int64, 256)
	InitGithubHttp(db, buildChan)
	CheckNoErr(http.ListenAndServe(opts.addr, nil))
}