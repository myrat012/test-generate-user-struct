package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"tstgeneration/internal/generator"
)

type Fields struct {
	First  string
	Second string
}

func main() {
	fmt.Println("started")
	stringTemplate, err := os.ReadFile("../../assets/templates/AddUpdateLogs")
	if err != nil {
		panic(err)
	}

	astInFile, err := parser.ParseFile(
		token.NewFileSet(),
		"../../internal/model/users.go",
		nil,
		parser.ParseComments,
	)
	if err != nil {
		log.Fatalf("parse file: %v", err)
	}

	var allFild []Fields

	for _, decl := range astInFile.Decls {
		if typeSpec, ok := decl.(*ast.GenDecl); ok && typeSpec.Tok == token.TYPE {
			for _, spec := range typeSpec.Specs {
				if typeDecl, ok := spec.(*ast.TypeSpec); ok {

					if structType, ok := typeDecl.Type.(*ast.StructType); ok {

						fmt.Printf("Structure Name: %s\n", typeDecl.Name.Name)

						for _, field := range structType.Fields.List {
							for _, name := range field.Names {
								fmt.Printf("Field Name: %s\n", name.Name)
								allFild = append(allFild, Fields{
									First:  strings.ToLower(name.Name),
									Second: name.Name,
								})
							}
						}
					}
				}
			}
		}
	}

	err = generator.UserGenerate(string(stringTemplate), "generated_user.go", struct {
		Fields []Fields
	}{
		Fields: allFild,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("ended")
}
