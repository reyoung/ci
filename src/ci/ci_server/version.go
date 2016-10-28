package ci_server

import "fmt"

const (
	Major = "0"
	Minor = "1"
	Patch = "0"
)

func Version() string {
	return fmt.Sprintf("%s.%s.%s", Major, Minor, Patch)
}
