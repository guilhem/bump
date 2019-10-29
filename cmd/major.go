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
	"fmt"
	"log"

	"github.com/guilhem/bump/pkg/git"
	"github.com/guilhem/bump/pkg/semver"
	"github.com/spf13/cobra"
)

// majorCmd represents the major command
var majorCmd = &cobra.Command{
	Use:   "major",
	Short: "Bump major version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		version := semver.New(currentTag)
		vInc := version.IncMajor()
		fmt.Println(vInc.Original())
		if !dryRun {
			g, err := git.New()
			if err != nil {
				log.Fatalf("not a git repository")
			}
			if err := g.CreateTag(vInc.Original()); err != nil {
				log.Fatalf("fail to create tag: %s", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(majorCmd)
}
