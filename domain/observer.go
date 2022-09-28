package domain

type Observer struct {
	atoms Atoms
}

func NewObserver(atoms Atoms) *Observer {
	return &Observer{atoms: atoms}
}

func (observer *Observer) CalcKineticEnergy() float64 {
	var k float64 = 0.0
	for _, a := range observer.atoms.atoms {
		k += a.px * a.px
		k += a.py * a.py
		k += a.pz * a.pz
	}
	k = k / float64(len(observer.atoms.atoms)) * (0.5)
	return k
}

func (observer *Observer) CalcPotentialEnergy() float64 {
	v := 0.0
	size := len(observer.atoms.atoms)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			dx := observer.atoms.atoms[j].x - observer.atoms.atoms[i].x
			dy := observer.atoms.atoms[j].y - observer.atoms.atoms[i].y
			dz := observer.atoms.atoms[j].z - observer.atoms.atoms[i].z
			adjustPeriodic(&dx, &dy, &dz)
			r2 := (dx*dx + dy*dy + dz*dz)
			if r2 > CL2 {
				continue
			}
			r6 := r2 * r2 * r2
			r12 := r6 * r6
			v += 4.0 * (1/r12 - 1/r6)
		}
	}
	v = v / float64(size)
	return v
}
