package main

import (
	"net/http"

	"aplicacaoWeb/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
