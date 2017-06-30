package commands

import (
	"github.com/ziemerz/gogobot/types"
	"github.com/ziemerz/gogobot/utility"
)

var catBaseUrl string = "http://thecatapi.com/api/images/get?api_key=" + utility.GetKey("cat") + "&format=xml&type="
var dogBaseUrl string = "https://api.giphy.com/v1"

type Random struct {
	subCommands map[string]types.Func
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

func (r *Random) SubCommands() map[string]types.Func {
	return r.subCommands
}

func NewRandom() *Random{
	subcmds := make(map[string]types.Func)
	subcmds["cat"] = randomCat
	subcmds["gif"] = randomGif
	subcmds["dog"] = randomDog
	return &Random{subcmds}
}

