package main

/*
import (
	// "fmt"
	"fmt"
	"math"
	"os"

	"github.com/fogleman/gg"
)

func quickSort(arr []int, low int, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func sortq(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

type LSystem struct {
	axiom string
	rules map[string]string
	depth int
}

func NewLSystem(axiom string, rules map[string]string, depth int) *LSystem {
	return &LSystem{axiom: axiom,
		rules: rules,
		depth: depth}
}

func (l *LSystem) Generate() string {
	current := l.axiom
	for l.depth > 0 {
		current = l.applyRules(current)
		l.depth--
	}
	return current
}

func (l *LSystem) applyRules(state string) string {
	nextState := ""
	for _, symbol := range state {
		if rule, ok := l.rules[string(symbol)]; ok {
			nextState += rule
		} else {
			nextState += string(symbol)
		}
	}
	return nextState
}

func main() { // Define the L-System grammar (Koch curve)
	rules := map[string]string{
		// "F": "F--F--F",
		"F": "F+F--F+F",
	}

	// Create the initial state (axiom) and depthaxiom := "F"depth := 4
	axiom := "F"
	axiom = "F--F--F"
	depth := 5 // Create a new L-System instance
	// Create a new L-System instance
	lsystem := NewLSystem(axiom, rules, depth)

	// Generate the resulting pattern
	result := lsystem.Generate()

	// Visualize the generated pattern
	img, err := visualizeKochCurve(result, 600, depth)
	if err != nil {
		fmt.Println("Error creating image:", err)
		os.Exit(1)
	} // Save the image to a file
	err = img.SavePNG("koch_curve.png")
	if err != nil {
		fmt.Println("Error saving image:", err)
		os.Exit(1)
	}
	fmt.Println("Image saved as koch_curve.png")
}

func visualizeKochCurve(pattern string, width int, depth int) (*gg.Context, error) {
	dc := gg.NewContext(width, width)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetRGBA(0.5, 1, 1, 0.5)
	x := 10.0
	y := 10.0
	angle := 0.0
	distance := 1.0
	step := 0.0
	dc.SetLineWidth(1)
	points := [][]float64{{x, y}}
	for _, symbol := range pattern {
		switch symbol {
		case 'F':
			x += distance * math.Cos(angle)
			y += distance * math.Sin(angle)
			points = append(points, []float64{x - distance*math.Cos(angle), y - distance*math.Sin(angle)})
			// dc.DrawLine(x, y, x-distance*math.Cos(angle), y-distance*math.Sin(angle))
			// dc.DrawLine(10, 10, 20, 50)
			dc.Stroke()
		case '+':
			angle -= math.Pi / 3.0
			step += 0.751
		case '-':
			angle += math.Pi / 3.0
			step -= 0.751
		default:
			return nil, fmt.Errorf("invalid symbol '%v'", symbol)
		}
	}
	fmt.Println(points)
	minx, maxx := math.Inf(1), math.Inf(-1)
	miny, maxy := math.Inf(1), math.Inf(-1)
	for _, point := range points {
		for i, val := range point {
			if i == 0 {
				if val < minx {
					minx = val
				}
				if val > maxx {
					maxx = val
				}
			}
			if i == 1 {
				if val < miny {
					miny = val
				}
				if val > maxy {
					maxy = val
				}
			}
		}
	}
	fmt.Println(minx, maxx, miny, maxy)
	for i := 1; i < len(points); i++ {

		x1 := (points[i-1][0] / maxx) * .9 * float64(width)
		y1 := (points[i-1][1] / maxy) * .9 * float64(width)
		x2 := (points[i][0] / maxx) * .9 * float64(width)
		y2 := (points[i][1] / maxy) * .9 * float64(width)

		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()

	}
	return dc, nil
}
*/
