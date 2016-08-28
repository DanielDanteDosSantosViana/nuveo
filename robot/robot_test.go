package robot

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Deve_Converter_Json_Do_Retorno_Http(t *testing.T) {
	str := `[{"nome": "daniel", "email": "danielvianarj@gmail.com", "sexo": "masculino", "idade" :"25"}]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(str))
	}))
	defer server.Close()
	robot := New()
	data, err := robot.GetData(server.URL)
	if err != nil {
		t.Errorf("%v", err)
	}
	for _, value := range data {
		if value.Nome != "daniel" || value.Email != "danielvianarj@gmail.com" || value.Idade != "25" || value.Sexo != "masculino" {
			t.Errorf(" Error : deveria converter corretamente todos os atributos do json : %s , porém retornou apenas : %v", str, value)
		}
	}
}

func Test_Deve_Converter_CSV_Do_Retorno_Http(t *testing.T) {
	str := `Nome,Email,Sexo,Idade
  daniel,danielvianarj@gmail.com,masculino,25`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/csv")
		w.Write([]byte(str))
	}))
	defer server.Close()
	robot := New()
	data, err := robot.GetData(server.URL)
	if err != nil {
		t.Errorf("%v", err)
	}
	for _, value := range data {
		if value.Nome != "daniel" || value.Email != "danielvianarj@gmail.com" || value.Idade != "25" || value.Sexo != "masculino" {
			t.Errorf(" Error : deveria converter corretamente todos os atributos do csv : %s , porém retornou apenas : %v", str, value)
		}
	}
}

func Test_Deve_Retornar_Error_Para_Content_type_Nao_Definido(t *testing.T) {
	str := `Nome,Email,Sexo,Idade
  daniel,danielvianarj@gmail.com,masculino,25`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		w.Write([]byte(str))
	}))
	defer server.Close()
	robot := New()
	_, err := robot.GetData(server.URL)
	if err == nil {
		t.Errorf("%v", err)
	}
	t.Logf("expectativa -> ErrorNotFoundTypeFile lançado -> : %v", err)

}
