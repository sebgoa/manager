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
	"fmt"
	"reflect"
	"strings"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	controllerutils "go.awsctrl.io/manager/controllers/utils"
	cfnencoder "go.awsctrl.io/manager/encoding/cloudformation"

	"github.com/awslabs/goformation/v4/cloudformation"
	"github.com/awslabs/goformation/v4/cloudformation/route53"
	"k8s.io/client-go/dynamic"
)

// GetNotificationARNs is an autogenerated deepcopy function, will return notifications for stack
func (in *RecordSetGroup) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *RecordSetGroup) GetTemplate(client dynamic.Interface) (string, error) {
	if client == nil {
		return "", fmt.Errorf("k8s client not loaded for template")
	}
	template := cloudformation.NewTemplate()

	template.Description = "AWS Controller - route53.RecordSetGroup (ac-{TODO})"

	template.Outputs = map[string]interface{}{
		"ResourceRef": map[string]interface{}{
			"Value": cloudformation.Ref("RecordSetGroup"),
		},
	}

	route53RecordSetGroup := &route53.RecordSetGroup{}

	if in.Spec.Comment != "" {
		route53RecordSetGroup.Comment = in.Spec.Comment
	}

	// TODO(christopherhein) move these to a defaulter
	route53RecordSetGroupHostedZoneItem := in.Spec.HostedZone.DeepCopy()

	if route53RecordSetGroupHostedZoneItem.ObjectRef.Kind == "" {
		route53RecordSetGroupHostedZoneItem.ObjectRef.Kind = "Deployment"
	}

	if route53RecordSetGroupHostedZoneItem.ObjectRef.APIVersion == "" {
		route53RecordSetGroupHostedZoneItem.ObjectRef.APIVersion = "apigateway.awsctrl.io/v1alpha1"
	}

	if route53RecordSetGroupHostedZoneItem.ObjectRef.Namespace == "" {
		route53RecordSetGroupHostedZoneItem.ObjectRef.Namespace = in.Namespace
	}

	in.Spec.HostedZone = *route53RecordSetGroupHostedZoneItem
	hostedZoneId, err := in.Spec.HostedZone.String(client)
	if err != nil {
		return "", err
	}

	if hostedZoneId != "" {
		route53RecordSetGroup.HostedZoneId = hostedZoneId
	}

	if in.Spec.HostedZoneName != "" {
		route53RecordSetGroup.HostedZoneName = in.Spec.HostedZoneName
	}

	route53RecordSetGroupRecordSets := []route53.RecordSetGroup_RecordSet{}

	for _, item := range in.Spec.RecordSets {
		route53RecordSetGroupRecordSet := route53.RecordSetGroup_RecordSet{}

		// TODO(christopherhein) move these to a defaulter
		route53RecordSetGroupRecordSetHostedZoneItem := item.HostedZone.DeepCopy()

		if route53RecordSetGroupRecordSetHostedZoneItem.ObjectRef.Kind == "" {
			route53RecordSetGroupRecordSetHostedZoneItem.ObjectRef.Kind = "Deployment"
		}

		if route53RecordSetGroupRecordSetHostedZoneItem.ObjectRef.APIVersion == "" {
			route53RecordSetGroupRecordSetHostedZoneItem.ObjectRef.APIVersion = "apigateway.awsctrl.io/v1alpha1"
		}

		if route53RecordSetGroupRecordSetHostedZoneItem.ObjectRef.Namespace == "" {
			route53RecordSetGroupRecordSetHostedZoneItem.ObjectRef.Namespace = in.Namespace
		}

		item.HostedZone = *route53RecordSetGroupRecordSetHostedZoneItem
		hostedZoneId, err := item.HostedZone.String(client)
		if err != nil {
			return "", err
		}

		if hostedZoneId != "" {
			route53RecordSetGroupRecordSet.HostedZoneId = hostedZoneId
		}

		if !reflect.DeepEqual(item.AliasTarget, RecordSetGroup_AliasTarget{}) {
			route53RecordSetGroupRecordSetAliasTarget := route53.RecordSetGroup_AliasTarget{}

			if item.AliasTarget.DNSName != "" {
				route53RecordSetGroupRecordSetAliasTarget.DNSName = item.AliasTarget.DNSName
			}

			if item.AliasTarget.EvaluateTargetHealth || !item.AliasTarget.EvaluateTargetHealth {
				route53RecordSetGroupRecordSetAliasTarget.EvaluateTargetHealth = item.AliasTarget.EvaluateTargetHealth
			}

			// TODO(christopherhein) move these to a defaulter
			route53RecordSetGroupRecordSetAliasTargetHostedZoneItem := item.AliasTarget.HostedZone.DeepCopy()

			if route53RecordSetGroupRecordSetAliasTargetHostedZoneItem.ObjectRef.Kind == "" {
				route53RecordSetGroupRecordSetAliasTargetHostedZoneItem.ObjectRef.Kind = "Deployment"
			}

			if route53RecordSetGroupRecordSetAliasTargetHostedZoneItem.ObjectRef.APIVersion == "" {
				route53RecordSetGroupRecordSetAliasTargetHostedZoneItem.ObjectRef.APIVersion = "apigateway.awsctrl.io/v1alpha1"
			}

			if route53RecordSetGroupRecordSetAliasTargetHostedZoneItem.ObjectRef.Namespace == "" {
				route53RecordSetGroupRecordSetAliasTargetHostedZoneItem.ObjectRef.Namespace = in.Namespace
			}

			item.AliasTarget.HostedZone = *route53RecordSetGroupRecordSetAliasTargetHostedZoneItem
			hostedZoneId, err := item.AliasTarget.HostedZone.String(client)
			if err != nil {
				return "", err
			}

			if hostedZoneId != "" {
				route53RecordSetGroupRecordSetAliasTarget.HostedZoneId = hostedZoneId
			}

			route53RecordSetGroupRecordSet.AliasTarget = &route53RecordSetGroupRecordSetAliasTarget
		}

		if item.Failover != "" {
			route53RecordSetGroupRecordSet.Failover = item.Failover
		}

		if len(item.ResourceRecords) > 0 {
			route53RecordSetGroupRecordSet.ResourceRecords = item.ResourceRecords
		}

		if item.Region != "" {
			route53RecordSetGroupRecordSet.Region = item.Region
		}

		if item.HostedZoneName != "" {
			route53RecordSetGroupRecordSet.HostedZoneName = item.HostedZoneName
		}

		if item.MultiValueAnswer || !item.MultiValueAnswer {
			route53RecordSetGroupRecordSet.MultiValueAnswer = item.MultiValueAnswer
		}

		if item.Name != "" {
			route53RecordSetGroupRecordSet.Name = item.Name
		}

		if item.TTL != "" {
			route53RecordSetGroupRecordSet.TTL = item.TTL
		}

		if !reflect.DeepEqual(item.GeoLocation, RecordSetGroup_GeoLocation{}) {
			route53RecordSetGroupRecordSetGeoLocation := route53.RecordSetGroup_GeoLocation{}

			if item.GeoLocation.ContinentCode != "" {
				route53RecordSetGroupRecordSetGeoLocation.ContinentCode = item.GeoLocation.ContinentCode
			}

			if item.GeoLocation.CountryCode != "" {
				route53RecordSetGroupRecordSetGeoLocation.CountryCode = item.GeoLocation.CountryCode
			}

			if item.GeoLocation.SubdivisionCode != "" {
				route53RecordSetGroupRecordSetGeoLocation.SubdivisionCode = item.GeoLocation.SubdivisionCode
			}

			route53RecordSetGroupRecordSet.GeoLocation = &route53RecordSetGroupRecordSetGeoLocation
		}

		if item.Comment != "" {
			route53RecordSetGroupRecordSet.Comment = item.Comment
		}

		if item.Type != "" {
			route53RecordSetGroupRecordSet.Type = item.Type
		}

		if item.SetIdentifier != "" {
			route53RecordSetGroupRecordSet.SetIdentifier = item.SetIdentifier
		}

		// TODO(christopherhein) move these to a defaulter
		route53RecordSetGroupRecordSetHealthCheckItem := item.HealthCheck.DeepCopy()

		if route53RecordSetGroupRecordSetHealthCheckItem.ObjectRef.Kind == "" {
			route53RecordSetGroupRecordSetHealthCheckItem.ObjectRef.Kind = "Deployment"
		}

		if route53RecordSetGroupRecordSetHealthCheckItem.ObjectRef.APIVersion == "" {
			route53RecordSetGroupRecordSetHealthCheckItem.ObjectRef.APIVersion = "apigateway.awsctrl.io/v1alpha1"
		}

		if route53RecordSetGroupRecordSetHealthCheckItem.ObjectRef.Namespace == "" {
			route53RecordSetGroupRecordSetHealthCheckItem.ObjectRef.Namespace = in.Namespace
		}

		item.HealthCheck = *route53RecordSetGroupRecordSetHealthCheckItem
		healthCheckId, err := item.HealthCheck.String(client)
		if err != nil {
			return "", err
		}

		if healthCheckId != "" {
			route53RecordSetGroupRecordSet.HealthCheckId = healthCheckId
		}

		if item.Weight != route53RecordSetGroupRecordSet.Weight {
			route53RecordSetGroupRecordSet.Weight = item.Weight
		}

	}

	if len(route53RecordSetGroupRecordSets) > 0 {
		route53RecordSetGroup.RecordSets = route53RecordSetGroupRecordSets
	}

	template.Resources = map[string]cloudformation.Resource{
		"RecordSetGroup": route53RecordSetGroup,
	}

	// json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	json, err := template.JSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *RecordSetGroup) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *RecordSetGroup) GenerateStackName() string {
	return strings.Join([]string{"route53", "recordsetgroup", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *RecordSetGroup) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *RecordSetGroup) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *RecordSetGroup) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *RecordSetGroup) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *RecordSetGroup) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *RecordSetGroup) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *RecordSetGroup) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *RecordSetGroup) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *RecordSetGroup) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *RecordSetGroup) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}
