package reports

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"

	"github.com/olblak/updateCli/pkg/core/result"
)

const reportsTpl string = `
=============================

REPORTS:

{{ range . }}
{{ if  .Err }}
{{- .Result }} {{ .Name -}}{{"\n"}}
{{ "\t"}}Error: {{ .Err}}
{{ else }}
{{- .Result }} {{ .Name -}}{{"\n"}}
{{- "\t"}}Source:
{{ "\t"}}{{"\t"}}{{- .Source.Result }}  {{ .Source.Name -}}({{- .Source.Kind -}}){{"\n"}}

{{- if .Conditions -}}
{{- "\t" }}Condition:
{{ range .Conditions }} 
{{- "\t" }}{{"\t"}}{{- .Result }}  {{ .Name -}}({{- .Kind -}}){{"\n"}}
{{- end -}}
{{- end -}}

{{- "\t" -}}Target:
{{ range .Targets }} 
{{- "\t" }}{{"\t"}}{{- .Result }}  {{ .Name -}}({{- .Kind -}}){{"\n"}}
{{- end }}
{{ end }}
{{ end }}
`

// Reports contains a list of report
type Reports []Report

// Show return a small reports of what has been changed
func (r *Reports) Show() error {
	t := template.Must(template.New("reports").Parse(reportsTpl))

	reports := ""

	buffer := new(bytes.Buffer)

	err := t.Execute(buffer, r)

	reports = buffer.String()

	if err != nil {
		return err
	}

	fmt.Println(reports)

	return nil
}

// Summary display a summary of
func (r *Reports) Summary() (int, int, int, error) {
	counter := 0
	successCounter := 0
	changedCounter := 0
	failedCounter := 0

	reports := *r

	for _, report := range reports {
		counter++
		if report.Result == result.SUCCESS {
			successCounter++
		} else if report.Result == result.FAILURE {
			failedCounter++
		} else if report.Result == result.CHANGED {
			changedCounter++
		} else {
			fmt.Printf("Unknown report result '%s'\n", report.Result)
		}
	}

	fmt.Printf("Run Summary\n")
	fmt.Printf("===========\n\n")
	fmt.Printf("%d job run\n", counter)
	fmt.Printf("%d job succeed\n", successCounter)
	fmt.Printf("%d job failed\n", failedCounter)
	fmt.Printf("%d job applied changes\n", changedCounter)

	if failedCounter > 0 {
		return successCounter, changedCounter, failedCounter, errors.New("a job has failed")
	}

	return successCounter, changedCounter, failedCounter, nil
}
