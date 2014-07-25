package tween

var Easing = NewEasingFuncs()

func NewEasingFuncs() *EasingFuncs {
	return &EasingFuncs{
		Linear: NewLinear(),
	}
}

type EasingFuncs struct {
	Linear    *Linear
	Quadratic *Quadratic
	Cubic     *Cubic
	Quartic   *Quartic
}

type Linear struct {
	None func(k float64) float64
}

func NewLinear() *Linear {
	return &Linear{
		None: LinearNone,
	}
}

func LinearNone(k float64) float64 {
	return k
}

type Quadratic struct {
	In    func(k float64) float64
	Out   func(k float64) float64
	InOut func(k float64) float64
}

type Cubic struct {
	In    func(k float64) float64
	Out   func(k float64) float64
	InOut func(k float64) float64
}

type Quartic struct {
	In    func(k float64) float64
	Out   func(k float64) float64
	InOut func(k float64) float64
}
