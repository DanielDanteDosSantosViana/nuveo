package converter

import (
	"github.com/DanielDanteDosSantosViana/nuveo/models"
	"strings"
	"testing"
)

func Test_Deve_Converter_Json_Para_Estrutura_Informada(t *testing.T) {

	str := `{"nome": "daniel", "email": "danielvianarj@gmail.com", "sexo": "masculino", "idade" :"25"}`

	converter := New()
	data := models.Data{}
	err := converter.ToJson([]byte(str), &data)
	if err != nil {
		t.Errorf(" Error : deveria converter corretamente o json : %s , porém retornou o error : %v", str, err)
	}
	if data.Nome != "daniel" || data.Email != "danielvianarj@gmail.com" || data.Idade != "25" || data.Sexo != "masculino" {
		t.Errorf(" Error : deveria converter corretamente todos os atributos do json : %s , porém retornou apenas : %v", str, data)
	}
}

func Test_Deve_Converter_CSV_Para_Estrutura_Informada(t *testing.T) {
	str := `Nome,Email,Sexo,Idade
	daniel,danielvianarj@gmail.com,masculino,25`
	converter := New()
	data := []models.Data{}
	err := converter.ToCsv(strings.NewReader(str), &data)
	if err != nil {
		t.Errorf(" Error : deveria converter corretamente o json : %s , porém retornou o error : %v", str, err)
	}
	for _, val := range data {
		if val.Nome != "daniel" || val.Email != "danielvianarj@gmail.com" || val.Idade != "25" || val.Sexo != "masculino" {
			t.Errorf(" Error : deveria converter corretamente todos os atributos do csv : %s , porém retornou apenas : %v", str, val)
		}

	}

}

func Test_Deve_Converter_Lista_Json(t *testing.T) {

	str := `[{"nome": "daniel", "email": "danielvianarj@gmail.com", "sexo": "masculino", "idade" :"25"},{"nome": "daniel", "email": "danielvianarj@gmail.com", "sexo": "masculino", "idade" :"25"}]`

	converter := New()
	data := []models.Data{}
	err := converter.ToJson([]byte(str), &data)
	if err != nil {
		t.Errorf(" Error : deveria converter corretamente o json : %s , porém retornou o error : %v", str, err)
	}

	for _, val := range data {
		if val.Nome != "daniel" || val.Email != "danielvianarj@gmail.com" || val.Idade != "25" || val.Sexo != "masculino" {
			t.Errorf(" Error : deveria converter corretamente todos os atributos do json : %s , porém retornou apenas : %v", str, val)
		}

	}

}

func Test_Deve_Converter_Lista_CSV(t *testing.T) {

	str := `Nome,Email,Sexo,Idade
	daniel,danielvianarj@gmail.com,masculino,25
	nuveo,nuveo@gmail.com,masculino,100`

	converter := New()
	data := []models.Data{}
	err := converter.ToCsv(strings.NewReader(str), &data)
	if err != nil {
		t.Errorf(" Error : deveria converter corretamente o json : %s , porém retornou o error : %v", str, err)
	}

	if len(data) < 2 || len(data) > 2 {
		t.Errorf("Error: deveria conveter uma lista com duas entidades, conforme o formato passado : %s , porém retornou apenas (%v) entidade(s)", str, len(data))
	}

}

func Test_Deve_Retornar_Error_Para_Json_Invalido(t *testing.T) {
	str := `[{"nome": }]`
	converter := New()
	data := []models.Data{}
	err := converter.ToJson([]byte(str), &data)
	if err == nil {
		t.Errorf(" Error : deveria retornar error para o json mal formatado : %s , porém retornou o OK com os seguintes dados : %v", str, data)
	}
}

func Test_Deve_Retornar_Error_Para_CSV_Invalido(t *testing.T) {
	str := `[{"nome": }]`
	converter := New()
	data := []models.Data{}
	err := converter.ToCsv(strings.NewReader(str), &data)
	if err == nil {
		t.Errorf(" Error : deveria retornar error para o csv mal formatado : %s , porém retornou o OK com os seguintes dados : %v", str, data)
	}
}

func Test_Deve_Retornar_Error_Para_CSV_Vazio(t *testing.T) {
	str := `Nome,Email,Sexo,Idade`
	converter := New()
	data := []models.Data{}
	err := converter.ToCsv(strings.NewReader(str), &data)
	if err == nil {
		t.Errorf(" expectativa -> ErrorEmptyFile | lançado ->: %v", err)
	}
	t.Logf("expectativa -> ErrorEmptyFile lançado -> : %v", err)

}

func Test_Deve_Retornar_Error_Para_CSV_Nao_Contendo_Header(t *testing.T) {
	str := `daniel,danielvianarj@gmail.com,masculino,25
	nuveo,nuveo@gmail.com,masculino,100`
	converter := New()
	data := []models.Data{}
	err := converter.ToCsv(strings.NewReader(str), &data)
	if err == nil {
		t.Errorf(" expectativa -> ErrorNotFoundAtrributesInHeader | lançado ->: %v", data)
	}

}
