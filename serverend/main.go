package main

func main() {
	server := InitServer()
	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
