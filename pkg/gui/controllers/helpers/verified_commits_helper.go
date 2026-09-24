package helpers

import (
	"github.com/jesseduffield/lazygit/pkg/gui/types"
	"github.com/jesseduffield/lazygit/pkg/verifiedcommits"
)

type VerifiedCommitsHelper struct {
	c *HelperCommon
}

func NewVerifiedCommitsHelper(c *HelperCommon) *VerifiedCommitsHelper {
	return &VerifiedCommitsHelper{c: c}
}

func (self *VerifiedCommitsHelper) Enabled() bool {
	return self.store().Enabled()
}

// Toggle marks the commits as verified, or unmarks them if they all are
// already, and redraws the commit lists.
func (self *VerifiedCommitsHelper) Toggle(hashes []string) error {
	verifiedCommits, err := self.store().Toggle(hashes)
	if err != nil {
		return err
	}

	self.c.Model().VerifiedCommits = verifiedCommits
	for _, context := range []types.Context{
		self.c.Contexts().LocalCommits,
		self.c.Contexts().SubCommits,
	} {
		self.c.PostRefreshUpdate(context)
	}
	return nil
}

func (self *VerifiedCommitsHelper) store() *verifiedcommits.Store {
	return verifiedcommits.NewStore(self.c.UserConfig().Gui.VerifiedCommitsFile)
}
