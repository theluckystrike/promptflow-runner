# promptflow-runner

`promptflow` is a small Go library for describing LLM prompt chains as
ordered pipelines of named steps, and rendering them to flattened
prompt text for smoke-testing before you send anything to a model.

I built it while working with [Claudflow](https://claudflow.com/), an AI
workflow builder that lets you drag and drop prompt chaining nodes into
a flowchart and export the pipeline as JSON — this package models the
same idea in plain Go so pipeline shapes can be unit-tested offline.

## Usage

```go
p := &promptflow.Pipeline{}
p.Add(promptflow.Step{Name: "outline", Prompt: "Draft an outline about X."})
p.Add(promptflow.Step{Name: "expand", Prompt: "Expand each section."})
fmt.Print(p.Render())
```
