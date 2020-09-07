package animals

import (
	"fmt"

	"github.com/gitchander/liteut/jsoncho"
)

type Animal struct {
	V isAnimal
}

type isAnimal interface {
	isAnimal()
}

func (Gopher) isAnimal() {}
func (Lion) isAnimal()   {}
func (Snake) isAnimal()  {}
func (Rabbit) isAnimal() {}

func (a Animal) MarshalJSON() ([]byte, error) {
	return jsoncho.Choice2.MarshalChoice(&choiceAnimal{&a})
}

func (a *Animal) UnmarshalJSON(data []byte) error {
	return jsoncho.Choice2.UnmarshalChoice(data, &choiceAnimal{a})
}

type choiceAnimal struct {
	*Animal
}

var _ jsoncho.Choice = &choiceAnimal{}

func (p *choiceAnimal) GetType() (t string, err error) {
	v := p.Animal.V
	switch v.(type) {
	case *Gopher:
		t = "gopher"
	case *Lion:
		t = "lion"
	case *Snake:
		t = "snake"
	case *Rabbit:
		t = "rabbit"
	default:
		err = fmt.Errorf("invalid type %T", v)
	}
	return t, err
}

func (p *choiceAnimal) SetType(t string) error {
	var v isAnimal
	switch t {
	case "gopher":
		v = new(Gopher)
	case "lion":
		v = new(Lion)
	case "snake":
		v = new(Snake)
	case "rabbit":
		v = new(Rabbit)
	default:
		return fmt.Errorf("invalid type %s", t)
	}
	p.Animal.V = v
	return nil
}

func (p *choiceAnimal) Value() interface{} {
	return p.Animal.V
}

type Gopher struct {
	A string
}

type Lion struct {
	B int
}

type Snake struct {
	C bool
}

type Rabbit struct {
	D []int
}
