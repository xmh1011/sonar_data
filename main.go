package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xmh1011/sonar_data/cmd"
)

func main() {
	mainCmd := GetCommand()

	SonarDataCmd := cmd.Sonar_data()
	mainCmd.AddCommand(SonarDataCmd)

	if err := mainCmd.Execute(); err != nil {
		fmt.Printf("Error : %+v\n", err)
	}
}
func GetCommand() *cobra.Command {
	c := &cobra.Command{
		Use:  "sonar_data",
		Long: "generate sonar_data for test",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true,
			DisableNoDescFlag:   true,
			DisableDescriptions: true},
	}
	return c
}
