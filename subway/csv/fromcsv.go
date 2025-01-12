package csv

import (
	"bufio"
	"fmt"
	"gosub/subway"
	"log"
	"os"
	"strconv"
	"strings"
)

func InsertNodesFromCSV(csvfile string) {
	file, err := os.Open(csvfile)
	if err != nil {
		log.Fatalf("[ERROR] :%s", err.Error())
	}
	input := bufio.NewScanner(file)
	input.Scan()
	var nodes []subway.Node
	for input.Scan() {
		line := input.Text()
		attrs := strings.Split(line, ",")
		if len(attrs) < 3 {
			log.Fatalf("[ERROR] Insufficient number of attrs\n")
		}
		n_id, err := strconv.Atoi(attrs[0])
		if err != nil {
			log.Fatalf("[ERROR]Invalid Id: %s", err.Error())
		}
		n_name := attrs[1]
		n_lane, err := strconv.Atoi(attrs[2])
		if err != nil {
			log.Fatalf("[ERROR]Invalid Lane Id: %s", err.Error())
		}
		node := subway.Node{
			Id:   n_id,
			Name: n_name,
			Lane: n_lane,
		}
		nodes = append(nodes, node)
	}

	for _, n := range nodes {
		_, err := subway.InsertOneNode(n)
		if err != nil {
			log.Fatalf("[ERROR] Failed inserting into DB: %s", err.Error())
		}
	}
}

func InsertLanesFromCSV(csvfile string) {
	file, err := os.Open(csvfile)
	if err != nil {
		log.Fatalf("[ERROR] :%s", err.Error())
	}
	input := bufio.NewScanner(file)
	input.Scan()
	var lanes []subway.Lane
	for input.Scan() {
		line := input.Text()
		attrs := strings.Split(line, ",")
		if len(attrs) < 2 {
			log.Fatalf("[ERROR] Insufficient number of attrs\n")
		}
		l_id, err := strconv.Atoi(attrs[0])
		if err != nil {
			log.Fatalf("[ERROR]Invalid Id: %s", err.Error())
		}
		l_name := attrs[1]
		lane := subway.Lane{
			Id:   l_id,
			Name: l_name,
		}
		lanes = append(lanes, lane)
	}

	for _, l := range lanes {
		_, err := subway.InsertOneLane(l)
		if err != nil {
			log.Fatalf("[ERROR] Failed inserting into DB: %s", err.Error())
		}
	}
}

func InsertEdgesFromCSV(csvfile string) {
	file, err := os.Open(csvfile)
	if err != nil {
		log.Fatalf("[ERROR] :%s", err.Error())
	}
	input := bufio.NewScanner(file)
	input.Scan()
	var edges []subway.Edge
	for input.Scan() {
		line := input.Text()
		attrs := strings.Split(line, ",")
		id, err := strconv.Atoi(attrs[0])
		if err != nil {
			log.Fatalf("[ERROR]Invalid Id: %s", err.Error())
		}
		src, err := strconv.Atoi(attrs[1])
        if err != nil {
            log.Fatalf("[ERROR]: %s", err.Error())
        }
        dest, err := strconv.Atoi(attrs[2])
        if err != nil {
            log.Fatalf("[ERROR]: %s", err.Error())
        }
		edge := subway.Edge{
			Id:   id,
            Src: src,
            Dest: dest,
		}
		edges = append(edges, edge)
	}

	for _, e := range edges {
		_, err := subway.InsertOneEdge(e)
		if err != nil {
			log.Fatalf("[ERROR] Failed inserting into DB: %s", err.Error())
		}
	}
}

func GenerateConnCsv(source string, dest string) {
	f, err := os.Open(source)
	if err != nil {
		log.Fatalf("[ERROR] :%s", err.Error())
	}
	input := bufio.NewScanner(f)
	input.Scan()
	count := 1
	var conns string
	for input.Scan() {
		line := input.Text()
		attrs := strings.Split(line, ",")

		s_name := attrs[0]
		s_lane, err := strconv.Atoi(attrs[1])
		if err != nil {
			log.Fatalf("[ERROR]: %s", err.Error())
		}

		src := subway.FindNodeByNameAndLane(s_name, s_lane)

		s_neig := attrs[2]
		neighs := strings.Split(s_neig, ";")
		for _, n := range neighs {
			d := subway.FindNodeByNameAndLane(n, s_lane)
			edge := fmt.Sprintf("%d,%d,%d\n", count, src.Id, d.Id)
			conns += edge
			count++
		}
	}
	f.Close()

	out, err := os.OpenFile(dest, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("[ERROR]: %s", err.Error())
	}
	defer out.Close()
	w := bufio.NewWriter(out)
	fmt.Printf("---\nWriting: \n%s", conns)
	n, _ := w.WriteString(conns)
	fmt.Printf("Wrote %d bytes\n---\n", n)
	err = w.Flush()
	if err != nil {
		log.Fatalf("[ERROR]Didn't flush shit: %s", err.Error())
	}
}
