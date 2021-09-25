package models

type Origin int
type SocialClass int
type Background int
type Profession int
type Ability int
type Focus int
type Talent int
type Specialization int
type Drive int

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

const (
	Accuracy Ability = iota
	Constitution
	Fighting
	Communication
	Dexterity
	Intelligence
	Perception
	Strength
	Willpower
)

const (
	Affluent Talent = iota
	Agility
	Artistry
	Attractive
	Burglary
	Carousing
	Command
	Contacts
	Doctor
	DualWeaponStyle
	Fringer
	ExpertiseFocus
	GrapplingStyle
	Hacking
	Improvisation
	Inspire
	Intrigue
	KnowItAll
	Knowledge
	Linguistics
	Maker
	Medic
	Misdirection
	Observation
	Oratory
	OverwhelmStyle
	Performance
	PilotTalent
	PinpointAccuracy
	QuickReflexes
	PistolStyle
	RifleStyle
	Protector
	Scouting
	SelfDefenseStyle
	TacticalAwareness
	SingleWeaponStyle
	ThrownWeaponStyle
	TwoHandedStyle
	StrikingStyle
)

const (
	AcademicSpecialization Specialization = iota
	AceSpecialization
	AgentSpecialization
	CommandoSpecialization
	ExecutiveSpecialization
	GunfighterSpecialization
	HackerSpecialization
	InvestigatorSpecialization
	MartialSpecialization
	ArtistSpecialization
	SniperSpecialization
	SocaliteSpecialization
	StarSpecialization
)

const (
	Brawler Profession = iota
	Survivalist
	Criminal
	Scavenger
	Fixer
	Artist
	Athlete
	Soldier
	Investigator
	Technician
	Clergy
	Negotiator
	Pilot
	Security
	Professional
	Scholar
	Merchant
	Politician
	Commander
	Explorer
	Dilettante
	Expert
	Executive
	Socialite
)

const (
	Bows Focus = iota
	Gunnery
	Pistols
	Rifles
	Throwing
	Bargaining
	Deception
	Disguise
	Etiquette
	Expression
	Gambling
	Investigation
	Leadership
	Performing
	Persuasion
	Seduction
	Running
	Stamina
	Swimming
	Tolerance
	Acrobatics
	Crafting
	Driving
	FreeFall
	Initiative
	Piloting
	SleightOfHand
	Stealth
	Brawling
	Grappling
	HeavyWeapons
	LightWeapons
	Art
	Business
	Cryptography
	CurrentAffairs
	Demolitions
	Engineering
	Evaluation
	Law
	Medicine
	Navigation
	Research
	Science
	SecurityFocus
	Tactics
	Technology
	Empathy
	Hearing
	Intuition
	Searching
	Seeing
	Smelling
	Survival
	Tasting
	Touching
	Tracking
	Climbing
	Intimidation
	Jumping
	Might
	Courage
	Faith
	SelfDiscipline
	History
)

type BenifitDefinition struct {
	Background      Background `json:"background"`
	Ability         Ability    `json:"ability"`
	AbilityModifier int        `json:"ability_mod"`
	FocusPool       []Focus    `json:"focus_pool"`
	TalentPool      []Talent   `json:"talent_pool"`
}

type Benifit struct {
	Name   string `json:"name"`
	Focus  Focus  `json:"focus"`
	Talent Talent `json:"talent"`
}

func AbilityScoreModifier(roll int) int {
	if roll <= 3 {
		return -2
	}
	if roll/3 <= 2 {
		return 0
	}
	if roll/3 <= 3 {
		return 1
	}
	if roll/4 <= 4 {
		return 2
	}
	if roll/5 <= 5 {
		return 3
	}
	if roll/6 <= 6 {
		return 4
	}
	return 0
}

