module go.awsctrl.io/manager

go 1.13

require (
	github.com/aws/aws-sdk-go v1.25.36
	github.com/awslabs/goformation/v4 v4.1.0
	github.com/davecgh/go-spew v1.1.1
	github.com/go-logr/logr v0.1.0
	github.com/gobuffalo/envy v1.8.1 // indirect
	github.com/gobuffalo/flect v0.2.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/huandu/xstrings v1.2.1 // indirect
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.8.1
	github.com/rogpeppe/go-internal v1.5.1 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v0.0.5
	go.awsctrl.io/generator v0.0.0-20200102111626-002d9016d6bc // indirect
	go.hein.dev/go-version v0.1.0
	golang.org/x/crypto v0.0.0-20191227163750-53104e6ec876 // indirect
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/sys v0.0.0-20191210023423-ac6580df4449 // indirect
	golang.org/x/tools v0.0.0-20200102165158-6ed6e6036ba4 // indirect
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0
	k8s.io/klog v1.0.0
	sigs.k8s.io/controller-runtime v0.4.0
)

replace github.com/appscode/jsonpatch => gomodules.xyz/jsonpatch v1.0.1
