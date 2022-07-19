package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

type Node struct {
	Name    string
	Type    int
	Attr    []*Attribute
	SubNode []*Node
}

type Attribute struct {
	kye, val string
}

var tmp string = `
Begin VB.Form frmMain
   BackColor       =   &H00008000&
   BorderStyle     =   1  '固定(実線)
   Caption         =   "録音検索"
   ClientHeight    =   13065
   ClientLeft      =   45
   ClientTop       =   435
   ClientWidth     =   17565
`

func main() {
	var nd Node
	var subnode Node
	// var isRoot boot

	scanner := bufio.NewScanner(strings.NewReader(tmp))
	for scanner.Scan() {

		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "Begin VB.Form") {
			nd = Node{}
			nd.Name = line
			// isRoot = true
		}

		if strings.HasPrefix(strings.TrimSpace(line), "Begin") {
			subnode = Node{}
			nd.SubNode = append(nd.SubNode, &subnode)
			subnode.Name = line
			// isRoot = false
		}

		a := &Attribute{}
		if strings.Contains(line, "=") {
			atts := strings.Split(line, "=")
			a.kye = atts[0]
			a.val = atts[1]
		}
		if true {
			nd.Attr = append(nd.Attr, a)
		} else {
			subnode.Attr = append(subnode.Attr, a)
		}
	}
	fmt.Printf("%v", nd.Attr)
	// fmt.Printf("%v", nd.SubNode[0].Attr[1].kye)
}

func countLeadingSpaces(line string) int {
	return len(line) - len(strings.TrimLeft(line, " "))
}

func countLeandings(line string) {
	re := regexp.MustCompile(`^\s+`)
	fmt.Println(len(re.FindAllString(line, -1)))
}
