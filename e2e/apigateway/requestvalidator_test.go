/*
Copyright © 2019 AWS Controller authors

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

package e2e_test

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	apigatewayv1alpha1 "go.awsctrl.io/manager/apis/apigateway/v1alpha1"
	cloudformationv1alpha1 "go.awsctrl.io/manager/apis/cloudformation/v1alpha1"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
)

// RunRequestValidatorSpecs allows all instance E2E tests to run
var _ = Describe("Run Apigateway RequestValidator Controller", func() {

	Context("Without RequestValidator{} existing", func() {

		It("Should create apigateway.RequestValidator{}", func() {
			var stackID string
			var stackName string
			var stack *cloudformationv1alpha1.Stack
			k8sclient := k8smanager.GetClient()
			Expect(k8sclient).ToNot(BeNil())

			restapi := &apigatewayv1alpha1.RestApi{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: "sample-restapi-",
					Namespace:    podnamespace,
				},
				Spec: apigatewayv1alpha1.RestApiSpec{
					Name:        "awsctrl",
					Description: "AWS Controller API Gateway API description",
				},
			}
			By("Creating new Apigateway RestApi")
			Expect(k8sclient.Create(context.Background(), restapi)).Should(Succeed())

			restapikey := types.NamespacedName{
				Name:      restapi.GetName(),
				Namespace: podnamespace,
			}

			instance := &apigatewayv1alpha1.RequestValidator{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: "sample-apikey-",
					Namespace:    podnamespace,
				},
				Spec: apigatewayv1alpha1.RequestValidatorSpec{
					Name:                      "test-validator",
					ValidateRequestBody:       true,
					ValidateRequestParameters: true,
					RestApi: metav1alpha1.ObjectReference{
						ObjectRef: metav1alpha1.ObjectRef{
							Kind:       "RestApi",
							APIVersion: "apigateway.awsctrl.io/v1alpha1",
							Name:       restapikey.Name,
							Key:        "ResourceRef",
						},
					},
				},
			}
			By("Creating new Apigateway RequestValidator")
			Expect(k8sclient.Create(context.Background(), instance)).Should(Succeed())

			key := types.NamespacedName{
				Name:      instance.GetName(),
				Namespace: podnamespace,
			}

			By("Expecting CreateComplete")
			Eventually(func() bool {
				By("Getting latest Apigateway RequestValidator")
				instance = &apigatewayv1alpha1.RequestValidator{}
				err := k8sclient.Get(context.Background(), key, instance)
				if err != nil {
					return false
				}

				stackID = instance.GetStackID()
				stackName = instance.GetStackName()

				return instance.Status.Status == metav1alpha1.CreateCompleteStatus ||
					(os.Getenv("USE_AWS_CLIENT") != "true" && instance.Status.Status != "")
			}, timeout, interval).Should(BeTrue())

			By("Checking object OwnerShip")
			Eventually(func() bool {
				stackkey := types.NamespacedName{
					Name:      stackName,
					Namespace: key.Namespace,
				}

				stack = &cloudformationv1alpha1.Stack{}
				err := k8sclient.Get(context.Background(), stackkey, stack)
				if err != nil {
					return false
				}

				expectedOwnerReference := v1.OwnerReference{
					Kind:       instance.Kind,
					APIVersion: instance.APIVersion,
					UID:        instance.UID,
					Name:       instance.Name,
				}

				ownerrefs := stack.GetOwnerReferences()
				Expect(len(ownerrefs)).To(Equal(1))

				return ownerrefs[0].Name == expectedOwnerReference.Name
			}, timeout, interval).Should(BeTrue())

			By("Deleting Apigateway RequestValidator")
			Expect(k8sclient.Delete(context.Background(), instance)).Should(Succeed())

			By("Deleting Apigateway RestApi")
			Expect(k8sclient.Delete(context.Background(), restapi)).Should(Succeed())

			By("Deleting RequestValidator Stack")
			Expect(k8sclient.Delete(context.Background(), stack)).Should(Succeed())

			By("Expecting metav1alpha1.DeleteCompleteStatus")
			Eventually(func() bool {
				if os.Getenv("USE_AWS_CLIENT") != "true" {
					return true
				}

				output, err := awsclient.GetClient("us-west-2").DescribeStacks(&cloudformation.DescribeStacksInput{StackName: aws.String(stackID)})
				Expect(err).To(BeNil())
				stackoutput := output.Stacks[0].StackStatus
				return *stackoutput == "DELETE_COMPLETE"
			}, timeout, interval).Should(BeTrue())
		})
	})
})
