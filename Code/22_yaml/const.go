package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type student struct {
	name string `yaml:"name"`
	age  int    `yaml:"age"`
}

func main() {
	var s student
	config, err := ioutil.ReadFile("student.yaml")
	if err!=nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal(config, &s)
	if err!=nil {
		println(err)
	}

	println(s.name)
	println(s.age)

}
