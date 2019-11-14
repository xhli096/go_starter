package sin

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	// 设置图片大小
	const picSize = 300
	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, picSize, picSize))
	// 遍历每个像素点，并填充为白色
	for x := 0; x < picSize; x++ {
		for y := 0; y < picSize; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	// 用来控制x轴的长度
	for x := 0; x < picSize; x++ {
		// 计算sin函数的定义域，x/size 为0 ~ 1的float数，乘以2π后正好为0~2π
		s := float64(x) / picSize * 2 * math.Pi
		// sin的幅度为一半的像素。向下偏移一半像素并翻转
		y := picSize/2 - math.Sin(s)*picSize/2
		// 用黑色轨迹画图
		pic.SetGray(x, int(y), color.Gray{0})
	}

	// 创建sin图形
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}

	// 以png格式将数据写入文件中
	png.Encode(file, pic)

	file.Close()
}
