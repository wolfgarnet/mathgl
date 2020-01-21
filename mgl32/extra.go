package mgl32

import "github.com/go-gl/mathgl"

func (direction Vec3) RotateSkipY(r float32) Vec3 {
	x := direction[0]*mathgl.Cos32(r) - direction[2]*mathgl.Sin32(r)
	y := direction[0]*mathgl.Sin32(r) + direction[2]*mathgl.Cos32(r)

	return Vec3{x, direction[1], y}
}

func (v Vec3) VectorMultiplication(other Vec3) float32 {
	return v[0]* other[0] + v[1]* other[1] + v[2]* other[2]
}

func (v Vec3) SquaredLength() float32 {
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


func (m Mat4) ColVec3(c int) (v Vec3) {
	v[0] = m[c*4]
	v[1] = m[c*4+1]
	v[2] = m[c*4+2]

	return
}