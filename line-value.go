package ichimoku

type ValueLine struct {
	valLine float64
	isNil   bool
}

func (n *ValueLine) Value() interface{} {
	if n.isNil {
		return nil
	}
	return n.valLine
}
func (n *ValueLine) SetValue(v float64) {
	n.valLine = v
	n.isNil = false
}

func NewValue(x float64) ValueLine {
	return ValueLine{x, false}
}

func NewValueLineNil() ValueLine {
	return ValueLine{0, true}
}
