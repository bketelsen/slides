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
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a local webserver to view slides.",
	Long: `The serve command starts a local web server on port 8080
where you can view your slides.`,
	Run: func(cmd *cobra.Command, args []string) {
		r := NewServer()
		port := "8080"
		log.Info("Started http://0.0.0.0:8080")
		r.Run(fmt.Sprintf(":%s", port))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// devCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// devCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func NewServer() *gin.Engine {
	r := gin.Default()
	r.Static("/", "./public")
	return r
}
