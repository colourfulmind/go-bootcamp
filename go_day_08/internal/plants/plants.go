package plants

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func DescribePlant(plant interface{}) error {
	TypeOfPlant := reflect.TypeOf(plant)

	if TypeOfPlant.Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("invalid data type: %v", TypeOfPlant))
	}

	ValueOfPlant := reflect.ValueOf(plant)
	for i := 0; i < TypeOfPlant.NumField(); i++ {
		field := TypeOfPlant.Field(i)
		fmt.Print(field.Name)

		tag := field.Tag
		if tag != "" {
			values := strings.Split(string(tag), ":")
			values[1] = strings.TrimSuffix(strings.TrimPrefix(values[1], "\""), "\"")
			fmt.Printf("(%s=%s)", values[0], values[1])
		}

		fmt.Print(":")
		fmt.Println(ValueOfPlant.Field(i))
	}

	return nil
}
