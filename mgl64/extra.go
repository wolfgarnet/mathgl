package mgl64

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

func (m Mat4) Mat4To32() mgl32.Mat4 {
	return mgl32.Mat4 {
		float32(m[0]), float32(m[1]), float32(m[2]), float32(m[3]),
		float32(m[4]), float32(m[5]), float32(m[6]), float32(m[7]),
		float32(m[8]), float32(m[9]), float32(m[10]), float32(m[11]),
		float32(m[12]), float32(m[13]), float32(m[14]), float32(m[15]),
	}
}

func (v Vec3) Vec3To32() mgl32.Vec3 {
	return mgl32.Vec3{float32(v[0]), float32(v[1]), float32(v[2])}
}

func (direction Vec3) RotateSkipY(r float64) Vec3 {
	x := direction[0]*math.Cos(r) - direction[2]*math.Sin(r)
	y := direction[0]*math.Sin(r) + direction[2]*math.Cos(r)

	return Vec3{x, direction[1], y}
}

func (v Vec3) VectorMultiplication(other Vec3) float64 {
	return v[0]* other[0] + v[1]* other[1] + v[2]* other[2]
}

func (v Vec3) SquaredLength() float64 {
	return v[0]* v[0] + v[1]* v[1] + v[2]* v[2]
}

func (m4 Mat4) MatrixMultiplyVector(v3 Vec3) Vec3 {
	r := Vec3{}
	r[0] = v3[0]*m4[0] + v3[1]*m4[1] + v3[2]*m4[2] + m4[3]
	r[1] = v3[0]*m4[4] + v3[1]*m4[5] + v3[2]*m4[6] + m4[7]
	r[2] = v3[0]*m4[8] + v3[1]*m4[9] + v3[2]*m4[10] + m4[11]

	return r
}

func (m *Mat4) TransformInverse(v Vec3) Vec3 {
	tmp := v
	tmp[0] -= m[12]
	tmp[1] -= m[13]
	tmp[2] -= m[14]

	return Vec3 {
		tmp[0]*m[0] + tmp[1]*m[1] + tmp[2]*m[2],
		tmp[0]*m[4] + tmp[1]*m[5] + tmp[2]*m[6],
		tmp[0]*m[8] + tmp[1]*m[9] + tmp[2]*m[10],
	}
}