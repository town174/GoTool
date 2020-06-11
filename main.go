package main

import (
	"flag"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"image/png"
	"os"
)

func main() {
	//获取输入信息
	var content = flag.String("content", "123456", "init content")
	var width = flag.Int("width", 350, "init width")
	var height = flag.Int("height", 70, "init height")
	flag.Parse()
	println("INFO HAS VALUE",*content)
	//生成一个一维码
	cs,_ := code128.Encode(*content)
	//创建一个输出文件
	file, _ := os.Create("qr3.png")
	defer file.Close()
	//保存为图片
	barcode,_:=barcode.Scale(cs,*width,*height)
	_ = png.Encode(file, barcode)
}
