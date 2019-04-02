package facade

import "log"

//API is facade interface of facade package
type API interface {
	Test()
}

//facade implement
type apiImpl struct {
	aModule AModuleAPI
	bModule BModuleAPI
}

func NewAPI() API {
	return &apiImpl{
		aModule: NewAModuleAPI(),
		bModule: NewBModuleAPI(),
	}
}

func (this *apiImpl) Test() {
	this.aModule.TestA()
	this.bModule.TestB()
}

//AModuleAPI ...
type AModuleAPI interface {
	TestA()
}

type aModuleImpl struct{}

//NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

func (*aModuleImpl) TestA() {
	log.Printf("A module running")
}

//BModuleAPI ...
type BModuleAPI interface {
	TestB()
}

type bModuleImpl struct{}

//NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

func (*bModuleImpl) TestB() {
	log.Printf("B module running")
}
