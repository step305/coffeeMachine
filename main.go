package main

import (
	"errors"
	"fmt"
)

const infoCoffeeMachineState string = "The coffee machine has:\n" +
	"%d ml of water\n" +
	"%d ml of milk\n" +
	"%d g of coffee beans\n" +
	"%d disposable cups\n" +
	"$%d of money\n"

const requestActionChoice string = "Write action (buy, fill, take):\n"
const requestCoffeeTypeChoice string = "What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:\n"
const requestAmountWaterAdd string = "Write how many ml of water you want to add:\n"
const requestAmountMilkAdd string = "Write how many ml of milk you want to add:\n"
const requestAmountCoffeeBeansAdd string = "Write how many grams of coffee beans you want to add:\n"
const requestAmountCupsAdd string = "Write how many disposable cups of coffee you want to add:\n"
const responseOnTake string = "I gave you $%d\n"

const errorInvalidInput string = "[Error] invalid input"
const errorUnknownCommand string = "[Error] unknown command"
const errorInvalidCoffeeType string = "[Error] invalid coffee type chosen"
const errorDivisionByZero string = "[Error] division by zero"
const errorNotEnoughIngredients string = "No, I can make only %d cups of coffee"

type recipe struct {
	milk  int
	water int
	beans int
	cost  int
}

type coffeeMachine struct {
	milk  int
	water int
	beans int
	cups  int
	money int
}

type action func(machine *coffeeMachine) error

var espressoRecipe = recipe{
	milk:  0,
	water: 250,
	beans: 16,
	cost:  4,
}

var latteRecipe = recipe{
	milk:  75,
	water: 350,
	beans: 20,
	cost:  7,
}

var cappuccinoRecipe = recipe{
	milk:  100,
	water: 200,
	beans: 12,
	cost:  6,
}

var knownCommands = map[string]action{
	"buy":  (*coffeeMachine).buy,
	"fill": (*coffeeMachine).fill,
	"take": (*coffeeMachine).take,
}

func (s *coffeeMachine) fill() error {
	fmt.Println("I am fill")
	return nil
}

func (s *coffeeMachine) take() error {
	fmt.Println("I am take")
	return nil
}

func (s *coffeeMachine) buy() error {
	fmt.Println("I am buy")
	return nil
}

func (s *coffeeMachine) state() string {
	return fmt.Sprintf(infoCoffeeMachineState, s.water, s.milk, s.beans, s.cups, s.money)
}

func (s *coffeeMachine) request() (string, error) {
	var command string
	err := requestUserInput(requestActionChoice, &command)
	if err != nil {
		return "", err
	}
	return command, nil
}

func requestUserInput[T any](prompt string, value T) error {
	fmt.Print(prompt)
	_, err := fmt.Scanln(value)
	if err != nil {
		return errors.New(errorInvalidInput)
	}
	return nil
}

func main() {
	var machine = coffeeMachine{
		milk:  540,
		water: 400,
		beans: 120,
		cups:  9,
		money: 550,
	}

	fmt.Print(machine.state())
	command, err := machine.request()
	if err != nil {
		fmt.Println(err)
		return
	}
	f, exist := knownCommands[command]
	if exist {
		err := f(&machine)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println(errorUnknownCommand)
		return
	}
	fmt.Print(machine.state())
}
