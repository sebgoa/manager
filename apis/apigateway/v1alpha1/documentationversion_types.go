/*
Copyright © 2019 AWS Controller authors

Licensed under the Apache License, Version 2.0 (the &#34;License&#34;);
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an &#34;AS IS&#34; BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DocumentationVersionSpec defines the desired state of DocumentationVersion
type DocumentationVersionSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// RestApi http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationversion.html#cfn-apigateway-documentationversion-restapiid
	RestApi metav1alpha1.ObjectReference `json:"restApi" cloudformation:"RestApiId,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationversion.html#cfn-apigateway-documentationversion-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// DocumentationVersion http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationversion.html#cfn-apigateway-documentationversion-documentationversion
	DocumentationVersion string `json:"documentationVersion" cloudformation:"DocumentationVersion,Parameter"`
}

// DocumentationVersionStatus defines the observed state of DocumentationVersion
type DocumentationVersionStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// DocumentationVersionOutput defines the stack outputs
type DocumentationVersionOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationversion.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;apigateway
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// DocumentationVersion is the Schema for the apigateway DocumentationVersion API
type DocumentationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DocumentationVersionSpec   `json:"spec,omitempty"`
	Status DocumentationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentationVersionList contains a list of Account
type DocumentationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []DocumentationVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DocumentationVersion{}, &DocumentationVersionList{})
}
