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
  Serialize() M

  takeDamage(amount int)
  maxHP() int
}

type Player struct {
  name string
  xp int
  hp int
}

func Unserialize(data M) Entity {
  var result Entity
  var kind string
  var ok bool

  if kind, ok = data["kind"].(string); ok {
    switch kind {
    case "player":
      var name string
      var xp, hp int

      player := &Player{}
      if name, ok = data["name"].(string); ok {
        player.name = name
      }
      if xp, ok = data["xp"].(int); ok {
        player.xp = xp
      }
      if hp, ok = data["hp"].(int); ok {
        player.hp = hp
      }
      result = player
    }
  }
  return result
}


func NewPlayer(name string) *Player {
  player := &Player{name: name}
  player.init()
  return player
}

func (p *Player) init() {
  p.hp = p.maxHP()
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

func (p *Player) Serialize() M {
  return M{"name": p.name, "xp": p.xp, "hp": p.hp, "kind": "player"}
}

func (p *Player) takeDamage(amount int) {
  p.hp -= amount
}

func (p *Player) maxHP() int {
  return p.Level() * 10
}
