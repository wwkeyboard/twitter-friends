// Copyright © 2019 Aaron lee <alee@aaronosaur.us>
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
	"fmt"
	"net/url"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A list of tweets",
	RunE: func(cmd *cobra.Command, args []string) error {
		api := anaconda.NewTwitterApiWithCredentials(
			accessToken,
			accessTokenSecret,
			consumerKey,
			consumerSecret,
		)

		v := url.Values{}
		v.Set("count", "100")
		timeline, err := api.GetHomeTimeline(v)
		if err != nil {
			return err
		}

		for _, t := range timeline {
			if t.User.FollowersCount > 10000 {
				continue
			}
			created, err := t.CreatedAtTime()
			if err != nil {
				fmt.Println(" *** ", err.Error())
				continue
			}
			ago := time.Now().Sub(created)
			if ago.Minutes() > 30 {
				continue
			}
			fmt.Printf("== %v  %v (%v)\n  %v \n", t.User.Name, ago, t.User.FollowersCount, t.Text)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
