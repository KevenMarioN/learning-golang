package formas_test

import (
	"math"
	. "test-avancado/formas"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{2, 10}
		expectedArea := float64(20)
		receivedArea := rectangle.Area()

		if expectedArea != receivedArea {
			t.Fatalf("Excpected %f.02, but recived %f.02", expectedArea, receivedArea)
		}
	})

	t.Run("Circule", func(t *testing.T) {
		circle := Circle{10}
		expectedArea := float64(math.Pi * 100)
		receivedArea := circle.Area()
		if expectedArea != receivedArea {
			t.Fatalf("Excpected %f.02, but recived %f.02", expectedArea, receivedArea)
		}
	})
}
