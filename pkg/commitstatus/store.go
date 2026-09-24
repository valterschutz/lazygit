package commitstatus

import (
	"errors"
	"io/fs"
	"os"
	"strings"

	"github.com/jesseduffield/lazygit/pkg/utils"
)

// Status is what the reviewer's hunk decisions add up to for one commit:
// reviewed once every hunk is decided, approved once every decision is accepted
// or fixed. A commit with an undecided hunk has no status.
type Status int

const (
	Unreviewed Status = iota
	Reviewed
	Approved
)

// Store reads the per-commit statuses hunk derives from its review file and
// writes as `<hash> <status>` lines, one commit per line. Lazygit only reads
// the file; every change to it is made by hunk. An empty path disables the
// store: it is then always empty.
type Store struct {
	path string
}

func NewStore(path string) *Store {
	return &Store{path: utils.ExpandTilde(path)}
}

func (self *Store) Enabled() bool {
	return self.path != ""
}

// Load returns the status of every commit the file names. A file that does
// not exist yet is an empty store; a line it cannot read is skipped.
func (self *Store) Load() (map[string]Status, error) {
	statuses := map[string]Status{}
	if !self.Enabled() {
		return statuses, nil
	}

	content, err := os.ReadFile(self.path)
	if errors.Is(err, fs.ErrNotExist) {
		return statuses, nil
	}
	if err != nil {
		return nil, err
	}

	for line := range strings.SplitSeq(string(content), "\n") {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		status, ok := parseStatus(fields[1])
		if !ok {
			continue
		}
		statuses[fields[0]] = status
	}
	return statuses, nil
}

func parseStatus(word string) (Status, bool) {
	switch word {
	case "reviewed", "verified":
		return Reviewed, true
	case "approved", "addressed":
		return Approved, true
	default:
		return Unreviewed, false
	}
}
