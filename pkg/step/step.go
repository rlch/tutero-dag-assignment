package step

import "tutero_assignment/pkg/src/graph"

type Stepper interface {
	// Step returns a prediction for the correct node; or an error if a prediction cannot be made.
	Step(graph graph.Graph) (graph.Node, error)
}

func New() *stepper {
	//* You may mutate this instantiation if necessary; but the function signature should not change.
	return &stepper{}
}

type stepper struct {
	//* You may add fields to this struct.
}

func (s *stepper) Step(graph graph.Graph) (graph.Node, error) {
	//* Implement the Step function.
	return graph.Nodes()[len(graph.Nodes())-1], nil // nieve solution -- returns a random node.
}
