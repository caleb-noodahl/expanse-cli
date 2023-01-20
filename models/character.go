package models

import (
	"bytes"
	"encoding/json"

	"fmt"

	"github.com/spf13/cobra"
)

type Character struct {
	Name            string           `json:"name"`
	Level           int              `json:"level"`
	Origin          Origin           `json:"origion"`
	Background      Background       `json:"background"`
	Talents         []Talent         `json:"talents"`
	Focus           []Focus          `json:"focus"`
	Specializations []Specialization `json:"specializations"`
	SocialClass     SocialClass      `json:"social_class"`
	Profession      Profession       `json:"profession"`
	Drive           Drive            `json:"drive"`
	Abilities       map[Ability]int  `json:"abilities"`
	Fortune         int              `json:"fortune"`
	Conditions      []Condition      `json:"conditions"`
}

func (c Character) JSON() ([]byte, error) {
	return json.Marshal(c)
}

func (c Character) CLIOutput() string {
	outbuf := bytes.Buffer{}
	outbuf.WriteString(fmt.Sprintf("name:         %s\n", c.Name))
	outbuf.WriteString(fmt.Sprintf("origin:       %s\n", c.Origin.String()))
	outbuf.WriteString(fmt.Sprintf("background:   %s\n", c.Background.String()))
	outbuf.WriteString(fmt.Sprintf("social class: %s\n", c.SocialClass.String()))
	outbuf.WriteString(fmt.Sprintf("profession:   %s\n", c.Profession.String()))
	outbuf.WriteString(fmt.Sprintf("drive:        %s\n", c.Drive.String()))
	outbuf.WriteString("\ntalents:\n")
	for _, talent := range c.Talents {
		outbuf.WriteString(fmt.Sprintf("-%s\n", talent.String()))
	}
	outbuf.WriteString("\nfocus:\n")
	for _, focus := range c.Focus {
		outbuf.WriteString(fmt.Sprintf("-%s\n", focus.String()))
	}
	outbuf.WriteString("\nabilities:\n")
	for ability, val := range c.Abilities {
		if ability == c.Background.BenefitDefinitions().Ability {
			outbuf.WriteString(fmt.Sprintf("-%s: %v + 1 (%v) *%s\n", ability.String(), val, AbilityScoreModifier(val+1), c.Background.String()))
		} else {
			outbuf.WriteString(fmt.Sprintf("-%s: %v (%v)\n", ability.String(), val, AbilityScoreModifier(val)))
		}

	}
	outbuf.WriteString("\nsecondary abilities:\n")
	outbuf.WriteString(fmt.Sprintf("-defense:   %v\n", 10+AbilityScoreModifier(c.Abilities[Dexterity])))
	outbuf.WriteString(fmt.Sprintf("-speed:     %v\n", 10+AbilityScoreModifier(c.Abilities[Dexterity])))
	outbuf.WriteString(fmt.Sprintf("-toughness: %v\n", AbilityScoreModifier(c.Abilities[Constitution])))
	outbuf.WriteString("\nfortune:\n")
	outbuf.WriteString(fmt.Sprintf("value: %v\n", c.Fortune))
	return outbuf.String()
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
