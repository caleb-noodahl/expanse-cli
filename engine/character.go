package engine

import (
	"fmt"
	"os"
	"strconv"

	"github.com/caleb-noodahl/expanse-cli/models"
	"github.com/caleb-noodahl/expanse-cli/utils"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CharacterWizard(cmd *cobra.Command, args []string) {
	var (
		selection string
		parsed    int
		err       error
	)
	c := models.Character{
		Fortune: 15,
	}
	fmt.Printf("beginning character wizard\n")
	var name string
	fmt.Printf("whats your characters name?\n")
	fmt.Scan(&name)
	c.Name = name
	fmt.Printf("\nstep #1: origin\n")
	fmt.Printf("[%v]:%s, [%v]:%s, [%v]:%s\n", 0, models.Origin(0).String(), 1, models.Origin(1).String(), 2, models.Origin(2).String())
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Origin = models.Origin(parsed)
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
		fmt.Printf("[%v]: %s\n", i, models.Background(i))
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Background = models.Background(parsed)

	fmt.Print("\nbackground Benefits\n")
	fmt.Printf("ability: %s + %v\n", c.Background.BenefitDefinitions().Ability.String(), c.Background.BenefitDefinitions().AbilityModifier)
	fmt.Printf("\nstep #3: background focus\n")
	for i, bkgrnd := range c.Background.BenefitDefinitions().FocusPool {
		fmt.Printf("[%v]: %s\n", i, bkgrnd)
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Focus = []models.Focus{c.Background.BenefitDefinitions().FocusPool[parsed]}
	fmt.Printf("\nstep #4: background talent\n")
	for i, talent := range c.Background.BenefitDefinitions().TalentPool {
		fmt.Printf("[%v]: %s\n", i, talent)
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Talents = []models.Talent{c.Background.BenefitDefinitions().TalentPool[parsed]}

	fmt.Printf("\nstep #5: profession\n")
	for i := 0; i <= 23; i++ {
		fmt.Printf("[%v]: %s\n", i, models.Profession(i).String())
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Profession = models.Profession(parsed)
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
		fmt.Printf("[%v]: %s\n", i, models.Drive(i).String())
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Drive = models.Drive(parsed)

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

func GenerateCharacter(cmd *cobra.Command, args []string) {
	fmt.Printf("generating new character:\n\n")
	profession := models.Profession(utils.Rand(23, 0))
	background := models.Background(utils.Rand(11, 0))
	focus := []models.Focus{
		background.BenefitDefinitions().FocusPool[utils.Rand(len(background.BenefitDefinitions().FocusPool), 0)],
		profession.BenefitDefinitions().FocusPool[utils.Rand(len(profession.BenefitDefinitions().FocusPool), 0)],
	}

	drive := models.Drive(utils.Rand(11, 0))
	talents := []models.Talent{
		background.BenefitDefinitions().TalentPool[utils.Rand(len(background.BenefitDefinitions().TalentPool), 0)],
		profession.BenefitDefinitions().TalentPool[utils.Rand(len(profession.BenefitDefinitions().TalentPool), 0)],
		drive.Talents()[utils.Rand(len(drive.Talents()), 0)],
	}
	fortune := 15
	if talents[1] == models.Fortune {
		fortune = 20
	}

	character := models.Character{
		Name:            uuid.NewString(),
		Origin:          models.Origin(utils.Rand(3, 0)),
		Background:      background,
		Level:           1,
		SocialClass:     profession.SocialClass(),
		Profession:      profession,
		Drive:           drive,
		Talents:         talents,
		Conditions:      []models.Condition{},
		Specializations: []models.Specialization{},
		Focus:           focus,
		Abilities: map[models.Ability]int{
			models.Accuracy:      utils.Rand(18, 3),
			models.Constitution:  utils.Rand(18, 3),
			models.Fighting:      utils.Rand(18, 3),
			models.Communication: utils.Rand(18, 3),
			models.Dexterity:     utils.Rand(18, 3),
			models.Intelligence:  utils.Rand(18, 3),
			models.Perception:    utils.Rand(18, 3),
			models.Strength:      utils.Rand(18, 3),
			models.Willpower:     utils.Rand(18, 3),
		},
		Fortune: fortune,
	}
	fmt.Printf("character:\n%+v", character.CLIOutput())
	j, _ := character.JSON()
	fmt.Printf("json:\n%+v\n", string(j))
}

func pointBuyPrompt(cmd *cobra.Command, args []string) map[models.Ability]int {
	out := map[models.Ability]int{}
	var input string
	total := 12
	for i := 0; i <= 8; i++ {
		fmt.Printf("points remaining: %v\n", total)
		fmt.Printf("spend how many on %s?(0 - 3)\n", models.Ability(i))
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
		out[models.Ability(i)] = abilityScoreValue
		total = total - parsed
	}

	return out
}

func rollAssignPrompt(cmd *cobra.Command, args []string) map[models.Ability]int {
	out := map[models.Ability]int{}
	fmt.Printf("\nability rolls:\n")
	abilityCnt := 8
	rolls := []int{}
	for i := 0; i <= abilityCnt; i++ {
		r := utils.Rand(18, 3)
		rolls = append(rolls, r)
		fmt.Printf("[%v]: %v: (%v)\n", i, r, models.AbilityScoreModifier(r))
	}
	var input string
	for i := 0; i <= abilityCnt; i++ {
		fmt.Printf("%s which roll?:\n", models.Ability(i).String())
		fmt.Scan(&input)
		parsed, _ := strconv.Atoi(input)
		out[models.Ability(i)] = parsed
		fmt.Printf("\n%s: %v\n", models.Ability(i).String(), rolls[parsed])
	}
	return out
}
