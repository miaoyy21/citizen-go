package main

import (
	"citizen/tools"
	"log"
)

func main() {
	srcAssets := "/Users/miaojingyi/Documents/dev/go/src/citizen/assets"
	dstAssets := "/Users/miaojingyi/Documents/dev/flutter/citizen/assets"
	if err := tools.RunSkills(srcAssets, dstAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}

	log.Println("自动化任务执行完成 ...")

	if err := tools.RunAudio(srcAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}

	log.Println("自动化任务执行完成 ...")

	if err := tools.RunSounds(srcAssets, dstAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}

	// 对图片进行高清处理
	if err := tools.ChangeDefinition(dstAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
