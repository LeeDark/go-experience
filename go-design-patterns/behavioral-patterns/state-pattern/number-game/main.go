package main

func main() {
	start := StartState{}
	game := GameContext{
		Next: &start,
	}
	for game.Next.executeState(&game) {
	}
}
