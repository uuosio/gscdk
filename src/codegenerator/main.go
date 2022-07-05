package main

func main() {
	err := GenerateCode(".", "generated.go", nil)
	if err != nil {
		panic(err)
	}
}
