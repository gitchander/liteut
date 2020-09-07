package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/gitchander/liteut/jsoncho/animals"
)

func main() {
	testJsonOneof()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testJsonOneof() {

	as := []animals.Animal{
		{
			V: &animals.Gopher{
				A: "gopher",
			},
		},
		{
			V: &animals.Lion{
				B: 107,
			},
		},
		{
			V: &animals.Snake{
				C: true,
			},
		},
		{
			V: &animals.Rabbit{
				D: []int{1, 2, 3},
			},
		},
		// {
		// 	&Rabbit{
		// 		D: []int{1, 2, 3},
		// 	},
		// },
	}

	for _, a := range as {

		//data, err := json.Marshal(a)
		data, err := json.MarshalIndent(a, "", "\t")
		checkError(err)
		fmt.Printf("%s\n", data)

		var b animals.Animal

		err = json.Unmarshal(data, &b)
		checkError(err)

		data, err = json.MarshalIndent(b, "", "\t")
		checkError(err)
		fmt.Printf("%s\n", data)

		if !reflect.DeepEqual(a, b) {
			fmt.Println("Not equal!")
		}

		fmt.Println("---------------------------------------")
	}
}

type Rabbit struct {
	Name string
}

func (Rabbit) isAnimal() {}