func (t Focus) String() string {
	switch t {
	case Bows:
		return "bows"
	case Gunnery:
		return "gunnery"
	case Pistols:
		return "pistols"
	case Rifles:
		return "rifles"
	case Throwing:
		return "throwing"
	case Bargaining:
		return "bargaining"
	case Deception:
		return "deception"
	case Disguise:
		return "disguise"
	case Etiquette:
		return "etiquette"
	case Expression:
		return "expression"
	case Gambling:
		return "gambling"
	case Investigation:
		return "investigation"
	case Leadership:
		return "leadership"
	case Performing:
		return "performing"
	case Persuasion:
		return "persuasion"
	case Seduction:
		return "seduction"
	case Running:
		return "running"
	case Stamina:
		return "stamina"
	case Swimming:
		return "swimming"
	case Tolerance:
		return "tolerance"
	case Acrobatics:
		return "acrobatics"
	case Crafting:
		return "crafting"
	case Driving:
		return "driving"
	case FreeFall:
		return "free_fall"
	case Initiative:
		return "initiative"
	case Piloting:
		return "piloting"
	case SleightOfHand:
		return "slight_of_hand"
	case Stealth:
		return "stealth"
	case Brawling:
		return "brawling"
	case Grappling:
		return "grappling"
	case HeavyWeapons:
		return "heavy_weapons"
	case LightWeapons:
		return "light_weapons"
	case Art:
		return "art"
	case Business:
		return "business"
	case Cryptography:
		return "cryptography"
	case CurrentAffairs:
		return "current_affairs"
	case Demolitions:
		return "demolitions"
	case Engineering:
		return "engineering"
	case Evaluation:
		return "evaluation"
	case Law:
		return "law"
	case Medicine:
		return "medicine"
	case Navigation:
		return "navigation"
	case Research:
		return "research"
	case Science:
		return "science"
	case SecurityFocus:
		return "security"
	case Tactics:
		return "tactics"
	case Technology:
		return "technology"
	case Empathy:
		return "empathy"
	case Hearing:
		return "hearing"
	case Intuition:
		return "intuition"
	case Searching:
		return "searching"
	case Seeing:
		return "seeing"
	case Smelling:
		return "smelling"
	case Survival:
		return "survival"
	case Tasting:
		return "tasting"
	case Touching:
		return "touching"
	case Tracking:
		return "tracking"
	case Climbing:
		return "climbing"
	case Intimidation:
		return "intimidation"
	case Jumping:
		return "jumping"
	case Might:
		return "might"
	case Courage:
		return "courage"
	case Faith:
		return "faith"
	case SelfDiscipline:
		return "self_discipline"
	}
	return ""
}

func (p Profession) String() string {
	switch p {
	case Brawler:
		return "brawler"
	case Survivalist:
		return "survivalist"
	case Criminal:
		return "criminal"
	case Scavenger:
		return "scavenger"
	case Fixer:
		return "fixer"
	case Artist:
		return "artist"
	case Athlete:
		return "athlete"
	case Soldier:
		return "soldier"
	case Investigator:
		return "investigator"
	case Technician:
		return "technician"
	case Clergy:
		return "clergy"
	case Negotiator:
		return "negotiator"
	case Pilot:
		return "pilot"
	case Security:
		return "security"
	case Professional:
		return "professional"
	case Scholar:
		return "scholar"
	case Merchant:
		return "merchant"
	case Politician:
		return "politician"
	case Commander:
		return "commander"
	case Explorer:
		return "explorer"
	case Dilettante:
		return "dilettante"
	case Expert:
		return "expert"
	case Executive:
		return "executive"
	case Socialite:
		return "socalite"
	}
	return ""
}

func (p Profession) SocialClass() SocialClass {
	if (p == Brawler) || (p == Survivalist) || (p == Criminal) || (p == Scavenger) || (p == Fixer) || (p == Artist) {
		return Outsider
	}
	if (p == Athlete) || (p == Soldier) || (p == Investigator) || (p == Technician) || (p == Clergy) || (p == Negotiator) {
		return LowerClass
	}
	if (p == Pilot) || (p == Security) || (p == Professional) || (p == Scholar) || (p == Merchant) || (p == Politician) {
		return MiddleClass
	}
	if (p == Commander) || (p == Explorer) || (p == Dilettante) || (p == Expert) || (p == Executive) || (p == Socialite) {
		return UppderClass
	}
	return Outsider
}

func (a Ability) String() string {
	switch a {
	case Accuracy:
		return "accuracy"
	case Constitution:
		return "constitution"
	case Fighting:
		return "fighting"
	case Communication:
		return "communication"
	case Dexterity:
		return "dexterity"
	case Intelligence:
		return "intelligence"
	case Perception:
		return "perception"
	case Strength:
		return "strength"
	case Willpower:
		return "willpower"
	}
	return ""
}

func (a Ability) Focus() []Focus {
	switch a {
	case Accuracy:
		return []Focus{Bows, Gunnery, Pistols, Rifles, Throwing}
	case Communication:
		return []Focus{
			Bargaining, Deception, Disguise,
			Etiquette, Expression, Gambling,
			Investigation, Leadership, Performing,
			Persuasion, Seduction,
		}
	case Constitution:
		return []Focus{Running, Stamina, Swimming, Tolerance}
	case Dexterity:
		return []Focus{Acrobatics, Crafting, Driving, FreeFall, Initiative, Piloting, SleightOfHand, Stealth}
	case Fighting:
		return []Focus{Brawling, Grappling, HeavyWeapons, LightWeapons}
	case Intelligence:
		return []Focus{
			Art, Business, Cryptography,
			CurrentAffairs, Demolitions, Engineering,
			Evaluation, Law, Medicine,
			Navigation, Research, Science,
			SecurityFocus, Tactics, Technology,
			History,
		}
	case Perception:
		return []Focus{
			Empathy, Hearing, Intuition,
			Searching, Seeing, Smelling,
			Survival, Tasting, Touching,
			Tracking,
		}
	case Strength:
		return []Focus{Climbing, Intimidation, Jumping, Might}
	case Willpower:
		return []Focus{Courage, Faith, SelfDiscipline}
	}
	return []Focus{}
}

