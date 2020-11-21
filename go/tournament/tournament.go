package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	rowFormat = "%-30s | %2v | %2v | %2v | %2v | %2v\n"
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
	allTeams := make(map[string]team)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '#' {
			continue
		}

		details := strings.Split(line, ";")
		if len(details) != 3 {
			return fmt.Errorf("failed to parse match details, expected 3 entries, got %d entries", len(details))
		}

		t1, t2, result := details[0], details[1], details[2]

		homeTeam := allTeams[t1]
		awayTeam := allTeams[t2]

		homeTeam.name = t1
		awayTeam.name = t2
		homeTeam.played++
		awayTeam.played++

		switch result {
		case "win":
			homeTeam.wins++
			homeTeam.points += 3
			awayTeam.losses++
		case "loss":
			homeTeam.losses++
			awayTeam.wins++
			awayTeam.points += 3
		case "draw":
			homeTeam.draws++
			homeTeam.points++
			awayTeam.draws++
			awayTeam.points++
		default:
			return fmt.Errorf("invalid result %s, expected one of [win, loss, draw]", result)
		}
		allTeams[t1] = homeTeam
		allTeams[t2] = awayTeam
	}

	var t []team
	for _, team := range allTeams {
		t = append(t, team)
	}
	sort.Slice(t, func(i, j int) bool {
		if t[i].points == t[j].points {
			return t[i].name < t[j].name
		}
		return t[i].points > t[j].points
	})

	_, err := fmt.Fprintf(w, rowFormat, "Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return fmt.Errorf("failed to print results header: %v", err)
	}

	for _, team := range t {
		_, err := fmt.Fprintf(w, rowFormat, team.name, team.played, team.wins, team.draws, team.losses, team.points)
		if err != nil {
			return fmt.Errorf("failed to print results data: %v", err)
		}
	}

	return nil
}
