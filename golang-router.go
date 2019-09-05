package main

import (
	"net/http"
	"encoding/json"
	"regexp"
	"strconv"
	"os"
)

type Usuario struct {
	Id int "json:\"id\""
	Nome string "json:\"nome\""
}

func usuariosHandler(w http.ResponseWriter, r *http.Request) {

	usuarios := []Usuario{
		Usuario{Id: 1, Nome: "Henrique"},
		Usuario{Id: 2, Nome: "Vieira"},
	}

	path := r.URL.Path
	re := regexp.MustCompile("^/usuarios/([0-9]*)$")

	if path == "/usuarios/" {
		usuariosJson, _ := json.Marshal(usuarios)
		w.Write([]byte(usuariosJson))
	} else if re.MatchString(path) {
		id := re.FindAllStringSubmatch(path, -1)[0][1]

		usuariosFiltro := make([]Usuario, 0)		

		for _, usuario := range usuarios {
			if strconv.Itoa(usuario.Id) == id {
				usuariosFiltro = append(usuariosFiltro, usuario)
				break
			}
		}
		usuariosJson, _ := json.Marshal(usuariosFiltro)
		w.Write([]byte(usuariosJson))
	}
}

func main() {
	porta := os.Getenv("PORT")
	http.HandleFunc("/usuarios/", usuariosHandler)
	http.ListenAndServe(":" + porta, nil)
}
