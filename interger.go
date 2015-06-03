package interger

const (
	MaxUint uint = ^uint(0)
	MinUint uint = 0
	MaxInt  int  = int(MaxUint >> 1)
	MinInt  int  = -MaxInt - 1
)
