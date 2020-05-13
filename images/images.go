package images

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/viniciusbds/navio/utilities"
)

var (
	images     = make(map[string]*Image)
	baseImages = make(map[string]*Image)
)

func init() {
	baseImages["alpine"] = NewImage("alpine", "alpine", "v3.11", "2.7M", "http://dl-cdn.alpinelinux.org/alpine/v3.11/releases/x86_64/alpine-minirootfs-3.11.6-x86_64.tar.gz")
	baseImages["busybox"] = NewImage("busybox", "busybox", "v4.0", "1.5M", "https://raw.githubusercontent.com/teddyking/ns-process/4.0/assets/busybox.tar")
	baseImages["ubuntu"] = NewImage("ubuntu", "ubuntu", "v20.04", "90.0M", "http://cloud-images.ubuntu.com/minimal/releases/focal/release/ubuntu-20.04-minimal-cloudimg-amd64-root.tar.xz")

	if utilities.FileExists(utilities.Imagescsv) {
		readImages()
	}
}

// Image ...
type Image struct {
	name    string
	base    string
	version string
	size    string
	url     string
}

// NewImage ...
func NewImage(name string, base string, version string, size string, url string) *Image {
	return &Image{
		name:    name,
		base:    base,
		version: version,
		size:    size,
		url:     url,
	}
}

// ToStr ...
func (i *Image) ToStr() string {
	return fmt.Sprintf("%s\t\t%s\t\t%s\t\t%s", i.name, i.base, i.version, i.size)
}

func getImage(name string) *Image {
	if images[name] != nil {
		return images[name]
	}
	if baseImages[name] != nil {
		return baseImages[name]
	}
	return nil
}

// InsertImage receive a containerName and the respective baseImage and creates a new
// Image and insert it on [images map]. Also update the csv file that store all containerImages
func InsertImage(containerName, baseImage string) {
	baseImg := getImage(baseImage)
	newImg := NewImage(containerName, baseImage, baseImg.version, baseImg.size, baseImg.url)
	images[containerName] = newImg
	updateImageCSV()
}

// RemoveImage receive a containerName and remove the respective image from the [images map].
// Also update the csv file that store all containerImages
func RemoveImage(containerName string) {
	delete(images, containerName)
	updateImageCSV()
}

// updateImageCSV read the entire current [images map] and store all Images
// in the utilities.Imagescsv file
func updateImageCSV() {
	err := os.RemoveAll(utilities.Imagescsv)
	if err != nil {
		log.Fatalln("Couldn't remove the csv file", err)
	}

	csvfile, err := os.Create(utilities.Imagescsv)
	if err != nil {
		log.Fatalln("Couldn't create the csv file", err)
	}

	w := csv.NewWriter(csvfile)

	for _, v := range images {
		newimage := []string{v.name, v.base, v.version, v.size, v.url}
		err = w.Write(newimage)
		if err != nil {
			fmt.Println("Erro aque")
		}
	}
	w.Flush()
	csvfile.Close()
}

// readImages read the entire utilities.Imagescsv file and store all available containerImages
// in the [images map]
func readImages() {
	images = make(map[string]*Image)

	csvfile, err := os.Open(utilities.Imagescsv)
	if err != nil {
		l.Log("Error", err.Error())
	}

	r := csv.NewReader(csvfile)
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		images[record[0]] = NewImage(record[0], record[1], record[2], record[3], record[4])
	}
}
