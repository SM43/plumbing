package validator

import (
	"strings"
	"testing"

	"github.com/tektoncd/plumbing/catlin/pkg/parser"
	"gotest.tools/v3/assert"
)

const taskWithInvalidImageRef = `
---
apiVersion: tekton.dev/v1beta1
kind: Task
metadata:
  name: invalid-image-refs
  labels:
    app.kubernetes.io/version: a,b,c
  annotations:
    tekton.dev/tags: a,b,c
    tekton.dev/pipelines.minVersion: "0.12"
    tekton.dev/displayName: My Example Task
spec:
  description: |-
    A summary of the resource
  
    A para about this valid task
  
  steps:
  - name: hello
    image: ubuntu
    command: [sleep, infinity]
  - name: foo
    image: abc.io/fedora:latest
  - name: bar
    image: abc.io/fedora:1.0@sha256:deadb33fdeadb33fdeadb33f
  - name: valid
    image: abc.io/ubuntu:1.0
`

func TestTaskValidator_InvalidImageRef(t *testing.T) {

	r := strings.NewReader(taskWithInvalidImageRef)
	parser := parser.ForReader(r)

	res, err := parser.Parse()
	assert.NilError(t, err)

	v := ForKind(res)
	result := v.Validate()

	assert.Equal(t, 3, result.Errors)
	assert.Equal(t, 3, len(result.Lints))
	assert.Equal(t, `Step "hello" has an invalid image: could not parse reference: ubuntu`, result.Lints[0].Message)
	assert.Equal(t, `Step "foo" image (abc.io/fedora:latest) must be tagged with a specific version`, result.Lints[1].Message)
	assert.Equal(t, `Step "bar" has an invalid image: could not parse reference: abc.io/fedora:1.0@sha256:deadb33fdeadb33fdeadb33f`, result.Lints[2].Message)
}
