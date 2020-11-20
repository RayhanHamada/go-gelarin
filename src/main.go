package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "gelarin",
		Short: "An awesome repository url storage",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Sukses")
		},
	}

	rootCmd.Execute()
}
