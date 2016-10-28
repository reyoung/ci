package ci_server

import (
	"database/sql"
	"net/http"
	"encoding/json"
)

type Repository struct {
	CloneUrl string `json:"clone_url,omitempty"`
}

// PushEvent represents the JSON payload from Github push
// Webhook. https://developer.github.com/v3/activity/events/types/#pushevent
type PushEvent struct {
	Ref string `json:"ref,omitempty"`
	Head string `json:"head,omitempty"`
	Repo Repository
}

func gh_on_push(w http.ResponseWriter, r *http.Request, db *CIDB, buildChan chan int64) {
	var push PushEvent
	CheckNoErr(json.NewDecoder(r.Body).Decode(&push))
	buildId, err := db.AddPushEvent(&push)
	CheckNoErr(err)
	buildChan <- buildId
	w.Write("OK.")
}

func InitGithubHttp(db *sql.DB, buildChan chan int64) {
	cidb := newCIDB(db)
	http.HandleFunc("/ci/",
		makeSafeHandler(func(w http.ResponseWriter, r *http.Request ) {
			event_type := r.Header["X-Github-Event"]
			if len(event_type) == 0 {
				http.Error(w, "Bad request, it seems that this is not github hook",
					http.StatusBadRequest)
			} else {
				switch event_type[0] {
				case "push":
					gh_on_push(w, r, cidb, buildChan)
				default:
					http.Error(w, "Github Event is not supported",
						http.StatusBadRequest)
				}
			}
		}))
}

