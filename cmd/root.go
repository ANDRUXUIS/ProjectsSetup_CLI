package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/fatih/color"
)

// SystemCommand Global Functions
func SystemCommand(command string)  {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}

}

// Function for get inputs from terminal

// func cliInputs() string  {
// 	prompt := promptui.Prompt{
// 		Label:    "Project Name",
// 	}

// 	result, err := prompt.Run()

// 	if err != nil {
// 		fmt.Printf("Prompt failed %v\n", err)
// 		os.Exit(1)
// 	}

// 	return result
// }

// 	this part is for the system command

func promptUI()  {
	// this part is for the prompt UI for select
	cellTemplate := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\U000027A4 {{ .| cyan | bold }}",
		Inactive: "  {{ .| white | bold }}",
		Selected: color.GreenString("\U00002713 ") + color.GreenString("Project setup for:  ") + "{{ .  | faint }}",
	}

	prompt := promptui.Select{
		Label: "Select A Project setup",
		Items: []string{"Nextjs", "React-Typescript", "ReactNative-Expo-Typescript", "Firebase Functions"},
		Templates: cellTemplate,	
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	
	switch (result) {
	case "React-Typescript":
		ReactTypescript()
	case "Nextjs":
		Nextjs()
	case "ReactNative-Expo-Typescript":
		// implement function later
	case "Firebase Functions":
		FirebaseFunctions()
	}

}

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pcli",
	Short: "A Cli for Project setup in various technologies",
	Long: `Nothing to see here`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		promptUI()
	},
}
// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".pcli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".pcli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
