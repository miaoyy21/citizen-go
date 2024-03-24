package main

import (
	"citizen/tools"
	"log"
)

func main() {
	//srcAssets := "/Users/miaojingyi/Documents/dev/go/src/citizen/assets"
	//dstAssets := "/Users/miaojingyi/Documents/dev/flutter/citizen/assets"
	//if err := tools.RunSkills(srcAssets, dstAssets); err != nil {
	//	log.Fatalf("%s", err.Error())
	//}
	//
	//log.Println("自动化任务执行完成 ...")

	dstAssets := "/Users/miaojingyi/Documents/dev/go/src/citizen/assets/audio"
	if err := tools.RunAudio(dstAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}

	log.Println("自动化任务执行完成 ...")

}
