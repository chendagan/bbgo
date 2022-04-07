package types

import "math"

type Float64Slice []float64

func (s *Float64Slice) Push(v float64) {
	*s = append(*s, v)
}

func (s *Float64Slice) Pop(i int64) (v float64) {
	v = (*s)[i]
	*s = append((*s)[:i], (*s)[i+1:]...)
	return v
}

func (s Float64Slice) Max() float64 {
	m := -math.MaxFloat64
	for _, v := range s {
		m = math.Max(m, v)
	}
	return m
}

func (s Float64Slice) Min() float64 {
	m := math.MaxFloat64
	for _, v := range s {
		m = math.Min(m, v)
	}
	return m
}

func (s Float64Slice) Sum() (sum float64) {
	for _, v := range s {
		sum += v
	}
	return sum
}

func (s Float64Slice) Mean() (mean float64) {
	return s.Sum() / float64(len(s))
}

func (s Float64Slice) Tail(size int) Float64Slice {
	length := len(s)
	if length <= size {
		win := make(Float64Slice, length)
		copy(win, s)
		return win
	}

	win := make(Float64Slice, size)
	copy(win, s[length-size:])
	return win
}

func (s Float64Slice) Diff() (values Float64Slice) {
	for i, v := range s {
		if i == 0 {
			values.Push(0)
			continue
		}
		values.Push(v - s[i-1])
	}
	return values
}

func (s Float64Slice) PositiveValuesOrZero() (values Float64Slice) {
	for _, v := range s {
		values.Push(math.Max(v, 0))
	}
	return values
}

func (s Float64Slice) NegativeValuesOrZero() (values Float64Slice) {
	for _, v := range s {
		values.Push(math.Min(v, 0))
	}
	return values
}

func (s Float64Slice) AbsoluteValues() (values Float64Slice) {
	for _, v := range s {
		values.Push(math.Abs(v))
	}
	return values
}

func (s Float64Slice) MulScalar(x float64) (values Float64Slice) {
	for _, v := range s {
		values.Push(v * x)
	}
	return values
}

func (s Float64Slice) DivScalar(x float64) (values Float64Slice) {
	for _, v := range s {
		values.Push(v / x)
	}
	return values
}

func (s Float64Slice) ElementwiseProduct(other Float64Slice) (values Float64Slice) {
	for i, v := range s {
		values.Push(v * other[i])
	}
	return values
}

func (s Float64Slice) Dot(other Float64Slice) float64 {
	return s.ElementwiseProduct(other).Sum()
}

func (s Float64Slice) Normalize() (values Float64Slice) {
	sum := s.Sum()
	for _, v := range s {
		values = append(values, v/sum)
	}
	return values
}
