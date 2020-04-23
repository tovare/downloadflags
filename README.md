# Download Flags of the world

[![Go Report Card](https://goreportcard.com/badge/github.com/tovare/downloadflags)

A small script to download all flags in svg-format from Wikipedia and optionally
convert them into small 40px wide png files for use as Shapes in Tableau.

The conversion relies on rsvg-convert which on a mac may be installed using

    brew install rsvg-convert.

By default it will download and convert the flags into the img/ folder, and create the folder if it doesn´t exist.

The countries of the world change their name from time to time, if you get an error just
google for the correct information and update the the csv-file.

