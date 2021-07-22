/*
Copyright © 2019 Guilhem Lettron <guilhem@barpilot.io>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"errors"
	"os"

	"github.com/bombsimon/logrusr/v2"
	"github.com/go-logr/logr"
	"github.com/guilhem/bump/pkg/git"
	"github.com/guilhem/bump/pkg/semver"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	allowDirty bool
	currentTag string
	latestTag  bool
	dryRun     bool
)

var rootCmd = &cobra.Command{
	Use:   "bump",
	Short: "Bump version",
	Long:  ``,

	SilenceUsage: true,

	PersistentPreRunE: preRun,
}

func Execute() {
	logrusLog := logrus.New()
	log := logrusr.New(logrusLog)

	ctx := logr.NewContext(context.Background(), log)

	g, err := git.New()
	if err != nil {
		log.Error(err, "git new")
		os.Exit(1)
	}

	ctx = context.WithValue(ctx, "git", g)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		log.Error(err, "ExecuteContext")
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().BoolVar(&allowDirty, "allow-dirty", false, "allow usage of bump on dirty git")
	rootCmd.PersistentFlags().BoolVar(&latestTag, "latest-tag", true, "use latest tag, prompt tags if false")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Don't touch git repository")
}

func preRun(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	log, err := logr.FromContext(ctx)
	if err != nil {
		cmd.PrintErrf("error getting log: %v\n", err)
		return err
	}

	g := ctx.Value("git").(*git.Git)

	if !allowDirty {
		if g.IsDirty() {
			err := errors.New("is dirty")
			log.Error(err, "test dirty")
			return err
		}
	}

	tags, err := g.Tags()
	if err != nil {
		log.Error(err, "get tags")
		return err
	}

	log = log.WithValues("tags", tags)

	if !latestTag {
		prompt := promptui.Select{
			Label: "Select Previous tag",
			Items: tags,
		}

		_, currentTag, err = prompt.Run()

		if err != nil {
			log.Error(err, "prompt run")
			return err
		}

	} else {
		currentTag, err = semver.Latest(tags)
		if err != nil {
			log.Error(err, "get latest tag")
			return err
		}
	}

	log.Info("tag choosed", "current tag", currentTag)

	return nil
}
