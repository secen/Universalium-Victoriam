package main

type law struct {
	effect      func(Country)
	name        string
	description string
	Type        uint32
}
type Country struct {
	Name                string
	Gdp                 uint32
	Money               uint32
	Debt                uint32
	Population          uint32
	InfrastructureScore uint32
	MilitaryScore       uint32
	CultureScore        uint32
	Code                uint32
	Income              uint32
	Expenses            uint32
	interest            int32
	laws                []law
	techs               []Technology
	nationalEffects     []nationalEffect
}

type GameData struct{
	pickedNation Country
	globalGoods []Good
	globalNations []Country
}
const(
	debugCountryFilename = "debugCountry.txt"
	countriesFilename = "countries.json"
	nationPickerDataFilename = "nationListings.json"
	consoleHelpDataFilename = "consoleHelpData.txt"
	nationPickerFilename = "\\consoleData\\nationPickerHelpData.txt"
	)
type CountryListing struct{
	CountryName string
	Subtitle string
	Description string
}
func createCountry(aCode uint32) Country {
	return Country{
		Name:                "",
		Gdp:                 0,
		Money:               0,
		Debt:                0,
		Population:          0,
		InfrastructureScore: 0,
		MilitaryScore:       0,
		CultureScore:        0,
		Code:                aCode,
		Income:              0,
		Expenses:             0,
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
	Country   uint32
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
	Name string
	Demand float64
	Supply float64
	Price float64
	Volatility float64
	VolatilityOfRateOfChange float64
	RateOfChange float64
}
func goodCalculateNextTick(gd Good)Good{
	gd.Price+=gd.Demand/gd.Supply*gd.Volatility+gd.RateOfChange*gd.VolatilityOfRateOfChange
	return gd
}
type building struct {
	ID        uint32
	condition bool
	isActive  bool
	effectID  uint32
}

type Technology struct {
	ID        uint32
	Category string
	Condition bool
	Taken     bool
	EffectID  uint32
	Name      string
}

type SecSpec int

const (
	NONE     SecSpec = 1
	SPY      SecSpec = 2
	DIPLOMAT SecSpec = 4
)

type RelationEntry struct {
	cnt1 Country
	cnt2 Country
	rel  SecRel
}

type politician struct {
	globalRank uint32
	party      string
	money      uint32
	influence  uint32
	partyRank  uint32
}

type Troop struct {
	Name string
	Cost                   uint32
	FirePower              uint32
	ManPower               uint32
	StrategicAdvantage     uint32
	TechnologicalAdvantage uint32
	TerrainAdvantage       uint32
	Cowardice              uint32
	Revanchism             uint32
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
