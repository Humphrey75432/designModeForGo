package factory

import "fmt"

func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil
	}

	if gunType == "musket" {
		return NewMusket(), nil
	}

	return nil, fmt.Errorf("wrong gun type passed.")
}
