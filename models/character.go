package models

import (
	"exp/utils"
	"fmt"

	"github.com/spf13/cobra"
)

type Character struct {
	Name        string          `json:"name"`
	Origin      Origin          `json:"origion"`
	Background  Background      `json:"background"`
	SocialClass SocialClass     `json:"social_class"`
	Profession  string          `json:"profession"`
	Drive       string          `json:"drive"`
	Abilities   map[Ability]int `json:"abilities"`
}

func (c *Character) GenerateCharacter(cmd *cobra.Command, args []string) {
	//origin
	abilites := map[Ability]int{
		Accuracy:      utils.Rand(18, 3),
		Constitution:  utils.Rand(18, 3),
		Fighting:      utils.Rand(18, 3),
		Communication: utils.Rand(18, 3),
		Dexterity:     utils.Rand(18, 3),
		Intelligence:  utils.Rand(18, 3),
		Perception:    utils.Rand(18, 3),
		Strength:      utils.Rand(18, 3),
		Willpower:     utils.Rand(18, 3),
	}
	highest := []Ability{}
	highestVal := -3
	for key, val := range abilites {
		if val > highestVal {
			highest = []Ability{key}
			continue
		}
		if val == highestVal {
			highest = append(highest, key)
			continue
		}
	}
	origin := Origin(utils.Rand(3, 0))
	profession := Profession(utils.Rand(23, 0))

	fmt.Printf("highest ability: %+v\n", highest)
	fmt.Printf(" (ability has following focus options: %+v)\n", highest[0].Focus())
	fmt.Printf("origin: %s\n", origin.String())
	fmt.Printf("profession: %s\n", profession.String())
	fmt.Printf("social class: %s\n", profession.SocialClass().String())
	fmt.Printf("abilities:\n")
	for ability, val := range abilites {
		fmt.Printf(" %s:%v (%v)\n", ability.String(), val, AbilityScoreModifier(val))
	}

}

func (c *Character) LoadCharacter(cmd *cobra.Command, args []string) {
	fmt.Printf("loading character: %s\n", args)

}

func (c *Character) UpdateCharacter(cmd *cobra.Command, args []string) {
	fmt.Printf("updating character: %s\n", args)

}

func (c *Character) AddCharacterNote(cmd *cobra.Command, args []string) {
	fmt.Printf("adding character note: %s\n", args)
}
