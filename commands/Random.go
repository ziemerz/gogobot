package commands

import (
	"github.com/ziemerz/gogobot/utility"
)

var catBaseUrl string = "http://thecatapi.com/api/images/get?api_key=" + utility.GetKey("cat") + "&format=xml&type="
var dogBaseUrl string = "https://api.giphy.com/v1"

type Random struct {
	availableCommands []string
}

type CatGif struct {
	Url 	string	`xml:"data>images>image>url"`
}

type DogData struct {
	Data DogGif `json:"data"`
}

type DogGif struct {
	Url 	string	`json:"image_url"`
}

func randomCat() string {
	catgif := new(CatGif);
	utility.GetFromXML(catBaseUrl + "gif", &catgif)
	return catgif.Url
}

func randomDog() string {
	dogGif := new(DogData)
	utility.GetJson(dogBaseUrl + "/gifs/random?api_key=" + utility.GetKey("dog") + "&tag=dog", dogGif)
	return dogGif.Data.Url
}

func randomGif() string {
	return "Random gif"
}

func (r *Random) AvailableCommands() []string {
	return r.availableCommands
}

func NewRandom() *Random{
	subcmds := []string {"cat", "gif", "dog"}
	return &Random{subcmds}
}

func (rand *Random) FireCommand(as []string) string {
	if len(as) >= 2 {
		return "Firing subcommand of random"
	}
	return "Firing random command"
}

