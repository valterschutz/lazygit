# Changelog of this fork

Personal changes to [jesseduffield/lazygit](https://github.com/jesseduffield/lazygit),
newest first. Upstream's own changes arrive by merging `upstream/master`.

## 2026-09-24: commit review status from hunk

- The commit graph draws each commit's review status as hunk derives it from
  the reviewer's per-hunk decisions: hollow while hunks are undecided, half
  filled (`◐`, `◑` for a merge) once every hunk is accepted or rejected but a
  rejection is still open, filled (`●`, `◉`) once nothing is left to address.
- The status comes from the `<hash> <status>` file hunk writes beside its
  review file, named by `gui.commitStatusFile`; lazygit only reads it.
- The manual `!` toggle, `gui.verifiedCommitsFile`, and the verified-commits
  store are gone; hunk is the only writer.

## Earlier

- A verified-commits marker: `!` toggled a filled circle for reviewed commits,
  kept in a synced plain-text file (superseded by the derived status above).
- Subprocesses only echo their command line when lazygit also stops for their
  output, so running hunk as a custom command leaves no `+ ...` lines in the
  terminal's scrollback.
