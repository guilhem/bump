package git

import (
	"fmt"
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
		return nil, fmt.Errorf("can't get current directory: %w", err)
	}

	opt := &git.PlainOpenOptions{
		DetectDotGit:          true,
		EnableDotGitCommonDir: false,
	}

	repo, err := git.PlainOpenWithOptions(dir, opt)
	if err != nil {
		return nil, fmt.Errorf("can't open git: %w", err)
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
}

// Thanks King'ori Maina @itskingori
// https://github.com/src-d/go-git/issues/1030#issuecomment-443679681

func (g *Git) Branches() ([]string, error) {
	var currentBranchesNames []string

	branchRefs, err := g.repository.Branches()
	if err != nil {
		return currentBranchesNames, fmt.Errorf("can't get branches: %w", err)
	}

	headRef, err := g.repository.Head()
	if err != nil {
		return currentBranchesNames, fmt.Errorf("can't get HEAD object: %w", err)
	}

	err = branchRefs.ForEach(func(branchRef *plumbing.Reference) error {
		if branchRef.Hash() == headRef.Hash() {
			currentBranchesNames = append(currentBranchesNames, branchRef.Name().Short())

			return nil
		}

		return nil
	})
	if err != nil {
		return currentBranchesNames, fmt.Errorf("can't parse branches: %w", err)
	}

	return currentBranchesNames, nil
}

func (g *Git) CurrentCommit() (string, error) {
	headRef, err := g.repository.Head()
	if err != nil {
		return "", fmt.Errorf("can't get HEAD object: %w", err)
	}

	return headRef.Hash().String(), nil
}

func (g *Git) Tags() ([]string, error) {
	return g.tags()
}

func (g *Git) tags() ([]string, error) {
	var tags []string

	tagsHash := make(map[plumbing.Hash][]string)

	tagRefs, err := g.repository.Tags()
	if err != nil {
		return tags, fmt.Errorf("can't get tags: %w", err)
	}

	err = tagRefs.ForEach(func(tagRef *plumbing.Reference) error {
		tagsHash[tagRef.Hash()] = append(tagsHash[tagRef.Hash()], tagRef.Name().Short())

		return nil
	})
	if err != nil {
		return tags, fmt.Errorf("can't parse tags: %w", err)
	}

	head, err := g.repository.Head()
	if err != nil {
		return tags, fmt.Errorf("can't get HEAD object: %w", err)
	}

	commitHead, err := g.repository.CommitObject(head.Hash())
	if err != nil {
		return tags, fmt.Errorf("can't get HEAD commits: %w", err)
	}

	iter := object.NewCommitIterCTime(commitHead, make(map[plumbing.Hash]bool), []plumbing.Hash{})
	err = iter.ForEach(func(c *object.Commit) error {
		tag, ok := tagsHash[c.ID()]
		if ok {
			tags = append(tags, tag...)
		}

		return nil
	})

	if err != nil {
		return tags, fmt.Errorf("can't parse commits: %w", err)
	}

	return tags, nil
}

func (g *Git) CreateTag(version string) error {
	head, err := g.repository.Head()
	if err != nil {
		return fmt.Errorf("can't get HEAD object: %w", err)
	}

	if _, err := g.repository.CreateTag(version, head.Hash(), nil); err != nil {
		return fmt.Errorf("can't create tag: %w", err)
	}

	return nil
}
