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
	"os"
	"strings"

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
	Run: runPP,
}

func runPP(cmd *cobra.Command, args []string) {
	data, err := os.ReadFile(args[0])
	check(err)
	prettyPrint(data)
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func prettyPrint(data []byte){
	colWidths, rows := getSizes(data)
	r := csv.NewReader(bytes.NewReader(data))
	totalLen := sum(colWidths)
	rowsShown := 15
	lineNum := 0
	fmt.Printf("+%s+\n", strings.Repeat("-", totalLen-2))
	for {
		record, err := r.Read()
		if err == io.EOF || lineNum > rowsShown {
			break
		}
		check(err)
		for i, val := range record {
			fmt.Printf("| %-*s ", colWidths[i]+2, val)
		}
		if lineNum == 0 {
			fmt.Printf("\n+%s+\n", strings.Repeat("-", totalLen-2))
		} else{
			fmt.Println()
		}
		lineNum++
	}
	fmt.Printf("+%s+\n", strings.Repeat("-", totalLen-2))
	fmt.Printf("%v rows not shown...\n", rows - rowsShown)
}

func getSizes(data []byte) ([]int, int) {
	r := csv.NewReader(bytes.NewReader(data))
	record, err := r.Read()
	check(err)
	var colWidths = make([]int,len(record))
	rows := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)
		for i, val := range record {
			if len(val) > colWidths[i] {
				colWidths[i] = len(val)
			}
		}
		rows++
	}
	return colWidths, rows
}

func init() {
	rootCmd.AddCommand(prettyCmd)
}
