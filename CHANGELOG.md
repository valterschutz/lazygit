# Changelog of this fork

Personal changes to [jesseduffield/lazygit](https://github.com/jesseduffield/lazygit),
newest first. Upstream's own changes arrive by merging `upstream/master`.

## 2026-09-25: semantic Catppuccin Mocha theme

- Rendering uses one Catppuccin Mocha palette consistently: Mauve for focus,
  Red for errors and failures, Yellow for in-progress work, Green for success,
  Blue and Pink for primary and secondary accents, Base for the application
  background, Mantle for separated surfaces, and Text for normal foregrounds.
- Statuses, progress indicators, commit and file states, pull requests, menus,
  authors, and graph highlights now choose colors by meaning instead of using
  unrelated terminal colors.

## 2026-09-24: commit review status from hunk

- The commit graph draws each commit's review status as hunk derives it from
  the reviewer's per-hunk decisions: hollow while hunks are undecided, half
  filled (`◐`, `◑` for a merge) once every hunk is reviewed but a rejection is
  still open, filled (`●`, `◉`) once every hunk is accepted or fixed.
- The status comes from the `<hash> <status>` file hunk writes beside its
  review file, named by `gui.commitStatusFile`; lazygit only reads it. The new
  words are `reviewed` and `approved`, with the old `verified` and `addressed`
  spellings retained for compatibility.
- The manual `!` toggle, `gui.verifiedCommitsFile`, and the verified-commits
  store are gone; hunk is the only writer.

## Earlier

- A verified-commits marker: `!` toggled a filled circle for reviewed commits,
  kept in a synced plain-text file (superseded by the derived status above).
- Subprocesses only echo their command line when lazygit also stops for their
  output, so running hunk as a custom command leaves no `+ ...` lines in the
  terminal's scrollback.
