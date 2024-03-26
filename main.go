package main

import (
	"citizen/tools"
	"log"
)

func main() {
	srcAssets := "/Users/miaojingyi/Documents/dev/go/src/citizen/assets"
	dstAssets := "/Users/miaojingyi/Documents/dev/flutter/citizen/assets"

	// 解析动画文件
	if err := tools.RunSkills(srcAssets, dstAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}

	// 解析Audio资源文件的时长信息
	//if err := tools.RunAudio(srcAssets); err != nil {
	//	log.Fatalf("%s", err.Error())
	//}

	// 解析声效文件
	if err := tools.RunSounds(srcAssets, dstAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}

	// 对图片进行抗锯齿处理
	if err := tools.ChangeDefinition(dstAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}

	log.Println("自动化任务执行完成 ...")
}
