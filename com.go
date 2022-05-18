package rng

const maxu uint64 = 1<<64 - 1

// Modn returns random integer from half-open interval [0, n) for n ≥ 2,
// or returns n-1 for n < 2. This is more uniform than Get() % n.
func Modn(n uint64) uint64 {
	k := n - 1
	if n&k == 0 { // n=0 or power of 2 ?
		if n > 1 {
			return Get() & k
		}
		return k
	}

	v := Get()

	if int64(n) < 0 { // n > 2^63 ?
		for v >= n {
			v = Get()
		}
		return v
	}

	// mostly avoid one division
	if v > maxu-n {
		// largest multiple of n < 2^64
		lastn := maxu - maxu%n
		for v >= lastn {
			v = Get()
		}
	}
	return v % n
}
