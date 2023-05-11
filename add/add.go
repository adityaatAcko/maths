package add

type Nums struct {
	Num1 float64
	Num2 float64
}

func (n Nums) Add() float64 {
	return (n.Num1 + n.Num2)
}
