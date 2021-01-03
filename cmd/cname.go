/*
Copyright © 2021 N Cole Summers <nsummers72@gmail.com>

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
	"net"

	"github.com/spf13/cobra"
)

// cnameCmd represents the cname command
var cnameCmd = &cobra.Command{
	Use:   "cname",
	Short: "Looks up the CNAME for a host",
	Long: `A means to look up a Canonical Name or CNAME record. CNAME records are typically used to map a subdomain such as www or mail to the domain hosting that subdomain's content.  
	Usage: net-toolbox cname host`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			fmt.Println("Getting CNAME records for:", arg)
			cname, err := net.LookupCNAME(arg)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(cname)
		}
	},
}

func init() {
	rootCmd.AddCommand(cnameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cnameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cnameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
