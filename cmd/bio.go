package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
	"gopkg.in/russross/blackfriday.v2"

	log "github.com/Sirupsen/logrus"
)
type Bio struct {
	Name     string
	Bio      string
	ShortBio string
	Links    string
	Images []string
}

func buildBio() error {
	bio := &Bio{}
	shortbb, err := fileAsHTML(filepath.Join("slides", "bio", "shortbio.md"))
	if err != nil {
		log.Println("Error getting short bio: ", err)
		return err
	}
	bio.ShortBio = string(blackfriday.Run(shortbb))

	biobb, err := fileAsHTML(filepath.Join("slides", "bio", "bio.md"))
	if err != nil {
		log.Println("Error getting bio: ", err)
		return err
	}
	bio.Bio = string(blackfriday.Run(biobb))

	linksbb, err := fileAsHTML(filepath.Join("slides", "bio", "links.md"))
	if err != nil {
		log.Println("Error getting links: ", err)
		return err
	}
	bio.Links = string(blackfriday.Run(linksbb))

	htmlpath := filepath.Join("public", "bio.html")
	hf, err := os.Create(htmlpath)
	if err != nil {
		log.Println("Error creating file: ", err)
		return err
	}
	imagespath := filepath.Join("slides", "images", "bio")
	images, err := ioutil.ReadDir(imagespath)
    if err != nil {
		log.Println("Error reading bio images directory: ", err)
		return err
	}
	for _, i := range images {
		imgsrc := filepath.Join("images", "bio", i.Name())
		bio.Images = append(bio.Images, imgsrc)
	}	

	templ := template.Must(template.New("").ParseFiles("templates/bio.tmpl"))
	err = templ.ExecuteTemplate(hf, "bio.tmpl", bio)
	if err != nil {
		log.Println("Error executing template:", err)
	}
	hf.Close()

	return nil
}

func fileAsHTML(path string) ([]byte, error) {
	sb, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println("Error opening bio file:", err)
		return []byte{}, err
	}
	bb, err := ioutil.ReadAll(sb)
	if err != nil {
		log.Println("Error reading bio file:", err)
		return []byte{}, err
	}
	return bb, err
}