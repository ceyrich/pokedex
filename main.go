package main

import( 

)

func main() {
	config := replConfig{
		commands: getCommands(),
	}
	StartRepl(&config)
}



