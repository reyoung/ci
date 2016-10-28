package ci_server

import "log"

func Main() {
	opts := ParseArgs()
	github := newGithubAPI(opts)
	db, err := newCIDB(opts.Database.User, opts.Database.Password, opts.Database.DatabaseName)
	CheckNoErr(err)
	buildChan := make(chan int64, 256)
	go func() { CheckNoErr(db.RecoverFromPreviousDown(buildChan)) }()
	builder, err := newBuilder(buildChan, opts, db, github)
	builder.Start()
	defer builder.Close()

	serv := newHttpServer(&opts.HTTP, db)
	serv.GoListenAndServe()

	for ev := range serv.EventQueue {
		switch ev.(type) {
		case *PushEvent:
			{
				event := ev.(*PushEvent)
				bid, err := db.AddPushEvent(event) // add event to db
				CheckNoErr(err)
				buildChan <- bid
			}
		}
	}
}

func CheckNoErr(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}
