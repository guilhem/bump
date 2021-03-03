package git

import (
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Git struct {
	repository git.Repository
}

func New() (*Git, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	opt := &git.PlainOpenOptions{DetectDotGit: true}
	repo, err := git.PlainOpenWithOptions(dir, opt)
	if err != nil {
		return nil, err
	}

	git := &Git{
		repository: *repo,
	}
	return git, nil
}

func IsGit() bool {
	_, err := New()
	return err == nil
}

func (g *Git) Revision(longSha bool) (string, error) {
	h, err := g.CurrentCommit()
	if longSha || err != nil {
		return h, err
	}
	return h[:7], err
}

func (g *Git) IsDirty() bool {
	w, err := g.repository.Worktree()
	if err != nil {
		return true
	}

	status, err := w.Status()
	if err != nil {
		return true
	}

	return !status.IsClean()

	// res, _ := oneliner("git", "status", "--porcelain")
	// return len(res) > 0
}

// Thanks King'ori Maina @itskingori
// https://github.com/src-d/go-git/issues/1030#issuecomment-443679681

func (g *Git) Branches() ([]string, error) {
	var currentBranchesNames []string

	branchRefs, err := g.repository.Branches()
	if err != nil {
		return currentBranchesNames, err
	}

	headRef, err := g.repository.Head()
	if err != nil {
		return currentBranchesNames, err
	}

	err = branchRefs.ForEach(func(branchRef *plumbing.Reference) error {
		if branchRef.Hash() == headRef.Hash() {
			currentBranchesNames = append(currentBranchesNames, branchRef.Name().Short())

			return nil
		}

		return nil
	})
	if err != nil {
		return currentBranchesNames, err
	}

	return currentBranchesNames, nil
}

func (g *Git) CurrentCommit() (string, error) {
	headRef, err := g.repository.Head()
	if err != nil {
		return "", err
	}
	headSha := headRef.Hash().String()

	return headSha, nil
}

func (g *Git) Tags() ([]string, error) {
	tags, _, err := g.tags()
	return tags, err
}

func (g *Git) LatestTag() (string, error) {
	_, tag, err := g.tags()
	return tag, err
}

func (g *Git) tags() ([]string, string, error) {
	var latestTagName string
	var tags []string
	var latestTagCommit *object.Commit

	tagRefs, err := g.repository.Tags()
	if err != nil {
		return tags, latestTagName, err
	}

	err = tagRefs.ForEach(func(tagRef *plumbing.Reference) error {
		revision := plumbing.Revision(tagRef.Name().String())
		tagCommitHash, err := g.repository.ResolveRevision(revision)
		if err != nil {
			return err
		}

		commit, err := g.repository.CommitObject(*tagCommitHash)
		if err != nil {
			return err
		}
		tags = append(tags, tagRef.Name().Short())

		if latestTagCommit == nil {
			latestTagCommit = commit
			latestTagName = tagRef.Name().Short()
		}

		if commit.Committer.When.After(latestTagCommit.Committer.When) {
			latestTagCommit = commit
			latestTagName = tagRef.Name().Short()
		}

		return nil
	})
	if err != nil {
		return tags, latestTagName, err
	}

	return tags, latestTagName, nil
}

func (g *Git) CreateTag(version string) error {
	head, err := g.repository.Head()
	if err != nil {
		return err
	}

	_, err = g.repository.CreateTag(version, head.Hash(), nil)
	return err

}
