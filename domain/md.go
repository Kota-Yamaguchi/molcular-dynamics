package domain

import (
	"fmt"
	"math"
)

type MdSystem struct {
	time float32
}

func NewMdSystem() *MdSystem {
	return &MdSystem{0.0}
}

func (md *MdSystem) makeconf() *Atoms {
	density := 0.50
	s := 1.0 / math.Pow(density*0.25, 1.0/3.0)
	hs := s * 0.5
	is := int(L / s)
	atoms := NewAtoms()
	for iz := 0.0; iz < float64(is); iz++ {
		for iy := 0.0; iy < float64(is); iy++ {
			for ix := 0.0; ix < float64(is); ix++ {
				atoms.AddAtom(ix*s, iy*s, iz*s)
				atoms.AddAtom(ix*s+hs, iy*s, iz*s)
				atoms.AddAtom(ix*s, iy*s+hs, iz*s)
				atoms.AddAtom(ix*s, iy*s, iz*s+hs)
			}
		}
	}
	atoms.SetInitialVelocity(1.0)
	return atoms
}

func (md *MdSystem) calculateForce(atoms *Atoms) {
	size := len(atoms.atoms)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			dx := atoms.atoms[j].x - atoms.atoms[i].x
			dy := atoms.atoms[j].y - atoms.atoms[i].y
			dz := atoms.atoms[j].z - atoms.atoms[i].z
			adjustPeriodic(&dx, &dy, &dz)
			r2 := (dx*dx + dy*dy + dz*dz)
			if r2 > CL2 {
				continue
			}
			r6 := r2 * r2 * r2
			r12 := r6 * r6
			// レナードジョーンズポテンシャルの微分が力になっている。ここではさらに方向ごとの力を計算する前処理分の1/rもかけている
			df := (24*r6 - 48) / (r12 * r2) * dt
			atoms.atoms[i].acceleration(df*dx, df*dy, df*dz)
			atoms.atoms[j].deceleration(df*dx, df*dy, df*dz)
		}
	}
}

func (md *MdSystem) updatePosition(atoms *Atoms) {
	dt2 := dt * 0.5
	for _, a := range atoms.atoms {
		a.changePosition(a.px*dt2, a.py*dt2, a.pz*dt2)
	}
}

func (md *MdSystem) periodic(atoms *Atoms) {
	for _, a := range atoms.atoms {
		if a.x < 0.0 {
			a.changeXPosition(L)
		}
		if a.y < 0.0 {
			a.changeYPosition(L)
		}
		if a.z < 0.0 {
			a.changeZPosition(L)
		}
		if a.x > L {
			a.changeXPosition(-L)
		}
		if a.y > L {
			a.changeYPosition(-L)
		}
		if a.z > L {
			a.changeZPosition(-L)
		}

	}
}

func (md *MdSystem) calculate(atoms *Atoms) {
	md.updatePosition(atoms)
	md.calculateForce(atoms)
	md.updatePosition(atoms)
	md.periodic(atoms)
	md.time = md.time + float32(dt)
}

func (md *MdSystem) Run() {

	atoms := md.makeconf()
	observer := NewObserver(*atoms)
	const STEP = 10000
	const EPOCH = 100
	for i := 0; i < STEP; i++ {
		if i%EPOCH == 0 {
			k := observer.CalcKineticEnergy()
			v := observer.CalcPotentialEnergy()
			fmt.Printf("STEP: %d \n", i)
			fmt.Printf("Kinetic Energy : %f \n", k)
			fmt.Printf("Potential Energy : %f \n", v)
			exportCdview(*atoms, i)
		}
		md.calculate(atoms)
	}
}
