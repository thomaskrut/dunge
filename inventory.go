package main

type inventory struct {
	items map[*item]bool
}

func newInventory() inventory {
	return inventory{
		items: make(map[*item]bool),
	}
}

func (i *inventory) add(item *item) {
	i.items[item] = true
}

func (i *inventory) remove(item *item) {
	delete(i.items, item)
}

func (i *inventory) size() int {
	return len(i.items)
}

func (i *inventory) all() map[*item]bool {
	return i.items
}

func (i *inventory) clear() {
	i.items = nil
}