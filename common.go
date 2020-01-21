package mathgl

import (
	//"github.com/go-gl/mathgl/mgl32"
	"math"
)

/*func Mat4To64(m mgl32.Mat4) mgl64.Mat4 {
	return mgl64.Mat4 {
		float64(m[0]), float64(m[1]), float64(m[2]), float64(m[3]),
		float64(m[4]), float64(m[5]), float64(m[6]), float64(m[7]),
		float64(m[8]), float64(m[9]), float64(m[10]), float64(m[11]),
		float64(m[12]), float64(m[13]), float64(m[14]), float64(m[15]),
	}
}

func Vec3To64(v mgl32.Vec3) mgl64.Vec3 {
	return mgl64.Vec3{float64(v[0]), float64(v[1]), float64(v[2])}
}*/

func Min(x, y float32) float32 {
	if y < x {
		return y
	}

	return x
}

func Max(x, y float32) float32 {
	if y > x {
		return y
	}

	return x
}

func Abs32(x float32) float32 {
	return float32(math.Abs(float64(x)))
}

func Mod32(x, y float32) float32 {
	return float32(math.Mod(float64(x), float64(y)))
}

func Atan232(x, y float32) float32 {
	return float32(math.Atan2(float64(x), float64(y)))
}

func Pow32(x, y float32) float32 {
	return float32(math.Pow(float64(x), float64(y)))
}

func Cos32(x float32) float32 {
	return float32(math.Cos(float64(x)))
}

func Sin32(x float32) float32 {
	return float32(math.Sin(float64(x)))
}
