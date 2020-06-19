package main

type country struct{
	gdp uint32
	money uint32
	debt uint32
	population uint32
	infrastructureScore uint32
	militaryScore uint32
	cultureScore uint32
	code uint32
};
func createCountry (a_code uint32) country {
 return country{
		gdp:                 0,
		money:               0,
		debt:                0,
		population:          0,
		infrastructureScore: 0,
		militaryScore:       0,
		cultureScore:        0,
		code:                a_code,
	}
}
const MAXGAMEBOARDSIZE uint32 =20000;
type event struct {
mmth uint32
dateby uint32
notBefore uint32
modifier uint32
country uint32
};
type nationalEffect struct
{ID uint32
 condition_ID uint32
 effect_ID uint32
};
type owner struct
{
	taxIncome uint32;
	income uint32;
	tradeNodes []tradeNode;
	tradeIncome uint32;
	taxEff uint32;
	tradeEff uint32;
};
type tradeNode struct {
owner owner;
mainTile string;
tradeIncome uint32;
}
type tile struct
{
height uint32
modifierTAG uint32
noOfTroops1 uint32
noOfTroops2 uint32
noOfTroops3 uint32
ownerOfTroops1 owner
ownerOfTroops2 owner
ownerOfTroops3 owner
occupationStatus uint32
vas owner
attacker owner
trade uint32
taxIncome uint32
tradeNode tradeNode;
manPower uint32
garrison uint32
fortLevel uint32
supplyLimit uint32
};
type decision struct
{
 ID uint32;
condition bool;
taken bool;
repeatable bool;
effect uint32;
effect2 uint32;
effect3 uint32;
};
type building struct
{
ID uint32 ;
 condition bool;
 isActive bool;
 effectID uint32;
};
type technology struct
{
ID uint32;
condition bool;
taken bool;
effectID uint32;
};
type SEC_SPEC int
const(
NONE SEC_SPEC = 1
SPY SEC_SPEC = 2
DIPLOMAT SEC_SPEC = 4
)
type RelationEntry struct {
	cnt1 country;
	cnt2 country;
	rel SEC_REL;
};
type politician struct
{
 globalRank uint32;
party string;
money uint32;
influence uint32;
partyRank uint32;
};
type troop struct
{
cost uint32;
firePower uint32;
manPower uint32;
strategicAdvantage uint32;
technologicalAdvantage uint32;
terrainAdvantage uint32;
cowardice uint32;
revanchism uint32;
ID uint32;
};
const CLEAR_COMMAND string = "cls";
type organization struct{
code uint32;
name string;
};
type government struct
{
	DODid uint32;
	CIAid uint32;
	ARMYid uint32;
	PROPid uint32;
	INTERid uint32;
	CHIEFid uint32;
	VICEid uint32;
	PARLid uint32;
	PEOPid uint32;
};
type person struct
{
id uint32;
portrait bool;
name string;
description string;
title string;
tag SEC_SPEC;

};
type SEC_REL int
const (
	PEACE SEC_REL = 1
 	WAR SEC_REL = 2
	CORDIAL SEC_REL = 4
	THREATENED SEC_REL = 8
	HISTORICAL SEC_REL = 16
)
type GAME_STATE int
const (
	MAIN_MENU GAME_STATE = 1
	GAME_START GAME_STATE = 2
	OPTIONS_MENU GAME_STATE = 3
	EXITING GAME_STATE = 4
)