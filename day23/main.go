// Advent of Code 2024, Day 23
package main

import (
	"fmt"
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/ghonzo/advent2024/common"
)

// Day 23:
// Part 1 answer: 1337
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2024, Day 23")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %s\n", part2(entries))
}

func part1(entries []string) int {
	connectedMap := make(map[string][]string)
	for _, line := range entries {
		connectedMap[line[0:2]] = append(connectedMap[line[0:2]], line[3:])
		connectedMap[line[3:]] = append(connectedMap[line[3:]], line[0:2])
	}
	allNetworks := mapset.NewThreadUnsafeSet[string]()
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

type cliqueHelper struct {
	largestClique mapset.Set[string]
	connectedSets map[string]mapset.Set[string]
}

func (ch *cliqueHelper) chooseMostConnected(choices mapset.Set[string]) string {
	var maxConnected string
	var maxConnectedSize int
	choices.Each(func(c string) bool {
		connected := ch.connectedSets[c]
		if connected.Cardinality() > maxConnectedSize {
			maxConnectedSize = connected.Cardinality()
			maxConnected = c
		}
		return true
	})
	return maxConnected
}

func part2(entries []string) string {
	connectedMap := make(map[string][]string)
	for _, line := range entries {
		connectedMap[line[0:2]] = append(connectedMap[line[0:2]], line[3:])
		connectedMap[line[3:]] = append(connectedMap[line[3:]], line[0:2])
	}
	connectedSets := make(map[string]mapset.Set[string])
	empty := mapset.NewThreadUnsafeSet[string]()
	for k, v := range connectedMap {
		connectedSets[k] = mapset.NewThreadUnsafeSet(v...)
	}
	ch := &cliqueHelper{mapset.NewThreadUnsafeSet[string](), connectedSets}
	findCliques(empty, mapset.NewThreadUnsafeSetFromMapKeys(connectedSets), empty, ch)
	allComputers := ch.largestClique.ToSlice()
	slices.Sort(allComputers)
	return strings.Join(allComputers, ",")
}

func findCliques(r, p, x mapset.Set[string], ch *cliqueHelper) {
	if p.IsEmpty() && x.IsEmpty() {
		fmt.Println("Clique found:", r)
		if r.Cardinality() > ch.largestClique.Cardinality() {
			ch.largestClique = r
		}
	} else {
		u := ch.chooseMostConnected(p.Union(x))
		neighborsU := ch.connectedSets[u]
		for v := range p.Difference(neighborsU).Iter() {
			newR := r.Clone()
			newR.Add(v)
			neighborsV := ch.connectedSets[v]
			findCliques(newR, p.Intersect(neighborsV), x.Intersect(neighborsV), ch)
			p.Remove(v)
			x.Add(v)
		}
	}
}
