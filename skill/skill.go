package skill

type Skill interface {
	Move() int
	SetMove(move int)

	CurrentPP() int
	SetCurrentPP(currentPP int)

	MaximumPP() int
	SetMaximumPP(maximumPP int)
}

type BasicSkill struct {
	move      int
	currentPP int
	maximumPP int
}

func (s *BasicSkill) Move() int {
	return s.move
}

func (s *BasicSkill) SetMove(move int) {
	s.move = move
}

func (s *BasicSkill) CurrentPP() int {
	return s.currentPP
}

func (s *BasicSkill) SetCurrentPP(currentPP int) {
	s.currentPP = currentPP
}

func (s *BasicSkill) MaximumPP() int {
	return s.maximumPP
}

func (s *BasicSkill) SetMaximumPP(maximumPP int) {
	s.maximumPP = maximumPP
}
