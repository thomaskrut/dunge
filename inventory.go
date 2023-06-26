package main

type inventory struct {
	items []item
}

func (i *inventory) add(item *item) {
	i.items = append(i.items, *item)
}

//TODO item not removed correctly
func (i *inventory) remove(index int) {
	i.items = append(i.items[:index], i.items[index+1:]...)
}

func (i *inventory) size() int {
	return len(i.items)
}

func (i *inventory) all() []item {
	return i.items
}

func (i *inventory) clear() {
	i.items = nil
}