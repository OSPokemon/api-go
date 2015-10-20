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
	value int
	ev    int
	iv    int
}

func (s *BasicStat) Value() int {
	return s.value
}

func (s *BasicStat) SetValue(value int) {
	s.value = value
}

func (s *BasicStat) EffortValue() int {
	return s.ev
}

func (s *BasicStat) SetEffortValue(value int) {
	s.ev = value
}

func (s *BasicStat) IndividualValue() int {
	return s.iv
}

func (s *BasicStat) SetIndividualValue(value int) {
	s.iv = value
}
