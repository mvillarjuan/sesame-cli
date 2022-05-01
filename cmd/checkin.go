/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkinCmd represents the checkin command
var checkinCmd = &cobra.Command{
	Use:   "checkin",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		sesameUrl := viper.GetString("SESAMEURL")
		sesameEmail := viper.GetString("SESAMEEMAIL")
		sesamePassword := viper.GetString("SESAMEPASSWORD")
		sessionId, err := getSessionId(sesameUrl, sesameEmail, sesamePassword)
		if err != nil {
			panic(err)
		}
		userId, err := getUserId(sesameUrl, sessionId)
		if err != nil {
			panic(err)
		}
		err = doCheckIn(sesameUrl, sessionId, userId)
		if err == nil {
			fmt.Println("Checked in")
		}
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
func doCheckIn(sesameUrl, sessionId, userId string) error {
	values := map[string]string{"origin": "web"}
	jsonData, _ := json.Marshal(values)
	req, _ := http.NewRequest(http.MethodPost, sesameUrl+"/api/v3/employees/"+userId+"/check-in", bytes.NewBuffer(jsonData))
	req.Header.Add("Authorization", "Bearer "+sessionId)
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return errors.New("error: could not connect to server")
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("error: unexpected server response")
	}
	defer res.Body.Close()
	return nil
}
