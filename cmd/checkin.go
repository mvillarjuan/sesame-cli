/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// checkinCmd represents the checkin command
var checkinCmd = &cobra.Command{
	Use:   "checkin",
	Short: "Check-In",
	Long: "Check-In",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("checkin called")
	},
}

func init() {
	rootCmd.AddCommand(checkinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
