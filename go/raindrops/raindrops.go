package raindrops

import "strconv"

const testVersion = 3

type Rain struct {
	number   int
	sound    string
}

func Convert(x int) string {
	rain := Rain{x, ""}

  rain.AppendSound(3, "Pling")
  rain.AppendSound(5, "Plang")
  rain.AppendSound(7, "Plong")

	if rain.sound == "" {
		rain.sound = strconv.Itoa(x)
	}

	return rain.sound
}

func (rain *Rain) AppendSound(factor int, pl string) {
  if rain.number % factor == 0 {
    rain.sound += pl
  }
}
