package subway

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Node struct {
	Id   int
	Name string
	Lane int
}

type Lane struct {
	Id   int
	Name string
}

type Edge struct {
	Id   int
	Src  int
	Dest int
}

const DB_DRIVER = "sqlite"
const DB_PATH   = "./db/sqlite3.db"
func Open() *sql.DB {
	db, err := sql.Open(DB_DRIVER, DB_PATH)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "Not Found"
}

func FindNodeByNameAndLane(n string, l int) Node {
	db := Open()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM nodes WHERE name = ? AND lane = ?", n, l)
	if err != nil {
		log.Fatalf("[ERROR] Failed Fetching DB: %s", err.Error())
	}
	var node Node
	if rows.Next() {
		rows.Scan(&node.Id, &node.Name, &node.Lane)
	}
	return node
}

func FindNodesByLane(l int) []Node {
	db := Open()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM nodes WHERE lane = ?", l)
	if err != nil {
		log.Fatalf("[ERROR] Failed Fetching DB: %s", err.Error())
	}
	var nodes []Node
	for rows.Next() {
		var node Node
		rows.Scan(&node.Id, &node.Name, &node.Lane)
		nodes = append(nodes, node)
	}
	return nodes
}

func ListAllNodes() []Node {
	db := Open()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM nodes")
	if err != nil {
		log.Fatalf("[ERROR] Failed Fetching DB: %s", err.Error())
	}
	var nodes []Node
	for rows.Next() {
		var node Node
		rows.Scan(&node.Id, &node.Name, &node.Lane)
		nodes = append(nodes, node)
	}
	return nodes
}

func MapAllNodes() map[int]Node {
	db := Open()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM nodes")
	if err != nil {
		log.Fatalf("[ERROR] Failed Fetching DB: %s", err.Error())
	}
    nodes := make(map[int]Node)
	for rows.Next() {
		var node Node
		rows.Scan(&node.Id, &node.Name, &node.Lane)
		nodes[node.Id] = node
	}
	return nodes
}

func GetAllNeighbors(s int) []Node {
	db := Open()
	defer db.Close()
	
    rows, err := db.Query("SELECT * FROM nodes WHERE id IN (SELECT dest FROM edges WHERE source = ?)", s)
	if err != nil {
		log.Fatalf("[ERROR] Failed Fetching DB: %s", err.Error())
	}

	var ns []Node
	for rows.Next() {
		var node Node
		rows.Scan(&node.Id, &node.Name, &node.Lane)
		ns = append(ns, node)
	}
	return ns
}

func MakeAdjacencyList() map[int][]Node {
	nodes := MapAllNodes()
	adj := make(map[int][]Node)
	for id := range nodes {
		adj[id] = GetAllNeighbors(id)
	}
	return adj
}

func InsertOneNode(value Node) (sql.Result, error) {
	db := Open()
	defer db.Close()
	res, err := db.Exec("INSERT INTO nodes (id, name, lane) VALUES (?, ?, ?)", value.Id, value.Name, value.Lane)
	return res, err
}

func InsertNodes(nodes []Node) (sql.Result, error) {
	db := Open()
	defer db.Close()
    exec := "INSERT INTO nodes (id, name, lane) VALUES "
    args := make([]any, 0, len(nodes)*3)
    for _, node := range nodes {
        exec += "(?, ?, ?),"
        args = append(args, node.Id, node.Name, node.Lane)
    }
    exec = exec[:len(exec)-1]
	res, err := db.Exec(exec, args...)
	return res, err
} 

func DeleteNodeByID(id int) (sql.Result, error) {
	db := Open()
	defer db.Close()
	res, err := db.Exec("DELETE FROM nodes WHERE id = ?", id)
	return res, err
}
func FindLane(field string, value string) []Lane {
	db := Open()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM lanes WHERE ? = ?")
	if err != nil {
		log.Fatalf("[ERROR] Failed Fetching DB: %s", err.Error())
	}
	var lanes []Lane
	for rows.Next() {
		var lane Lane
		rows.Scan(&lane.Id, &lane.Name)
		lanes = append(lanes, lane)
	}
	return lanes
}

func ListAllLanes() []Lane {
	db := Open()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM lanes")
	if err != nil {
		log.Fatalf("[ERROR] Failed Fetching DB: %s", err.Error())
	}
	var lanes []Lane
	for rows.Next() {
		var lane Lane
		rows.Scan(&lane.Id, &lane.Name)
		lanes = append(lanes, lane)
	}
	return lanes
}


func MapAllLanes() map[int]Lane {
	db := Open()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM lanes")
	if err != nil {
		log.Fatalf("[ERROR] Failed Fetching DB: %s", err.Error())
	}
    lanes := make(map[int]Lane)
	for rows.Next() {
		var lane Lane
		rows.Scan(&lane.Id, &lane.Name)
		lanes[lane.Id] = lane
	}
	return lanes
}

func InsertOneLane(value Lane) (sql.Result, error) {
	db := Open()
	defer db.Close()
	res, err := db.Exec("INSERT INTO lanes (id, name) VALUES (?, ?)", value.Id, value.Name)
	return res, err
}

func InsertOneEdge(value Edge) (sql.Result, error) {
	db := Open()
	defer db.Close()
	res, err := db.Exec("INSERT INTO edges (id, source, dest) VALUES (?, ?, ?)", value.Id, value.Src, value.Dest)
	return res, err
}
