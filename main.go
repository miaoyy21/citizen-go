package main

import (
	"citizen/tools"
	"log"
)

func main() {
	srcAssets := "/Users/miaojingyi/Documents/dev/go/src/citizen/assets"
	dstAssets := "/Users/miaojingyi/Documents/dev/flutter/citizen/assets"
	if err := tools.Run(srcAssets, dstAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}
	log.Println("自动化任务执行完成 ...")
}
