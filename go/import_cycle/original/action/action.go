// Package action provides ...
package action

type Action struct {
	Names  []string
	Action func(character.Character) int
}
