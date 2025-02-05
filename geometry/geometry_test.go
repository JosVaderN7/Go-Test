package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	}

// 	t.Run("Rectangle Area", func(t *testing.T) {
// 		rect := Rectangle{12.0, 6.0}
// 		checkArea(t, rect, 72.0)
// 	})

// 	t.Run("Circle Area", func(t *testing.T) {
// 		circ := Circle{10}
// 		checkArea(t, circ, 314.1592653589793)
// 	})
// }

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2g want %.2g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
