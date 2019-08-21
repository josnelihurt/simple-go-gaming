package engine

import "math"

/// Circle describes a circle
type Circle struct {
	Center *Vector
	Radius float64
}

func collides(c1, c2 Circle) bool {
	dist := math.Sqrt(math.Pow(c2.Center.X-c1.Center.X, 2) + math.Pow(c2.Center.Y-c1.Center.Y, 2))
	return dist <= c1.Radius+c2.Radius
}

func CheckColisions(pool *ElementPool) error {
	elements := pool.elements // I don't like it

	for i := 0; i < len(elements)-1; i++ {
		for j := 0; j < len(elements); j++ {
			for _, currenCollision1 := range elements[i].Collisions {
				for _, currenCollision2 := range elements[j].Collisions {
					if elements[i].Active && elements[j].Active && collides(currenCollision1, currenCollision2) {
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
