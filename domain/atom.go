package domain

import (
	"math"
	"math/rand"
	"strconv"
)

type Atom struct {
	x, y, z    float64
	px, py, pz float64
}

func (a *Atom) deceleration(vx, vy, vz float64) {
	a.px -= vx
	a.py -= vy
	a.pz -= vz
}
func (a *Atom) acceleration(vx, vy, vz float64) {
	a.px += vx
	a.py += vy
	a.pz += vz
}

func (a *Atom) changePosition(x, y, z float64) {
	a.x += x
	a.y += y
	a.z += z
}
func (a *Atom) changeXPosition(x float64) {
	a.x += x
}
func (a *Atom) changeYPosition(y float64) {
	a.y += y
}
func (a *Atom) changeZPosition(z float64) {
	a.z += z
}

func (a *Atom) toString(count int) string {
	return "C " + strconv.FormatFloat(a.x, 'f', 2, 64) + " " + strconv.FormatFloat(a.y, 'f', 2, 64) + " " + strconv.FormatFloat(a.z, 'f', 2, 64) + "\n"
}

type Atoms struct {
	atoms []*Atom
}

func NewAtoms() *Atoms {
	var atom []*Atom
	atoms := Atoms{atom}
	return &atoms
}

func (as *Atoms) AddAtom(x, y, z float64) {
	a := Atom{x, y, z, 0.0, 0.0, 0.0}
	as.atoms = append(as.atoms, &a)
}

func (as *Atoms) SetInitialVelocity(v0 float64) {

	//原子全体の平均速度をまず初期化する。
	aveVx := 0.0
	aveVy := 0.0
	aveVz := 0.0

	for _, a := range as.atoms {
		ud := rand.Float64()
		var z float64 = ud*2.0 - 1.0
		var phi float64 = 2.0 * ud * math.Pi
		vx := v0 * math.Sqrt(1.0-z*z) * math.Cos(phi)
		vy := v0 * math.Sqrt(1.0-z*z) * math.Sin(phi)
		vz := v0 * z
		a.acceleration(vx, vy, vz)

		//原仕事の速度を足して最後に平均する。
		aveVx += vx
		aveVy += vy
		aveVz += vz
	}
	//atomの速度足して平均と原子全体の速度が出る
	size := len(as.atoms)
	aveVx = aveVx / float64(size)
	aveVy = aveVy / float64(size)
	aveVz = aveVz / float64(size)

	for _, a := range as.atoms {
		a.deceleration(aveVx, aveVy, aveVz)
	}

}
