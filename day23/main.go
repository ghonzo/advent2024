// Advent of Code 2024, Day 23
package main

import (
	"fmt"
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/ghonzo/advent2024/common"
)

// Day 23: LAN Party
// Part 1 answer: 1337
// Part 2 answer: aw,fk,gv,hi,hp,ip,jy,kc,lk,og,pj,re,sr
func main() {
	fmt.Println("Advent of Code 2024, Day 23")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %s\n", part2(entries))
}

func part1(entries []string) int {
	connectedMap := make(map[string][]string)
	for _, line := range entries {
		a, b := line[0:2], line[3:]
		connectedMap[a] = append(connectedMap[a], b)
		connectedMap[b] = append(connectedMap[b], a)
	}
	allNetworks := empty()
	for k, v := range connectedMap {
		if k[0] == 't' {
			for i, a := range v[:len(v)-1] {
				for _, b := range v[i+1:] {
					if slices.Contains(connectedMap[a], b) {
						computers := []string{k, a, b}
						slices.Sort(computers)
						allNetworks.Add(strings.Join(computers, "-"))
					}
				}
			}
		}
	}
	return allNetworks.Cardinality()
}

func empty() mapset.Set[string] {
	return mapset.NewThreadUnsafeSet[string]()
}

type cliqueHelper struct {
	largestClique mapset.Set[string]
	connectedMap  map[string]mapset.Set[string]
}

func part2(entries []string) string {
	// computer pointing to all connected computers
	connectedMap := make(map[string]mapset.Set[string])
	for _, line := range entries {
		a, b := line[0:2], line[3:]
		if c, ok := connectedMap[a]; ok {
			c.Add(b)
		} else {
			connectedMap[a] = mapset.NewThreadUnsafeSet(b)
		}
		if c, ok := connectedMap[b]; ok {
			c.Add(a)
		} else {
			connectedMap[b] = mapset.NewThreadUnsafeSet(a)
		}
	}
	ch := &cliqueHelper{empty(), connectedMap}
	findCliques(empty(), mapset.NewThreadUnsafeSetFromMapKeys(connectedMap), empty(), ch)
	largestClique := ch.largestClique.ToSlice()
	slices.Sort(largestClique)
	return strings.Join(largestClique, ",")
}

// Recursive Bron-Kerbosch algorithm
func findCliques(r, p, x mapset.Set[string], ch *cliqueHelper) {
	if p.IsEmpty() && x.IsEmpty() {
		if r.Cardinality() > ch.largestClique.Cardinality() {
			ch.largestClique = r
		}
	} else {
		for _, v := range p.ToSlice() {
			newR := r.Clone()
			newR.Add(v)
			neighborsV := ch.connectedMap[v]
			newP := p.Intersect(neighborsV)
			newX := x.Intersect(neighborsV)
			findCliques(newR, newP, newX, ch)
			p.Remove(v)
			x.Add(v)
		}
	}
}
