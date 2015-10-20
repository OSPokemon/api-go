package ospokemon

type ExperienceCurve interface {
	Id() int
	GetExperienceForLevel(level int) int
}
