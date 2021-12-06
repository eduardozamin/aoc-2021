/*
Copyright © 2021 Eduardo Zamin

*/
package cmd

import (
	"aoc-2021/internal/pkg/solutions"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve <day>",
	Args:  cobra.ExactArgs(1),
	Short: "Executes the solutions for the Advent of Code 2021.",
	Long: `Executes the solutions for the Advent of Code 2021.

This command expects exactly one argument corresponding to the solution
to be executed. For example:

	day05    Runs the solution for day 5.
	day06    Runs the solution for day 6.
	... and so on!

`,
	Run: func(cmd *cobra.Command, args []string) {
		m := map[string]func() error{
			"day04": solutions.SolveDay04,
		}
		d := args[0]
		solver, exists := m[d]
		if !exists {
			log.Fatalf("error: solution does not exist for %v\n", d)
			os.Exit(1)
		}
		if err := solver(); err != nil {
			log.Fatalf("error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// solveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// solveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
