package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// team holds the statistics for a given team.
type team struct {
	name   string
	played int
	wins   int
	draws  int
	losses int
	points int
}

// Tally reads a scorecard and writes the results.
func Tally(r io.Reader, w io.Writer) error {
	teams := make(map[string]*team)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '#' {
			continue
		}

		details := strings.Split(line, ";")
		if len(details) != 3 {
			return errors.New("failed to parse details from line")
		}

		t1, t2, result := details[0], details[1], details[2]

		if _, ok := teams[t1]; !ok {
			teams[t1] = &team{name: t1}
		}

		if _, ok := teams[t2]; !ok {
			teams[t2] = &team{name: t2}
		}

		teams[t1].played++
		teams[t2].played++

		switch result {
		case "win":
			teams[t1].wins++
			teams[t2].losses++
		case "loss":
			teams[t1].losses++
			teams[t2].wins++
		case "draw":
			teams[t1].draws++
			teams[t2].draws++
		default:
			return errors.New("invalid result")
		}
	}

	var all []*team
	for _, team := range teams {
		all = append(all, team)
	}
	sort.Slice(all, func(i, j int) bool {
		p1 := 3*all[i].wins + 1*all[i].draws
		p2 := 3*all[j].wins + 1*all[j].draws
		if p1 == p2 {
			return all[i].name < all[j].name
		}
		return p1 > p2
	})

	_, err := fmt.Fprintf(w, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return fmt.Errorf("failed to print header: %v", err)
	}

	for _, team := range all {
		points := 3*team.wins + 1*team.draws
		_, err := fmt.Fprintf(w, "%-30s | %2d | %2d | %2d | %2d | %2d\n", team.name, team.played, team.wins, team.draws, team.losses, points)
		if err != nil {
			return fmt.Errorf("failed to print data: %v", err)
		}
	}

	return nil
}
