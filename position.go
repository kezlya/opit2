package main

type Position struct {
	Burn         int
	Step         int
	Hole         int
	Damage       int
	LowY         int
	HighY        int
	Score        int
	ColumnsAfter Picks
	FieldAfter   Field
	Moves        string
}

func (p *Position) Init(picks, picksAfter Picks, holes, fixable []Cell, burn int, s Strategy) {
	_, _, highY, step, hole := picks.Damage(picksAfter, holes)
	////
	damage := len(fixable)
	////
	p.Burn = burn
	p.Step = step
	p.Hole = hole
	p.Damage = damage
	p.HighY = highY
	p.Score = damage*s.DamageK + highY*s.PostyK + step*s.StepK - burn*s.BurnK + hole
}
