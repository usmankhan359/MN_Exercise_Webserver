package main
     // Importing Go libraries
import (
	"code.google.com/p/gorest"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/sethvargo/go-password/password"
	"net/http"
	"os"
)
     //  Defining restful service
type RandPasswordGenerator struct {
	gorest.RestService `root:"/" consumes:"application/json" produces:"application/json"`

	// End-Point level configs: Field names must be the same as the corresponding method names,
	// but not-exported (starts with lowercase)
	getPassword gorest.EndPoint `method:"GET" path:"password/{length:int}" output:"string"`
}
     // Random password generator
func (serv RandPasswordGenerator) GetPassword(length int) string {
	fmt.Println("received request_id:")
	equalNumber := length / 2
	res, err := password.Generate(length, equalNumber, equalNumber, true, true)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return res
}
func main() {
	fmt.Println("starting services http")
	// Instantiating of struct RandPasswordGenerator
	r := RandPasswordGenerator{}
	// Registering API in RandPasswordGenerator
	gorest.RegisterService(&r)
	// Running handler to run the API
	res := gorest.Handle()
	http.Handle("/", handlers.LoggingHandler(os.Stdout, res))

	err := http.ListenAndServe(":8110", nil)

	if err != nil {
		fmt.Println(err)
	}
}
