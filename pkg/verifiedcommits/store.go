package verifiedcommits

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/jesseduffield/generics/set"
	"github.com/jesseduffield/lazygit/pkg/utils"
)

// Store persists the hashes of the commits the user has marked as verified in
// a plain text file, one hash per line, so that any file syncing tool can
// carry the list between machines. An empty path disables the store: it is
// then always empty and cannot be written to.
type Store struct {
	path string
}

func NewStore(path string) *Store {
	return &Store{path: utils.ExpandTilde(path)}
}

func (self *Store) Enabled() bool {
	return self.path != ""
}

// Load returns the stored hashes. A file that does not exist yet is an empty
// store.
func (self *Store) Load() (*set.Set[string], error) {
	hashes := set.New[string]()
	if !self.Enabled() {
		return hashes, nil
	}

	content, err := os.ReadFile(self.path)
	if errors.Is(err, fs.ErrNotExist) {
		return hashes, nil
	}
	if err != nil {
		return nil, err
	}

	for line := range strings.SplitSeq(string(content), "\n") {
		if line = strings.TrimSpace(line); line != "" {
			hashes.Add(line)
		}
	}
	return hashes, nil
}

// Toggle marks the commit as verified if it is not, and unmarks it otherwise,
// and returns the resulting hashes. The file is re-read first so that entries
// added on another machine in the meantime are kept.
func (self *Store) Toggle(hash string) (*set.Set[string], error) {
	if !self.Enabled() {
		panic("cannot toggle a verified commit in a disabled store")
	}
	if hash == "" {
		panic("commit hash must not be empty")
	}

	hashes, err := self.Load()
	if err != nil {
		return nil, err
	}

	if hashes.Includes(hash) {
		hashes.Remove(hash)
	} else {
		hashes.Add(hash)
	}

	if err := self.save(hashes); err != nil {
		return nil, err
	}
	return hashes, nil
}

// The hashes are written sorted, so that the file only changes when the set
// does, and via a rename, so that a sync tool never sees a half-written file.
func (self *Store) save(hashes *set.Set[string]) error {
	lines := hashes.ToSlice()
	slices.Sort(lines)
	content := ""
	if len(lines) > 0 {
		content = strings.Join(lines, "\n") + "\n"
	}

	if err := os.MkdirAll(filepath.Dir(self.path), 0o755); err != nil {
		return err
	}
	tmpPath := self.path + ".tmp"
	if err := os.WriteFile(tmpPath, []byte(content), 0o644); err != nil {
		return err
	}
	return os.Rename(tmpPath, self.path)
}
