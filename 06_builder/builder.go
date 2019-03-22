package builder

import (
	"fmt"
)

type Product interface {
	Show()
}

//Builder 是生成器接口
type Builder interface {
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetProduct() Product
}

type Director struct {
	builder Builder
}

// NewDirector ...
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

//Construct Product
func (d *Director) Construct() Product {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	d.builder.BuildPart3()
	return d.builder.GetProduct()
}

type Product1 struct {
	result string
}

func (p *Product1) Show() {
	fmt.Println(p.result)
}

type Builder1 struct {
	product Product1
}

func (b *Builder1) BuildPart1() {
	b.product.result += "1"
}

func (b *Builder1) BuildPart2() {
	b.product.result += "2"
}

func (b *Builder1) BuildPart3() {
	b.product.result += "3"
}

func (b *Builder1) GetProduct() Product {
	return &b.product
}

type Product2 struct {
	result int
}

func (p *Product2) Show() {
	fmt.Println(p.result)
}

type Builder2 struct {
	product Product2
}

func (b *Builder2) BuildPart1() {
	b.product.result += 1
}

func (b *Builder2) BuildPart2() {
	b.product.result += 2
}

func (b *Builder2) BuildPart3() {
	b.product.result += 3
}

func (b *Builder2) GetProduct() Product {
	return &b.product
}
