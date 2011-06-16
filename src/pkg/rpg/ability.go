package rpg

type Ability interface {
  Name() string
  Classes() []string  // which classes can use it
  MinimumLevel() int
  Use(target Entity)
}
