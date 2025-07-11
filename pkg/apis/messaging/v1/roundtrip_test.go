/*
Copyright 2020 The Knative Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"testing"

	"k8s.io/apimachinery/pkg/runtime/serializer"
	"knative.dev/eventing/pkg/apis/messaging"
	"sigs.k8s.io/randfill"

	"k8s.io/apimachinery/pkg/api/apitesting/fuzzer"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	pkgfuzzer "knative.dev/pkg/apis/testing/fuzzer"
	"knative.dev/pkg/apis/testing/roundtrip"
)

// FuzzerFuncs includes fuzzing funcs for knative.dev/messaging v1 types
//
// For other examples see
// https://github.com/kubernetes/apimachinery/blob/master/pkg/apis/meta/fuzzer/fuzzer.go
var FuzzerFuncs = fuzzer.MergeFuzzerFuncs(
	func(codecs serializer.CodecFactory) []interface{} {
		return []interface{}{
			func(ch *Channel, c randfill.Continue) {
				c.FillNoCustom(ch) // fuzz the Channel
				if ch != nil {
					if ch.Annotations == nil {
						ch.Annotations = make(map[string]string)
					}
					ch.Annotations[messaging.SubscribableDuckVersionAnnotation] = "v1"
				}
				// Clear the random fuzzed condition
				ch.Status.SetConditions(nil)

				// Fuzz the known conditions except their type value
				ch.Status.InitializeConditions()
				pkgfuzzer.FuzzConditions(&ch.Status, c)
			},
			func(imc *InMemoryChannel, c randfill.Continue) {
				c.FillNoCustom(imc) // fuzz the InMemoryChannel
				if imc != nil {
					if imc.Annotations == nil {
						imc.Annotations = make(map[string]string)
					}
					imc.Annotations[messaging.SubscribableDuckVersionAnnotation] = "v1"
				}
				// Clear the random fuzzed condition
				imc.Status.SetConditions(nil)

				// Fuzz the known conditions except their type value
				imc.Status.InitializeConditions()
				pkgfuzzer.FuzzConditions(&imc.Status, c)
			},
			func(s *SubscriptionStatus, c randfill.Continue) {
				c.FillNoCustom(s) // fuzz the status object

				// Clear the random fuzzed condition
				s.Status.SetConditions(nil)

				// Fuzz the known conditions except their type value
				s.InitializeConditions()
				pkgfuzzer.FuzzConditions(&s.Status, c)
			},
		}
	},
)

func TestMessagingRoundTripTypesToJSON(t *testing.T) {
	scheme := runtime.NewScheme()
	utilruntime.Must(AddToScheme(scheme))

	fuzzerFuncs := fuzzer.MergeFuzzerFuncs(
		pkgfuzzer.Funcs,
		FuzzerFuncs,
	)
	roundtrip.ExternalTypesViaJSON(t, scheme, fuzzerFuncs)
}
