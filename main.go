package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"

	"github.com/eiannone/keyboard"
)

const (
	width  = 80
	height = 40
)

var (
	charmap characterMapper
	world   dungeon
	lvl     *level
	pl      player

	currentState keyProcessor
	messages     messagePrompt
	gameplay     gamePlay

	persistance  persist

	arrows arrowQueue

	gridOverlay  []string
	menuItems    []*item
	selectedItem int

	seed *int
)

const (
	obstacle byte = 0
	empty    byte = 1 << iota
	visited
	lit
	room
)

func init() {

	seed = flag.Int("seed", 0, "seed for random number generation")
	flag.Parse()
	setRandomSource(*seed)

	persistance.register(&pl, &world)

	charmap = initCharMap()

	savedStateExists := persistance.loadState("save.sav")

	if !savedStateExists {
		pl = newPlayer('@')
		world = newDungeon()
		world.CurrentDepth = 1
		generateLevel(world.CurrentDepth)
	} else {
		lvl = world.Levels[world.CurrentDepth]
	}

}

func generateLevel(depth int) {
	lvl = world.newLevel(depth, width, height)
	lvl.excavate()
	pl.setPosition(lvl.getPointInRoom())
	lvl.generateDoors((width + height) / 10)
	lvl.generateStairs()
	lvl.generateItems(readItemsTemplate(), 50)
	lvl.generateMonsters(readMonsterTemplate(), 50)
}


func checkPosition() {

	if f, ok := lvl.Features[pl.Position]; ok {
		messages.push("There is "+f.Description+" here", gameplay)
	}

	if i, ok := lvl.Items[pl.Position]; ok {
		if len(i) == 1 {
			messages.push("There is "+i[0].Prefix+" "+i[0].Name+" here, press 5 to pick up", gameplay)
		} else if len(i) > 1 {
			messages.push("There are some things there, press 5 to examine", gameplay)
		}

	}
}

func generateOverlay(menu bool, verb string) {

	gridOverlay = nil
	menuItems = nil
	cursor := "| "

	if verb == "pick up" {

		for _, item := range lvl.Items[pl.Position] {
			itemToAdd := item
			menuItems = append(menuItems, itemToAdd)
		}

	} else {

		if pl.Items.count() == 0 {
			messages.push("Inventory empty", gameplay)
			currentState = gameplay
			return
		}

		for item := range pl.Items.all() {
			for _, v := range item.Verbs {
				if v == verb {
					itemToAdd := item
					menuItems = append(menuItems, itemToAdd)
					break
				}
			}
		}
	}

	sort.SliceStable(menuItems, func(i, j int) bool {
		return menuItems[i].Name < menuItems[j].Name
	})

	if len(menuItems) == 0 {
		messages.push("No items to "+verb, gameplay)
		currentState = gameplay
		return
	}

	gridOverlay = append(gridOverlay, "_______________________________")

	if menu {
		gridOverlay = append(gridOverlay, fmt.Sprintf("%-30s%v", "|Select an item to "+verb+":", " |"))
	} else {
		gridOverlay = append(gridOverlay, fmt.Sprintf("%-30s%v", "|Inventory", " |"))
	}

	gridOverlay = append(gridOverlay, fmt.Sprintf("%-30s%v", "|", " |"))

	for index, item := range menuItems {
		if menu {
			if index == selectedItem {
				cursor = "| > "
			} else {
				cursor = "|   "
			}
		}
		gridOverlay = append(gridOverlay, fmt.Sprintf("%-30s%v", cursor+strconv.Itoa(index)+": "+item.Prefix+" "+item.Name, " |"))
	}

	gridOverlay = append(gridOverlay, "|______________________________|")

}

func printDungeon() {
	gridToPrint := render(lvl, pl, &arrows, gridOverlay, 40, 40)
	//gridToPrint := renderAll(lvl, pl, &arrows, lvl.Monsters, lvl.Items, lvl.Features)
	fmt.Println()
	fmt.Println(string(gridToPrint))
}

func printStats() {
	fmt.Println("HP:", pl.Hp, "Turn:", world.Turn, "Depth:", world.CurrentDepth*10)
}

func printMessages() {
	switch {
	case len(messages.messageQueue) == 1:
		fmt.Print(messages.pop())
		currentState = messages.revertToState

	case len(messages.messageQueue) > 1:
		fmt.Print(messages.pop())
		fmt.Print(" (press space for more...)")
		currentState = messages

	}
}

func main() {

	currentState = gameplay

	pl.attemptMove(None)

	for {

		validKeyPressed := false

		printDungeon()
		printStats()
		printMessages()

		for !validKeyPressed {
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}
			validKeyPressed = currentState.processKey(char)
		}

	}

}
