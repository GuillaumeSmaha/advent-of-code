package p1

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type Cart struct {
	ID           int
	Direction    int // up: 0, down: 1, left: 2, right: 3
	X            int
	Y            int
	Intersection int
}

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2018 14[a|b] 'input'")
	}

	switch args[0] {
	case "14a", "14":
		fmt.Println(Part1(args[1]))
	case "14b":
		fmt.Println(Part2(args[1]))
	}
}

func load(s string) []string {
	res := make([]string, 0)
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		// s = strings.TrimSpace(string(f))
		s = string(f)

		for _, line := range strings.Split(s, "\n") {
			if len(line) == 0 {
				continue
			}
			res = append(res, line)
		}
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func deleteFromSlice(list []string, element string) []string {
	for i, v := range list {
		if v == element {
			list = append(list[:i], list[i+1:]...)
			break
		}
	}

	return list
}

func loadData(s string) ([][]rune, []*Cart) {
	rails := make([][]rune, 0)
	carts := make([]*Cart, 0)

	cartID := 0
	for y, line := range load(s) {
		for x, r := range line {
			switch r {
			case '^':
				carts = append(carts, &Cart{
					ID:        cartID,
					X:         x,
					Y:         y,
					Direction: 0,
				})
				line = line[:x] + "|" + line[x+1:]
				cartID++
			case 'v':
				carts = append(carts, &Cart{
					ID:        cartID,
					X:         x,
					Y:         y,
					Direction: 1,
				})
				line = line[:x] + "|" + line[x+1:]
				cartID++
			case '<':
				carts = append(carts, &Cart{
					ID:        cartID,
					X:         x,
					Y:         y,
					Direction: 2,
				})
				line = line[:x] + "-" + line[x+1:]
				cartID++
			case '>':
				carts = append(carts, &Cart{
					ID:        cartID,
					X:         x,
					Y:         y,
					Direction: 3,
				})
				line = line[:x] + "-" + line[x+1:]
				cartID++
			}
		}
		rails = append(rails, []rune(line))
	}
	return rails, carts
}

func checkCollision(carts []*Cart) (*Cart, *Cart) {

	for _, c1 := range carts {
		for _, c2 := range carts {
			if c1.ID != c2.ID && c1.X == c2.X && c1.Y == c2.Y {
				return c1, c2
			}
		}
	}
	return nil, nil
}

func moveStepWithCheckCollision(rails [][]rune, carts []*Cart) *Cart {
	for _, c := range carts {

		cc1, _ := checkCollision(carts)
		if cc1 != nil {
			return cc1
		}

		// Move carts
		switch c.Direction {
		case 0:
			c.Y--
			break
		case 1:
			c.Y++
			break
		case 2:
			c.X--
			break
		case 3:
			c.X++
			break
		}

		// Determine new direction
		switch rails[c.Y][c.X] {
		case '|':
			switch c.Direction {
			case 0:
				// don't change
				break
			case 1:
				// don't change
				break
			case 2:
				panic("not possible | and - direction")
				break
			case 3:
				panic("not possible | and - direction")
				break
			}
			break
		case '-':
			switch c.Direction {
			case 0:
				panic("not possible | and - direction")
				break
			case 1:
				panic("not possible | and - direction")

				break
			case 2:
				// don't change
				break
			case 3:
				// don't change
				break
			}
			break
		case '\\':
			switch c.Direction {
			case 0:
				// go to left
				c.Direction = 2
				break
			case 1:
				// go to right
				c.Direction = 3
				break
			case 2:
				// go to up
				c.Direction = 0
				break
			case 3:
				// go to down
				c.Direction = 1
				break
			}
			break
		case '/':
			switch c.Direction {
			case 0:
				// go to right
				c.Direction = 3
				break
			case 1:
				// go to left
				c.Direction = 2
				break
			case 2:
				// go to down
				c.Direction = 1
				break
			case 3:
				// go to up
				c.Direction = 0
				break
			}
			break

		case '+':
			switch c.Intersection % 3 {
			case 0: // turn left
				switch c.Direction {
				case 0:
					// go to left
					c.Direction = 2
					break
				case 1:
					// go to right
					c.Direction = 3
					break
				case 2:
					// go to down
					c.Direction = 1
					break
				case 3:
					// go to up
					c.Direction = 0
					break
				}
				break
				break
			case 1: // go straight
				// Don't change
				break
			case 2: // turn right
				switch c.Direction {
				case 0:
					// go to right
					c.Direction = 3
					break
				case 1:
					// go to left
					c.Direction = 2
					break
				case 2:
					// go to up
					c.Direction = 0
					break
				case 3:
					// go to down
					c.Direction = 1
					break
				}
				break
			}
			c.Intersection++
		}
	}

	return nil
}

func Part1(s string) string {
	rails, carts := loadData(s)

	i := 0
	for {
		c := moveStepWithCheckCollision(rails, carts)

		if c != nil {
			fmt.Printf("End at step: %v\n", i)
			fmt.Println(c)
			return fmt.Sprintf("%v,%v", c.X, c.Y)
		}
		i++
	}

	return ""
}

func deleteCartByID(carts []*Cart, ID int) []*Cart {
	for i, v := range carts {
		if v.ID == ID {
			carts = append(carts[:i], carts[i+1:]...)
			break
		}
	}

	return carts
}

func deleteCart(carts []*Cart, cart *Cart) []*Cart {
	for i, v := range carts {
		if v.ID == cart.ID {
			carts = append(carts[:i], carts[i+1:]...)
			break
		}
	}

	return carts
}

func cartInSlice(a *Cart, list []*Cart) bool {
	for _, b := range list {
		if b.ID == a.ID {
			return true
		}
	}
	return false
}

func checkCollisionMultiple(carts []*Cart) []*Cart {
	ccs := make([]*Cart, 0)
	for _, c1 := range carts {
		for _, c2 := range carts {
			if c1.ID != c2.ID && c1.X == c2.X && c1.Y == c2.Y {
				ccs = append(ccs, c1)
				ccs = append(ccs, c2)
			}
		}
	}
	return ccs
}

func moveStepWithCheckCollisionByRemoving(rails [][]rune, carts []*Cart) []*Cart {
	// previousCarts := carts
	cartsToRemove := make([]*Cart, 0)

	ccs := checkCollisionMultiple(carts)
	if len(ccs) != 0 {
		cartsToRemove = append(cartsToRemove, ccs...)
		for _, cs := range ccs {
			fmt.Printf("Collision at start with %v: %v\n", cs.ID, cs)
			cartsToRemove = append(cartsToRemove, cs)
		}
	}

	// Sort by coord
	sort.Slice(carts, func(i, j int) bool {
		if carts[i].X < carts[j].X {
			return true
		}
		if carts[i].X > carts[j].X {
			return false
		}
		return carts[i].Y < carts[j].Y
	})

	for _, c := range carts {

		if cartInSlice(c, cartsToRemove) {
			continue
		}

		// Move carts
		switch c.Direction {
		case 0:
			c.Y--
			break
		case 1:
			c.Y++
			break
		case 2:
			c.X--
			break
		case 3:
			c.X++
			break
		}

		// Determine new direction
		switch rails[c.Y][c.X] {
		case '|':
			switch c.Direction {
			case 0:
				// don't change
				break
			case 1:
				// don't change
				break
			case 2:
				panic("not possible | and - direction")
				break
			case 3:
				panic("not possible | and - direction")
				break
			}
			break
		case '-':
			switch c.Direction {
			case 0:
				panic("not possible | and - direction")
				break
			case 1:
				panic("not possible | and - direction")

				break
			case 2:
				// don't change
				break
			case 3:
				// don't change
				break
			}
			break
		case '\\':
			switch c.Direction {
			case 0:
				// go to left
				c.Direction = 2
				break
			case 1:
				// go to right
				c.Direction = 3
				break
			case 2:
				// go to up
				c.Direction = 0
				break
			case 3:
				// go to down
				c.Direction = 1
				break
			}
			break
		case '/':
			switch c.Direction {
			case 0:
				// go to right
				c.Direction = 3
				break
			case 1:
				// go to left
				c.Direction = 2
				break
			case 2:
				// go to down
				c.Direction = 1
				break
			case 3:
				// go to up
				c.Direction = 0
				break
			}
			break

		case '+':
			switch c.Intersection % 3 {
			case 0: // turn left
				switch c.Direction {
				case 0:
					// go to left
					c.Direction = 2
					break
				case 1:
					// go to right
					c.Direction = 3
					break
				case 2:
					// go to down
					c.Direction = 1
					break
				case 3:
					// go to up
					c.Direction = 0
					break
				}
				break
				break
			case 1: // go straight
				// Don't change
				break
			case 2: // turn right
				switch c.Direction {
				case 0:
					// go to right
					c.Direction = 3
					break
				case 1:
					// go to left
					c.Direction = 2
					break
				case 2:
					// go to up
					c.Direction = 0
					break
				case 3:
					// go to down
					c.Direction = 1
					break
				}
				break
			}
			c.Intersection++
		}

		// fmt.Println(c)
		//Check collision after move
		ccs := checkCollisionMultiple(carts)
		if len(ccs) != 0 {
			for _, cs := range ccs {
				if !cartInSlice(cs, cartsToRemove) {
					cartsToRemove = append(cartsToRemove, cs)
					fmt.Printf("Collision with %v: %v\n", cs.ID, cs)
				}
			}
		}
	}

	// Remove carts
	for _, c := range cartsToRemove {
		fmt.Printf("Remove cart: %v\n", c)
		carts = deleteCart(carts, c)
	}

	return carts
}

func Part2(s string) string {
	rails, carts := loadData(s)

	for _, c := range carts {
		fmt.Println(c)
	}

	i := 0
	prevLen := len(carts)
	for {
		// fmt.Printf("Step: %v\n", i)
		carts = moveStepWithCheckCollisionByRemoving(rails, carts)
		if prevLen != len(carts) {
			fmt.Printf("Step: %v\n", i+1)
		}

		if len(carts) == 1 {
			fmt.Printf("End at step: %v\n", i)
			fmt.Println(carts[0])
			return fmt.Sprintf("%v,%v", carts[0].X, carts[0].Y)
		}
		i++
		prevLen = len(carts)
	}

	return ""
}
