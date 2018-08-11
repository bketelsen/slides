// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new slide repository.",
	Long: `Init creates a new slide repository by cloning
the slide assets required for building and serving the
slide decks.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := doInit(args[0])
		if err != nil {
			log.Println("Error creating slide repository:", err)
			return
		}
	},
}

func doInit(directoryName string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	newRepoPath := filepath.Join(wd, directoryName)
	log.Println("New slide repository location:", newRepoPath)
	if err := checkNewPath(newRepoPath); err != nil {
		return err
	}
	if err := cloneTemplate(newRepoPath); err != nil {
		return err
	}
	return nil
}

func cloneTemplate(target string) error {
	cmd := "git"
	// TODO: consider making this a defaulted parameter?
	args := []string{"clone", "https://github.com/bketelsen/slides-template", target}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		return err
	}
	gitDir := filepath.Join(target, ".git")
	if err := os.RemoveAll(gitDir); err != nil {
		log.Println("Error removing .git directory:", err)
		return err
	}
	return nil
}

func checkNewPath(dir string) error {
	fi, err := os.Stat(dir)
	// this is the happy path, directory doesn't exist
	if os.IsNotExist(err) {
		return nil
	}
	if !fi.IsDir() {
		return errors.New("file with same name exists")
	}
	if fi.IsDir() {
		return errors.New("directory exists")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
