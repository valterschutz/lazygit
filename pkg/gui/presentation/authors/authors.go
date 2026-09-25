package authors

import (
	"strings"

	"github.com/jesseduffield/lazygit/pkg/gui/style"
	"github.com/jesseduffield/lazygit/pkg/theme"
	"github.com/jesseduffield/lazygit/pkg/utils"
	"github.com/rivo/uniseg"
)

type authorNameCacheKey struct {
	authorName string
	truncateTo int
}

// if these being global variables causes trouble we can wrap them in a struct
// attached to the gui state.
var (
	authorInitialCache = make(map[string]string)
	authorNameCache    = make(map[authorNameCacheKey]string)
	authorStyleCache   = make(map[string]*style.TextStyle)
)

const authorNameWildcard = "*"

func ShortAuthor(authorName string) string {
	if value, ok := authorInitialCache[authorName]; ok {
		return value
	}

	initials := getInitials(authorName)
	if initials == "" {
		return ""
	}

	value := AuthorStyle(authorName).Sprint(initials)
	authorInitialCache[authorName] = value

	return value
}

func LongAuthor(authorName string, length int) string {
	cacheKey := authorNameCacheKey{authorName: authorName, truncateTo: length}
	if value, ok := authorNameCache[cacheKey]; ok {
		return value
	}

	paddedAuthorName := utils.WithPadding(authorName, length, utils.AlignLeft)
	truncatedName := utils.TruncateWithEllipsis(paddedAuthorName, length)
	value := AuthorStyle(authorName).Sprint(truncatedName)
	authorNameCache[cacheKey] = value

	return value
}

// AuthorWithLength returns a representation of the author that fits into a
// given maximum length:
// - if the length is less than 2, it returns an empty string
// - if the length is 2, it returns the initials
// - otherwise, it returns the author name truncated to the maximum length
func AuthorWithLength(authorName string, length int) string {
	if length < 2 {
		return ""
	}

	if length == 2 {
		return ShortAuthor(authorName)
	}

	return LongAuthor(authorName, length)
}

func AuthorStyle(authorName string) *style.TextStyle {
	if value, ok := authorStyleCache[authorName]; ok {
		return value
	}

	// use the unified style whatever the author name is
	if value, ok := authorStyleCache[authorNameWildcard]; ok {
		return value
	}

	value := theme.Semantic.SecondaryAccent
	authorStyleCache[authorName] = &value
	return &value
}

func getInitials(authorName string) string {
	if authorName == "" {
		return authorName
	}

	firstChar, _, width, _ := uniseg.FirstGraphemeClusterInString(authorName, -1)
	if width > 1 {
		return firstChar
	}

	split := strings.Split(authorName, " ")
	if len(split) == 1 {
		return utils.LimitStr(authorName, 2)
	}

	return utils.LimitStr(split[0], 1) + utils.LimitStr(split[1], 1)
}

func SetCustomAuthors(customAuthorColors map[string]string) {
	authorStyleCache = utils.SetCustomColors(customAuthorColors)
}
