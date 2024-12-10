package abstract

import (
	"fmt"
	"math"
)

type Circle struct {
	Shape
	ColoredShape
	Resizable
	radius float64
	//Анонимная функция.
	//test   func(t string) string
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	Shape
	ColoredShape
	Resizable
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Shape interface {
	Area() float64
}

func PrintArea(shape Shape) {
	fmt.Println(shape.Area())
}

func PrintColor(coloredShape ColoredShape, shape Shape) {
	color, area := coloredShape.Describe(shape)
	fmt.Println(color, area)
}

type ColoredShape struct {
	Shape
	color string
}

// Task six.
func (coloredShape ColoredShape) Describe(shape Shape) (string, float64) {
	return coloredShape.color, shape.Area()
}

func Test(t string) string {
	fmt.Println("test")
	fmt.Println(t)
	return t + "TRRR"
}

// Task seven.
type ComplexShape struct {
	common []Shape
}

func (complexShape ComplexShape) TotalArea() float64 {
	i := 0.0
	for _, y := range complexShape.common {
		i += y.Area()
	}
	return i
}

func PrintTotalArea(complexShape ComplexShape) {
	fmt.Println(complexShape.TotalArea())
}

// Task eight.
type Resizable interface {
	Resize(factor float64)
}

func (c *Circle) Resize(factor float64) {
	c.radius *= factor
}

func (r *Rectangle) Resize(factor float64) {
	r.width *= factor
	r.height *= factor
}

func Run() {
	circle := &Circle{radius: 1, ColoredShape: ColoredShape{color: "blue"}}
	rectangle := &Rectangle{width: 5, height: 3, ColoredShape: ColoredShape{color: "red"}}
	common := ComplexShape{common: []Shape{circle, rectangle}}
	//Ниже два примера для анонимной функций.
	//circle.test = Test
	//fmt.Println(circle.test("fgf"))
	PrintArea(circle)
	PrintArea(rectangle)
	PrintColor(circle.ColoredShape, circle)
	PrintColor(rectangle.ColoredShape, rectangle)
	fmt.Println(common.TotalArea())

	//Task eight. Добавляем значения.
	circle.Resize(2)
	rectangle.Resize(1.1)

	PrintArea(circle)
	PrintArea(rectangle)
}
