package main

import "fmt"

type ratings struct {
	score []string
}

// creating a struct and embeding another struct
type favMovie struct {
	movieName []string
	song      []string
	favMeter  ratings
}

func main() {
	// assigning value to a struct
	fav := favMovie{
		movieName: []string{"99 Songs", "la la land"},
		song:      []string{"The Oracle", "city of stars"},
		favMeter: ratings{
			score: []string{"3.5", "4"},
		},
	}

	fav.updateFavList("minnale", "bgm", "5")
	fmt.Println(fav)
}

// receiving value and that value type should match favMovie
func (movieL *favMovie) updateFavList(newMovie string, newSong string, newScore string) {
	// fetching the value from the memory address using pointer
	(*movieL).movieName = append((*movieL).movieName, newMovie)
	(*movieL).song = append((*movieL).song, newSong)
	(*movieL).favMeter.score = append((*movieL).favMeter.score, newScore)
}
