package main

import (
	"fmt"

	ibpt "github.com/joaopaulo-bertoncini/go-ibpt"
)

func main() {
	clientProduct, err := ibpt.NewClientProduct()
	if err != nil {
		panic(err)
	}
	request := &ibpt.Request{
		Token:           "your_token",
		CNPJ:            "your_cnpj",
		Code:            "21050010",
		UF:              "SP",
		EX:              0,
		InternalCode:    "245",
		Description:     "Sorvete",
		UnitMeasurement: "PT",
		Value:           10,
		Gtin:            "5445",
	}

	response, _ := clientProduct.Send(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response)
}
