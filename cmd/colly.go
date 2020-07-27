package cmd

import "github.com/spf13/cobra"

var (
	CollyCommand = &cobra.Command{
		Use: "colly",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)
