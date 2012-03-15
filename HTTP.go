package main

import "fmt"
import "http"


func handler (w http.ResponseWriter, r *http.Request){
 fmt.Fprint(w, "Hello, world!")
}

func main(){
 http.HandleFunc("/",handler)
 http.ListenAndServe(":8120",nil)
}
