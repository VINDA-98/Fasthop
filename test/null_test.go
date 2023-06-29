package test

import (
	"fmt"
	"math"
	"testing"
)

// @Title  test
// @Description  MyGO
// @Author  WeiDa  2023/6/29 11:43
// @Update  WeiDa  2023/6/29 11:43

type CalculatorVolume interface {
	calc() float64
}

// 正方体
type cube struct {
	// 边长
	length float64
}

// 正方体的体积计算
func (c *cube) calc() float64 {
	return c.length * c.length * c.length
}

// 长方体
type cuboid struct {
	// 长
	length float64
	// 宽
	width float64
	// 高
	height float64
}

// 长方体的体积计算
func (c *cuboid) calc() float64 {
	return c.length * c.width * c.height
}

// 圆柱体
type cylinder struct {
	// 直径
	diameter float64
	// 高度
	height float64
}

// 圆柱体的体积计算
func (c *cylinder) calc() float64 {
	return math.Pi * (c.diameter / 2) * (c.diameter / 2) * c.height
}

// 计算某个物体的体积
func calculateVolume(c CalculatorVolume) float64 {
	return c.calc()
}

func TestNull(t *testing.T) {
	truckSize := 0.0
	// 声明空接口类型变量materials，存放各种不同体积的家具
	var materials []CalculatorVolume
	materials = append(materials, &cube{12.5})
	materials = append(materials, &cuboid{25, 13, 60})
	materials = append(materials, &cylinder{5, 25.3})
	// 遍历materials切片，依次计算每个家具的体积，并相加求和
	for _, singleMaterial := range materials {
		truckSize += calcSize(singleMaterial)
	}
	fmt.Println(truckSize)
}

// 计算某个物体的体积
func calcSize(c CalculatorVolume) float64 {
	return c.calc()
}
