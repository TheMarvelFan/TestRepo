package tournament

import (
    "bufio"
    "errors"
    "fmt"
    "io"
    "sort"
    "strings"
)

type Scores struct {
    Mp int
    W int
    D int
    L int
    P int
}

type KeyVal struct {
    Key string
    Val Scores
}

func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
    scoreBoard := map[string]Scores{}

    var line string
    var team string
	var data []string
    var score Scores
    
    for scanner.Scan() {
        line = scanner.Text()
        data = strings.Split(line, ";")

        if len(data) < 2 {
            continue
        }

        if len(data) == 2 || len(data) > 3 {
            return errors.New("Wrong input format")
        }

        teamOne, presentOne := scoreBoard[data[0]]
        teamTwo, presentTwo := scoreBoard[data[1]]

        if data[2] == "draw" {
            if !presentOne {
                scoreBoard[data[0]] = Scores{
                    Mp: 1,
                    W: 0,
                    D: 1,
                    L: 0,
                    P: 1,
                }
            } else {
                teamOne.Mp += 1
                teamOne.D += 1
                teamOne.P += 1

                scoreBoard[data[0]] = teamOne
            }

            if !presentTwo {
                scoreBoard[data[1]] = Scores{
                    Mp: 1,
                    W: 0,
                    D: 1,
                    L: 0,
                    P: 1,
                }
            } else {
                teamTwo.Mp += 1
                teamTwo.D += 1
                teamTwo.P += 1

                scoreBoard[data[1]] = teamTwo
            }
        } else if data[2] == "win" {
            if !presentOne {
                scoreBoard[data[0]] = Scores{
                    Mp: 1,
                    W: 1,
                    D: 0,
                    L: 0,
                    P: 3,
                }
            } else {
                teamOne.Mp += 1
                teamOne.W += 1
                teamOne.P += 3

                scoreBoard[data[0]] = teamOne
            }

            if !presentTwo {
                scoreBoard[data[1]] = Scores{
                    Mp: 1,
                    W: 0,
                    D: 0,
                    L: 1,
                    P: 0,
                }
            } else {
                teamTwo.Mp += 1
                teamTwo.L += 1

                scoreBoard[data[1]] = teamTwo
            }
        } else if data[2] == "loss" {
            if !presentOne {
                scoreBoard[data[0]] = Scores{
                    Mp: 1,
                    W: 0,
                    D: 0,
                    L: 1,
                    P: 0,
                }
            } else {
                teamOne.Mp += 1
                teamOne.L += 1

                scoreBoard[data[0]] = teamOne
            }

            if !presentTwo {
                scoreBoard[data[1]] = Scores{
                    Mp: 1,
                    W: 1,
                    D: 0,
                    L: 0,
                    P: 3,
                }
            } else {
                teamTwo.Mp += 1
                teamTwo.W += 1
                teamTwo.P += 3

                scoreBoard[data[1]] = teamTwo
            }
        } else {
            return errors.New("Wrong input format")
        }
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    entries := []KeyVal{}

    for k, v := range scoreBoard {
        entries = append(entries, KeyVal{
            Key: k,
            Val: v,
        })
    }

    sort.Slice(entries, func(i, j int) bool {
        if entries[i].Val.P == entries[j].Val.P {
            return entries[i].Key < entries[j].Key
        }

        return entries[i].Val.P > entries[j].Val.P
    })

    fmt.Fprintf(writer, "Team                           | MP |  W |  D |  L |  P\n")

    for _, entry := range entries {
		team = entry.Key
        score = entry.Val
        
        fmt.Fprintf(writer, 
                    "%-31s|  %d |  %d |  %d |  %d |  %d\n",
                   	team, score.Mp, score.W, score.D, score.L, score.P)
    }
    
    return nil
}
