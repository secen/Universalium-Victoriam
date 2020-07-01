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
	relations           []RelationEntry
}
type GameData struct {
	pickedNation  Country
	globalGoods   []Good
	globalNations []Country
}

const (
	debugCountryFilename     = "debugCountry.txt"
	countriesFilename        = "countries.json"
	nationPickerDataFilename = "nationListings.json"
	consoleHelpDataFilename  = "consoleHelpData.txt"
	nationPickerFilename     = "\\consoleData\\nationPickerHelpData.txt"
)

type CountryListing struct {
	CountryName string
	Subtitle    string
	Description string
}

type queue []func()

func appendQueue(q queue, f func()) queue { q = append(q, f); return q } // Enqueue
func deQueue(q queue) queue               { q = q[1:]; return q }
func popQueue(q queue) (func(), queue)    { var aux = q[0]; q = deQueue(q); return aux, q }

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

type Good struct {
	Name                     string
	Demand                   float64
	Supply                   float64
	Price                    float64
	Volatility               float64
	VolatilityOfRateOfChange float64
	RateOfChange             float64
}

func goodCalculateNextTick(gd Good) Good {
	gd.Price += gd.Demand/gd.Supply*gd.Volatility + gd.RateOfChange*gd.VolatilityOfRateOfChange
	return gd
}

type Technology struct {
	ID        uint32
	Category  string
	Condition bool
	Taken     bool
	EffectID  uint32
	Name      string
}

type RelationEntry struct {
	Cnt1         Country
	Cnt2         Country
	Rel          SecRel
	Opinionvalue int32
}

type Troop struct {
	Name                   string
	Cost                   uint32
	FirePower              uint32
	ManPower               uint32
	StrategicAdvantage     uint32
	TechnologicalAdvantage uint32
	TerrainAdvantage       uint32
	Cowardice              uint32
	Revanchism             uint32
	ID                     uint32
	LocationID             uint32
}

type SecRel int
type CONTINENT uint32

const (
	//Continental Regions
	//EUROPE
	EUROPE_NORTH   CONTINENT = 0
	EUROPE_CENTRAL CONTINENT = 1
	EUROPE_SOUTH   CONTINENT = 2 << 0
	EUROPE_WEST    CONTINENT = 2 << 1
	EUROPE_EAST    CONTINENT = 2 << 2
	//ASIA
	ASIA_WEST    = 2 << 3
	ASIA_EAST    = 2 << 4
	ASIA_CENTRAL = 2 << 5
	ASIA_NORTH   = 2 << 6
	ASIA_SOUTH   = 2 << 7
	//AFRICA
	//MIDDLE EAST
	//NORTH AMERICA
	//SOUTH AMERICA
	//AUSTRALIA
	//OCEANIA
)
const (
	//Regions
	//KOREA
	KOREA_NORTH = iota
	KOREA_SOUTH
	//JAPAN
	JAPAN_HOKKAIDO
	JAPAN_TOHOKU
	JAPAN_CHUBU
	JAPAN_KANTO
	JAPAN_KINKI
	JAPAN_CHUGOKU
	JAPAN_SHIKOKU
	JAPAN_KYUSHU
	JAPAN_OKINAWA
	//CHINA
	CHINA_HEILONGJIANG
	CHINA_JILIN
	CHINA_LIAONING
	CHINA_BEIJING
	CHINA_HEBEI
	CHINA_SHANDONG
	CHINA_JIANGSU
	CHINA_SHANGHAI
	CHINA_ZHEJIANG
	CHINA_FUJIAN
	CHINA_GUANGDONG
	CHINA_MACAO
	CHINA_HONGKONG
	CHINA_HAINAN
	CHINA_GUANGXI
	CHINA_TUNNAN
	CHINA_GUIZHOU
	CHINA_CHONGQING
	CHINA_SHAANXI
	CHINA_NINGXIA
	CHINA_INNERMONGOLIA
	CHINA_GANSU
	CHINA_QINGHAI
	CHINA_XIANJIANG
	CHINA_TIBET
	//MONGOLIA
	MONGOLIA_OUTERMONGOLIA
)
const (
	PEACE      SecRel = 0
	WAR        SecRel = 1 << 0
	CORDIAL    SecRel = 2 << 1
	THREATENED SecRel = 2 << 2
	HISTORICAL SecRel = 2 << 3
	ALLIED     SecRel = 2 << 4
)

type GameState int

const (
	MainMenu    GameState = 1
	GameStart   GameState = 2
	OptionsMenu GameState = 3
	EXITING     GameState = 4
)
