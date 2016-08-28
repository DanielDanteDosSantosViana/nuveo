# Robot Nuveo
[![Build Status](https://api.travis-ci.org/DanielDanteDosSantosViana/nuveo.svg)](https://api.travis-ci.org/DanielDanteDosSantosViana/nuveo.svg?branch=master)

# Excecução

Para utilizar o Robot Nuveo é necessário passar em linha de comando apenas 1 argumento que é a URL de destino:
```
    go nuveo [url]
```
    
# Teste

Para rodar todos os teste é ncessário o seguinte  comando :
```
    go test -v ./...

`=== RUN   Test_Deve_Converter_Json_Para_Estrutura_Informada
--- PASS: Test_Deve_Converter_Json_Para_Estrutura_Informada (0.00s)
=== RUN   Test_Deve_Converter_CSV_Para_Estrutura_Informada
--- PASS: Test_Deve_Converter_CSV_Para_Estrutura_Informada (0.00s)
=== RUN   Test_Deve_Converter_Lista_Json
--- PASS: Test_Deve_Converter_Lista_Json (0.00s)
=== RUN   Test_Deve_Converter_Lista_CSV
--- PASS: Test_Deve_Converter_Lista_CSV (0.00s)
=== RUN   Test_Deve_Retornar_Error_Para_Json_Invalido
--- PASS: Test_Deve_Retornar_Error_Para_Json_Invalido (0.00s)
=== RUN   Test_Deve_Retornar_Error_Para_CSV_Invalido
--- PASS: Test_Deve_Retornar_Error_Para_CSV_Invalido (0.00s)
=== RUN   Test_Deve_Retornar_Error_Para_CSV_Vazio
--- PASS: Test_Deve_Retornar_Error_Para_CSV_Vazio (0.00s)
	converter_test.go:109: expectativa -> ErrorEmptyFile lançado -> : O Arquivo csv eśtá vazio!
=== RUN   Test_Deve_Retornar_Error_Para_CSV_Nao_Contendo_Header
--- PASS: Test_Deve_Retornar_Error_Para_CSV_Nao_Contendo_Header (0.00s)
PASS
ok  	github.com/DanielDanteDosSantosViana/nuveo/converter	0.002s
=== RUN   Test_Deve_Retornar_Error_Para_Url_Invalida
--- PASS: Test_Deve_Retornar_Error_Para_Url_Invalida (0.00s)
=== RUN   Test_Deve_Retornar_Error_Para_Qualquer_Retorno_De_Status_Code_Diferente_De_OK
--- PASS: Test_Deve_Retornar_Error_Para_Qualquer_Retorno_De_Status_Code_Diferente_De_OK (0.00s)
=== RUN   Test_Deve_Retornar_Status_Code_OK_Para_Uma_URL_Valida
--- PASS: Test_Deve_Retornar_Status_Code_OK_Para_Uma_URL_Valida (0.00s)
PASS
ok  	github.com/DanielDanteDosSantosViana/nuveo/http	0.005s
?   	github.com/DanielDanteDosSantosViana/nuveo/models	[no test files]
=== RUN   Test_Deve_Converter_Json_Do_Retorno_Http
--- PASS: Test_Deve_Converter_Json_Do_Retorno_Http (0.00s)
=== RUN   Test_Deve_Converter_CSV_Do_Retorno_Http
--- PASS: Test_Deve_Converter_CSV_Do_Retorno_Http (0.00s)
=== RUN   Test_Deve_Retornar_Error_Para_Content_type_Nao_Definido
--- PASS: Test_Deve_Retornar_Error_Para_Content_type_Nao_Definido (0.00s)
	robot_test.go:64: expectativa -> ErrorNotFoundTypeFile lançado -> : Não foi encontrado o tipo do arquivo !
PASS
ok  	github.com/DanielDanteDosSantosViana/nuveo/robot	0.006s
`
```
Obs:
```
    Para conversão do JSON foi considerado json inicializando como list, conforme a seguir:
   ` [{
	"nome": "daniel",
	"email": "danielvianarj@gmail.com",
	"sexo": "masculino",
	"idade": "25"
}, {
	"nome": "daniel",
	"email": "danielvianarj@gmail.com",
	"sexo": "masculino",
	"idade": "25"
}]`

Para o CSV foi considerado o seguinte Header:
`Nome,Email,Sexo,Idade
	daniel,danielvianarj@gmail.com,masculino,25`
```



