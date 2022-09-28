package domain

const L float64 = 10.0
const dt float64 = 0.01
const CUTOFF float64 = 2.0
const CL2 float64 = CUTOFF * CUTOFF
const RC2 float64 = 1 / CL2
const RC6 float64 = RC2 * RC2 * RC2
const RC12 float64 = RC6 * RC6
const C0 float64 = -4.0 * (RC12 - RC6)

func adjustPeriodic(dx, dy, dz *float64) {
	const LH float64 = L * 0.5
	if *dx < -LH {
		*dx += L
	}
	if *dx > LH {
		*dx -= L
	}
	if *dy < -LH {
		*dy += L
	}
	if *dy > LH {
		*dy -= L
	}
	if *dz < -LH {
		*dz += L
	}
	if *dz > LH {
		*dz -= L
	}
}
