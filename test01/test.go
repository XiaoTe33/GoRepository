package main

import (
	"fmt"
	"image"
	"io"

	//"image/png"
	"os"
)

func main() {
	r := io.Reader()
	image.Decode()
	f, err := os.Open("屏幕截图 2022-11-17 150123.png")
	if err != nil {
		panic(err)
	}
	img, formatName, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(formatName)
	fmt.Println(img.Bounds())
	fmt.Println(img.ColorModel())
}

//func main01() {
//	r := gin.Default()
//	r.GET("/user", func(c *gin.Context) {
//		fin, err := os.Open("屏幕截图 2022-11-17 150123.png") //打开文件
//		if err != nil {
//			fmt.Println(err)
//		}
//		defer fin.Close()
//
//		buf := make([]byte, 1024) //开辟1024个字节的slice作为缓冲
//		for {
//			n, _ := fin.Read(buf) //读文件
//			if n == 0 {           //0表示已经到文件结束
//				break
//			}
//		}
//
//		c.JSON(200, string(buf))
//	})
//	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
//}
