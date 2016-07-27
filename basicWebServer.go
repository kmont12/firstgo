package main

import ("net/http"
	"log"
	"io/ioutil"
)

func main(){
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	path := r.URL.Path[1:]
	log.Println(path)

	if path==""{
		path="index.html"
	}

	 data, err := ioutil.ReadFile("docs/"+string(path))


	if err==nil{
		w.Write(data)
	} else{
		w.WriteHeader(404)
		w.Write([]byte("404 this didnt work" + http.StatusText(404)))
	}
}
