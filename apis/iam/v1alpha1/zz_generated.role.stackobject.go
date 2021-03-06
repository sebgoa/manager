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
	"encoding/json"
	"fmt"
	"strings"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	controllerutils "go.awsctrl.io/manager/controllers/utils"
	cfnencoder "go.awsctrl.io/manager/encoding/cloudformation"

	"github.com/awslabs/goformation/v4/cloudformation"
	"github.com/awslabs/goformation/v4/cloudformation/iam"
	"k8s.io/client-go/dynamic"
)

// GetNotificationARNs is an autogenerated deepcopy function, will return notifications for stack
func (in *Role) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *Role) GetTemplate(client dynamic.Interface) (string, error) {
	if client == nil {
		return "", fmt.Errorf("k8s client not loaded for template")
	}
	template := cloudformation.NewTemplate()

	template.Description = "AWS Controller - iam.Role (ac-{TODO})"

	template.Outputs = map[string]interface{}{
		"ResourceRef": map[string]interface{}{
			"Value": cloudformation.Ref("Role"),
		},
		"RoleId": map[string]interface{}{
			"Value": cloudformation.GetAtt("Role", "RoleId"),
		},
		"Arn": map[string]interface{}{
			"Value": cloudformation.GetAtt("Role", "Arn"),
		},
	}

	iamRole := &iam.Role{}

	if in.Spec.PermissionsBoundary != "" {
		iamRole.PermissionsBoundary = in.Spec.PermissionsBoundary
	}

	if len(in.Spec.ManagedPolicyArns) > 0 {
		iamRole.ManagedPolicyArns = in.Spec.ManagedPolicyArns
	}

	if in.Spec.Path != "" {
		iamRole.Path = in.Spec.Path
	}

	if in.Spec.AssumeRolePolicyDocument != "" {
		iamRoleJSON := make(map[string]interface{})
		err := json.Unmarshal([]byte(in.Spec.AssumeRolePolicyDocument), &iamRoleJSON)
		if err != nil {
			return "", err
		}
		iamRole.AssumeRolePolicyDocument = iamRoleJSON
	}

	if in.Spec.Description != "" {
		iamRole.Description = in.Spec.Description
	}

	if in.Spec.RoleName != "" {
		iamRole.RoleName = in.Spec.RoleName
	}

	// TODO(christopherhein): implement tags this could be easy now that I have the mechanims of nested objects
	iamRolePolicies := []iam.Role_Policy{}

	for _, item := range in.Spec.Policies {
		iamRolePolicy := iam.Role_Policy{}

		if item.PolicyDocument != "" {
			iamRolePolicyJSON := make(map[string]interface{})
			err := json.Unmarshal([]byte(item.PolicyDocument), &iamRolePolicyJSON)
			if err != nil {
				return "", err
			}
			iamRolePolicy.PolicyDocument = iamRolePolicyJSON
		}

		if item.PolicyName != "" {
			iamRolePolicy.PolicyName = item.PolicyName
		}

	}

	if len(iamRolePolicies) > 0 {
		iamRole.Policies = iamRolePolicies
	}
	if in.Spec.MaxSessionDuration != iamRole.MaxSessionDuration {
		iamRole.MaxSessionDuration = in.Spec.MaxSessionDuration
	}

	template.Resources = map[string]cloudformation.Resource{
		"Role": iamRole,
	}

	// json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	json, err := template.JSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *Role) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *Role) GenerateStackName() string {
	return strings.Join([]string{"iam", "role", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *Role) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *Role) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *Role) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *Role) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *Role) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *Role) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *Role) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *Role) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *Role) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *Role) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}
