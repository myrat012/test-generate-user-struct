package generator

import (
	"os"
	"path/filepath"
	"text/template"
)

func UserGenerate(tmpl, outFilePath string, fields interface{}) error {
	t, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		return err
	}
	file, err := os.Create(filepath.Join("../../internal/model", outFilePath))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = t.Execute(file, fields)
	if err != nil {
		return err
	}

	return nil
}
