package converter

import (
	"errors"
	"reflect"
	"strings"
)

var (
	ErrorAttributesNotInspected     = errors.New("O tipo não tem atributo que seja inspecionável, apenas é aceito objeto e não uma estrutura!")
	ErrorNotFoundAtrributesInHeader = errors.New("Não foram encontrados os atributos do header na estrutura de dados informada!")
)

func getAttributes(m interface{}) (map[string]reflect.Type, error) {

	typ := reflect.TypeOf(m)
	if isSlice(m) {
		typ = getTypeInSlice(m)
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	attrs := make(map[string]reflect.Type)
	if typ.Kind() != reflect.Struct {
		return nil, ErrorAttributesNotInspected
	}
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			attrs[p.Name] = p.Type
		}
	}
	return attrs, nil
}

func convert(entitys []*Entity, out interface{}) error {
	var countsAttr int = 0
	slice := reflect.Indirect(reflect.ValueOf(out))
	for _, ent := range entitys {
		newData := reflect.New(getTypeInSlice(out)).Elem()
		for _, value := range ent.Fields {
			mapa, _ := getAttributes(out)
			for nome, _ := range mapa {
				if value.Key == nome {
					countsAttr++
					newData.FieldByName(nome).SetString(strings.TrimSpace(value.Value))
				}
			}
		}
		slice.Set(reflect.Append(slice, newData))
	}

	if countsAttr == 0 {
		return ErrorNotFoundAtrributesInHeader
	}
	return nil

}

func getValueAndType(in interface{}) (reflect.Value, reflect.Type) {
	value := reflect.ValueOf(in)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value, value.Type()
}

func isSlice(in interface{}) bool {
	value := reflect.ValueOf(in)
	if value.Elem().Kind() != reflect.Slice {
		return false
	}
	return true

}

func getTypeInSlice(in interface{}) reflect.Type {
	return reflect.TypeOf(in).Elem().Elem()
}
