package main


import(
	"net/http"
)

func main(){
	server := http.Server{
		Addr:		"127.0.0.2:8080",
		Handler:	nil,
	}
	server.ListenAndServeTLS("cert.perm", "key.perm")
}