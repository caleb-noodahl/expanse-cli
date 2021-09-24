package models

type Character struct {
	Name        string `json:"name"`
	Origin      string `json:"origion"`
	Background  string `json:"background"`
	SocialClass string `json:"social_class"`
	Profession  string `json:"profession"`
	Drive       string `json:"drive"`
	Stats       map[string]Stat
}

func (c *Character) Generate(name, origin, background, social, prof, drive string, stats map[string]Stat) {

}

func (c *Character) Load(path string) {

}
