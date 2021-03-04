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
	"log"
	"os"

	"github.com/guilhem/bump/pkg/git"
	"github.com/guilhem/bump/pkg/semver"
	"github.com/manifoldco/promptui"
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

	PersistentPreRun: preRun,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().BoolVar(&allowDirty, "allow-dirty", false, "allow usage of bump on dirty git")
	rootCmd.PersistentFlags().BoolVar(&latestTag, "latest-tag", true, "use latest tag, prompt tags if false")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Don't touch git repository")
}

func preRun(cmd *cobra.Command, args []string) {
	g, err := git.New()
	if err != nil {
		log.Fatalf("not a git repository")
	}

	if !allowDirty {
		if g.IsDirty() {
			log.Fatalf("is dirty")
		}
	}

	tags, err := g.Tags()

	if !latestTag {
		if err != nil {
			log.Fatalf("error tags: %s", err)
		}

		prompt := promptui.Select{
			Label: "Select Previous tag",
			Items: tags,
		}

		_, currentTag, err = prompt.Run()

		if err != nil {
			log.Printf("Prompt failed %v\n", err)

			return
		}

		log.Printf("You choose %q\n", currentTag)
	} else {
		currentTag, err = semver.Latest(tags)
		if err != nil {
			log.Fatalf("Can't get latest tag: %s", err)
		}
	}
}
