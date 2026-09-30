// Package promptflow is a tiny library for chaining LLM prompt steps
// into a linear pipeline and running them in order.
package promptflow

import "fmt"

// Step is one stage in a prompt pipeline.
type Step struct {
	Name    string
	Prompt  string
	Depends []string
}

// Pipeline is an ordered set of steps.
type Pipeline struct {
	Steps []Step
}

// Add appends a step and returns the pipeline for chaining.
func (p *Pipeline) Add(s Step) *Pipeline {
	p.Steps = append(p.Steps, s)
	return p
}

// Render returns the flattened prompt text for a pipeline, with each
// step's name as a heading. It is useful for smoke-testing that a
// chain is well-formed before sending it to a model.
func (p *Pipeline) Render() string {
	out := ""
	for _, s := range p.Steps {
		out += fmt.Sprintf("## %s\n%s\n\n", s.Name, s.Prompt)
	}
	return out
}
