package main

import "fmt"

/**
 * Thanks https://en.wikipedia.org/wiki/Round-robin_tournament#Scheduling_algorithm
 *
 * @param int weeks
 * @param array teams
 * @return array
 */
// added return type to the end
func schedule(weeks int, teams []int) [][][]int {

	//Check if team count is odd, if so we add a 'bye'
	if len(teams)%2 != 0 {
		teams = append(teams, len(teams)+1)
	}

	halfTeams := len(teams) / 2

	mTeams := teams[1:]
	count := len(mTeams)

	sched := make([][][]int, weeks)

	for i := 0; i < weeks; i++ {
		teamIndex := i % count

		sched[i] = make([][]int, halfTeams)

		sched[i][0] = make([]int, 2)

		sched[i][0][0] = teams[0]
		sched[i][0][1] = mTeams[teamIndex]

		for j := 1; j < halfTeams; j++ {
			teamA := (i + j) % count
			teamB := (i + count - j) % count

			sched[i][j] = make([]int, 2)

			sched[i][j][0] = mTeams[teamA]
			sched[i][j][1] = mTeams[teamB]
		}
	}

	return sched
}

func main() {
	weeks := 13
	teams := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	fmt.Println(schedule(weeks, teams))
}
