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

// byPoints implements sort.Interface for []*team based on the points
// each team has then by the team name if points are the same.
type byPoints []*team

func (p byPoints) Len() int      { return len(p) }
func (p byPoints) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p byPoints) Less(i, j int) bool {
	if p[i].points == p[j].points {
		return p[i].name < p[j].name
	}
	return p[i].points > p[j].points
}

// Tally reads a scorecard and writes the results.
func Tally(r io.Reader, w io.Writer) error {
	allTeams := make(map[string]*team)
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

		home, away, result := details[0], details[1], details[2]

		if _, ok := allTeams[home]; !ok {
			allTeams[home] = &team{name: home}
		}

		if _, ok := allTeams[away]; !ok {
			allTeams[away] = &team{name: away}
		}

		allTeams[home].played++
		allTeams[away].played++

		switch result {
		case "win":
			allTeams[home].wins++
			allTeams[home].points += 3
			allTeams[away].losses++
		case "loss":
			allTeams[home].losses++
			allTeams[away].wins++
			allTeams[away].points += 3
		case "draw":
			allTeams[home].draws++
			allTeams[home].points += 1
			allTeams[away].draws++
			allTeams[away].points += 1
		default:
			return errors.New("invalid result")
		}
	}

	var t byPoints
	for _, team := range allTeams {
		t = append(t, team)
	}
	sort.Sort(byPoints(t))

	_, err := fmt.Fprintf(w, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return fmt.Errorf("failed to print header: %v", err)
	}

	for _, team := range t {
		_, err := fmt.Fprintf(w, "%-30s | %2d | %2d | %2d | %2d | %2d\n", team.name, team.played, team.wins, team.draws, team.losses, team.points)
		if err != nil {
			return fmt.Errorf("failed to print data: %v", err)
		}
	}

	return nil
}
