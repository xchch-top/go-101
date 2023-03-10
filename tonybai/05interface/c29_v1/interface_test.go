package c29_v1

import "testing"

func Test_interface(t *testing.T) {
	animals := []QuackableAnimal{new(Duck), new(Dog), new(Bird)}
	for _, animal := range animals {
		AnimalQuackInForest(animal)
	}
}

type QuackableAnimal interface {
	Quack()
}

type Duck struct{}

func (Duck) Quack() {
	println("duck quack!")
}

type Dog struct{}

func (Dog) Quack() {
	println("dog quack!")
}

type Bird struct{}

func (Bird) Quack() {
	println("bird quack!")
}

func AnimalQuackInForest(a QuackableAnimal) {
	a.Quack()
}
