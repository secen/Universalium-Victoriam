package main

type law struct {
	effect      func(country)
	name        string
	description string
	Type        uint32
}
type country struct {
	gdp                 uint32
	money               uint32
	debt                uint32
	population          uint32
	infrastructureScore uint32
	militaryScore       uint32
	cultureScore        uint32
	code                uint32
	income              uint32
	expenses            uint32
	interest            int32
	laws                []law
	techs 				[]technology
	nationalEffects 	[]nationalEffect
}

const(
	debugCountriesFilename = "debugCountries.txt"
	nationPickerDataFilename = "nationPickerData.txt"
	consoleHelpDataFilename = "consoleHelpData.txt"
)

func createCountry(aCode uint32) country {
	return country{
		gdp:                 0,
		money:               0,
		debt:                0,
		population:          0,
		infrastructureScore: 0,
		militaryScore:       0,
		cultureScore:        0,
		code:                aCode,
		income:              0,
		expenses:            0,
	}
}

type queue []func()

func appendQueue(q queue,f func())queue{q=append(q, f); return q } // Enqueue
func deQueue(q queue)queue {q=q[1:]; return q }
func popQueue(q queue)(func(),queue){var aux = q[0]; q = deQueue(q); return aux, q}
const MAXGAMEBOARDSIZE uint32 = 20000

type event struct {
	mmth      uint32
	dateby    uint32
	notBefore uint32
	modifier  uint32
	country   uint32
}

type nationalEffect struct {
	ID          uint32
	conditionId uint32
	effectId    uint32
}

type owner struct {
	taxIncome   uint32
	income      uint32
	tradeNodes  []tradeNode
	tradeIncome uint32
	taxEff      uint32
	tradeEff    uint32
}

type tradeNode struct {
	owner       owner
	mainTile    string
	tradeIncome uint32
}
type tile struct {
	height           uint32
	modifierTAG      uint32
	noOfTroops1      uint32
	noOfTroops2      uint32
	noOfTroops3      uint32
	ownerOfTroops1   owner
	ownerOfTroops2   owner
	ownerOfTroops3   owner
	occupationStatus uint32
	vas              owner
	attacker         owner
	trade            uint32
	taxIncome        uint32
	tradeNode        tradeNode
	manPower         uint32
	garrison         uint32
	fortLevel        uint32
	supplyLimit      uint32
}

type decision struct {
	ID         uint32
	condition  bool
	taken      bool
	repeatable bool
	effect     uint32
	effect2    uint32
	effect3    uint32
}
type Good struct{
	demand float64
	supply float64
	price float64
	volatility float64
	volatilityOfRateOfChange float64
	rateOfChange float64
}
func goodCalculateNextTick(gd Good)Good{
	gd.price+=gd.demand/gd.supply*gd.volatility+gd.rateOfChange*gd.volatilityOfRateOfChange
	return gd
}
type building struct {
	ID        uint32
	condition bool
	isActive  bool
	effectID  uint32
}

type technology struct {
	ID        uint32
	condition bool
	taken     bool
	effectID  uint32
}

type SecSpec int

const (
	NONE     SecSpec = 1
	SPY      SecSpec = 2
	DIPLOMAT SecSpec = 4
)

type RelationEntry struct {
	cnt1 country
	cnt2 country
	rel  SecRel
}

type politician struct {
	globalRank uint32
	party      string
	money      uint32
	influence  uint32
	partyRank  uint32
}

type troop struct {
	cost                   uint32
	firePower              uint32
	manPower               uint32
	strategicAdvantage     uint32
	technologicalAdvantage uint32
	terrainAdvantage       uint32
	cowardice              uint32
	revanchism             uint32
	ID                     uint32
}

const ClearCommand string = "cls"

type organization struct {
	code uint32
	name string
}

type government struct {
	DODid   uint32
	CIAid   uint32
	ARMYid  uint32
	PROPid  uint32
	INTERid uint32
	CHIEFid uint32
	VICEid  uint32
	PARLid  uint32
	PEOPid  uint32
}

type person struct {
	id          uint32
	portrait    bool
	name        string
	description string
	title       string
	tag         SecSpec
}

type SecRel int

const (
	PEACE      SecRel = 1
	WAR        SecRel = 2
	CORDIAL    SecRel = 4
	THREATENED SecRel = 8
	HISTORICAL SecRel = 16
)

type GameState int

const (
	MainMenu    GameState = 1
	GameStart   GameState = 2
	OptionsMenu GameState = 3
	EXITING     GameState = 4
)
