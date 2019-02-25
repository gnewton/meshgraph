package main

import (
	"github.com/gnewton/mesh2sqlite3/lib"
	"text/template"
)

type DotNode struct {
	NodeName string
	Color    string
	Tooltip  string
	Label    string
	Url      string
}

const node = `"{{.NodeName}}"[color={{.Color}}, tooltip="{{.Tooltip}}",label=<{{.Label}}>,URL="{{.Url}}"];
`

func dot() *template.Template {
	return template.Must(template.New("node").Parse(node))
}

const UrlBase = "https://meshb.nlm.nih.gov/record/ui?ui="

func tree2dot(tree *lib.MeshTree) *DotNode {
	dn := new(DotNode)
	dn.NodeName = tree.Tree

	c := string(tree.Tree[0])

	switch c {
	case "A":
		dn.Color = "salmon"

	case "B":
		dn.Color = "orange"

	case "C":
		dn.Color = "yellow"

	case "D":
		dn.Color = "purple"

	case "E":
		dn.Color = "pink"

	case "F":
		dn.Color = "lemonchiffon"

	case "G":
		dn.Color = "cadetblue1"

	case "H":
		dn.Color = "darkgoldenrod1"

	case "I":
		dn.Color = "darkolivegreen1"
	case "J":
		dn.Color = "lavender"
	case "K":
		dn.Color = "orchid2"

	case "L":
		dn.Color = "seagreen1"
	case "M":
		dn.Color = "salmon2"
	case "N":
		dn.Color = "sienna1"
	case "V":
		dn.Color = "tomato1"
	case "Z":
		dn.Color = "thistle"

	}
	dn.Label = tree.DescriptorName + " " + tree.Tree
	dn.Tooltip = tree.DescriptorUI + " " + tree.DescriptorName
	dn.Url = UrlBase + tree.DescriptorUI

	return dn
}
