package routes

import (
	"net/http"

	c "../controllers"
)

//CarregaRotas : carregaRotasC
func CarregaRotas() {
	http.HandleFunc("/", c.Index)
	http.HandleFunc("/new", c.New)
	http.HandleFunc("/insert", c.Insert)
	http.HandleFunc("/delete", c.Delete)
	http.HandleFunc("/edit", c.Edit)
	http.HandleFunc("/update", c.Update)
}
