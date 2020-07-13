package main

import (
	"fmt"
	"math"
)

func main() {
	var t Triangle = Triangle{3, 4, 5}
	fmt.Println(t.peri())
	fmt.Println(t.area())
	fmt.Println(t.a, t.b, t.c)

	var c Circle = Circle{4}
	fmt.Println(c.peri())
	fmt.Println(c.area())
	fmt.Println(c.radius)
	var s1 Shape
	s1 = t
	fmt.Println(s1.peri())
	fmt.Println(s1.area())

	var s2 Shape
	s2 = c
	fmt.Println(s2.peri())
	fmt.Println(s2.area())

	testShape(t)
	testShape(c)
	testShape(s1)

	getType(t)
	getType(c)
	getType(s1)

	getType2(t)
	getType2(c)
	getType2(s1)
}

func getType2(s Shape) {
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("三角形", ins.a, ins.b, ins.c)
	case Circle:
		fmt.Println("圆形", ins.radius)
	default:
		fmt.Println("不知道什么类型")
	}

}

func getType(s Shape) {
	if ins, ok := s.(Triangle); ok {
		fmt.Println("是三角形，三边是:", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Printf("是圆形,半径是%.2f\n", ins.radius)
	} else {
		fmt.Println("不知道是什么类型")
	}
}
func testShape(s Shape) {
	fmt.Printf("周长：%.2f,面积:%.2f\n", s.peri(), s.area())
}

//1、定义一个接口

type Shape interface {
	peri() float64 //形状的周长
	area() float64 //形状的面积

}

type Triangle struct {
	a, b, c float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

type Circle struct {
	radius float64
}

func (c Circle) peri() float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}
