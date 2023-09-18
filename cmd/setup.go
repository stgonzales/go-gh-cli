/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up a project",
	Long: `This command will clone a repo from github and install dependencies.
	
it need to be passed a full repo url, like: stgonzales/gh-cli`,
	Run: func(cmd *cobra.Command, args []string) {

		repo, _ := cmd.Flags().GetString("repo")
		path, _ := cmd.Flags().GetString("path")

		// fmt.Printf("cloning %s into %s ...", repo, path)

		fullRepoUrl := fmt.Sprintf("https://github.com/%s.git", repo)

		out, err := Commander(fmt.Sprintf("git clone %s %s", fullRepoUrl, path))

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(out[:]))

		_, err = Commander("vtest -v")

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	setupCmd.Flags().StringP("repo", "r", "", "The repo to clone, like: stgonzales/gh-cli")
	setupCmd.Flags().StringP("path", "p", "./", "The path to clone the repo, like: /home/user/projects (default: current directory)")

	setupCmd.MarkFlagRequired("repo")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Commander(c string) ([]byte, error) {

	out, err := exec.Command("bash", "-c", c).Output()

	return out, err
}
