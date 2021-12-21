// Copyright 2021 Google Inc. All rights reserved.
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

package bp2build

import (
	"testing"

	"android/soong/android"
	"android/soong/java"
)

func runJavaLibraryTestCase(t *testing.T, tc bp2buildTestCase) {
	t.Helper()
	(&tc).moduleTypeUnderTest = "java_library"
	(&tc).moduleTypeUnderTestFactory = java.LibraryFactory
	runBp2BuildTestCase(t, func(ctx android.RegistrationContext) {}, tc)
}

func TestJavaLibrary(t *testing.T) {
	runJavaLibraryTestCase(t, bp2buildTestCase{
		description: "java_library with srcs, exclude_srcs and libs",
		blueprint: `java_library {
    name: "java-lib-1",
    srcs: ["a.java", "b.java"],
    exclude_srcs: ["b.java"],
    libs: ["java-lib-2"],
    bazel_module: { bp2build_available: true },
}

java_library {
    name: "java-lib-2",
    srcs: ["b.java"],
    bazel_module: { bp2build_available: true },
}`,
		expectedBazelTargets: []string{
			makeBazelTarget("java_library", "java-lib-1", attrNameToString{
				"srcs": `["a.java"]`,
				"deps": `[":java-lib-2"]`,
			}),
			makeBazelTarget("java_library", "java-lib-2", attrNameToString{
				"srcs": `["b.java"]`,
			}),
		},
	})
}