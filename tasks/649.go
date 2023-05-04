package tasks

// 649. Dota2 Senate

import (
	"log"
	"strings"
)

func Main() {
	log.Println(predictPartyVictory("RDD"))
}

func predictPartyVictory(senate string) string {
	for strings.Contains(senate, "R") && strings.Contains(senate, "D") {
		for i := 0; i < len(senate); i++ {
			if []rune(senate)[i] == []rune("D")[0] {
				if strings.Contains(senate[i:], "R") {
					senate = senate[:i] + strings.Replace(senate[i:], "R", "", 1)
				} else {
					senate = strings.Replace(senate, "R", "", 1)
				}
			} else {
				if strings.Contains(senate[i:], "D") {
					senate = senate[:i] + strings.Replace(senate[i:], "D", "", 1)
				} else {
					senate = strings.Replace(senate, "D", "", 1)
				}
			}
		}
	}

	if strings.Contains(senate, "R") {
		return "Radiant"
	} else {
		return "Dire"
	}
}
