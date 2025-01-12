package main

import (
    "flag"
	"fmt"
	"gosub/subway"
	subcsv "gosub/subway/csv"
	"log"
	"net/http"
	"os"
	"strings"
)

const ADDR = "0.0.0.0"
const PORT = 8000

func main() {
    bs_flag := flag.Bool("boostrapdb", false, "Bootstraps the database with values readen from csv files")
    flag.Parse()
    log.Printf("Flags Parsed. BS Flag value = %t\n", *bs_flag)

    initdb()
    if *bs_flag {
        bootstrapdb()
    }
    
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
	log.Printf("[INFO] Executing db.sql\n")
	execfile("./db/db.sql")
	log.Printf("[INFO] Executing upt.sql\n")
	execfile("./db/upt.sql")
}

func bootstrapdb() {
	log.Printf("[INFO] Bootstraping db\n")
    file , err := os.ReadFile("./csvfiles.txt")
    if err != nil {
        panic("bootstrapdb: couldn't open csv files")
    }
	log.Printf("[INFO] Successfully read csvfiles.txt\n")
    names := strings.Split(string(file), ",")
    if len(names) != 3 {
        panic(fmt.Sprintf("bootstrapdb: wrong amount of csv files. want 3 got %d", len(names))) 
    }
    subcsv.InsertNodesFromCSV(strings.TrimSpace(names[0]))
    subcsv.InsertLanesFromCSV(strings.TrimSpace(names[1]))
    subcsv.InsertEdgesFromCSV(strings.TrimSpace(names[2]))
	log.Printf("[INFO] All Insert Functions were Called\n")
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
