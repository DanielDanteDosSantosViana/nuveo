# Robot Nuveo

# Excecução

Para utilizar o Robot Nuveo é necessário passar em linha de comando apenas 1 argumento que é a URL de destino:
```
    go nuveo [url]
```
    
# Teste

Para rodar todos os teste é ncessário o seguinte  comando :
```
    go test -v ./...
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



