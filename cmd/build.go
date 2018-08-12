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
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Prepare your slides for publishing",
	Long: `Build prepares your slides for publishing by 
merging them into html templates and publishing in the
/public directory which is suitable for publishing with
a static web server or CDN.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Building Slides")
		templ := template.Must(template.New("").ParseFiles("templates/publish.tmpl"))
		files, err := ioutil.ReadDir("slides")
		if err != nil {
			log.Println("Error loading slides", err)
			return
		}
		var slides []Slide
		for _, f := range files {
			if strings.Contains(f.Name(), ".md") {
				// copy file to /public
				path := f.Name()
				htmlpath := strings.Replace(path, ".md", ".html", -1)
				err := copy(filepath.Join("slides", path), filepath.Join("public", path))
				if err != nil {
					log.Println("Error copying markdown", err)
					return
				}
				publicMarkdown, err := os.OpenFile(filepath.Join("public", path), os.O_RDWR|os.O_CREATE, 0755)
				if err != nil {
					return
				}

				slide, err := FromReader(publicMarkdown)
				slide.HTMLFile = htmlpath
				slide.MarkdownFile = path
				htmlpath = filepath.Join("public", htmlpath)
				hf, err := os.Create(htmlpath)
				if err != nil {
					log.Println("Error creating file: ", err)
					return
				}

				_, err = publicMarkdown.WriteAt([]byte(slide.Source), 0)
				if err != nil {
					log.Println("Error writing file: ", err)
					return
				}

				slides = append(slides, *slide)
				err = templ.ExecuteTemplate(hf, "publish.tmpl", slide)
				if err != nil {
					log.Println("Error executing template:", err)
				}
				hf.Close()
			}
		}
		// copy images
		err = copyfolder(filepath.Join("slides", "images"), filepath.Join("public", "images"))
		if err != nil {
			log.Println("Error copying images: ", err)
			return
		}
		// now make index
		for _, slide := range slides {
			fmt.Printf("%s: %s\n", slide.MarkdownFile, slide.Title)
		}
		log.Println("Building Index")

		templ = template.Must(template.New("").ParseFiles("templates/root.tmpl"))
		htmlpath := filepath.Join("public", "index.html")
		hf, err := os.Create(htmlpath)
		if err != nil {
			log.Println("Error creating file: ", err)
			return
		}
		pretty.Println(slides)
		err = templ.ExecuteTemplate(hf, "root.tmpl", slides)
		if err != nil {
			log.Println("Error executing template:", err)
		}
		log.Println("Build completed")
		return
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func copy(src string, dst string) error {
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	return err
}
func copyfolder(source string, dest string) (err error) {

	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			err = copyfolder(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = copyfile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func copyfile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

func faReplace(text string) string {
	re := regexp.MustCompile(`@fa\[([a-zA-Z0-9\s]*)\]`)
	s := re.ReplaceAllString(text, "<i class='fas fa-$1'></i>")
	return s
}
