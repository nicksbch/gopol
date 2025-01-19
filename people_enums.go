package main

type Influence int

const (
	LowInfluence Influence = iota
	MediumInfluence
	HighInfluence
	UnknownInfluence
)

var InfluenceName = map[Influence]string{
	LowInfluence:     "low",
	MediumInfluence:  "medium",
	HighInfluence:    "high",
	UnknownInfluence: "unknown",
}

func (i Influence) String() string {
	return InfluenceName[i]
}

type Level int

const (
	ContractorLevel Level = iota
	EmployeeLevel
	SeniorEmployeeLevel
	ManagerLevel
	ExecutiveLevel
)

var LevelName = map[Level]string{
	ContractorLevel:     "contractor",
	EmployeeLevel:       "employee",
	SeniorEmployeeLevel: "senior employee",
	ManagerLevel:        "manager",
	ExecutiveLevel:      "executive",
}

func (l Level) String() string {
	return LevelName[l]
}
