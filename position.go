package main

type Position struct {
	Rotation     int
	X            int
	Burn         int
	Step         int
	Hole         int
	Damage       int
	LowY         int
	HighY        int
	Score        int
	ColumnsAfter Picks
	FieldAfter   Field
	Type         string
}

func (p *Position) InitTop(picks Picks, fieldAfter Field, holes []Cell, s Strategy) {
	burn := fieldAfter.WillBurn()
	picksAfter := fieldAfter.Picks()
	damage, _, highY, step, hole := picks.Damage(picksAfter, holes)
	p.Type = "top"
	p.Burn = burn
	p.Step = step
	p.Hole = hole
	p.Damage = damage
	p.HighY = highY
	p.Score = damage*s.DamageK + highY*s.PostyK + step*s.StepK - burn*s.BurnK + hole
	p.FieldAfter = fieldAfter
}

func (p *Position) InitLeft(fieldAfter Field, s Strategy) {
	//burn := fieldAfter.WillBurn()
	//picksAfter := fieldAfter.Picks()
	//damage, _, highY, step, hole := picks.Damage(picksAfter, holes)
	//p.Rotation =   r
	//p.X=          i
	p.Type = "left"
	/*p.Burn = burn
	p.Step = step
	p.Hole = hole
	p.Damage = damage
	p.HighY = highY
	p.Score = damage*s.DamageK + highY*s.PostyK + step*s.StepK - burn*s.BurnK + hole*/
	p.FieldAfter = fieldAfter
}

func (p *Position) InitRight(fieldAfter Field, s Strategy) {
	//burn := fieldAfter.WillBurn()
	//picksAfter := fieldAfter.Picks()
	//damage, _, highY, step, hole := picks.Damage(picksAfter, holes)
	//p.Rotation =   r
	//p.X=          i
	p.Type = "right"
	/*p.Burn = burn
	p.Step = step
	p.Hole = hole
	p.Damage = damage
	p.HighY = highY
	p.Score = damage*s.DamageK + highY*s.PostyK + step*s.StepK - burn*s.BurnK + hole*/
	p.FieldAfter = fieldAfter
}
