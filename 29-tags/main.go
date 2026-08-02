package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// implement custom validator using tags
// TAGS use REFLECTION -> examine and modify behaviour at runtime
// `key:"value"`
type User struct {
	Name  string `validate:"min=2,max=32"`
	Email string `validate:"required,email"`
}

func validate(val any) error {
	v := reflect.ValueOf(val) // get the VALUE

	fmt.Println(v) // value of our fields

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := v.Type().Field(i).Tag.Get("validate")

		if tag == "" { // no validation on current field
			continue
		}

		rules := strings.Split(tag, ",") // split the tag to get the rules for validation
		fieldName := v.Type().Field(i).Name
		fieldValue := field.String()
		for _, rule := range rules {
			// TODO: make the validation rules modularised and reusable
			switch {
			case strings.HasPrefix(rule, "min="):
				min, _ := strconv.Atoi(strings.TrimPrefix(rule, "min="))
				// field FOO
				if len(fieldValue) < min {
					return fmt.Errorf("%s must be at-least %d characters long", fieldName, min)
				}
			case strings.HasPrefix(rule, "max="):
				max, _ := strconv.Atoi(strings.TrimPrefix(rule, "max="))
				// field FOO
				if len(fieldValue) > max {
					return fmt.Errorf("%s must be at-max %d characters long", fieldName, max)
				}
			case rule == "required":
				if fieldValue == "" {
					return fmt.Errorf("%s is a required field", fieldName)
				}
			case rule == "email":
				emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
				if !emailRegex.MatchString(fieldValue) {
					return fmt.Errorf("%s must be a valid email address", fieldName)
				}
			}
		}
	}

	return nil

}

func main() {
	user := User{
		Name:  "Foo",
		Email: "abc@def.com",
	}

	fmt.Println(validate(user))

	// t := reflect.TypeOf(user)
	// fmt.Println(t)
	// fmt.Println("Name:", t.Name()) // User
	// fmt.Println("Kind:", t.Kind()) // struct

	// fmt.Println("Num Field", t.NumField()) // field count

	// for i := 0; i < t.NumField(); i++ {
	// 	field := t.Field(i)

	// 	fmt.Printf("i: %d Field: %+v\n", i, field)

	// 	tag := field.Tag.Get("validate")

	// 	fmt.Printf("%d. %s (%s), tag: %v\n", i, field.Name, field.Type, tag)
	// }
}
