package handler

import (
	"fmt"
	"net/http"
)

func listener(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 PAGE NOT FOUND!", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Get received!")
		request_id := r.FormValue("request_id")
		ip := r.FormValue("ip ")
		url_package := r.FormValue("url_package")

		fmt.Fprintf(w, request_id)
		fmt.Fprintf(w, ip)
		fmt.Fprintf(w, url_package)

	case "POST":
		fmt.Fprintf(w, "Post received!")
		// if err := r.ParseFrom(); err != nil{
		//   fmt.Fprintf(w, "ParseForm() err : %v", err)
		//   return
		// }
		// fmt.Fprintf(w, "Post from website r.postfrom = %v\n", r.PostFrom)
		// {
		//   name := r.FormValue("name")
		//   address := r.FormValue("address")
		//
		//   fmt.Fprintf(w, "Name = %s\n", name)
		//   fmt.Fprintf(w, "Adress = %s\n", name)
		// }
	default:
		fmt.Fprintf(w, "Only get and post!")
	}
}
