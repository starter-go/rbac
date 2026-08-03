package checkers

type CheckerRegistration struct {
	Label string

	Order int

	Enabled bool

	Checker Checker
}

type CheckerRegistry interface {
	Registration() *CheckerRegistration
}
