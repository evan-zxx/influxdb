// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: values.gen.go.tmpl

package gen

type floatConstantValuesSequence struct {
	v float64
}

func NewFloatConstantValuesSequence(v float64) FloatValuesSequence {
	return &floatConstantValuesSequence{
		v: v,
	}
}

func (g *floatConstantValuesSequence) Reset() {
}

func (g *floatConstantValuesSequence) Write(vs []float64) {
	for i := 0; i < len(vs); i++ {
		vs[i] = g.v
	}
}

type integerConstantValuesSequence struct {
	v int64
}

func NewIntegerConstantValuesSequence(v int64) IntegerValuesSequence {
	return &integerConstantValuesSequence{
		v: v,
	}
}

func (g *integerConstantValuesSequence) Reset() {
}

func (g *integerConstantValuesSequence) Write(vs []int64) {
	for i := 0; i < len(vs); i++ {
		vs[i] = g.v
	}
}

type unsignedConstantValuesSequence struct {
	v uint64
}

func NewUnsignedConstantValuesSequence(v uint64) UnsignedValuesSequence {
	return &unsignedConstantValuesSequence{
		v: v,
	}
}

func (g *unsignedConstantValuesSequence) Reset() {
}

func (g *unsignedConstantValuesSequence) Write(vs []uint64) {
	for i := 0; i < len(vs); i++ {
		vs[i] = g.v
	}
}

type stringConstantValuesSequence struct {
	v string
}

func NewStringConstantValuesSequence(v string) StringValuesSequence {
	return &stringConstantValuesSequence{
		v: v,
	}
}

func (g *stringConstantValuesSequence) Reset() {
}

func (g *stringConstantValuesSequence) Write(vs []string) {
	for i := 0; i < len(vs); i++ {
		vs[i] = g.v
	}
}

type booleanConstantValuesSequence struct {
	v bool
}

func NewBooleanConstantValuesSequence(v bool) BooleanValuesSequence {
	return &booleanConstantValuesSequence{
		v: v,
	}
}

func (g *booleanConstantValuesSequence) Reset() {
}

func (g *booleanConstantValuesSequence) Write(vs []bool) {
	for i := 0; i < len(vs); i++ {
		vs[i] = g.v
	}
}

type floatArrayValuesSequence struct {
	v  []float64
	vi int
}

func NewFloatArrayValuesSequence(v []float64) FloatValuesSequence {
	return &floatArrayValuesSequence{
		v: v,
	}
}

func (g *floatArrayValuesSequence) Reset() {
	g.vi = 0
}

func (g *floatArrayValuesSequence) Write(vs []float64) {
	var (
		v  = g.v
		vi = g.vi
	)
	for i := 0; i < len(vs); i++ {
		if vi >= len(v) {
			vi = 0
		}
		vs[i] = v[vi]
		vi += 1
	}
	g.vi = vi
}

type integerArrayValuesSequence struct {
	v  []int64
	vi int
}

func NewIntegerArrayValuesSequence(v []int64) IntegerValuesSequence {
	return &integerArrayValuesSequence{
		v: v,
	}
}

func (g *integerArrayValuesSequence) Reset() {
	g.vi = 0
}

func (g *integerArrayValuesSequence) Write(vs []int64) {
	var (
		v  = g.v
		vi = g.vi
	)
	for i := 0; i < len(vs); i++ {
		if vi >= len(v) {
			vi = 0
		}
		vs[i] = v[vi]
		vi += 1
	}
	g.vi = vi
}

type unsignedArrayValuesSequence struct {
	v  []uint64
	vi int
}

func NewUnsignedArrayValuesSequence(v []uint64) UnsignedValuesSequence {
	return &unsignedArrayValuesSequence{
		v: v,
	}
}

func (g *unsignedArrayValuesSequence) Reset() {
	g.vi = 0
}

func (g *unsignedArrayValuesSequence) Write(vs []uint64) {
	var (
		v  = g.v
		vi = g.vi
	)
	for i := 0; i < len(vs); i++ {
		if vi >= len(v) {
			vi = 0
		}
		vs[i] = v[vi]
		vi += 1
	}
	g.vi = vi
}

type stringArrayValuesSequence struct {
	v  []string
	vi int
}

func NewStringArrayValuesSequence(v []string) StringValuesSequence {
	return &stringArrayValuesSequence{
		v: v,
	}
}

func (g *stringArrayValuesSequence) Reset() {
	g.vi = 0
}

func (g *stringArrayValuesSequence) Write(vs []string) {
	var (
		v  = g.v
		vi = g.vi
	)
	for i := 0; i < len(vs); i++ {
		if vi >= len(v) {
			vi = 0
		}
		vs[i] = v[vi]
		vi += 1
	}
	g.vi = vi
}

type booleanArrayValuesSequence struct {
	v  []bool
	vi int
}

func NewBooleanArrayValuesSequence(v []bool) BooleanValuesSequence {
	return &booleanArrayValuesSequence{
		v: v,
	}
}

func (g *booleanArrayValuesSequence) Reset() {
	g.vi = 0
}

func (g *booleanArrayValuesSequence) Write(vs []bool) {
	var (
		v  = g.v
		vi = g.vi
	)
	for i := 0; i < len(vs); i++ {
		if vi >= len(v) {
			vi = 0
		}
		vs[i] = v[vi]
		vi += 1
	}
	g.vi = vi
}