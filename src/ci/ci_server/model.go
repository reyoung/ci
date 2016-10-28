package ci_server

import "database/sql"

type CIDB struct {
	DB *sql.DB
}

func newCIDB(db *sql.DB) *CIDB {
	return &CIDB {
		DB: db,
	}
}

func (db *CIDB) AddPushEvent(event *PushEvent) (build_id int64, err error) {
	addPushEventStmt, err := db.DB.Prepare("select new_push_event(?, ?, ?)")
	if err != nil {
		return
	}
	defer addPushEventStmt.Close()
	err = addPushEventStmt.QueryRow(event.Head, event.Ref, event.Repo.CloneUrl).Scan(&build_id)
	return
}