func (b Background) BenifitDefinitions() BenifitDefinition {
	switch b {
	case Academic:
		return BenifitDefinition{
			Background:      Academic,
			Ability:         Intelligence,
			AbilityModifier: 1,
			FocusPool:       Intelligence.Focus(),
			TalentPool:      []Talent{Knowledge, Linguistics},
		}
	case Aristocratic:
		return BenifitDefinition{
			Background:      Aristocratic,
			Ability:         Communication,
			AbilityModifier: 1,
			FocusPool:       []Focus{Etiquette, History},
			TalentPool:      []Talent{Affluent, Contacts},
		}
	case Bohemian:
		return BenifitDefinition{
			Background:      Bohemian,
			Ability:         Communication,
			AbilityModifier: 1,
			FocusPool:       Intelligence.Focus(),
			TalentPool:      []Talent{Affluent, Contacts},
		}
	case Corporate:
		return BenifitDefinition{
			Background:      Corporate,
			Ability:         Communication,
			AbilityModifier: 1,
			FocusPool:       []Focus{Bargaining, Business},
			TalentPool:      []Talent{Contacts, Intrigue},
		}
	case Cosmopolitan:
		return BenifitDefinition{
			Background:      Cosmopolitan,
			Ability:         Perception,
			AbilityModifier: 1,
			FocusPool:       []Focus{Etiquette, CurrentAffairs},
			TalentPool:      []Talent{Knowledge, Observation},
		}
	case Exile:
		return BenifitDefinition{
			Background: Exile,
			Ability:    Constitution,
			FocusPool:  []Focus{Brawling, SelfDiscipline},
			TalentPool: []Talent{Affluent, Fringer},
		}
	case Military:
		return BenifitDefinition{
			Background:      Military,
			Ability:         Fighting,
			AbilityModifier: 1,
			FocusPool:       []Focus{Pistols, Tactics},
			TalentPool: []Talent{
				DualWeaponStyle, GrapplingStyle, GrapplingStyle, PistolStyle, RifleStyle, SelfDefenseStyle,
				SingleWeaponStyle, StrikingStyle, ThrownWeaponStyle, TwoHandedStyle, Observation,
			},
		}
	case Outcast:
		return BenifitDefinition{
			Background:      Outcast,
			Ability:         Perception,
			AbilityModifier: 1,
			FocusPool:       []Focus{Deception, Stealth},
			TalentPool:      []Talent{Fringer, Misdirection},
		}
	case Laborer:
		return BenifitDefinition{
			Background:      Laborer,
			Ability:         Constitution,
			AbilityModifier: 1,
			FocusPool:       []Focus{Crafting, Might},
			TalentPool:      []Talent{GrapplingStyle, SelfDefenseStyle, StrikingStyle, TwoHandedStyle, Carousing},
		}
	case Suburban:
		return BenifitDefinition{
			Background:      Suburban,
			Ability:         Communication,
			AbilityModifier: 1,
			FocusPool:       []Focus{Etiquette, CurrentAffairs},
			TalentPool:      []Talent{Affluent, Contacts},
		}
	case Trade:
		return BenifitDefinition{
			Background:      Trade,
			Ability:         Dexterity,
			AbilityModifier: 1,
			FocusPool:       []Focus{Crafting, Engineering},
			TalentPool:      []Talent{Improvisation, Maker},
		}
	case Urban:
		return BenifitDefinition{
			Background:      Urban,
			Ability:         Dexterity,
			AbilityModifier: 1,
			FocusPool:       []Focus{Persuasion, Stamina},
			TalentPool:      []Talent{Agility, Misdirection},
		}
	}
	return BenifitDefinition{}
}

func (b Background) String() string {
	switch b {
	case Bohemian:
		return "bohemian"
	case Exile:
		return "exile"
	case Outcast:
		return "outcast"
	case Military:
		return "military"
	case Laborer:
		return "laborer"
	case Urban:
		return "urban"
	case Academic:
		return "academic"
	case Suburban:
		return "suburban"
	case Trade:
		return "trade"
	case Aristocratic:
		return "aristocratic"
	case Corporate:
		return "corporate"
	case Cosmopolitan:
		return "cosmopolitan"
	}
	return ""
}

func (b Background) SocialClass() SocialClass {
	if (b == Bohemian) || (b == Exile) || (b == Outcast) {
		return Outsider
	}
	if (b == Military) || (b == Laborer) || (b == Urban) {
		return LowerClass
	}
	if (b == Academic) || (b == Suburban) || (b == Trade) {
		return MiddleClass
	}
	if (b == Aristocratic) || (b == Corporate) || (b == Cosmopolitan) {
		return UppderClass
	}
	return Outsider
}

func (o Origin) String() string {
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

func ParseOrigin(i int) Origin {
	return Origin(i)
}

func (s SocialClass) String() string {
	switch s {
	case Outsider:
		return "outsider"
	case LowerClass:
		return "lower_class"
	case MiddleClass:
		return "middle_class"
	case UppderClass:
		return "upper_class"
	}
	return ""
}