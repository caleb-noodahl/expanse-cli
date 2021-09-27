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
exp.exe roll        
> 1d6: 4
```

## commands

### character

character commands are routed by an _action_ flag `--a=[flag]`

**generate a character**:

a level 1 character is generated and it's character sheet is displayed along with a json representation. 

```bash
# generate a new character
exp.exe char --a=gen

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

```bash
```