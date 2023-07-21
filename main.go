package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"

	"github.com/eiannone/keyboard"
)

const (
	width  = 200
	height = 140
)

var (
	charmap characterMapper
	world   dungeon
	level   *levelMap
	p       player

	currentState keyProcessor
	messages     messagePrompt
	gameplay     gamePlay
	persistance  persist

	arrows arrowQueue

	validKeyPressed bool

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

	persistance.register(&p, &world)

	charmap = initCharMap()

	savedStateExists := persistance.loadState("save.sav")

	if !savedStateExists {
		p = newPlayer('@')
		world = newDungeon()
		world.CurrentDepth = 1
		generateLevel(world.CurrentDepth)
	} else {
		level = world.Levels[world.CurrentDepth]
	}

}

func generateLevel(depth int) {
	level = world.newLevel(depth, width, height)
	generateDungeon()
	p.setPosition(level.getPointInRoom())
	level.generateDoors((width + height) / 10)
	level.generateStairs()
	level.generateItems(readItemsTemplate(), 50)
	level.generateMonsters(readMonsterTemplate(), 50)
}

func moveMonsters() {

	for i := 0; i < p.Speed; i++ {

		for i, m := range level.Monsters {

			if m.readyToMove() {

				if items, ok := level.Items[m.Position]; ok && m.CarriesItems && randomNumber(20) > m.Speed {
					if level.read(m.Position)&lit == lit {
						messages.push("The "+m.Name+" picked up "+items[len(items)-1].Prefix+" "+items[len(items)-1].Name, gameplay)
					}
					m.Items.add(items[len(items)-1])
					level.Items[m.Position] = level.Items[m.Position][:len(level.Items[m.Position])-1]
					if len(level.Items[m.Position]) == 0 {
						delete(level.Items, m.Position)
					}
					continue
				}

				var newDirection direction
				newDirection.connect(m.getPosition(), p.getPosition())

				if !m.Aggressive {
					newDirection = newDirection.opposite()
				}

				if !m.MovesDiagonally {
					newDirection.toNonDiagonal()
				}
				for i := 0; !m.attemptMove(newDirection) && i < 10; i++ {
					newDirection = randomDirection(newDirection, false, m.MovesDiagonally)
				}
				delete(level.Monsters, i)
				level.Monsters[m.Position] = m
			}

		}

	}
}

func checkPosition() {

	if f, ok := level.Features[p.Position]; ok {
		messages.push("There is "+f.Description+" here", gameplay)
	}

	if i, ok := level.Items[p.Position]; ok {
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

		for _, item := range level.Items[p.Position] {
			itemToAdd := item
			menuItems = append(menuItems, itemToAdd)
		}

	} else {

		if p.Items.count() == 0 {
			messages.push("Inventory empty", gameplay)
			currentState = gameplay
			return
		}

		for item := range p.Items.all() {
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

func generateDungeon() {

	var previousRoom point
	var nextRoom point
	var err error

	for {
		previousRoom, err = level.newRoom(level.getRandomPoint(), 20, 20)
		if err != nil {
			continue
		}
		break
	}

	for i := 0; i < 10; i++ {

		for {
			nextRoom, err = level.newRoom(level.getRandomPoint(), 20, 20)
			if err != nil {
				continue
			}
			break
		}

		level.newCorridor(previousRoom, nextRoom)
		previousRoom = nextRoom

	}

}

func printDungeon() {
	gridToPrint := render(level, p, &arrows, gridOverlay, 60, 40, level.Monsters, level.Items, level.Features)
	//gridToPrint := renderAll(&d, p, &arrows, gridOverlay, monstersOnMap, itemsOnMap, featuresOnMap)
	fmt.Println()
	fmt.Println(string(gridToPrint))
}

func printStats() {
	fmt.Println("HP:", p.Hp, "Turn:", world.Turn, "Depth:", world.CurrentDepth*10)
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

	p.attemptMove(None)

	for {

		validKeyPressed = false

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
