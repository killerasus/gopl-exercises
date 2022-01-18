// Surface calcula uma renderização SVG de uma função de superfície 3D.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // tamanho do canvas em pixels
	cells         = 100                 // número de células da grade
	xyrange       = 30.0                // intervalos dos eixos (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels por unidade x ou y
	zscale        = height * 0.4        // pixels por unidade z
	angle         = math.Pi / 6         // ângulo dos eixos x, y (=30°)
	min           = -0.6
	max           = 1
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // seno(30°), cosseno(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; strokewidth: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, t := corner(i+1, j+1)

			r := uint8(Lerp(0x00, 0xFF, t))
			g := uint8(0x00)
			b := uint8(Lerp(0xFF, 0x00, t))

			//c := color.RGBA{r, g, b, uint8(0)}
			color := fmt.Sprintf("%02x%02x%02x", r, g, b)

			if IsValid([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style=\"fill:#%s\"/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) { // Encontra o ponto (x,y) no canto da célula (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5) // Calcula a altura z da superfície
	z := f(x, y)                            // Faz uma projeção isométrica de (x,y,z) sobre (sx,sy) do canvas SVG 2D

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, GetTLerp(min, max, z)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distância de (0,0)
	return math.Sin(r) / r
}

// Gets the value in the interval [0, 1], where 0 = v0, 1 = v1
func Lerp(v0, v1, t float64) float64 {
	return (1-t)*v0 + t*v1
}

// Gets the value of t in the [0,1] interval, where x is a value in the [v0, v1] interval
func GetTLerp(v0, v1, x float64) float64 {
	return (x - v0) / (v1 - v0)
}

func IsValid(s []float64) bool {
	for _, y := range s {
		if math.IsNaN(y) || math.IsInf(y, 0) {
			return false
		}
	}
	return true
}
