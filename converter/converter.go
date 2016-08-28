package converter

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
)

var (
	ErrorDontHaveHeader = errors.New("Não foi informado o header no arquivo csv!")
	ErrorEmptyFile      = errors.New("O Arquivo csv eśtá vazio!")
)

type Converter struct{}

type Entity struct {
	Fields []*Field
}
type Field struct {
	Key   string
	Value string
}

func New() *Converter {
	return &Converter{}
}

func (r *Converter) ToJson(data []byte, out interface{}) error {
	err := json.Unmarshal(data, out)
	if err != nil {
		return err
	}
	return nil
}

func (r *Converter) ToCsv(data io.Reader, out interface{}) error {
	read := csv.NewReader(data)

	csvRows, err := read.ReadAll()
	if err != nil {
		return err
	}
	entitys, err := getEntitys(csvRows)
	if err != nil {
		return err
	}
	err = convert(entitys, out)
	if err != nil {
		return err
	}
	return nil
}

func getEntitys(rowns [][]string) ([]*Entity, error) {
	header := rowns[0]
	lenghtHeader := len(header)

	body := rowns[1:]
	numberLinesBody := len(body)

	if lenghtHeader == 0 {
		return nil, ErrorDontHaveHeader
	}
	if numberLinesBody == 0 {
		return nil, ErrorEmptyFile
	}

	entitys := []*Entity{}
	for i := 0; i < numberLinesBody; i++ {
		fields := []*Field{}
		for j := 0; j < lenghtHeader; j++ {
			f := &Field{}
			f.Value = body[i][j]
			f.Key = header[j]
			fields = append(fields, f)
		}
		entitys = append(entitys, &Entity{fields})
	}

	if len(entitys) == 0 {
		return nil, ErrorEmptyFile
	}

	return entitys, nil

}
