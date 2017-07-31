package character

type Character struct {
	Inventory []item.Item
}

// Do implement the interface ActionTarget
func (c Character) Do() int {
}
