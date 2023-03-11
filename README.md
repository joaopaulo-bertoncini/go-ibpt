# Módulo go-ibpt

Este módulo auxilia na busca dos impostos incidentes sobre um determinado produto, utilizando o recurso gratuito oferecido pelo IBPT.

A Lei do Imposto na Nota (Lei nº 12.741/12, de 8 de dezembro de 2012) nasceu com o intuito de informar ao cidadão o quanto representa a parcela dos tributos que paga a cada compra realizada.

Assim, todo estabelecimento que efetuar vendas diretamente ao consumidor final está obrigado a incluir nos documentos fiscais ou equivalentes os impostos pagos, valores aproximados e percentuais.

Como consumidores finais incluem-se as pessoas físicas ou jurídicas que adquirem produtos ou serviços, por exemplo, para consumo próprio, materiais de uso ou consumo e ativo imobilizado.

As Microempresas e Empresas de Pequeno Porte optantes do Simples Nacional podem informar apenas a alíquota a que se encontram sujeitas nos termos do referido regime. Além disso, devem somar eventual incidência tributária anterior (IPI, substituição tributária, por exemplo).

## Instalação

importe no seu codigo 

```
import (
	ibpt "github.com/joaopaulo-bertoncini/go-ibpt"
)
```

Atualizar as dependências

```
go mod tidy
```

## Exemplo de uso

```
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
```

