package minitask1

import "math"

func HitungLuas(r float32) float32 {
	return math.Pi * r * r
}

func HitungKeliling(r float32) float32 {
	return 2 * math.Pi * r
}

func LuasKeliling(r float32) (float32, float32) {
	var luas float32 = math.Pi * r * r
	var keliling float32 = 2 * math.Pi * r
	return luas, keliling
}
