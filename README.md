# expanse-cli

_cli dungemaster's toolbox for the expanse tabletop rpg_

## instructions

clone the repo

```bash
# pull back the repo
git clone https://github.com/caleb-noodahl/expanse-cli.git

# nav into the repo
cd expanse-cli

# make a build
go build

# have fun!
expanse-cli.exe roll        
> 1d6: 4
```

## commands

### character

character commands are routed by an _action_ flag `--a=[flag]`

**generate a character**:

a level 1 character is generated and it's character sheet is displayed along with a json representation.

```bash
# generate a new character
expanse-cli.exe char --a=gen

>character:
name:         6f273a65-85ac-46bd-a747-ab26938bd3ac
origin:       martian
background:   exile
social class: upper_class
profession:   expert
drive:        builder

talents:
-affluent
-know_it_all
-maker

focus:
-self_discipline
-

abilities:
-communication: 3 (-2)
-willpower: 14 (2)
-fighting: 5 (-1)
-constitution: 10 + 1 (1) *exile
-dexterity: 14 (2)
-intelligence: 11 (1)
-perception: 10 (1)
-strength: 6 (0)
-accuracy: 5 (-1)

secondary abilities:
-defense:   12
-speed:     12
-toughness: 1

fortune:
value: 15
json:
{"name":"6f273a65-85ac-46bd-a747-ab26938bd3ac","level":1,"origion":1,"background":1,"talents":[0,17,20],"focus":[63,64],"specializations":[],"social_class":3,"profession":21,"drive":0,"abilities":{"0":5,"1":10,"2":5,"3":3,"4":14,"5":11,"6":10,"7":6,"8":14},"fortune":15,"conditions":[]}
```

**character builder**:

walks the user through character creation wizard.
choices can be selected by sending their corresponding index, eg. `[0], [1], [2]...`

```bash
expanse-cli.exe char --a=wiz
beginning character wizard

step #1: origin
[0]:earther, [1]:martian, [2]:belter
1

step #2: abilities
point buy?
[0]: no, [1]: yes
1
points remaining: 12
spend how many on accuracy?(0 - 3)
0
points remaining: 12
spend how many on constitution?(0 - 3)
2
points remaining: 10
spend how many on fighting?(0 - 3)
0
points remaining: 10
spend how many on communication?(0 - 3)
3
points remaining: 7
spend how many on dexterity?(0 - 3)
2
points remaining: 5
spend how many on intelligence?(0 - 3)
3
points remaining: 2
spend how many on perception?(0 - 3)
2
points remaining: 0
spend how many on strength?(0 - 3)
0
points remaining: 0
spend how many on willpower?(0 - 3)
0


step #3: background
[0]: bohemian
[1]: exile
[2]: outcast
[3]: military
[4]: laborer
[5]: urban
[6]: academic
[7]: suburban
[8]: trade
[9]: aristocratic
[10]: corporate
[11]: cosmopolitan
3

background Benefits
ability: fighting + 1

step #3: background focus
[0]: pistols
[1]: tactics
1

step #4: background talent
[0]: dual_weapon_style
[1]: grappling_style
[2]: grappling_style
[3]: pistol_style
[4]: rifle_style
[5]: self_defense_style
[6]: single_weapon_style
[7]: striking_style
[8]: thrown_weapon_style
[9]: two_handed_style
[10]: observation
4

step #5: profession
[0]: brawler
[1]: survivalist
[2]: criminal
[3]: scavenger
[4]: fixer
[5]: artist
[6]: athlete
[7]: soldier
[8]: investigator
[9]: technician
[10]: clergy
[11]: negotiator
[12]: pilot
[13]: security
[14]: professional
[15]: scholar
[16]: merchant
[17]: politician
[18]: commander
[19]: explorer
[20]: dilettante
[21]: expert
[22]: executive
[23]: socalite
13
socal class (derived from profession): middle_class

step #6: profession focus
[0]: security
[1]: empathy
[2]: intuition
[3]: seeing
3

step #7: profession talent
[0]: dual_weapon_style
[1]: grappling_style
[2]: overwhelm_style
[3]: pistol_style
[4]: rifle_style
[5]: self_defense_style
[6]: single_weapon_style
[7]: striking_style
[8]: thrown_weapon_style
4

step #8: drive
[0]: builder
[1]: caregiver
[2]: ecstatic
[3]: judge
[4]: leader
[5]: networker
[6]: penitent
[7]: protector
[8]: rebel
[9]: survivor
[10]: visionary
[11]: fortune
6

step #9: drive talent
[0]: fringer
[1]: know_it_all
0


character: 
name:
origin:       martian
background:   military
social class: middle_class
profession:   security
drive:        penitent

talents:
-rifle_style
-rifle_style
-fringer

focus:
-tactics
-seeing

abilities:
-dexterity: 14 (2)
-intelligence: 17 (3)
-perception: 14 (2)
-accuracy: 8 (0)
-fighting: 8 + 1 (1) *military
-communication: 17 (3)
-constitution: 14 (2)
-strength: 8 (0)
-willpower: 8 (0)

secondary abilities:
-defense:   12
-speed:     12
-toughness: 2

fortune:
value: 15
```
