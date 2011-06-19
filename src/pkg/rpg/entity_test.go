package rpg

import "testing"

func newPlayerFromLevel(name string, level int) *Player {
  xp := XPMultiplier * (level - 1) * (level) / 2
  return NewPlayer(name, xp)
}

func TestNewPlayer(t *testing.T) {
  player := NewPlayer("nooblet", 0)
  if player.Name() != "nooblet" || player.XP() != 0 {
    t.Fail()
  }
}

func TestPlayerLevel(t *testing.T) {
  var player *Player

  player = NewPlayer("nooblet", 0)
  if player.Level() != 1 {
    t.Error("level should have been 1, but was ", player.Level())
  }

  player = NewPlayer("nooblet", 10)
  if player.Level() != 2 {
    t.Error("level should have been 2, but was ", player.Level())
  }

  player = NewPlayer("nooblet", 35)
  if player.Level() != 3 {
    t.Error("level should have been 3, but was ", player.Level())
  }
}

func TestPlayerHP(t *testing.T) {
  var player *Player

  player = newPlayerFromLevel("nooblet", 1)
  if player.HP() != 10 {
    t.Error("HP should have been 10, but was ", player.HP())
  }

  player = newPlayerFromLevel("nooblet", 2)
  if player.HP() != 20 {
    t.Error("HP should have been 20, but was ", player.HP())
  }

  player = newPlayerFromLevel("nooblet", 100)
  if player.HP() != 1000 {
    t.Error("HP should have been 1000, but was ", player.HP())
  }
}

func TestPlayerAttack(t *testing.T) {
  p1 := newPlayerFromLevel("nooblet", 1)
  p2 := newPlayerFromLevel("newbie", 2)

  p1.Attack(p2)
  if p2.HP() != 18 {
    t.Error("p2's HP should have been", 18, "but was", p2.HP())
  }

  p2.Attack(p1)
  if p1.HP() != 6 {
    t.Error("p1's HP should have been", 6, "but was", p1.HP())
  }
}

func TestPlayerDeath(t *testing.T) {
  p1 := newPlayerFromLevel("nooblet", 1)
  p2 := newPlayerFromLevel("newbie", 2)

  p2.Attack(p1)
  p2.Attack(p1)
  p2.Attack(p1)

  if !p1.IsDead() {
    t.Error("p1 should have died, but their HP was", p1.HP())
  }
}

func TestPlayerXPGaining(t *testing.T) {
  p1 := newPlayerFromLevel("nooblet", 1)
  p2 := newPlayerFromLevel("newbie", 2)

  if p1.XP() != 0 {
    t.Error("p1 should have 0 XP to begin with, but had", p1.XP())
  }

  for !p2.IsDead() {
    p1.Attack(p2)
  }

  if p1.XP() != 2 {
    t.Error("p1 should have 2 XP, but had", p1.XP())
  }
}

func TestPlayerLevelingUp(t *testing.T) {
  p1 := NewPlayer("nooblet", 8)
  p2 := newPlayerFromLevel("newbie", 2)

  for !p2.IsDead() {
    p1.Attack(p2)
  }

  if p1.XP() != 10 {
    t.Error("p1 should have 10 XP, but had", p1.XP())
  }

  if p1.Level() != 2 {
    t.Error("p1 should now be level 2, but is level", p1.Level())
  }
}

func TestPlayerSerialize(t *testing.T) {
  p1 := NewPlayer("nooblet", 123)
  result := p1.Serialize()
  if !(len(result) == 4 && result["name"] == "nooblet") {
    t.Error("result was incorrect:", result)
  }
}

func TestUnserializePlayer(t *testing.T) {
  player := NewPlayer("nooblet", 123)
  playerS := player.Serialize()
  playerU := Unserialize(playerS)
  if playerU.Name() != player.Name() || playerU.XP() != player.XP() || playerU.HP() != player.HP() {
    t.Error("unserialization failed:", playerU)
  }
}
