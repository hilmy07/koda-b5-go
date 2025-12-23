package internals

import "math"

func hitungLuas(r float64) float64 {
	return math.Pi * r * r
}

func hitungKeliling(r float64) float64 {
	return 2 * math.Pi * r
}

func HitungLuasdanKeliling(r float64) (float64, float64) {
	luas := hitungLuas(r)
	keliling := hitungKeliling(r)
	return luas, keliling
}