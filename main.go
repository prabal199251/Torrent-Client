package main

import (
	"log"
	"os"

	torrentfile "github.com/prabal199251/Torrent-Client/torrentFile"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	TorrentFile, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = TorrentFile.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
