package verifiedcommits

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisabledStoreIsEmpty(t *testing.T) {
	store := NewStore("")

	assert.False(t, store.Enabled())
	hashes, err := store.Load()
	assert.NoError(t, err)
	assert.Equal(t, 0, hashes.Len())
}

func TestLoadMissingFile(t *testing.T) {
	store := NewStore(filepath.Join(t.TempDir(), "verified-commits"))

	assert.True(t, store.Enabled())
	hashes, err := store.Load()
	assert.NoError(t, err)
	assert.Equal(t, 0, hashes.Len())
}

func TestLoadIgnoresBlankLines(t *testing.T) {
	path := filepath.Join(t.TempDir(), "verified-commits")
	assert.NoError(t, os.WriteFile(path, []byte("\nabc\n  \ndef  \n"), 0o644))

	hashes, err := NewStore(path).Load()
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"abc", "def"}, hashes.ToSlice())
}

func TestToggle(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nested", "verified-commits")
	store := NewStore(path)

	hashes, err := store.Toggle([]string{"def"})
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"def"}, hashes.ToSlice())

	hashes, err = store.Toggle([]string{"abc"})
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"abc", "def"}, hashes.ToSlice())

	content, err := os.ReadFile(path)
	assert.NoError(t, err)
	assert.Equal(t, "abc\ndef\n", string(content))

	hashes, err = store.Toggle([]string{"def"})
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"abc"}, hashes.ToSlice())

	reloaded, err := NewStore(path).Load()
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"abc"}, reloaded.ToSlice())
}

func TestToggleManyMarksAllUnlessAllMarked(t *testing.T) {
	store := NewStore(filepath.Join(t.TempDir(), "verified-commits"))

	_, err := store.Toggle([]string{"abc"})
	assert.NoError(t, err)

	hashes, err := store.Toggle([]string{"abc", "def", "ghi"})
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"abc", "def", "ghi"}, hashes.ToSlice())

	hashes, err = store.Toggle([]string{"abc", "def"})
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"ghi"}, hashes.ToSlice())
}

func TestToggleRejectsBadInput(t *testing.T) {
	store := NewStore(filepath.Join(t.TempDir(), "verified-commits"))

	assert.Panics(t, func() { _, _ = store.Toggle(nil) })
	assert.Panics(t, func() { _, _ = store.Toggle([]string{"abc", ""}) })
	assert.Panics(t, func() { _, _ = NewStore("").Toggle([]string{"abc"}) })
}

func TestToggleKeepsExternalChanges(t *testing.T) {
	path := filepath.Join(t.TempDir(), "verified-commits")
	store := NewStore(path)

	_, err := store.Toggle([]string{"abc"})
	assert.NoError(t, err)

	assert.NoError(t, os.WriteFile(path, []byte("abc\nsynced-from-elsewhere\n"), 0o644))

	hashes, err := store.Toggle([]string{"abc"})
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"synced-from-elsewhere"}, hashes.ToSlice())
}

func TestToggleExpandsTilde(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	store := NewStore("~/verified-commits")

	_, err := store.Toggle([]string{"abc"})
	assert.NoError(t, err)

	content, err := os.ReadFile(filepath.Join(os.Getenv("HOME"), "verified-commits"))
	assert.NoError(t, err)
	assert.Equal(t, "abc\n", string(content))
}
