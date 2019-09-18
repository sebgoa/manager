package event

import "testing"

func newEvent() string {
	return "StackId='arn:aws:cloudformation:us-west-2:915347744415:stack/test-ror/c769c5e0-8c24-11e9-8e2e-02cb67b6aa16'\nTimestamp='2019-06-11T08:44:19.385Z'\nEventId='WebServerSecurityGroup-CREATE_COMPLETE-2019-06-11T08:44:19.385Z'\nLogicalResourceId='WebServerSecurityGroup'\nNamespace='915347744415'\nPhysicalResourceId='sg-0f3131359f75d68b6'\nResourceProperties='{\"GroupDescription\":\"Enable HTTP access locked down to the load balancer + SSH access\",\"VpcId\":\"vpc-244f4e43\",\"SecurityGroupIngress\":[{\"FromPort\":\"80\",\"ToPort\":\"80\",\"IpProtocol\":\"tcp\",\"SourceSecurityGroupId\":\"sg-b80c64c3\"},{\"CidrIp\":\"0.0.0.0/0\",\"FromPort\":\"22\",\"ToPort\":\"22\",\"IpProtocol\":\"tcp\"}]}'\nResourceStatus='CREATE_COMPLETE'\nResourceStatusReason=''\nResourceType='AWS::EC2::SecurityGroup'\nStackName='test-ror'\nClientRequestToken='Console-CreateStack-1b1012af-5ffe-7195-4dc8-9e5709e230b7'\n"
}

func TestUnmarshal(t *testing.T) {
	event := &Event{}

	err := Unmarshal(newEvent(), event)
	if err != nil {
		t.Errorf("Error unmarshaling event '%s'", err.Error())
	}

	if event.StackId != "arn:aws:cloudformation:us-west-2:915347744415:stack/test-ror/c769c5e0-8c24-11e9-8e2e-02cb67b6aa16" {
		t.Errorf("Stack ID does not equal expected '%s'", event.StackId)
	}

	if event.ResourceStatus != "CREATE_COMPLETE" {
		t.Errorf("Resource Status does not equal expected '%s'", event.ResourceStatus)
	}
}