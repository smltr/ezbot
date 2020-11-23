package main

//Function main is the bot, it uses Parse to create a map of instructions, and then uses Driver to execute each instruction.
func main() {

	f := Parse("instructions.txt")
	Driver(f)

}
