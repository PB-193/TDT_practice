package model

type NameModel struct {
	FamilyName string
	LastName   string
}

func NewNameModel(familyName string, lastName string) *NameModel {
	return &NameModel{
		FamilyName: familyName,
		LastName:   lastName,
	}
}
