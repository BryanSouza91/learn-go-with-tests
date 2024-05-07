package dependencyinjection

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s!", name)
}

// Cannot run the server as this is not a "main" package
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	if err := http.ListenAndServe(":9191", http.HandlerFunc(MyGreeterHandler)); err != nil {
		log.Fatal(err)
	}
}
