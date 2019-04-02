package decorator

import "log"

type Calc interface {
	Add(a, b int) int
}

type ConcreteCalc struct{}

func (*ConcreteCalc) Add(a, b int) int {
	return a + b
}

type LogDecorator struct {
	Calc
}

func NewLogDecorator(c Calc) Calc {
	return &LogDecorator{
		Calc: c,
	}
}

func (d *LogDecorator) Add(a, b int) int {
	log.Printf("Before adding operation")
	sum := d.Calc.Add(a, b)
	log.Printf("After adding operation")
	return sum
}
