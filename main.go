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

const requestActionChoice string = "Write action (buy, fill, take, remaining, exit):\n"
const requestCoffeeTypeChoice string = "What do you want to buy? 1 - espresso, " +
	"2 - latte, 3 - cappuccino, 4 - marocchino:\n"
const requestAmountWaterAdd string = "Write how many ml of water you want to add:\n"
const requestAmountMilkAdd string = "Write how many ml of milk you want to add:\n"
const requestAmountCoffeeBeansAdd string = "Write how many grams of coffee beans you want to add:\n"
const requestAmountCupsAdd string = "Write how many disposable cups of coffee you want to add:\n"
const responseOnTake string = "I gave you $%d\n"

const errorInvalidInput string = "[Error] invalid input"
const errorUnknownCommand string = "[Error] unknown command"
const errorInvalidCoffeeType string = "[Error] invalid coffee type chosen"
const errorNotEnoughIngredients string = "[Error] out of ingredients"

var notSoErrorExit = errors.New("exiting")

type recipe struct {
	milk  uint
	water uint
	beans uint
	cost  uint
}

type coffeeMachine struct {
	milk  uint
	water uint
	beans uint
	cups  uint
	money uint
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

var marocchinoRecipe = recipe{
	milk:  25,
	water: 250,
	beans: 16,
	cost:  8,
}

var knownRecipes = []recipe{
	espressoRecipe,
	latteRecipe,
	cappuccinoRecipe,
	marocchinoRecipe,
}

var knownCommands = map[string]action{
	"buy":       (*coffeeMachine).buy,
	"fill":      (*coffeeMachine).fill,
	"take":      (*coffeeMachine).take,
	"remaining": (*coffeeMachine).state,
	"exit":      (*coffeeMachine).exit,
}

func (s *coffeeMachine) exit() error {
	return notSoErrorExit
}

func (s *coffeeMachine) fill() error {
	var water, milk, beans, cups uint
	err := requestUserInput(requestAmountWaterAdd, &water)
	if err != nil {
		return err
	}
	err = requestUserInput(requestAmountMilkAdd, &milk)
	if err != nil {
		return err
	}
	err = requestUserInput(requestAmountCoffeeBeansAdd, &beans)
	if err != nil {
		return err
	}
	err = requestUserInput(requestAmountCupsAdd, &cups)
	if err != nil {
		return err
	}
	s.water += water
	s.milk += milk
	s.beans += beans
	s.cups += cups
	return nil
}

func (s *coffeeMachine) take() error {
	fmt.Printf(responseOnTake, s.money)
	s.money = 0
	return nil
}

func (s *coffeeMachine) validateBuy(coffee recipe) bool {
	if s.water >= coffee.water && s.milk >= coffee.milk && s.beans >= coffee.beans && s.cups > 0 {
		return true
	}
	return false
}

func (s *coffeeMachine) buy() error {
	var coffeeType int
	err := requestUserInput(requestCoffeeTypeChoice, &coffeeType)
	if err != nil {
		return err
	}

	if coffeeType < 1 || coffeeType > len(knownRecipes) {
		return errors.New(errorInvalidCoffeeType)
	}
	coffeeType--
	if s.validateBuy(knownRecipes[coffeeType]) {
		s.water -= knownRecipes[coffeeType].water
		s.milk -= knownRecipes[coffeeType].milk
		s.beans -= knownRecipes[coffeeType].beans
		s.cups--
		s.money += knownRecipes[coffeeType].cost
	} else {
		return errors.New(errorNotEnoughIngredients)
	}
	return nil
}

func (s *coffeeMachine) state() error {
	fmt.Printf(infoCoffeeMachineState, s.water, s.milk, s.beans, s.cups, s.money)
	return nil
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

	for {
		command, err := machine.request()
		if err != nil {
			fmt.Println(err)
			continue
		}
		f, exist := knownCommands[command]
		if exist {
			err := f(&machine)
			if err != nil {
				if errors.Is(err, notSoErrorExit) {
					return
				}
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println(errorUnknownCommand)
			continue
		}
	}
}
