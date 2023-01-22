package models

import (
	"bytes"
	"encoding/json"

	"fmt"

	"github.com/caleb-noodahl/expanse-cli/utils"
	"github.com/google/uuid"
)

type Meta struct {
	AbilityPool int    `json:"ability_pool"`
	LastAbility string `json:"last_ability"`
	FocusPool   int    `json:"focus_pool"`
	LastFocus   string `json:"last_focus"`
}

type Character struct {
	Name            string                 `json:"name"`
	Level           int                    `json:"level"`
	Origin          Origin                 `json:"origion"`
	Background      Background             `json:"background"`
	Talents         map[Talent]int         `json:"talents"`
	Focus           map[Focus]int          `json:"focus"`
	Specializations map[Specialization]int `json:"specializations"`
	SocialClass     SocialClass            `json:"social_class"`
	Profession      Profession             `json:"profession"`
	Drive           Drive                  `json:"drive"`
	Abilities       map[Ability]int        `json:"abilities"`
	Fortune         int                    `json:"fortune"`
	Conditions      []Condition            `json:"conditions"`
	Meta            Meta                   `json:"meta"`
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
	for talent, _ := range c.Talents {
		outbuf.WriteString(fmt.Sprintf("-%s\n", talent.String()))
	}
	outbuf.WriteString("\nfocus:\n")
	for focus, val := range c.Focus {
		outbuf.WriteString(fmt.Sprintf("-%s +%v\n", focus.String(), val))
	}
	outbuf.WriteString("\nabilities:\n")
	for ability, val := range c.Abilities {
		if ability == c.Background.BenefitDefinitions().Ability {
			outbuf.WriteString(fmt.Sprintf("-%s: %v + 1 (%v) *%s\n", ability.String(), val, AbilityScoreModifier(val+1), c.Background.String()))
			continue
		}
		if ability != Defense && ability != Toughness && ability != Speed {
			outbuf.WriteString(fmt.Sprintf("-%s: %v (%v)\n", ability.String(), val, AbilityScoreModifier(val)))
		}
	}
	outbuf.WriteString("\nsecondary abilities:\n")
	outbuf.WriteString(fmt.Sprintf("-defense:   %v\n", c.Abilities[Defense]))
	outbuf.WriteString(fmt.Sprintf("-speed:     %v\n", c.Abilities[Speed]))
	outbuf.WriteString(fmt.Sprintf("-toughness: %v\n", c.Abilities[Toughness]))
	outbuf.WriteString("\nfortune:\n")
	outbuf.WriteString(fmt.Sprintf("value: %v\n", c.Fortune))
	outbuf.WriteString("\nmeta:\n")
	outbuf.WriteString(fmt.Sprintf("ability pool: %v\n", c.Meta.AbilityPool))
	outbuf.WriteString(fmt.Sprintf("last ability increased: %s\n\n", c.Meta.LastAbility))

	return outbuf.String()
}

func GenerateRandomCharacter() Character {
	profession := Profession(utils.Rand(23, 0))
	background := Background(utils.Rand(11, 0))
	focus := map[Focus]int{
		background.BenefitDefinitions().FocusPool[utils.Rand(len(background.BenefitDefinitions().FocusPool), 0)]: 1,
		profession.BenefitDefinitions().FocusPool[utils.Rand(len(profession.BenefitDefinitions().FocusPool), 0)]: 1,
	}

	drive := Drive(utils.Rand(11, 0))
	talents := map[Talent]int{
		background.BenefitDefinitions().TalentPool[utils.Rand(len(background.BenefitDefinitions().TalentPool), 0)]: 1,
		profession.BenefitDefinitions().TalentPool[utils.Rand(len(profession.BenefitDefinitions().TalentPool), 0)]: 1,
		drive.Talents()[utils.Rand(len(drive.Talents()), 0)]:                                                       1,
	}
	fortune := 15
	if _, ok := talents[Fortune]; ok {
		fortune = 20
	}
	character := Character{
		Name:            uuid.NewString(),
		Origin:          Origin(utils.Rand(3, 0)),
		Background:      background,
		Level:           1,
		Profession:      profession,
		SocialClass:     profession.SocialClass(),
		Drive:           drive,
		Conditions:      []Condition{},
		Specializations: map[Specialization]int{},
		Focus:           focus,
		Fortune:         fortune,
		Abilities: map[Ability]int{
			Accuracy:      utils.Rand(18, 3),
			Constitution:  utils.Rand(18, 3),
			Fighting:      utils.Rand(18, 3),
			Communication: utils.Rand(18, 3),
			Dexterity:     utils.Rand(18, 3),
			Intelligence:  utils.Rand(18, 3),
			Perception:    utils.Rand(18, 3),
			Strength:      utils.Rand(18, 3),
			Willpower:     utils.Rand(18, 3),
		},
	}
	character.Abilities[Defense] = 10 + AbilityScoreModifier(character.Abilities[Dexterity])
	character.Abilities[Speed] = 10 + AbilityScoreModifier(character.Abilities[Dexterity])
	character.Abilities[Toughness] = AbilityScoreModifier(character.Abilities[Constitution])
	return character
}
