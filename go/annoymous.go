package main

func main() {
	var str = &struct {
		settings struct {
			name string
		}
	}{
		settings: struct{ name string }{
			name: "fuck",
		},
	}
}
