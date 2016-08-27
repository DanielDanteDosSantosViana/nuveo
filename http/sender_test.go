package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Deve_Retornar_Error_Para_Url_Invalida(t *testing.T) {
	_, err := Send("wwwtesterrorcom")
	if err != ErrorInvalidUrl {
		t.Errorf(" Error : deveria retornar um error de url inválida, mas retornou:  %v", err)
	}
}
func Test_Deve_Retornar_Error_Para_Qualquer_Retorno_De_Status_Code_Diferente_De_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	resp, err := Send(server.URL)
	if err != ErrorBadRequest {
		t.Errorf(" Error : deveria retornar error para status code diferente de OK(200) , mas retornou Status Code %v )", resp.StatusCode)
	}
}
func Test_Deve_Retornar_Status_Code_OK_Para_Uma_URL_Valida(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	resp, _ := Send(server.URL)
	if resp.StatusCode != http.StatusOK {
		t.Errorf(" Error : dado uma url válida deve retornar status code OK(200), porém retorno do Status Code %v )", resp.StatusCode)
	}
}
