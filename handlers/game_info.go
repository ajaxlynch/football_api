package handlers

import (
  "fmt"
  "encoding/json"
  "net/http"
  "api/database"
)

type gameinfo struct{
  Id            string
  HomeTeam      string
  AwayTeam      string
  KickOff       string
  DateOfFixture string
  Favourite     string
  Venue         string
}

type gameInfoStats struct {
  GameStats     []gameinfo
}

func GameHandler(w http.ResponseWriter, r *http.Request){
  fmt.Println("In Gameinfo")
  gameStats := gameInfoStats{}

  err := queryGames(&gameStats)
  if err != nil{
    http.Error(w, err.Error(), 500)
    return
  }

  out, err := json.Marshal(gameStats)
  if err != nil{
    http.Error(w, err.Error(), 500)
    return
  }

  fmt.Fprintf(w, string(out))
}

// fetches the data from the DB
func queryGames(gameStats *gameInfoStats) error {
  db := database.InitDb()

  defer db.Close()

  rows, err := db.Query(`SELECT
                          id,
                          home_team,
                          away_team,
                          kickoff,
                          date_of_fixture,
                          favourite,
                          venue
                        FROM data.fixtures`)
  if err != nil{
    return err
  }
  defer rows.Close()

  for rows.Next(){
    gameinfo := gameinfo{}
    err = rows.Scan(
      &gameinfo.Id,
      &gameinfo.HomeTeam,
      &gameinfo.AwayTeam,
      &gameinfo.KickOff,
      &gameinfo.DateOfFixture,
      &gameinfo.Favourite,
      &gameinfo.Venue)
    if err != nil{
      return err
    }
    gameStats.GameStats = append(gameStats.GameStats, gameinfo)
  }

  err = rows.Err()
  if err != nil {
    return err
  }

  return nil
}
