package engine

import "math"

type Circle struct {
	center *Vector
	radius float64
}

func collides(c1, c2 Circle) bool {
	dist := math.Sqrt(math.Pow(c2.center.x-c1.center.x, 2) + math.Pow(c2.center.y-c1.center.y, 2))
	return dist <= c1.radius+c2.radius
}

func CheckColisions(pool *ElementPool) error {
	elements := pool.elements // I don't like it

	for i := 0; i < len(elements)-1; i++ {
		for j := 0; j < len(elements); j++ {
			for _, currenCollision1 := range elements[i].collisions {
				for _, currenCollision2 := range elements[j].collisions {
					if elements[i].active && elements[j].active && collides(currenCollision1, currenCollision2) {
						if err := elements[i].Collision(elements[j]); err != nil {
							return err
						}
						if err := elements[j].Collision(elements[i]); err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}
