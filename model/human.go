package model

import (
	"errors"
)

type humanModel struct {
	Name string
	Age  int
}

func NewHumanModel(
	name string,
	age int) (*humanModel, error) {

	// name が空文字の時はエラー
	if name == "" {
		return nil, errors.New("nameは空文字にできません。")
	}

	// age は 0 以上の値でなければならない
	if age < 0 {
		return nil, errors.New("ageは0以上の値でなければなりません。")
	}

	res := &humanModel{
		Name: name,
		Age:  age,
	}

	return res, nil
}
