package hello

const testVersion = 2

// Greets the user by name, or by saying "Hello, World!" if no name is given
func HelloWorld(name string) string {
  if name == "" {
    name = "World"
  }
	return "Hello, " + name + "!"
}
