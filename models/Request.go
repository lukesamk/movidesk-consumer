package models

/*
	Funções responsáveis por executarem as requisições à API
	do Movidesk e retornarem a resposta convertida para o tipo
	de dados incluídos em um chamado

	@author Lucas Samuel
	@version 1.0.0
	@since 17.07.19
*/

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
	Estrutura de dados contidos em uma requisição
*/
type Request struct {
	/*
		Contém a URL onde será efetuada a requisição. O token,
		os campos e os filtros são incluídos aqui

		@example https://api.movidesk.com/public/v1/tickets?token=TOKEN&$select=id,subject&$filters=createdDate eq 2019-01-01
	*/
	URL string

	/*
		Indica o método que será utilizado na requisição

		@example GET
	*/
	// TODO: Implementar funções para o método POST
	Method string

	/*
		Contém a resposta da requisição de acordo com a
		estrutura de dados de uma resposta de requisição
	*/
	Response Response
}

/*
	Estrutura de dados de uma resposta de requisição
*/
type Response struct {
	/*
		Corpo da resposta em bytes
	*/
	Body []byte

	/*
		Resposta convertida de bytes para um vetor de tickets
	*/
	Data []Ticket
}

/*
	New inicia uma nova requisição à API

	Essa função valida se a URL e o método informados são vazios e retorna
	um erro se forem. Se não forem vazios, define a URL e o método da
	requisição de acordo com os parâmetros da função.

	@param url string -> Recebe a URL que será utilizada para a requisição
	@param method string -> Indica o tipo de requisição que será utilizado
*/
func (r *Request) New(url string, method string) error {
	// Se a URL é vazia, retorna um erro
	if url == "" {return errors.New("Erro ao criar uma requisição: a URL não pode ser vazia")}

	// Se o método é vazio, retorna um erro
	if method == "" {return errors.New("Erro ao criar uma requisição: o método não pode ser vazio")}

	// Define a URL da requisição
	r.URL = url

	// Define o método da requisição
	r.Method = method

	return nil
}

/*
	Run executa a requisição à API

	O método e a URL da requisição são obtidos dos dados que foram
	informados através do início da requisição (função New).

	Se houver um erro durante a execução da requisição, ele é retornado
	como erro fatal, encerrando a aplicação. O corpo da resposta à
	requisição é salvo e convertido para um vetor de tickets.
*/
func (r *Request) Run() error {
	// Define o timeout da requisição em 5 segundos
	client := http.Client{Timeout: time.Second * 10}

	// Define a URL e o método da requisição de acordo com o que foi passado pela função New
	request, clientErr := http.NewRequest(r.Method, r.URL, nil)

	// Valida erro na requisição
	if clientErr != nil {return errors.New("Erro ao definir a requisição HTTP: " + clientErr.Error())}

	// Define o cabeçalho da requisição
	request.Header.Set("User-Agent", "movidesk-api")

	// Executa a requisição
	response, requestErr := client.Do(request)

	// Valida erro na execução da requisição
	if requestErr != nil {return errors.New("Erro ao executar a requisição HTTP: " + requestErr.Error())}

	// Recebe o corpo da resposta à requisição
	body, readErr := ioutil.ReadAll(response.Body)

	// Valida erro na resposta
	if readErr != nil {return errors.New("Erro ao ler a resposta da requisição HTTP: " + readErr.Error())}

	// Se o corpo da resposta for vazio, retorna um erro
	if body == nil {return errors.New("Erro ao ler o corpo da resposta da requisição HTTP: " + string(body))}

	// Salva o corpo da resposta
	r.Response.Body = body

	// Converte a resposta em um vetor de tickets
	responseErr := r.ReadResponse()

	// Valida erro na conversão
	if responseErr != nil {return errors.New("Erro ao converter a resposta da requisição HTTP: " + responseErr.Error())}

	return nil
}

/*
	ReadResponse converte o corpo da resposta à requisição em
	um vetor de tickets e retorna todos os dados da resposta.

	Os itens dentro do JSON retornado pela API são convertidos
	em atributos de acordo com a estrutura de dados de um Ticket

	@return Response -> retorna os dados da resposta da requisição
*/
func (r *Request) ReadResponse() error {
	// Instancia a estrutura de dados Ticket
	tickets := []Ticket{}
	ticket := Ticket{}

	// Converte o JSON de acordo com os dados da estrutura Ticket
	var jsonErrArray error
	var jsonErrStruct error

	jsonErrArray = json.Unmarshal(r.Response.Body, &tickets)

	if jsonErrArray != nil {
		jsonErrStruct = json.Unmarshal(r.Response.Body, &ticket)
	} else {
		// Salva o vetor de tickets da resposta
		r.Response.Data = tickets
	}

	if (jsonErrArray != nil) && (jsonErrStruct != nil) {
		log.Printf("Não foi possível decodificar a resposta: %v", jsonErrStruct)

		if e, ok := jsonErrStruct.(*json.SyntaxError); ok {
			log.Printf("Erro de sintaxe no byte %d", e.Offset)
		}

		log.Printf("Resposta do Movidesk: %q", r.Response.Body)
	}

	return nil
}