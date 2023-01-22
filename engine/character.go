package engine

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/caleb-noodahl/expanse-cli/models"
	"github.com/caleb-noodahl/expanse-cli/utils"

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
		Level:           1,
		Specializations: make(map[models.Specialization]int),
		Conditions:      make([]models.Condition, 0),
		Fortune:         15,
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
	//default abilities
	c.Abilities[models.Toughness] = 10 + models.AbilityScoreModifier(c.Abilities[models.Dexterity])
	c.Abilities[models.Speed] = 10 + models.AbilityScoreModifier(c.Abilities[models.Dexterity])
	c.Abilities[models.Toughness] = models.AbilityScoreModifier(c.Abilities[models.Constitution])

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
	c.Focus = map[models.Focus]int{
		c.Background.BenefitDefinitions().FocusPool[parsed]: 1,
	}
	fmt.Printf("\nstep #4: background talent\n")
	for i, talent := range c.Background.BenefitDefinitions().TalentPool {
		fmt.Printf("[%v]: %s\n", i, talent)
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Talents = map[models.Talent]int{
		c.Background.BenefitDefinitions().TalentPool[parsed]: 1,
	}

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
	c.Focus[c.Profession.BenefitDefinitions().FocusPool[parsed]]++
	fmt.Printf("\nstep #7: profession talent\n")
	for i, talent := range c.Profession.BenefitDefinitions().TalentPool {
		fmt.Printf("[%v]: %s\n", i, talent.String())
	}
	fmt.Scan(&selection)
	parsed, _ = strconv.Atoi(selection)
	c.Talents[c.Profession.BenefitDefinitions().TalentPool[parsed]]++

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
	c.Talents[c.Drive.Talents()[parsed]]++

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

func LevelUpCharacterWizard(cmd *cobra.Command, args []string) {
	fmt.Printf("leveling up character, %s\n", args)
	if len(args) > 0 {
		selection := ""
		charBytes := utils.ReadJSONFile(args[0])
		c := models.Character{}

		if err := json.Unmarshal(charBytes, &c); err != nil {
			log.Panic(err)
		}
		c.Level++
		//fortune
		if c.Level >= 2 && c.Level <= 10 {
			c.Fortune += 3
		} else {
			c.Fortune += 2
		}
		//toughness or defense
		if c.Level%4 == 0 {
			fmt.Printf("\n[0]: defense or [1]: toughness?\n")
			fmt.Scan(&selection)
			parsed, _ := strconv.Atoi(selection)
			if parsed == 0 {
				c.Abilities[models.Defense]++
			}
			if parsed == 1 {
				c.Abilities[models.Toughness]++
			}
		}
		fmt.Printf("ability to increase\n")
		for i := 0; i <= 8; i++ {
			fmt.Printf("[%v] %s\n", i, models.Ability(i).String())
		}
		fmt.Scan(&selection)
		parsed, _ := strconv.Atoi(selection)

		ability := models.Ability(parsed)
		//the same ability can't be increased twice just add it to the ability pool

		//disqualifiers
		if c.Meta.LastAbility == ability.String() {
			c.Meta.AbilityPool++
		} else {
			val := c.Abilities[ability]
			if val <= 5 {
				c.Abilities[ability]++
			}
			if val >= 6 && val <= 8 {
				if c.Meta.AbilityPool+1 >= 2 {
					c.Abilities[ability]++
				} else {
					fmt.Printf("not enough to advance ability - adding to ability pool\n")
					c.Meta.AbilityPool++
				}
			}
			if val >= 9 && val <= 12 {
				if c.Meta.AbilityPool+1 >= 3 {
					c.Abilities[ability]++
				} else {
					fmt.Printf("not enough to advance ability - adding to ability pool\n")
					c.Meta.AbilityPool++
				}
			}
		}
		//ability focus
		fmt.Printf("\nability focus\n")
		for i := 0; i <= 65; i++ {
			fmt.Printf("[%v] %s\n", i, models.Focus(i).String())
		}

		fmt.Scan(&selection)
		parsed, _ = strconv.Atoi(selection)
		//disqualifiers
		var (
			//might be a bug here - todo:review
			focusCap  bool = c.Focus[models.Focus(parsed)] > 1 && c.Level < 11
			lastFocus bool = c.Meta.LastAbility == models.Focus(parsed).String()
		)
		if focusCap || lastFocus {
			fmt.Printf("warning! cannot apply focus: level cap: %v, last focus: %s:%v\n", focusCap, c.Meta.LastFocus, lastFocus)
			c.Meta.FocusPool++
		} else {
			c.Focus[models.Focus(parsed)]++
		}

		//talent improvement
		var (
			specalizationChosen  bool
			specalizationAllowed bool = c.Level >= 4 && c.Level <= 16 && c.Level%2 == 0
		)
		fmt.Printf("specalizationAllowed: %v", specalizationAllowed)
		fmt.Printf("%v, %v, %v, %v", c.Level >= 4, c.Level <= 16, c.Level%2 == 0, c.Level)
		if specalizationAllowed {
			fmt.Printf("\ntake specalization?\n[0]: yes [1]: no\n")
			fmt.Scan(&selection)
			parsed, _ = strconv.Atoi(selection)
			specalizationChosen = parsed == 0

			if specalizationChosen {
				fmt.Printf("\nspecalization improvement\n")
				for i := 0; i <= 12; i++ {
					fmt.Printf("[%v] %s\n", i, models.Specialization(i).String())
				}
				fmt.Scan(&selection)
				parsed, _ = strconv.Atoi(selection)
				c.Specializations[models.Specialization(parsed)]++
			}
		}

		fmt.Printf("\ntalent improvement\n")
		if !specalizationChosen {
			for i := 0; i <= 39; i++ {
				fmt.Printf("[%v] %s\n", i, models.Focus(i).String())
			}
			fmt.Scan(&selection)
			parsed, _ = strconv.Atoi(selection)
			c.Talents[models.Talent(parsed)]++
		}

		//income

		//goals

		//save character
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

	}
}

func GenerateCharacter(cmd *cobra.Command, args []string) {
	fmt.Printf("generating new character:\n\n")
	character := models.GenerateRandomCharacter()

	fmt.Printf("character:\n%+v", character.CLIOutput())
	j, _ := character.JSON()
	fmt.Printf("\njson:\n%+v\n", string(j))
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

func LoadCharacter(cmd *cobra.Command, args []string) {
	fmt.Printf("loading character: %s\n", args)

}

func UpdateCharacter(cmd *cobra.Command, args []string) {
	fmt.Printf("updating character: %s\n", args)

}

func AddCharacterNote(cmd *cobra.Command, args []string) {
	fmt.Printf("adding character note: %s\n", args)
}
