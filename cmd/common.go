package cmd

import (
	"errors"

	"github.com/go-logr/logr"
	"github.com/guilhem/bump/pkg/git"
	"github.com/guilhem/bump/pkg/semver"
	"github.com/spf13/cobra"
)

func inc(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	log, err := logr.FromContext(ctx)
	if err != nil {
		cmd.PrintErrf("error getting log: %v\n", err)
		return err
	}

	log = log.WithName(cmd.CommandPath())

	log = log.WithValues("Current Tag", currentTag)

	version := semver.New(currentTag)

	switch cmd.CommandPath() {
	case "bump patch":
		version.IncPatch()
	case "bump minor":
		version.IncMinor()
	case "bump major":
		version.IncMajor()
	default:
		log.Error(errors.New("command not matched"), "increment", "commandPath", cmd.CommandPath())
		return err
	}

	log = log.WithValues("New  Tag", version.StringFull())

	log.Info("values")

	if !dryRun {
		g := ctx.Value("git").(*git.Git)

		if err := g.CreateTag(version.StringFull()); err != nil {
			log.Error(err, "Create Tag")
			return err
		}
	}

	return nil
}
