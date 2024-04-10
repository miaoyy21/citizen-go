package main

import (
	"citizen/tools"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("%s \n", err.Error())
	}

	srcAssets := filepath.Join(dir, "assets")
	dstAssets := strings.Replace(filepath.Join(dir, "assets"), "go/src", "flutter", -1)

	// 文件类型转换
	//convert(srcAssets)

	// 动作解析
	generate(srcAssets, dstAssets)

	log.Println("自动化任务执行完成 ...")
}

func convert(srcAssets string) {
	// 将MAP4文件转为GIF文件
	if err := tools.Mp4SwfToGif(srcAssets); err != nil {
		log.Fatalf("%s \n", err.Error())
	}

	// 将OGG文件转为WAV文件
	if err := tools.OggToWav(srcAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}

	// 解析Audio资源文件的时长信息
	if err := tools.RunAudio(srcAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

func generate(srcAssets, dstAssets string) {
	// 解析声效文件
	if err := tools.RunSounds(srcAssets, dstAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}

	// 解析动画文件
	if err := tools.RunSkills(srcAssets, dstAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}

	// 对图片进行抗锯齿处理
	//if err := tools.ChangeDefinition(dstAssets); err != nil {
	//	log.Fatalf("%s", err.Error())
	//}
}
