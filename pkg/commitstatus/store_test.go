package commitstatus

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisabledStoreIsEmpty(t *testing.T) {
	store := NewStore("")

	assert.False(t, store.Enabled())
	statuses, err := store.Load()
	assert.NoError(t, err)
	assert.Empty(t, statuses)
}

func TestLoadMissingFile(t *testing.T) {
	store := NewStore(filepath.Join(t.TempDir(), "commit-status"))

	assert.True(t, store.Enabled())
	statuses, err := store.Load()
	assert.NoError(t, err)
	assert.Empty(t, statuses)
}

func TestLoadParsesStatusesAndSkipsWhatItCannotRead(t *testing.T) {
	path := filepath.Join(t.TempDir(), "commit-status")
	content := "\nabc reviewed\n  \ndef  approved  \nghi unknown\nlonely\n"
	assert.NoError(t, os.WriteFile(path, []byte(content), 0o644))

	statuses, err := NewStore(path).Load()
	assert.NoError(t, err)
	assert.Equal(t, map[string]Status{"abc": Reviewed, "def": Approved}, statuses)
}

func TestLoadParsesLegacyStatusNames(t *testing.T) {
	path := filepath.Join(t.TempDir(), "commit-status")
	assert.NoError(t, os.WriteFile(path, []byte("abc verified\ndef addressed\n"), 0o644))

	statuses, err := NewStore(path).Load()
	assert.NoError(t, err)
	assert.Equal(t, map[string]Status{"abc": Reviewed, "def": Approved}, statuses)
}

func TestLoadExpandsTilde(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	path := filepath.Join(os.Getenv("HOME"), "commit-status")
	assert.NoError(t, os.WriteFile(path, []byte("abc approved\n"), 0o644))

	statuses, err := NewStore("~/commit-status").Load()
	assert.NoError(t, err)
	assert.Equal(t, map[string]Status{"abc": Approved}, statuses)
}
