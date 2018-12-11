// Copyright 2016 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

import (
	"fmt"
)

func someProcessingFunction(val int) error {
	if val == 1 {
		// New and Reason automatically include stack information.
		return Reason("bad number: %d", val).Err()
	}
	if err := someProcessingFunction(val - 1); err != nil {
		// correctly handles recursion
		return Annotate(err, "").InternalReason("val(%d)", val).Err()
	}
	return nil
}

func someLibFunc(vals ...int) error {
	for i, v := range vals {
		if err := someProcessingFunction(v); err != nil {
			return Annotate(err, "processing %d", v).
				InternalReason("secret(%s)/i(%d)", "value", i).Err()
		}
	}
	return nil
}

type MiscWrappedError struct{ error }

func (e *MiscWrappedError) Error() string     { return fmt.Sprintf("super wrapper(%s)", e.error.Error()) }
func (e *MiscWrappedError) InnerError() error { return e.error }

func errorWrapper(err error) error {
	if err != nil {
		err = &MiscWrappedError{err}
	}
	return err
}

func someIntermediateFunc(vals ...int) error {
	errch := make(chan error)
	go func() {
		defer close(errch)
		errch <- Annotate(errorWrapper(someLibFunc(vals...)), "could not process").Err()
	}()
	me := MultiError(nil)
	for err := range errch {
		if err != nil {
			me = append(me, err)
		}
	}
	if me != nil {
		return Annotate(me, "while processing %v", vals).Err()
	}
	return nil
}

func ExampleAnnotate() {
	if err := someIntermediateFunc(3); err != nil {
		err = Annotate(err, "top level").Err()
		fmt.Println("Public-facing error:\n ", err)
		fmt.Println("\nfull error:")
		for _, l := range FixForTest(RenderStack(err, "runtime", "_test")) {
			fmt.Println(l)
		}
	}

	// Output:
	// Public-facing error:
	//   top level: while processing [3]: could not process: super wrapper(processing 3: bad number: 1)
	//
	// full error:
	// GOROUTINE LINE
	// #? github.com/TriggerMail/luci-go/common/errors/annotate_example_test.go:24 - errors.someProcessingFunction()
	//   reason: bad number: 1
	//
	// #? github.com/TriggerMail/luci-go/common/errors/annotate_example_test.go:26 - errors.someProcessingFunction()
	//   internal reason: val(2)
	//
	// #? github.com/TriggerMail/luci-go/common/errors/annotate_example_test.go:26 - errors.someProcessingFunction()
	//   internal reason: val(3)
	//
	// From frame 2 to 3, the following wrappers were found:
	//   unknown wrapper *errors.MiscWrappedError
	//
	// #? github.com/TriggerMail/luci-go/common/errors/annotate_example_test.go:35 - errors.someLibFunc()
	//   reason: processing 3
	//   internal reason: secret(value)/i(0)
	//
	// From frame 3 to 4, the following wrappers were found:
	//   internal reason: MultiError 1/1: following first non-nil error.
	//
	// #? github.com/TriggerMail/luci-go/common/errors/annotate_example_test.go:59 - errors.someIntermediateFunc.func1()
	//   reason: could not process
	//
	// ... skipped SOME frames in pkg "runtime"...
	//
	// GOROUTINE LINE
	// #? github.com/TriggerMail/luci-go/common/errors/annotate_example_test.go:68 - errors.someIntermediateFunc()
	//   reason: while processing [3]
	//
	// #? github.com/TriggerMail/luci-go/common/errors/annotate_example_test.go:74 - errors.ExampleAnnotate()
	//   reason: top level
	//
	// #? testing/example.go:XXX - testing.runExample()
	// #? testing/example.go:XXX - testing.runExamples()
	// #? testing/testing.go:XXX - testing.(*M).Run()
	// #? ./_testmain.go:XXX - main.main()
	// ... skipped SOME frames in pkg "runtime"...
}
