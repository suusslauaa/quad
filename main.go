package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func quad(width, height int, top_left_corner, between_top_corners, top_right_corner, side, bottom_left_corner, between_bottom_corners, bottom_right_corner string) {
	if width <= 0 || height <= 0 {
		return
	}
	for i := 1; i <= width; i++ {
		if i == 1 {
			fmt.Print(top_left_corner)
			continue
		}
		if i == width {
			fmt.Print(top_right_corner)
			continue
		}
		fmt.Print(between_top_corners)
	}
	fmt.Println()
	for j := 1; j <= height-2; j++ {
		for i := 1; i <= width; i++ {
			if i == 1 || i == width {
				fmt.Print(side)
				continue
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	if height != 1 {
		for i := 1; i <= width; i++ {
			if i == 1 {
				fmt.Print(bottom_left_corner)
				continue
			}
			if i == width {
				fmt.Print(bottom_right_corner)
				continue
			}
			fmt.Print(between_bottom_corners)
		}
		fmt.Println()
	}
}

func main() {
	if len(os.Args) != 10 {
		fmt.Print("Usage of quad:\nTo run this program, you need to enter all 9 arguments as shown below\ngo run . width height top_left top top_right sides bottom_left bottom bottom_right")
		return
	}

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Print(err)
	}

	y, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Print(err)
	}

	quad(x, y, os.Args[3], os.Args[4], os.Args[5], os.Args[6], os.Args[7], os.Args[8], os.Args[9])
}
