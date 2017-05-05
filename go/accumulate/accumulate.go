package accumulate

const testVersion = 1

func Accumulate(data []string, function func(string) string) []string {
	accumulation := make([]string, len(data))

	for i := 0; i < len(data); i++ {
		accumulation[i] = function(data[i])
	}

	return accumulation
}
