package getratings

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"net/http"
	"strings"
)

//GetTrailer takes in a movie name as argument and searches youtube for "<movie name> movie trailer"
//and returns the first result
func GetTrailer(mname string) string {
	MovieName := strings.Replace(mname, " ", "_", 9)
	urlis := "https://www.googleapis.com/youtube/v3/search?part=snippet&q=" + MovieName + "+movie+trailer&key=AIzaSyCXqzqRJUDGsQUO44y-G8PM_qxVpgd66GE"
	var AllVideoArray []string
	res, err := http.Get(urlis)
	if err != nil {
		fmt.Println(err)
	}
	v, err := jason.NewObjectFromReader(res.Body)
	if err != nil {
		fmt.Println("An error occured!")
	}
	ItemsArr, _ := v.GetObjectArray("items")
	for _, item := range ItemsArr {
		id, _ := item.GetObject("id")
		Vid, _ := id.GetString("videoId")
		AllVideoArray = append(AllVideoArray, Vid)
	}
	MovieTrailerUrl := "https://www.youtube.com/embed/" + AllVideoArray[0]
	return MovieTrailerUrl
}
