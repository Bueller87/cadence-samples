package main

import "github.com/uber-common/cadence-samples/new_samples/template"

func main() {
	// Define the data for Data samples
	// Note: worker.go is hand-written (not generated) because this sample
	// requires a custom DataConverter in worker options
	data := template.TemplateData{
		SampleName: "Data",
		Workflows:  []string{"LargeDataConverterWorkflow"},
		Activities: []string{"LargeDataConverterActivity"},
	}

	template.GenerateAll(data)
}

// Implement custom generator below
