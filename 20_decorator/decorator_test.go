package decorator

import "testing"

func TestLogDecorator(t *testing.T) {
	var c Calc = &ConcreteCalc{}
	d := NewLogDecorator(c)
	res := d.Add(3, 5)
	t.Logf("Result: %v", res)
}
