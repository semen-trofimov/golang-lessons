package main
import(
	"rmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Reqest){
		fmt.Printf(w, "Hello world!")
	})
	http.ListenAndServe(":80", nill)
}