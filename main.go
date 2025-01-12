package main

import (
	"fmt"
	"gosub/subway"
	"log"
	"net/http"
	"os"
	"strings"
)

const ADDR = "127.0.0.1"
const PORT = 8000

func main() {
    initdb()
    
    router := http.NewServeMux()

	router.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	for _, route := range subway.SubwayRouter.Routes {
		router.HandleFunc(route.Path, route.Handler)
	}

	url := fmt.Sprintf("%v:%v", ADDR, PORT)
	server := http.Server{
		Addr: url,
		Handler: router,
	}

	log.Printf("[INFO] Listening on %v...\n", url)
	log.Fatal("[CRITICAL] Server failed :\n", server.ListenAndServe())

}

func initdb() {
	execfile("./db/db.sql")
	execfile("./db/upt.sql")
}

func execfile(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		log.Fatal("Could not Open DB: ", err)
	}

	db := subway.Open()
	defer db.Close()
	if err != nil {
		log.Fatal("Could not Open DB: ", err)
	}

	cmds := strings.Split(string(file), ";")
	for _, cmd := range cmds {
        _, err := db.Exec(cmd)
        if err != nil {
            log.Printf("error execfile(%s) \"%s\": %s", name, cmd, err)
        }
	}
}
