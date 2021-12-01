/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"bytes"
	"fmt"
	"encoding/csv"
	"io"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// prettyCmd represents the pretty command
var prettyCmd = &cobra.Command{
	Use:   "pretty",
	Short: "Pretty print csv",
	Long: `Pretty printin, spendin gs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pretty called")
		data, err := os.ReadFile(args[0])
		check(err)
		r := csv.NewReader(bytes.NewReader(data))

		record, err := r.Read()
		var colLens = make([]int,len(record))
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			for i, val := range record {
				if len(val) > colLens[i] {
					colLens[i] = len(val)
				}
			}
			fmt.Println(colLens)
		}

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', 0)
			for i, val := range record {
				if len(val) > colLens[i] {
					colLens[i] = len(val)
				}
			}
			fmt.Println(colLens)
		}
	},
}

func init() {
	rootCmd.AddCommand(prettyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prettyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prettyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
