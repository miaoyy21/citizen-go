package main

import (
	"citizen/tools"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	srcAssets := "/Users/miaojingyi/Documents/dev/go/src/citizen/assets"
	dstAssets := "/Users/miaojingyi/Documents/dev/flutter/citizen/assets"
	if err := tools.ParseAnimations(srcAssets, dstAssets); err != nil {
		log.Fatalf("%s", err.Error())
	}
	log.Println("解析动作帧动画成功 ...")

	sss()
}

func sss() {
	// 打开原始PNG文件
	file, err := os.Open("original.png")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// 解码PNG文件
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("failed to decode image: %s", err)
	}

	// 在图片的 (x, y) 处修改颜色
	x := 100
	y := 100
	rgba := img.(*image.NRGBA)
	rgba.Set(x, y, color.RGBA{R: 255, G: 0, B: 0, A: 255})

	// 创建一个新的PNG文件并保存修改后的图片
	outFile, err := os.Create("modified.png")
	if err != nil {
		log.Fatalf("failed to create output file: %s", err)
	}
	defer outFile.Close()

	// 将修改后的图片编码为PNG格式并保存到文件
	if err := png.Encode(outFile, img); err != nil {
		log.Fatalf("failed to encode image: %s", err)
	}

	log.Println("Image successfully modified and saved as modified.png")
}
