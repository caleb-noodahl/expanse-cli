package models

type Character struct {
	Name        string      `json:"name"`
	Origin      Origin      `json:"origion"`
	Background  Background  `json:"background"`
	SocialClass SocialClass `json:"social_class"`
	Profession  string      `json:"profession"`
	Drive       string      `json:"drive"`
	Stats       map[string]Stat
}

type Origin int
type SocialClass int
type Background int

const (
	Earther Origin = iota
	Martian
	Belter
)

const (
	Outsider SocialClass = iota
	LowerClass
	MiddleClass
	UppderClass
)

const (
	Bohemian Background = iota
	Exile
	Outcast
	Military
	Laborer
	Urban
	Academic
	Suburban
	Trade
	Aristocratic
	Corporate
	Cosmopolitan
)

func (o Origin) ToString() string {
	switch o {
	case Earther:
		return "earther"
	case Martian:
		return "martian"
	case Belter:
		return "belter"
	}
	return ""
}

func (c *Character) Generate(name, origin, background, social, prof, drive string, stats map[string]Stat) {

}

func (c *Character) Load(path string) {

}
