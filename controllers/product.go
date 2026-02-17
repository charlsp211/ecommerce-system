package controllers

import "errors"

func ValidateProduct(name string) error {
	if name == "" {
		return errors.New("producto sin nombre")
	}
	return nil
}
