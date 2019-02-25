package main

import (
	"fmt"
	"github.com/gnewton/mesh2sqlite3/lib"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {

	t := dot()
	db, err := gorm.Open("sqlite3", "mesh2019_sqlite3.db")
	if err != nil {
		log.Fatal(err)

	}

	// Write nodes
	count := 0
	//depth=14
	depth := 10
	for i := 0; i < depth; i++ {
		fmt.Printf("\n\nsubgraph cluster%d{\n", i)
		fmt.Println("rank=same;")
		fmt.Println("style=invis;")
		b
		var trees []*lib.MeshTree
		db.Where("depth = ?", i).Find(&trees)
		count += len(trees)
		for j := 0; j < len(trees); j++ {
			dotNode := tree2dot(trees[j])
			err = t.Execute(os.Stdout, dotNode)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("}")
	}

	// Write edges
	var trees []*lib.MeshTree
	db.Find(&trees)
	m := make(map[string]string, 0)
	for i := 0; i < len(trees); i++ {
		t := trees[i]
		if t.Depth < depth {
			if t.T1 != nil {
				s := fmt.Sprintf("\n\"%s\"->\"%s\";", *t.T0, *t.T0+*t.T1)
				if _, ok := m[s]; !ok {
					fmt.Print(s)
					m[s] = s
				}
				if t.T2 != nil {
					s = fmt.Sprintf("\n\"%s\"->\"%s\";", *t.T0+*t.T1, *t.T0+*t.T1+"."+*t.T2)
					if _, ok := m[s]; !ok {
						fmt.Print(s)
						m[s] = s
					}

					if t.T3 != nil {
						s = fmt.Sprintf("\n\"%s\"->\"%s\";", *t.T0+*t.T1+"."+*t.T2, *t.T0+*t.T1+"."+*t.T2+"."+*t.T3)
						if _, ok := m[s]; !ok {
							fmt.Print(s)
							m[s] = s
						}
						if t.T4 != nil {
							s = fmt.Sprintf("\n\"%s\"->\"%s\";", *t.T0+*t.T1+"."+*t.T2+"."+*t.T3, *t.T0+*t.T1+"."+*t.T2+"."+*t.T3+"."+*t.T4)
							if _, ok := m[s]; !ok {
								fmt.Print(s)
								m[s] = s
							}
							if t.T5 != nil {
								s = fmt.Sprintf("\n\"%s\"->\"%s\";", *t.T0+*t.T1+"."+*t.T2+"."+*t.T3+"."+*t.T4, *t.T0+*t.T1+"."+*t.T2+"."+*t.T3+"."+*t.T4+"."+*t.T5)
								if _, ok := m[s]; !ok {
									fmt.Print(s)
									m[s] = s
								}

								if t.T6 != nil {
									s = fmt.Sprintf("\n\"%s\"->\"%s\";", *t.T0+*t.T1+"."+*t.T2+"."+*t.T3+"."+*t.T4+"."+*t.T5, *t.T0+*t.T1+"."+*t.T2+"."+*t.T3+"."+*t.T4+"."+*t.T5+"."+*t.T6)
									if _, ok := m[s]; !ok {
										fmt.Print(s)
										m[s] = s
									}
									if t.T7 != nil {
										s = fmt.Sprintf("\n\"%s\"->\"%s\";", *t.T0+*t.T1+"."+*t.T2+"."+*t.T3+"."+*t.T4+"."+*t.T5+"."+*t.T6, *t.T0+*t.T1+"."+*t.T2+"."+*t.T3+"."+*t.T4+"."+*t.T5+"."+*t.T6+"."+*t.T7)
										if _, ok := m[s]; !ok {
											fmt.Print(s)
											m[s] = s
										}
										if t.T8 != nil {
											s = fmt.Sprintf("\n\"%s\"->\"%s\";", *t.T0+*t.T1+"."+*t.T2+"."+*t.T3+"."+*t.T4+"."+*t.T5+"."+*t.T6+"."+*t.T7, *t.T0+*t.T1+"."+*t.T2+"."+*t.T3+"."+*t.T4+"."+*t.T5+"."+*t.T6+"."+*t.T7+"."+*t.T8)
											if _, ok := m[s]; !ok {
												fmt.Print(s)
												m[s] = s
											}
											if t.T9 != nil {
												s = fmt.Sprintf("\n\"%s\"->\"%s\";", *t.T0+*t.T1+"."+*t.T2+"."+*t.T3+"."+*t.T4+"."+*t.T5+"."+*t.T6+"."+*t.T7+"."+*t.T8, *t.T0+*t.T1+"."+*t.T2+"."+*t.T3+"."+*t.T4+"."+*t.T5+"."+*t.T6+"."+*t.T7+"."+*t.T8+"."+*t.T9)
												if _, ok := m[s]; !ok {
													fmt.Print(s)
													m[s] = s
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}

		}
	}

}
