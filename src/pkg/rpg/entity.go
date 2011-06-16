package rpg

import "math"

const XPMultiplier = 10

type Entity interface {
  Name() string
  XP() int
  Level() int
  HP() int
  Attack(target Entity)
  IsDead() bool

  takeDamage(amount int)
  maxHP() int
}

type Player struct {
  name string
  xp int
  hp int
}

func levelFromXP(xp int) int {
}

func NewPlayer(name string, xp int) *Player {
  player := &Player{name: name, xp: xp}
  player.hp = player.maxHP()
  return player
}

func (p *Player) Name() string {
  return p.name
}

func (p *Player) XP() int {
  return p.xp
}

func (p *Player) Level() int {
  // each level costs an additional XPMultiplier points
  // when XPMultiplier is 10: 10, 30, 60, 100, 150
  //
  // xpRequiredForLevel = XPMultiplier * level * (level + 1) / 2
  // level = sqrt(2 * xp / XPMultiplier + 0.25) - 0.5
  return int(math.Trunc(math.Sqrt(2 * float64(p.xp) / XPMultiplier + 0.25) - 0.5)) + 1
}

func (p *Player) HP() int {
  return p.hp
}

func (p *Player) Attack(target Entity) {
  target.takeDamage(p.Level() * 2)
  if target.IsDead() {
    p.xp += target.Level()
  }
}

func (p *Player) IsDead() bool {
  return p.hp <= 0
}

func (p *Player) takeDamage(amount int) {
  p.hp -= amount
}

func (p *Player) maxHP() int {
  return p.Level() * 10
}
