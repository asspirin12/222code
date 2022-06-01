package main

type GChannel[C any] <-chan C

func main() {

	var c = GChannel[int]()
	print(c)
}
