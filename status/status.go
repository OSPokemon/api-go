package status

type T uint8

const (
	Asleep T = 1 << iota
	Poisoned
	Paralyzed
	Burned
	Frozen
	Confused
)

func New() T {
	return T(0)
}

func (t *T) Has(status T) bool { return *t&status > 0 }
func (t *T) Set(status T)      { *t = *t | status }
func (t *T) Clear(status T)    { *t = *t &^ status }

func (t *T) IsAsleep() bool    { return t.Has(Asleep) }
func (t *T) IsPoisoned() bool  { return t.Has(Poisoned) }
func (t *T) IsParalyzed() bool { return t.Has(Paralyzed) }
func (t *T) IsBurned() bool    { return t.Has(Burned) }
func (t *T) IsFrozen() bool    { return t.Has(Frozen) }
func (t *T) IsConfused() bool  { return t.Has(Confused) }

func (t *T) SetAsleep()    { t.Set(Asleep) }
func (t *T) SetPoisoned()  { t.Set(Poisoned) }
func (t *T) SetParalyzed() { t.Set(Paralyzed) }
func (t *T) SetBurned()    { t.Set(Burned) }
func (t *T) SetFrozen()    { t.Set(Frozen) }
func (t *T) SetConfused()  { t.Set(Confused) }

func (t *T) ClearAsleep()    { t.Clear(Asleep) }
func (t *T) ClearPoisoned()  { t.Clear(Poisoned) }
func (t *T) ClearParalyzed() { t.Clear(Paralyzed) }
func (t *T) ClearBurned()    { t.Clear(Burned) }
func (t *T) ClearFrozen()    { t.Clear(Frozen) }
func (t *T) ClearConfused()  { t.Clear(Confused) }
