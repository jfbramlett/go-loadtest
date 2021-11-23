package steps

import (
	"context"

	wr "github.com/mroth/weightedrand"
)

type WeightedStepDef struct {
	Step   Step
	Weight uint
}

func NewWightedStepDef(step Step, weight uint) WeightedStepDef {
	return WeightedStepDef{Step: step, Weight: weight}
}

type weightedStep struct {
	steps *wr.Chooser
}

func (r *weightedStep) Execute(ctx context.Context) error {
	step := r.steps.Pick().(Step)
	err := step.Execute(ctx)
	return err
}

func NewWeightedStep(steps ...WeightedStepDef) Step {
	choices := make([]wr.Choice, len(steps))
	for idx, step := range steps {
		choices[idx] = wr.Choice{Item: step.Step, Weight: step.Weight}
	}

	chooser, _ := wr.NewChooser(
		choices...,
	)

	return &weightedStep{steps: chooser}
}
