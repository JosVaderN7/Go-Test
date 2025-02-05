package smi

import (
	"math"
	"testing"
)
func TestPerimeter(t *testing.T) {
	rectangle := Rectangulo{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangulo{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
type Triangle struct {
	Base   float64
	Height float64
}
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5c
}

type Shape interface {
	Area() float64
}

type Circle struct{
	Radius float64
}

func (c Circle) Area()float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangulo struct{
	width float64 
	height float64
}

func (r Rectangulo) Area() float64 {
	return r.width * r.height
}

func Area (rect Rectangulo)float64{
    
	return rect.height * rect.width
}

func Perimeter (rect Rectangulo) float64{
    
	return  2* (rect.height + rect.width)
}

