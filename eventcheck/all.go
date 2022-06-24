package eventcheck

import (
	"github.com/copyco6628/lachesis-base/eventcheck/basiccheck"
	"github.com/copyco6628/lachesis-base/eventcheck/epochcheck"
	"github.com/copyco6628/lachesis-base/eventcheck/parentscheck"
	"github.com/copyco6628/lachesis-base/inter/dag"
)

// Checkers is collection of all the checkers
type Checkers struct {
	Basiccheck   *basiccheck.Checker
	Epochcheck   *epochcheck.Checker
	Parentscheck *parentscheck.Checker
}

// Validate runs all the checks except Lachesis-related
func (v *Checkers) Validate(e dag.Event, parents dag.Events) error {
	if err := v.Basiccheck.Validate(e); err != nil {
		return err
	}
	if err := v.Epochcheck.Validate(e); err != nil {
		return err
	}
	if err := v.Parentscheck.Validate(e, parents); err != nil {
		return err
	}
	return nil
}
