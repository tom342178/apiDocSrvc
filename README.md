# apiDocSrvc
Provide embeded doc server

## Attribution

Forked from https://github.com/flowchartsman/swaggerui  
Modified swaggerui.go to support gorilla mux.  
Moved all handler code from main to the module.  

## Updated the embedded swaggerui code

Clone the code and go run generate.go  
This will update the embed directory with the latest swaggerui code.  
Push the results back to the repository.  

## Example usage

'package main

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/flowchartsman/swaggerui"
)

//go:embed app-api-spec.yaml
var spec []byte

func main() {

// If using the openApi generate=, the following code is 
// generated in the routes.go file and will have been called already.

router := mux.NewRouter().StrictSlash(true)

// The following calls out to a module that addes a fileserver
// and therefore must be placed after all other paths are declared. 

apiDocSrvc.apiDoc(spec, *router)

log.Fatal(http.ListenAndServe(":8080", router))  
  
}'  







