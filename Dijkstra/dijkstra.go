package main

import (
	"fmt"
	"time"
)

func Dijkstra(a, b *Point, g *Graph) int {

	if g.CountPoints() == 0 {
		panic("Invalid points in graph")
	}

	// init L
	l := make(map[string]map[string]int)
	for _, p := range g.Points {
		l[p.Name] = map[string]int{}
	}
	// 1<<16 means it's very far
	for _, p1 := range g.Points {
		for _, p2 := range g.Points {
			l[p1.Name][p2.Name] = 1 << 16
		}
	}
	// when has edge
	for k1, tmp := range g.Edges {
		for k2, w := range tmp {
			l[k1][k2] = w
		}
	}

	// init S
	s := []string{}
	for !inSlice(b.Name, s) {
		minLP := g.Points[0]
		for _, p := range g.Points {
			if a.Name == p.Name || inSlice(p.Name, s) {
				continue
			}
			if l[a.Name][p.Name] < l[a.Name][minLP.Name] {
				minLP = p
			}
		}

		time.Sleep(time.Second)

		if inSlice(minLP.Name, s) {
			continue
		}
		s = append(s, minLP.Name)
		fmt.Printf("s: %v\n", s)
		fmt.Printf("%s in s: %v\n", b.Name, inSlice(b.Name, s))
		for _, p := range g.Points {
			if inSlice(p.Name, s) {
				continue
			}
			fmt.Printf("l[%s][%s] = %d\t", a.Name, minLP.Name, l[a.Name][minLP.Name])
			fmt.Printf("l[%s][%s] = %d\t", minLP.Name, p.Name, l[minLP.Name][p.Name])
			fmt.Printf("l[%s][%s] = %d\n", a.Name, p.Name, l[a.Name][p.Name])
			if l[a.Name][minLP.Name]+l[minLP.Name][p.Name] < l[a.Name][p.Name] {
				l[a.Name][p.Name] = l[a.Name][minLP.Name] + l[minLP.Name][p.Name]
			}
		}
	}
	return l[a.Name][b.Name]
}

func inSlice(s string, ss []string) bool {
	for _, i := range ss {
		if s == i {
			return true
		}
	}
	return false
}
