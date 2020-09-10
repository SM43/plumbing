package validator

var (
	validTask = `
---
apiVersion: tekton.dev/v1beta1
kind: Task
metadata:
  name: valid
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
    image: abc.io/ubuntu:1.0
    command: [sleep, infinity]
  - name: foo-bar
    image: abc.io/fedora:1.0@sha256:deadb33fdeadb33fdeadb33fdeadb33fdeadb33fdeadb33fdeadb33fdeadb33f
`

	taskWithInvalidImageRef = `
---
apiVersion: tekton.dev/v1beta1
kind: Task
metadata:
  name: valid
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

	validPipeline = `
---
apiVersion: tekton.dev/v1beta1
kind: Pipeline
metadata:
  name: valid
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

  tasks:
  - name: hello
    taskRef:
      name: hello
`
)
