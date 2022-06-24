package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"nba-scores/utils"
)

func main() {
	key := os.Getenv("API_KEY")
	fmt.Print("Enter date in format (2022-01-01): ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("error while reading input. Please try again", err)
	}
	date := strings.TrimSuffix(input, "\n")
	games, err := utils.GetGames(date, key)
	if err != nil {
		log.Fatalf("unable to retrieve game information")
	}

	fmt.Printf("There were %d games played on %s\n", games.Results, date)
	for _, game := range games.Info {
		fmt.Printf("Home Team: %s (%d) vs Visitor Team: %s (%d)\n", game.Team.Home.Nickname, game.Scores.Home.Points, game.Team.Visitor.Nickname, game.Scores.Visitor.Points)

	}

}
