package models

import (
	"bytes"
	"encoding/json"
	"exp/utils"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func (c *Character) GenerateCharacter(cmd *cobra.Command, args []string) {
	//origin
	fmt.Printf("generating new character:\n\n")
	profession := Profession(utils.Rand(23, 0))
	background := Background(utils.Rand(11, 0))
	focus := []Focus{
		background.BenefitDefinitions().FocusPool[utils.Rand(len(background.BenefitDefinitions().FocusPool), 0)],
		profession.BenefitDefinitions().FocusPool[utils.Rand(len(profession.BenefitDefinitions().FocusPool), 0)],
	}

	drive := Drive(utils.Rand(11, 0))
	talents := []Talent{
		background.BenefitDefinitions().TalentPool[utils.Rand(len(background.BenefitDefinitions().TalentPool), 0)],
		profession.BenefitDefinitions().TalentPool[utils.Rand(len(profession.BenefitDefinitions().TalentPool), 0)],
		drive.Talents()[utils.Rand(len(drive.Talents()), 0)],
	}
	fortune := 15
	if talents[1] == Fortune {
		fortune = 20
	}

	character := Character{
		Name:            uuid.NewString(),
		Origin:          Origin(utils.Rand(3, 0)),
		Background:      background,
		Level:           1,
		SocialClass:     profession.SocialClass(),
		Profession:      profession,
		Drive:           drive,
		Talents:         talents,
		Conditions:      []Condition{},
		Specializations: []Specialization{},
		Focus:           focus,
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
		Fortune: fortune,
	}
	fmt.Printf("character:\n%+v", character.CLIOutput())
	j, _ := character.JSON()
	fmt.Printf("json:\n%+v\n", string(j))
}

func (c *Character) LoadCharacter(cmd *cobra.Command, args []string) {
	fmt.Printf("loading character: %s\n", args)

}

func pointBuyPrompt(cmd *cobra.Command, args []string) map[Ability]int {
	out := map[Ability]int{}
	var input string
	total := 12
	for i := 0; i <= 8; i++ {
		fmt.Printf("points remaining: %v\n", total)
		fmt.Printf("spend how many on %s?(0 - 3)\n", Ability(i))
		fmt.Scan(&input)
		parsed, _ := strconv.Atoi(input)
		abilityScoreValue := 0
		switch parsed {
		case 0:
			abilityScoreValue = 8
		case 1:
			abilityScoreValue = 11
		case 2:
			abilityScoreValue = 14
		case 3:
			abilityScoreValue = 17
		}
		out[Ability(i)] = abilityScoreValue
		total = total - parsed
	}

	return out
}

func rollAssignPrompt(cmd *cobra.Command, args []string) map[Ability]int {
	out := map[Ability]int{}
	fmt.Printf("\nability rolls:\n")
	abilityCnt := 8
	rolls := []int{}
	for i := 0; i <= abilityCnt; i++ {
		r := utils.Rand(18, 3)
		rolls = append(rolls, r)
		fmt.Printf("[%v]: %v: (%v)\n", i, r, AbilityScoreModifier(r))
	}
	var input string
	for i := 0; i <= abilityCnt; i++ {
		fmt.Printf("%s which roll?:\n", Ability(i).String())
		fmt.Scan(&input)
		parsed, _ := strconv.Atoi(input)
		out[Ability(i)] = parsed
		fmt.Printf("\n%s: %v\n", Ability(i).String(), rolls[parsed])
	}
	return out
}

func (c *Character) Wizard(cmd *cobra.Command, args []string) {
	var (
		selection string
		parsed    int
		err       error
	)
	c = &Character{
		Fortune: 15,
	}
	fmt.Printf("beginning character wizard\n")
	var name string
	fmt.Printf("whats your characters name?\n")
	fmt.Scan(&name)
	c.Name = name
	fmt.Printf("\nstep #1: origin\n")
	fmt.Printf("[%v]:%s, [%v]:%s, [%v]:%s\n", 0, Origin(0).String(), 1, Origin(1).String(), 2, Origin(2).String())
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Origin = Origin(parsed)
	if err != nil {
		fmt.Printf("error:\n%s", err)
	}
	fmt.Printf("\nstep #2: abilities\n")
	fmt.Printf("point buy?\n[0]: no, [1]: yes\n")
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	if parsed == 1 {
		c.Abilities = pointBuyPrompt(cmd, args)
	} else {
		c.Abilities = rollAssignPrompt(cmd, args)
	}

	fmt.Print("\nstep #3: background\n")
	for i := 0; i <= 11; i++ {
		fmt.Printf("[%v]: %s\n", i, Background(i))
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Background = Background(parsed)

	fmt.Print("\nbackground Benefits\n")
	fmt.Printf("ability: %s + %v\n", c.Background.BenefitDefinitions().Ability.String(), c.Background.BenefitDefinitions().AbilityModifier)
	fmt.Printf("\nstep #3: background focus\n")
	for i, bkgrnd := range c.Background.BenefitDefinitions().FocusPool {
		fmt.Printf("[%v]: %s\n", i, bkgrnd)
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Focus = []Focus{c.Background.BenefitDefinitions().FocusPool[parsed]}
	fmt.Printf("\nstep #4: background talent\n")
	for i, talent := range c.Background.BenefitDefinitions().TalentPool {
		fmt.Printf("[%v]: %s\n", i, talent)
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Talents = []Talent{c.Background.BenefitDefinitions().TalentPool[parsed]}

	fmt.Printf("\nstep #5: profession\n")
	for i := 0; i <= 23; i++ {
		fmt.Printf("[%v]: %s\n", i, Profession(i).String())
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Profession = Profession(parsed)
	c.SocialClass = c.Profession.SocialClass()
	fmt.Printf("socal class (derived from profession): %s\n", c.Profession.SocialClass().String())
	fmt.Print("\nstep #6: profession focus\n")
	for i, focus := range c.Profession.BenefitDefinitions().FocusPool {
		fmt.Printf("[%v]: %s\n", i, focus.String())
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Focus = append(c.Focus, c.Profession.BenefitDefinitions().FocusPool[parsed])
	fmt.Printf("\nstep #7: profession talent\n")
	for i, talent := range c.Profession.BenefitDefinitions().TalentPool {
		fmt.Printf("[%v]: %s\n", i, talent.String())
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Talents = append(c.Talents, c.Profession.BenefitDefinitions().TalentPool[parsed])

	fmt.Printf("\nstep #8: drive\n")
	for i := 0; i <= 11; i++ {
		fmt.Printf("[%v]: %s\n", i, Drive(i).String())
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Drive = Drive(parsed)

	fmt.Printf("\nstep #9: drive talent\n")
	for i, talent := range c.Drive.Talents() {
		fmt.Printf("[%v]: %s\n", i, talent.String())
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Talents = append(c.Talents, c.Drive.Talents()[parsed])
	fmt.Printf("\n\ncharacter: %+v\n", c.CLIOutput())
	root := viper.GetString("data.root")
	characters := viper.GetString("data.characters")
	fmt.Printf("\nsaving to %s%s%s.json\n", root, characters, c.Name)
	file, err := os.OpenFile(fmt.Sprintf("%s%s%s.json", root, characters, c.Name), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("\nerror:%s\n", err)
		return
	}
	defer file.Close()
	raw, _ := c.JSON()
	file.Write(raw)
	fmt.Printf("\nwrote %v bytes to %s%s\n", len(raw), root, characters)
}

func (c *Character) UpdateCharacter(cmd *cobra.Command, args []string) {
	fmt.Printf("updating character: %s\n", args)

}

func (c *Character) AddCharacterNote(cmd *cobra.Command, args []string) {
	fmt.Printf("adding character note: %s\n", args)
}
