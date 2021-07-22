package cmd

import (
	"errors"

	"github.com/go-logr/logr"
	"github.com/guilhem/bump/pkg/git"
	"github.com/guilhem/bump/pkg/semver"
	"github.com/spf13/cobra"
)

func inc(cmd *cobra.Command, args []string) {
	ctx := cmd.Context()

	log, err := logr.FromContext(ctx)
	if err != nil {
		cmd.PrintErrf("error getting log: %v\n", err)
		return
	}

	log = log.WithName(cmd.CommandPath())

	log = log.WithValues("Current Tag", currentTag)

	version := semver.New(currentTag)

	var vInc semver.Bump

	switch cmd.CommandPath() {
	case "bump patch":
		vInc = version.IncPatch()
	case "bump minor":
		vInc = version.IncMinor()
	case "bump major":
		vInc = version.IncMajor()
	default:
		log.Error(errors.New("command not matched"), "increment", "commandPath", cmd.CommandPath())
		return
	}

	log = log.WithValues("Original", vInc.Original())

	log.Info("values")

	if !dryRun {
		g := ctx.Value("git").(*git.Git)

		if err := g.CreateTag(vInc.Original()); err != nil {
			log.Error(err, "Create Tag")
		}
	}
}
