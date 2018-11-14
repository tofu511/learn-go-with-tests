package dependency_injection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//func MyGreetHandler(w http.ResponseWriter, r *http.Request)  {
//	Greet(w, "world")
//}
//
//func main() {
//	Greet(os.Stdout, "world")
//	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))
//}
