package main

import "fmt"

func getXPos(vX int, steps int) (int, int) {
	pos := 0
	for steps != 0 {
		pos += vX
		if vX > 0 {
			vX--
		} else if vX < 0 {
			vX++
		}
		steps--
	}

	return pos, vX
}

func getYPos(vY int, steps int) int {
	var pos int = 0
	for steps != 0 {
		pos += vY
		vY--
		steps--
	}

	return pos
}

func searchX(xBegin, xEnd int) []int {
	possibilities := []int{}

	for v := 0; v <= xEnd; v++ {
		// at least you cannot cross the range by one step so vX<=xEnd
		// if xBegin=0, then this puzzle becomes meaningless, I just to shot the ball as high as I can
		// so no need to worry that.. and we need at least one step to enter the range otherwise.. meaningless

		for s := 1; ; s++ {

			// travel these steps
			xPos := 0
			steps := s
			vX := v
			for steps != 0 {
				xPos += vX
				if vX > 0 {
					vX--
				} else if vX < 0 {
					vX++
				}
				steps--
			}

			if vX == 0 && xPos < xBegin {
				// this won't reach the range
				break
			}

			if xPos >= xBegin && xPos <= xEnd {
				// this becomes one possibilities
				possibilities = append(possibilities, v)
				if vX == 0 {
					break
				}
			}

			if xPos > xEnd {
				// crossed the range without entering it
				break
			}
		}

	}

	return possibilities
}

func searchY(yLow, yHigh int) []int {

	possibilities := []int{}

	// the key is to find out the ending condition for vY
	// firstly from the example, I thought on first fall thru it would never recover
	// but it ain't true
	// then notice the trajectory, going up then coming down and it will absolutely pass 0 again
	// and when it reaches 0 again, its speed is vY+1 exactly
	// so if it crosses the area by one strike, it shall be the ending condition
	// vY+1 < 0-yLow + 1 (it can be vY+1 = 0-yLow)

	// also the vY lower bound should be yLow, you cannot cross it by one strik
	for vY := yLow; vY < 0-yLow; vY++ {
		enterRange := false
		for steps := 1; ; steps++ {
			yPos := getYPos(vY, steps)
			fmt.Printf("yPos: %v\n", yPos)

			//fmt.Printf("vY:%v, steps:%v, yPos:%v\n", vY, steps, yPos)

			if yPos >= yLow && yPos <= yHigh {
				possibilities = append(possibilities, vY)
				enterRange = true

				fmt.Printf("Possible: %v\n", vY)
			}

			if yPos < yLow {
				if enterRange {
					break
				}
				// pass thru
				// no more vY would be possible at this point
				//break YLOOP

				break
			}

		}
	}

	return possibilities
}

func searchXandY(xLeft, xRight, yLow, yHigh int) {
	count := 0
	for vX := 0; vX <= xRight; vX++ {
	YLOOP:
		for vY := yLow; vY < 0-yLow; vY++ {
			for steps := 1; ; steps++ {
				xPos, leftVX := getXPos(vX, steps)
				yPos := getYPos(vY, steps)

				if leftVX == 0 && xPos < xLeft {
					break YLOOP
				}

				if xPos > xRight {
					break YLOOP
				}

				if yPos > yHigh || xPos < xLeft {
					continue
				}

				if yPos < yLow {
					break
				}

				fmt.Printf("Starting vX:%v, vY:%v, after steps:%v, reaching xPos:%v, yPos:%v\n", vX, vY, steps, xPos, yPos)
				count += 1
				break
			}
		}
	}

	fmt.Println(count)
}

func trickShotDriver() {
	/*
		for i := 0; i < 6; i++ {
			fmt.Printf("After %v steps, the position is at %v:%v\n", i, getXPos(7, i), getYPos(2, i))
		}
	*/
	xPoss := searchX(124, 174)
	fmt.Printf("%v\n", xPoss)

	yPoss := searchY(-123, -86)
	fmt.Printf("%v\n", yPoss)

	searchXandY(20, 30, -10, -5)
	searchXandY(124, 174, -123, -86)
}
