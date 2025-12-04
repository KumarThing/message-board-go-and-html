package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)


var tmpl = template.Must(template.ParseFiles("template/index.html"))
var message [] string

func main() {
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		log.Println("Web site visited.")

		if r.Method == http.MethodPost{
			r.ParseForm()


		newMessage := r.FormValue("message")

			if newMessage != "" {
				message = append(message, newMessage)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		data:= struct{
			Message[] string
		} {
			Message : message,
		}

		tmpl.Execute(w, data)
		
	})


	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodPost{
			r.ParseForm()

			delStr := r.FormValue("delete")

			i, err := strconv.Atoi(delStr)
			if err == nil && i >=0 && i < len(message)  {
				message = append(message[:i], message[i+1:]... )
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return

		}
	})

	fmt.Println("Your server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}