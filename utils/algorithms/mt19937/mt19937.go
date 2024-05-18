package mt19937

// Thanks github.com/seehuhn/mt19937
const (
	n         = 312
	m         = 156
	notSeeded = n + 1

	hiMask uint64 = 0xffffffff80000000
	loMask uint64 = 0x000000007fffffff

	matrixA uint64 = 0xB5026F5AA96619E9
)

// MT19937 is the structure to hold the state of one instance of the
// Mersenne Twister PRNG.  New instances can be allocated using the
// mt19937.New() function.  MT19937 implements the rand.Source
// interface and rand.New() from the math/rand package can be used to
// generate different distributions from a MT19937 PRNG.
//
// This class is not safe for concurrent accesss by different
// goroutines.  If more than one goroutine accesses the PRNG, the
// callers must synchronise access using sync.Mutex or similar.
type MT19937 struct {
	state []uint64
	index int
}

// New allocates a new instance of the 64bit Mersenne Twister.
// A seed can be set using the .Seed() or .SeedFromSlice() methods.
func New() *MT19937 {
	res := &MT19937{
		state: make([]uint64, n),
		index: notSeeded,
	}
	return res
}

// Seed uses the given 64bit value to initialise the generator state.
// This method is part of the rand.Source interface.
func (mt *MT19937) Seed(seed uint64) {
	x := mt.state
	x[0] = uint64(seed)
	for i := uint64(1); i < n; i++ {
		x[i] = 6364136223846793005*(x[i-1]^(x[i-1]>>62)) + i
	}
	mt.index = n
}

// NextUint64 generates a (pseudo-)random 64bit value.  The output can be
// used as a replacement for a sequence of independent, uniformly
// distributed samples in the range 0, 1, ..., 2^64-1.
func (mt *MT19937) NextUint64() uint64 {
	x := mt.state
	if mt.index >= n {
		if mt.index == notSeeded {
			mt.Seed(5489) // default seed, as in mt19937-64.c
		}
		for i := 0; i < n-m; i++ {
			y := (x[i] & hiMask) | (x[i+1] & loMask)
			x[i] = x[i+m] ^ (y >> 1) ^ ((y & 1) * matrixA)
		}
		for i := n - m; i < n-1; i++ {
			y := (x[i] & hiMask) | (x[i+1] & loMask)
			x[i] = x[i+(m-n)] ^ (y >> 1) ^ ((y & 1) * matrixA)
		}
		y := (x[n-1] & hiMask) | (x[0] & loMask)
		x[n-1] = x[m-1] ^ (y >> 1) ^ ((y & 1) * matrixA)
		mt.index = 0
	}
	y := x[mt.index]
	y ^= (y >> 29) & 0x5555555555555555
	y ^= (y << 17) & 0x71D67FFFEDA60000
	y ^= (y << 37) & 0xFFF7EEE000000000
	y ^= (y >> 43)
	mt.index++
	return y
}
