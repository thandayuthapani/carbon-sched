module admiralty.io/multicluster-scheduler

go 1.17

require (
	admiralty.io/multicluster-service-account v0.6.1
	github.com/go-test/deep v1.0.8
	github.com/hashicorp/go-multierror v1.1.1
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/virtual-kubelet/virtual-kubelet v1.6.0
	k8s.io/api v0.23.3
	k8s.io/apimachinery v0.23.3
	k8s.io/client-go v0.23.3
	k8s.io/code-generator v0.23.3
	k8s.io/component-base v0.23.3
	k8s.io/klog v1.0.0
	k8s.io/kubernetes v1.23.3
	k8s.io/sample-controller v0.23.3
	sigs.k8s.io/controller-runtime v0.11.1
	sigs.k8s.io/yaml v1.3.0
)

//replace google.golang.org/grpc => google.golang.org/grpc v1.29.1

replace k8s.io/api => k8s.io/api v0.23.3

replace k8s.io/apimachinery => k8s.io/apimachinery v0.23.3

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.23.3

replace k8s.io/apiserver => k8s.io/apiserver v0.23.3

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.23.3

replace k8s.io/client-go => k8s.io/client-go v0.23.3

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.23.3

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.23.3

replace k8s.io/code-generator => k8s.io/code-generator v0.23.3

replace k8s.io/component-base => k8s.io/component-base v0.23.3

replace k8s.io/cri-api => k8s.io/cri-api v0.23.3

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.23.3

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.23.3

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.23.3

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.23.3

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.23.3

replace k8s.io/kubectl => k8s.io/kubectl v0.23.3

replace k8s.io/kubelet => k8s.io/kubelet v0.23.3

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.23.3

replace k8s.io/metrics => k8s.io/metrics v0.23.3

replace k8s.io/node-api => k8s.io/node-api v0.23.3

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.23.3

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.23.3

replace k8s.io/sample-controller => k8s.io/sample-controller v0.23.3

replace k8s.io/component-helpers => k8s.io/component-helpers v0.23.3

replace k8s.io/controller-manager => k8s.io/controller-manager v0.23.3

replace k8s.io/mount-utils => k8s.io/mount-utils v0.23.3

replace k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.23.3
