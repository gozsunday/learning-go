package main

import (
	"cmp"
	"fmt"
	"io"
	"slices"
)

type Ranker interface {
	Ranking() []string
}

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	// Ensure the team names passed belong to valid teams in the League.
	// This will also handle the scenario where l is nil, or where the Teams map wasn't
	// instantiated, since you can't read from the Teams map unless l is a valid League
	// and/or the Teams map was instantiated.
	if _, ok := l.Teams[team1]; !ok {
		return
	}
	if _, ok := l.Teams[team2]; !ok {
		return
	}
	if score1 == score2 {
		return
	}

	// update the teams wins count in League
	if score1 > score2 {
		l.Wins[team1]++
	} else {
		l.Wins[team2]++
	}
}

func (l *League) Ranking() []string {
	// use `make` to create a slice with the exact capacity required
	names := make([]string, 0, len(l.Teams))

	// add teams' name in the Teams map to the slice
	for t := range l.Teams {
		names = append(names, t)
	}

	// sort the teams in the array from highest no of wins to lowest
	slices.SortFunc(names, func(a, b string) int {
		return cmp.Compare(l.Wins[b], l.Wins[a])
	})
	return names
}

func RankPrinter(r Ranker, w io.Writer) {
	for _, v := range r.Ranking() {
		fmt.Fprintf(w, "%s\n", v)
		// io.WriteString(w, fmt.Sprintf("%s\n", v))
	}
}
