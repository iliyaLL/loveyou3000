package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	//clear terminal and hide cursor
	fmt.Print("\x1b[2J\x1b[?25l")
	ascii := " .,-~:;!*o$#0@"

	t := 0.0
	for {
		zb := make([]float64, 40*100)
		var maxz, c, s float64 = 0, math.Cos(t), math.Sin(t)

		for y := -0.5; y <= 0.5; y += 0.01 {
			r := 0.4 + 0.05*math.Pow(0.5+0.5*math.Sin(t*6+y*2), 8)

			for x := -0.5; x <= 0.5; x += 0.01 {
				z := -x*x - math.Pow(1.2*y-math.Abs(x)*2/3, 2) + r*r
				if z < 0 {
					continue
				}
				z = math.Sqrt(z) / (2 - y)

				for tz := -z; tz <= z; tz += z / 6 {
					nx := x*c - tz*s
					nz := x*s + tz*c

					p := 1 + nz/2
					vx := int(math.Round((nx*p+0.5)*80 + 10))
					vy := int(math.Round((-y*p+0.5)*39 + 2))

					idx := vx + vy*100
					if zb[idx] <= nz {
						zb[idx] = nz
						if maxz <= nz {
							maxz = nz
						}
					}
				}
			}
		}

		fmt.Print("\x1b[H")

		for i := 0; i < 40*100; i++ {
			if i%100 == 0 {
				fmt.Print("\n")
			} else {
				fmt.Print(string(ascii[int(math.Round(zb[i]/maxz*13))]))
			}
		}

		t += 0.003
		time.Sleep(3 * time.Millisecond)
	}
}
