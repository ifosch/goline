// Copyright © 2016 Ignasi Fosch <ignasi.fosch@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/ifosch/goline/goline"
	"github.com/spf13/cobra"
)

// timelineCmd represents the timeline command
var timelineCmd = &cobra.Command{
	Use:   "timeline",
	Short: "Visualizes a Twitter timeline on your terminal",
	Long: `This command runs goline visualizing your timeline on your
 terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		goline.GetTwitterAPI()
	},
}

func init() {
	RootCmd.AddCommand(timelineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timelineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timelineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
