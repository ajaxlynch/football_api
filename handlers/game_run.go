package handlers

import (
  "fmt"
  "encoding/json"
  "net/http"
  "api/database"
)

type gameRun struct{
  GameId          int
  Winner          *string
  HomeScore       *int
  AwayScore       *int
  Postponed       *bool
  PosponedReason  *string
}

type gameRuns struct {
  Game     []gameRun
}

func GameRunHandler(w http.ResponseWriter, r *http.Request){
  fmt.Println("In game run ")
  runs := gameRuns{}

  err := queryGameRuns(&runs)
  if err != nil{
    http.Error(w, err.Error(), 500)
    return
  }

  out, err := json.Marshal(runs)
  if err != nil{
    http.Error(w, err.Error(), 500)
    return
  }

  fmt.Fprintf(w, string(out))
}

// fetches the data from the DB
func queryGameRuns(runs *gameRuns) error {
  db := database.InitDb()

  defer db.Close()

  rows, err := db.Query(`SELECT
                          game_id,
                          winner,
                          home_score,
                          away_score,
                          postponed,
                          postponed_reason
                        FROM data.game_info`)
  if err != nil{
    return err
  }
  defer rows.Close()

  for rows.Next(){
    gameRun := gameRun{}
    err = rows.Scan(
      &gameRun.GameId,
      &gameRun.Winner,
      &gameRun.HomeScore,
      &gameRun.AwayScore,
      &gameRun.Postponed,
      &gameRun.PosponedReason)

    if err != nil{
      return err
    }
    runs.Game = append(runs.Game, gameRun)
  }

  err = rows.Err()
  if err != nil {
    return err
  }

  return nil
}
