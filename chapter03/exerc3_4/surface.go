// Surface calcula uma renderização SVG de uma função de superfície 3D.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	color         = "grey"              // drawing color
	cells         = 100                 // grid cells
	xyrange       = 30.0                // axis interval (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per unity x ou y
	zscale        = height * 0.4        // pixels per unity z
	angle         = math.Pi / 6         // x, y axis angle (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cosin(30°)

func main() {
	http.HandleFunc("/", createImage)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func createImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	m_width, _ := strconv.Atoi(r.FormValue("width"))
	m_height, _ := strconv.Atoi(r.FormValue("height"))
	m_color := r.FormValue("color")

	if m_width <= 0 {
		m_width = width
	}
	if m_height <= 0 {
		m_height = height
	}
	if len(m_color) == 0 {
		m_color = color
	}

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; strokewidth: 0.7' "+
		"width='%d' height='%d'>\n", m_color, m_width, m_height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if IsValid([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}

	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64) { // Finds (x,y) of cell (i,j) corner
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5) // Computes surface z height
	z := f(x, y)                            // Isometric projection of (x,y,z) on (sx,sy) from SVG 2D canvas
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func IsValid(s []float64) bool {
	for _, y := range s {
		if math.IsNaN(y) || math.IsInf(y, 0) {
			return false
		}
	}
	return true
}
