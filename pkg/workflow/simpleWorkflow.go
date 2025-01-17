package workflow

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type SimpleWorkflow struct {
	steps        []Step
	status       Status
	dependencies map[Step][]Step
}

var stepOutputContextKey struct{}

// hasCycle check if the workflow has cycle starting from currentStep.
// visited contains the list of already visited steps.
// secureCycles contains the list of all steps that have already been checked and can be skipped.
func (s *SimpleWorkflow) hasCycle(visited []Step, secureCycles []Step, currentStep Step) (bool, []Step) {
	if sliceContainsStep(secureCycles, currentStep) {
		// The step has already been checked and has no cycle
		return false, []Step{}
	}
	if sliceContainsStep(visited, currentStep) {
		// The step has already been visited, this is a cycle
		return true, []Step{}
	}
	visited = append(visited, currentStep)

	currentStepSecureCycles := make([]Step, 0, len(s.steps))

	// Check for each step if its dependencies contains cycles
	if dependencies, ok := s.dependencies[currentStep]; ok {
		for _, dependency := range dependencies {
			if hasCycle, dependencySecureCycles := s.hasCycle(visited, secureCycles, dependency); !hasCycle {
				currentStepSecureCycles = append(currentStepSecureCycles, dependencySecureCycles...)
			} else {
				// A dependency has a cycle, return true and no secure steps
				return true, []Step{}
			}
		}
	}
	// this dependency is safe
	currentStepSecureCycles = append(currentStepSecureCycles, currentStep)
	return false, currentStepSecureCycles
}

func (s *SimpleWorkflow) Init(_ context.Context) error {
	visitedSteps := make([]Step, 0, len(s.steps))
	secureSteps := make([]Step, 0, len(s.steps))

	for _, currentStep := range s.steps {
		if hasCycles, stepDependency := s.hasCycle(visitedSteps, secureSteps, currentStep); !hasCycles {
			secureSteps = append(secureSteps, stepDependency...)
		} else {
			return fmt.Errorf("cycle detected")
		}
	}

	return nil
}

// MakeSimpleWorkflow creates a SimpleWorkflow instance
func MakeSimpleWorkflow(stepCount uint) *SimpleWorkflow {
	return &SimpleWorkflow{
		steps:        make([]Step, 0, stepCount),
		status:       CREATED,
		dependencies: make(map[Step][]Step),
	}
}

func (s *SimpleWorkflow) GetStatus() Status {
	return s.status
}

func (s *SimpleWorkflow) getNextSteps() []Step {
	nextSteps := make([]Step, 0, len(s.steps))
	for _, step := range s.steps {
		if !step.GetStatus().IsFinished() && step.GetStatus() != RUNNING && areRequirementsFullFilled(step, s.dependencies) {
			nextSteps = append(nextSteps, step)
		}
	}
	return nextSteps
}

func (s *SimpleWorkflow) executeStep(ctx context.Context, step Step) {
	requirementsOutputs := s.getRequirementsOutputs(step)
	stepContext := context.WithValue(ctx, stepOutputContextKey, requirementsOutputs)
	_, err := step.Execute(stepContext)
	if err != nil {
		log.Printf("[WARN]: step %v ended with error : %v", step.GetId(), err)
	}
}

func (s *SimpleWorkflow) cancelNextSteps(lastStep Step) error {
	errorsLst := createErrorList(len(s.steps))

	concernedSteps := getStepThatRequires(lastStep, s.dependencies)

	for _, step := range concernedSteps {
		if step.GetStatus().IsCancellable() {
			err := step.Cancel()
			if err != nil {
				errorsLst = append(errorsLst, fmt.Errorf("fail to cancel step %v : %v", step.GetId(), err))
			}
		}
	}
	if len(errorsLst) != 0 {
		return joinErrorList(errorsLst)
	}
	return nil
}

func (s *SimpleWorkflow) getRequirementsOutputs(step Step) map[string]map[string]string {
	res := make(map[string]map[string]string)
	stepDependencies := s.dependencies[step]

	for _, dependency := range stepDependencies {
		res[dependency.GetId()] = dependency.GetOutput()
		res = mergeStringStringStringMaps(res, s.getRequirementsOutputs(dependency))
	}

	return res
}

func (s *SimpleWorkflow) getOutput() map[string]map[string]string {
	output := make(map[string]map[string]string)
	for _, step := range s.steps {
		output[step.GetId()] = step.GetOutput()
	}
	return output
}

func (s *SimpleWorkflow) hasUnfinishedSteps() bool {
	for _, step := range s.steps {
		if !step.GetStatus().IsFinished() {
			return true
		}
	}
	return false
}

func (s *SimpleWorkflow) Execute(ctx context.Context) (map[string]map[string]string, error) {
	closingSteps := make(chan Step, len(s.steps))
	errorLst := make([]error, 0, len(s.steps))

	activeSteps := &sync.WaitGroup{}

	s.startNextSteps(ctx, activeSteps, closingSteps)

	for s.hasUnfinishedSteps() {
		select {
		case <-ctx.Done():
			errorLst = s.cancelRemainingSteps(errorLst)
			close(closingSteps)

		case step, ok := <-closingSteps:
			if ok {

			}
			log.Printf("Step %v terminated with status %v", step.GetId(), step.GetStatus().GetName())
			// A step as ended
			if step.GetStatus() == ERROR || step.GetStatus() == CANCELLED {
				errorLst = appendErrorList(errorLst, s.cancelNextSteps(step))
			} else {
				s.startNextSteps(ctx, activeSteps, closingSteps)
			}
		}
	}

	if len(errorLst) > 0 {
		return nil, joinErrorList(errorLst)
	}

	activeSteps.Wait()

	return s.getOutput(), nil
}

func (s *SimpleWorkflow) cancelRemainingSteps(errorLst []error) []error {
	for _, step := range s.steps {
		if step.GetStatus().IsCancellable() {
			err := step.Cancel()
			if err != nil {
				errorLst = append(errorLst, err)
			}
		}
	}
	return errorLst
}

func (s *SimpleWorkflow) startNextSteps(ctx context.Context, activeSteps *sync.WaitGroup, closingSteps chan Step) {
	nextSteps := s.getNextSteps()
	for _, step := range nextSteps {
		activeSteps.Add(1)

		go func(step Step) {
			s.executeStep(ctx, step)
			closingSteps <- step
			activeSteps.Done()
		}(step)
	}
}

func (s *SimpleWorkflow) debug() {
	for _, step := range s.steps {
		log.Printf("[DEBUG]: %v : %v", step.GetId(), step.GetStatus().GetName())
	}
}

func (s *SimpleWorkflow) AddStep(step Step, dependencies []Step) error {
	for _, previousStep := range s.steps {
		if previousStep.GetId() == step.GetId() {
			return fmt.Errorf("step [%s] is already present in workflow", step.GetId())
		}
	}

	s.steps = append(s.steps, step)
	s.dependencies[step] = dependencies
	return nil
}
