package accumulate

const testVersion = 1

func Accumulate(poop []string, crap func(string) string) []string {
	accumulation := make([]string, len(poop))

	for i := 0; i < len(poop); i++ {
		accumulation[i] = crap(poop[i])
	}

	return accumulation
}
