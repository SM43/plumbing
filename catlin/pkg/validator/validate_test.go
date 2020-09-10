// Copyright © 2020 The Tekton Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/tektoncd/plumbing/catlin/pkg/parser"
	"github.com/tektoncd/plumbing/catlin/pkg/test"
)

func TestContentValidator_Task(t *testing.T) {
	tc := test.New()

	r := strings.NewReader(validTask)
	parser := parser.ForReader(r)

	res, err := parser.Parse()
	assert.NilError(t, err)

	v := NewContentValidator(tc, res)
	result := v.Validate()

	assert.Equal(t, 0, result.Errors)
	assert.Equal(t, 0, len(result.Lints))
}

func TestContentValidator_Pipeline(t *testing.T) {
	tc := test.New()

	r := strings.NewReader(validPipeline)
	parser := parser.ForReader(r)

	res, err := parser.Parse()
	assert.NilError(t, err)

	v := NewContentValidator(tc, res)
	result := v.Validate()

	assert.Equal(t, 0, result.Errors)
	assert.Equal(t, 0, len(result.Lints))
}

func TestValidatorForKind_Task(t *testing.T) {

	r := strings.NewReader(validTask)
	parser := parser.ForReader(r)

	res, err := parser.Parse()
	assert.NilError(t, err)

	v := ForKind(res)
	result := v.Validate()

	assert.Equal(t, 0, result.Errors)
	assert.Equal(t, 0, len(result.Lints))
}

func TestValidatorForKind_Task_InvalidImageRef(t *testing.T) {

	r := strings.NewReader(taskWithInvalidImageRef)
	parser := parser.ForReader(r)

	res, err := parser.Parse()
	assert.NilError(t, err)

	v := ForKind(res)
	result := v.Validate()

	assert.Equal(t, 3, result.Errors)
	assert.Equal(t, 3, len(result.Lints))
	assert.Equal(t, "Invalid Image Reference: could not parse reference: ubuntu", result.Lints[0].Message)
	assert.Equal(t, "Task image (abc.io/fedora:latest) must be tagged with a specific version", result.Lints[1].Message)
	assert.Equal(t, "Invalid Image Reference: could not parse reference: abc.io/fedora:1.0@sha256:deadb33fdeadb33fdeadb33f", result.Lints[2].Message)
}
