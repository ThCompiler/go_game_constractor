package genders

type Gender string

const (
	MALE   = Gender(`male`)   // мужской род
	FEMALE = Gender(`female`) // женский род
	NEUTER = Gender(`neuter`) // средний род
)
