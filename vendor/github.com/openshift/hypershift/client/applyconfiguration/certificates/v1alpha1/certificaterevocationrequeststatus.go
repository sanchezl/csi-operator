/*


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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// CertificateRevocationRequestStatusApplyConfiguration represents an declarative configuration of the CertificateRevocationRequestStatus type for use
// with apply.
type CertificateRevocationRequestStatusApplyConfiguration struct {
	RevocationTimestamp *v1.Time                             `json:"revocationTimestamp,omitempty"`
	PreviousSigner      *corev1.LocalObjectReference         `json:"previousSigner,omitempty"`
	Conditions          []metav1.ConditionApplyConfiguration `json:"conditions,omitempty"`
}

// CertificateRevocationRequestStatusApplyConfiguration constructs an declarative configuration of the CertificateRevocationRequestStatus type for use with
// apply.
func CertificateRevocationRequestStatus() *CertificateRevocationRequestStatusApplyConfiguration {
	return &CertificateRevocationRequestStatusApplyConfiguration{}
}

// WithRevocationTimestamp sets the RevocationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RevocationTimestamp field is set to the value of the last call.
func (b *CertificateRevocationRequestStatusApplyConfiguration) WithRevocationTimestamp(value v1.Time) *CertificateRevocationRequestStatusApplyConfiguration {
	b.RevocationTimestamp = &value
	return b
}

// WithPreviousSigner sets the PreviousSigner field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreviousSigner field is set to the value of the last call.
func (b *CertificateRevocationRequestStatusApplyConfiguration) WithPreviousSigner(value corev1.LocalObjectReference) *CertificateRevocationRequestStatusApplyConfiguration {
	b.PreviousSigner = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *CertificateRevocationRequestStatusApplyConfiguration) WithConditions(values ...*metav1.ConditionApplyConfiguration) *CertificateRevocationRequestStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}