package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type ctr struct {
	Name  NodeName
	Attrs []NodeAttr
}

type NodeName struct {
	Type string
	Name string
}

type NodeAttr struct {
	Name  string
	Value string
}

//current begin
var c *ctr
var ctrSlice []*ctr

func main() {

	ctrfile := "control.log"

	// fmt.Scan(&ctrfile)
	ctrSlice = make([]*ctr, 5)

	//ファイル名読込
	f, err := os.Open(ctrfile)

	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(f)

	var line string
	for scan.Scan() {
		// read line
		line = scan.Text()
		// trim space
		s := strings.TrimSpace(line)
		//start with begin -> create node
		if strings.HasPrefix(s, "Begin") {
			cp := ctr{}

			ctrName := strings.Split(s, " ")

			if len(ctrName) > 2 {
				cp.Name.Name = ctrName[2]
				cp.Name.Type = ctrName[1]
			} else {
				cp.Name.Name = ctrName[1]
				cp.Name.Type = ctrName[0]
			}

			ctrSlice = append(ctrSlice, &cp)
			// set current node
			c = &cp
		}
		replaceLine(s, c)
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	for _, r := range ctrSlice {
		if r != nil {
			sb := strings.Builder{}
			for _, a := range r.Attrs {
				sb.WriteString(a.Name)
				sb.WriteString("\t")
				sb.WriteString(a.Value)
			}
			fmt.Printf("%s\t%s\t%s\n", r.Name.Name, r.Name.Type, sb.String())
		}
	}
}

func replaceLine(str string, cc *ctr) {
	keys := [...]string{"Enabled ", "Caption ", "IMEMode ", "BackColor ", "Text "}
	for _, v := range keys {
		if strings.Contains(str, v) {
			at := strings.Split(str, "=")
			n := NodeAttr{}
			n.Name = strings.TrimSpace(at[0])
			n.Value = strings.TrimSpace(at[1])
			cc.Attrs = append(cc.Attrs, n)
		}
	}
}
