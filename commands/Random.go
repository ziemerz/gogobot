package commands

import (
	"github.com/ziemerz/gogobot/utility"
	"github.com/ziemerz/gogobot/types"
)

var catBaseUrl string = "http://thecatapi.com/api/images/get?api_key=" + utility.GetKey("cat") + "&format=xml&type="
var dogBaseUrl string = "https://api.giphy.com/v1"

type RandomWrapper struct {
	Random Random	`json:"random"`
}
type Random struct {
	SubCommands []types.SubCmd `json:"subcommands"`
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

func (r *Random) AvailableCommands() []types.SubCmd {
	return r.SubCommands
}

func NewRandom() *Random{
	var rw RandomWrapper
	utility.GetJsonFromFile("commands.json", &rw)
	return &rw.Random
}

func (rand *Random) FireCommand(as []string) string {
	if len(as) >= 2 {
		switch as[1] {
		case "cat":
			return randomCat()
		case "dog":
			return randomDog()
		}
	}
	return "Firing random command"
}

