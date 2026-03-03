package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

type German struct {}
type Italian struct {}
type Portuguese struct {}

func (g German) LanguageName() string {
    return "German"
}

func (g German) Greet(name string) string {
    return fmt.Sprintf("Hallo %s!", name)
}

func (i Italian) LanguageName() string {
    return "Italian"
}

func (i Italian) Greet(name string) string {
    return fmt.Sprintf("Ciao %s!", name)
}

func (p Portuguese) LanguageName() string {
    return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
    return fmt.Sprintf("Olá %s!", name)
}

func SayHello(name string, langGreeter Greeter) string {
    return fmt.Sprintf("I can speak %s: %s", langGreeter.LanguageName(), langGreeter.Greet(name))
}
