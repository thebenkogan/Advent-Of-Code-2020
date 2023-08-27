package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	day := os.Args[1]
	var inputFileName string
	if len(os.Args) > 2 && os.Args[2] == "test" {
		inputFileName = "test"
	} else {
		inputFileName = "in"
	}
	pwd, _ := os.Getwd()
	path := fmt.Sprintf("%s/cmd/%s/%s.txt", pwd, day, inputFileName)
	b, _ := os.ReadFile(path)
	input := string(b)

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v", part2(input))
}

type Food struct {
	ingredients map[string]bool // set of ingredients in this food
	allergens   map[string]bool // set of allergens in this food
}

func parseFoods(input string) []Food {
	foods := make([]Food, 0)
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line[:len(line)-1], " (contains ")
		ingredients := make(map[string]bool)
		for _, i := range strings.Split(split[0], " ") {
			ingredients[i] = true
		}
		allergens := make(map[string]bool)
		for _, a := range strings.Split(split[1], ", ") {
			allergens[a] = true
		}
		foods = append(foods, Food{ingredients, allergens})
	}
	return foods
}

// returns a map of ingredient to the set of allergens it could have.
// empty set if that ingredient could not have any allergen
func getIngredientsWithAllergens(foods []Food) map[string]map[string]bool {
	ingredientsPossibleAllergens := make(map[string]map[string]bool)
	for i, food := range foods {
		for ingredient := range food.ingredients {
			if _, ok := ingredientsPossibleAllergens[ingredient]; !ok {
				ingredientsPossibleAllergens[ingredient] = make(map[string]bool)
			}
			for allergen := range food.allergens {
				valid := true
				for j, otherFood := range foods {
					if i == j {
						continue
					}
					if !otherFood.ingredients[ingredient] && otherFood.allergens[allergen] {
						valid = false
						break
					}
				}
				if valid {
					ingredientsPossibleAllergens[ingredient][allergen] = true
				}
			}
		}
	}

	return ingredientsPossibleAllergens
}

func part1(input string) int {
	foods := parseFoods(input)
	ingredientsPossibleAllergens := getIngredientsWithAllergens(foods)

	badIngredients := make([]string, 0)
	for ingredient, allergens := range ingredientsPossibleAllergens {
		if len(allergens) == 0 {
			badIngredients = append(badIngredients, ingredient)
		}
	}

	badCount := 0
	for _, food := range foods {
		for _, badIngredient := range badIngredients {
			if food.ingredients[badIngredient] {
				badCount++
			}
		}
	}

	return badCount
}

func part2(input string) string {
	foods := parseFoods(input)
	ingredientsPossibleAllergens := getIngredientsWithAllergens(foods)

	ingredientToAllergen := make(map[string]string)
	for len(ingredientToAllergen) < len(ingredientsPossibleAllergens) {
		var knownAllergen string
		for ingredient, allergens := range ingredientsPossibleAllergens {
			if len(allergens) == 1 {
				for allergen := range allergens {
					knownAllergen = allergen
					ingredientToAllergen[ingredient] = allergen
				}
			}
		}
		if knownAllergen == "" {
			break
		}
		for _, allergens := range ingredientsPossibleAllergens {
			delete(allergens, knownAllergen)
		}
	}

	ingredients := make([]string, 0, len(ingredientToAllergen))
	for ingredient := range ingredientToAllergen {
		ingredients = append(ingredients, ingredient)
	}

	slices.SortFunc(ingredients, func(i1 string, i2 string) bool {
		return ingredientToAllergen[i1] < ingredientToAllergen[i2]
	})

	return strings.Join(ingredients, ",")
}
