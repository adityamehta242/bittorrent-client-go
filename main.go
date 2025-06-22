package main

import (
	"log"
	"os"

	"bittorrent-client/torrentfile"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	// Open the torrent file
	file, err := os.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse the torrent file
	tf, err := torrentfile.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	// Download the torrent
	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
