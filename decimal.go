package decimal

type NumberFlt float64

type Number struct {
	Value     NumberFlt
	Precision int
	Hi        int
	Lo        int
}

func New(val NumberFlt, precision int) *Number {
	hi := int(val * NumberFlt(precision*10))
	return &Number{
		Value:     val,
		Precision: precision,
		Hi:        hi,
		Lo:        0,
	}
}

func (n *Number) Decimal() float64 {
	return float64(n.Hi) / float64(n.Precision*10)
}
