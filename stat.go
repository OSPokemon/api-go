package ospokemon

type Stat interface {
	Value() int
	SetValue(value int)

	EffortValue() int
	SetEffortValue(value int)

	IndividualValue() int
	SetIndividualValue(value int)
}

type BasicStat struct {
	VALUE int
	EV    int
	IV    int
}

func (s *BasicStat) Value() int {
	return s.VALUE
}

func (s *BasicStat) SetValue(value int) {
	s.VALUE = value
}

func (s *BasicStat) EffortValue() int {
	return s.EV
}

func (s *BasicStat) SetEffortValue(value int) {
	s.EV = value
}

func (s *BasicStat) IndividualValue() int {
	return s.IV
}

func (s *BasicStat) SetIndividualValue(value int) {
	s.IV = value
}
