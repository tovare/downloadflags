package main

/*
 * A small script to download all flags in svg-format from Wikipedia and optionally
 * convert them into small 40px wide png files for use as Shapes in Tableau.
 *
 * The conversion relies on rsvg-convert which on a mac may be installed using
 * brew install rsvg-convert.
 */

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/tovare/stringtable"
)

func main() {

	var (
		destinationFolderParam string
		convertIcon            bool
	)

	flag.StringVar(&destinationFolderParam, "destination", "img/", "Destination directory")
	flag.BoolVar(&convertIcon, "converticon", true, "Convert to icon, relies on rsvg-convert, install this first (brew install rsvg-convert)")
	flag.Parse()

	_ = os.Mkdir(destinationFolderParam, 0700)

	table, err := stringtable.ReadCSV("Country_Flags.csv", ',')
	if err != nil {
		return
	}
	err = downloadAll(table, destinationFolderParam)
	if err != nil {
		log.Fatal(err)
	}

	if convertIcon {
		err = convertAll(table, destinationFolderParam)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// downloadAll flag files from Wikipedia to disk.
func downloadAll(table stringtable.Table, destinationFolder string) (err error) {
	cm := table.Colmap()
	for _, row := range table[1:] {
		imageURL := row[cm["ImageURL"]]
		destinationFileName := destinationFolder + row[cm["Images File Name"]]
		//fmt.Printf("Reading %v and writing to %v", imageURL, destinationFileName)

		resp, err := http.Get(imageURL)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			fmt.Printf("Failed to load %v with status code %v \n", imageURL, resp.StatusCode)
		}
		defer resp.Body.Close()
		out, err := os.Create(destinationFileName)
		if err != nil {
			return err
		}
		defer out.Close()
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}
	}
	return
}

// convertAll svg files into small 40px wide png-files.
func convertAll(table stringtable.Table, folder string) (err error) {
	cm := table.Colmap()
	for _, row := range table[1:] {
		sourceFile := folder + row[cm["Images File Name"]]
		destinationFile := strings.Replace(sourceFile, ".svg", ".png", -1)
		//	rsvg-convert -w 40 Flag_of_Norway.svg -o Flag_og_Norway.png
		cmd := exec.Command("rsvg-convert", "-w", "40", sourceFile, "-o", destinationFile)
		err = cmd.Run()
		if err != nil {
			return
		}
	}
	return
}
