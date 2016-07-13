package rest

import (
	"testing"
	"net/http"
	"fmt"
	"log"
	"strings"
)

func TestMake(t *testing.T) {

	go func () {
		routes := Routes{
			Route{
				"index",
				http.MethodGet,
				"/testing/{lester}/{category}",
				GetRequest,
			},
		}
		NewServer("/Users/home/Code/Go/src/bitbucket.org/matchmove/rest/server",routes)
	}()

	request, _ := http.NewRequest("GET", "http://localhost:8085/testing/lester/pogi", strings.NewReader("")) //Create request with JSON body



	response, _ := http.DefaultClient.Do(request)

	log.Println(fmt.Sprint(response.Body))

}


func GetRequest(w http.ResponseWriter, r *http.Request, data map[string]string) {
	fmt.Fprint(w, "Welcome!\n")
	fmt.Fprint(w, data)

	log.Println(data)
}
