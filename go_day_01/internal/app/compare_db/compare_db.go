// Package compare_db compares two databases and outputs the difference.
package compare_db

import (
	"fmt"
	"go_day_01/internal/app/read_db"
)

// CompareDB compares cakes in two different databases
func CompareDB(OldDB, NewDB *read_db.DB) {
	CompareCakes(OldDB, NewDB)
	for _, OldCake := range OldDB.Cakes {
		for _, NewCake := range NewDB.Cakes {
			if OldCake.Name == NewCake.Name {
				CompareTime(OldCake, NewCake)
				break
			}
		}
	}
}

// CompareCakes compares titles of the cakes
func CompareCakes(OldCakes, NewCakes *read_db.DB) {
	for _, NewCake := range NewCakes.Cakes {
		Flag := true
		for _, OldCake := range OldCakes.Cakes {
			if NewCake.Name == OldCake.Name {
				Flag = false
				break
			}
		}
		if Flag {
			fmt.Println("ADDED cake \"" + NewCake.Name + "\"")
		}
	}
	for _, OldCake := range OldCakes.Cakes {
		Flag := true
		for _, NewCake := range NewCakes.Cakes {
			if OldCake.Name == NewCake.Name {
				Flag = false
				break
			}
		}
		if Flag {
			fmt.Println("REMOVED cake \"" + OldCake.Name + "\"")
		}
	}
}

// CompareTime compares cooking time of the same cake
func CompareTime(OldCake, NewCake read_db.Cake) {
	if OldCake.Time != NewCake.Time {
		fmt.Println("CHANGED cooking time for cake \"" + OldCake.Name +
			"\" - \"" + NewCake.Time + "\" instead of \"" + OldCake.Time + "\"")
	}
	CompareIngredients(OldCake, NewCake)
}

// CompareIngredients compares ingredients of the same cake
func CompareIngredients(OldCake, NewCake read_db.Cake) {
	for _, NewIn := range NewCake.CakeIngredients {
		Flag := true
		for _, OldIn := range OldCake.CakeIngredients {
			if NewIn.IngredientName == OldIn.IngredientName {
				Flag = false
				break
			}
		}
		if Flag {
			fmt.Println("ADDED ingredient \"" + NewIn.IngredientName +
				"\" for cake \"" + NewCake.Name + "\"")
		}
	}
	for _, OldIn := range OldCake.CakeIngredients {
		Flag := true
		for _, NewIn := range NewCake.CakeIngredients {
			if OldIn.IngredientName == NewIn.IngredientName {
				Flag = false
				break
			}
		}
		if Flag {
			fmt.Println("REMOVED ingredient \"" + OldIn.IngredientName +
				"\" for cake \"" + OldCake.Name + "\"")
		}
	}
	for _, OldIn := range OldCake.CakeIngredients {
		for _, NewIn := range NewCake.CakeIngredients {
			if OldIn.IngredientName == NewIn.IngredientName {
				CompareUnits(OldIn, NewIn, OldCake.Name)
				break
			}
		}
	}
}

// CompareUnits compares units of the same ingredient
func CompareUnits(OldIngredients, NewIngredients read_db.Ingredients, Cake string) {
	if OldIngredients.IngredientUnit != NewIngredients.IngredientUnit {
		if NewIngredients.IngredientUnit == "" {
			fmt.Println("REMOVED unit \"" + OldIngredients.IngredientUnit +
				"\" for ingredient \"" + OldIngredients.IngredientName +
				"\" for cake  \"" + Cake + "\"")
		} else {
			fmt.Println("CHANGED unit for ingredient \"" + OldIngredients.IngredientName +
				"\" for cake  \"" + Cake +
				"\" - \"" + NewIngredients.IngredientUnit +
				"\" instead of \"" + OldIngredients.IngredientUnit + "\"")
		}
	} else if OldIngredients.IngredientCount != NewIngredients.IngredientCount {
		fmt.Println("CHANGED unit count for ingredient \"" + OldIngredients.IngredientName +
			"\" for cake  \"" + Cake + "\" - \"" +
			NewIngredients.IngredientCount + "\" instead of \"" +
			OldIngredients.IngredientCount + "\"")
	}
}
