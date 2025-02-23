package adapter

import (
	"fmt"
)

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) GetWeather() {
	fmt.Println("Nice wheather")
}
