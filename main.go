package main

import (
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
)

func main() {
	log.Println("Process Started ...")

	filename := "002"
	err := ffmpeg.Input(fmt.Sprintf("./%s.mp4", filename)).
		Output(fmt.Sprintf("./%s.gif", filename), ffmpeg.KwArgs{"r": "12"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		log.Fatalf("ERROR :: %s \n", err.Error())
	}
	log.Println("Process Finished ...")
}
