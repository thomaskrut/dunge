package main

type inventory struct {
	Items map[*item]bool
}

func newInventory() inventory {
	return inventory{
		Items: make(map[*item]bool),
	}
}

func (i *inventory) add(item *item) {
	i.Items[item] = true
}

func (i *inventory) remove(item *item) {
	delete(i.Items, item)
}

func (i *inventory) count() int {
	return len(i.Items)
}

func (i *inventory) all() map[*item]bool {
	return i.Items
}

func (i *inventory) clear() {
	i.Items = nil
}
