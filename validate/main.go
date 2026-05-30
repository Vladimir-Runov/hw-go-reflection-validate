package main

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type User struct {
	Name  string `validate:"min=3"`
	Age   int    `validate:"min=18;max=65"`
	Email string `validate:"regexp=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"`
}

func Validate(v interface{}) error {
	valu := reflect.ValueOf(v) //предоставляет доступ к значению и операциям с ним
	if valu.Kind() != reflect.Struct {
		return errors.New("ожидается валидация структуры, с 'validate' тегами в полях")
	}

	typ := reflect.TypeOf(v) //  возвращает информацию о типе
	for i := 0; i < valu.NumField(); i++ {
		fieldType := typ.Field(i)

		validateTag := fieldType.Tag.Get("validate")
		if validateTag == "" {
			return errors.New("поле " + fieldType.Name + " не имеет тега 'validate'")
		}
		for _, rule := range strings.Split(validateTag, ";") {
			if err := applyRuleToValofFiled(valu.Field(i), rule, fieldType.Name); err != nil {
				return err
			}
		}
	}
	return nil
}

func applyRuleToValofFiled(field reflect.Value, rule string, fieldName string) error {
	switch {
	case strings.HasPrefix(rule, "min="):
		min, err := strconv.Atoi(strings.TrimPrefix(rule, "min="))
		if err != nil {
			return fmt.Errorf("%s tag,invalid min value for field %s", rule, fieldName)
		}
		switch field.Kind() {
		case reflect.String:
			if len(field.String()) < min {
				return fmt.Errorf("%s tag, field %s length =%d must be at least %d characters", rule, fieldName, len(field.String()), min)
			}

		case reflect.Int:
			if field.Int() < int64(min) {
				return fmt.Errorf("%s tag, field %s must be at least %d", rule, fieldName, min)
			}
		}

	case strings.HasPrefix(rule, "max="):
		max, err := strconv.Atoi(strings.TrimPrefix(rule, "max="))
		if err != nil {
			return fmt.Errorf("%s tag,invalid max value for field %s", rule, fieldName)
		}
		switch field.Kind() {
		case reflect.String:
			if len(field.String()) > max {
				return fmt.Errorf("%s tag, field %s length =%d must be at most %d characters", rule, fieldName, len(field.String()), max)
			}

		case reflect.Int:
			if field.Int() > int64(max) {
				return fmt.Errorf("field %s must be at most %d", fieldName, max)
			}
		}

	case rule == "required":
		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				return fmt.Errorf("field %s is required", fieldName)
			}
		}

	case strings.HasPrefix(rule, "regexp="):
		if field.Kind() == reflect.String {
			regRes := regexp.MustCompile(strings.TrimPrefix(rule, "regexp="))
			if !regRes.MatchString(field.String()) {
				return fmt.Errorf("field %s Regexp does not match the required format: %s", fieldName, field.String())
			}
		}
	}

	return nil
}

func main() {
	usersw := []User{User{Name: "Jo", Age: 17, Email: "valid@email.su"},
		User{Name: "John A", Age: 97, Email: "valid@email.it"},
		User{Name: "John A", Age: 27, Email: "invalid-email"},
		User{Name: "John A", Age: 17, Email: "a.valid@email.com"},
		User{Name: "Jo Ok", Age: 52, Email: "a.valid@email.su"}}

	for i, user := range usersw {
		if err := Validate(user); err != nil {
			fmt.Printf("user #%d Validation error:    %s : %v\n", i+1, err, user)
		} else {
			fmt.Printf("user #%d Validation Ok!\n", i+1)
		}
	}
}