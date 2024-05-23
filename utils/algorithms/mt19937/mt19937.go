package mt19937

// Thanks github.com/seehuhn/mt19937
const (
	n         = 312
	m         = 156
	matrixA   = 0xB5026F5AA96619E9
	upperMask = 0xFFFFFFFF80000000
	lowerMask = 0x7FFFFFFF
)

type MT19937 struct {
	mt  [n]uint64
	mti int
}

func New() *MT19937 {
	return &MT19937{
		mt:  [312]uint64{},
		mti: n + 1,
	}
}

func (mt *MT19937) Seed(seed uint64) {
	mt.mt[0] = seed
	for i := 1; i < n; i++ {
		mt.mt[i] = (6364136223846793005*(mt.mt[i-1]^(mt.mt[i-1]>>62)) + uint64(i))
	}
	mt.mti = n
}

func (mt *MT19937) NextUint64() uint64 {
	mag01 := [2]uint64{0x0, matrixA}
	var x uint64

	if mt.mti >= n {
		if mt.mti == n+1 {
			mt.Seed(5489)
		}

		for i := 0; i < n-m; i++ {
			x = (mt.mt[i] & upperMask) | (mt.mt[i+1] & lowerMask)
			mt.mt[i] = mt.mt[i+m] ^ (x >> 1) ^ mag01[x&0x1]
		}
		for i := n - m; i < n-1; i++ {
			x = (mt.mt[i] & upperMask) | (mt.mt[i+1] & lowerMask)
			mt.mt[i] = mt.mt[i+(m-n)] ^ (x >> 1) ^ mag01[x&0x1]
		}
		x = (mt.mt[n-1] & upperMask) | (mt.mt[0] & lowerMask)
		mt.mt[n-1] = mt.mt[m-1] ^ (x >> 1) ^ mag01[x&0x1]

		mt.mti = 0
	}

	x = mt.mt[mt.mti]
	mt.mti++

	x ^= (x >> 29) & 0x5555555555555555
	x ^= (x << 17) & 0x71D67FFFEDA60000
	x ^= (x << 37) & 0xFFF7EEE000000000
	x ^= (x >> 43)

	return x
}
