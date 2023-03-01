module CloudDevKubernetes

go 1.18

//debug
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.1.0

require (
	github.com/liushuochen/gotable v0.0.0-20210703140901-b0faa25d33c8
	github.com/pkg/profile v1.7.0
	github.com/spf13/cobra v1.6.1
	k8s.io/api v0.24.2
	k8s.io/apimachinery v0.24.2
	k8s.io/client-go v0.24.2
	k8s.io/kubectl v0.24.3
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/emicklei/go-restful/v3 v3.9.0 // indirect
	github.com/felixge/fgprof v0.9.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20211214055906-6f57359322fd // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/onsi/ginkgo/v2 v2.4.0 // indirect
	github.com/onsi/gomega v1.23.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gin-gonic/gin v1.8.1
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/imdario/mergo v0.3.6 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	golang.org/x/net v0.3.1-0.20221206200815-1e63c2f08a10 // indirect
	golang.org/x/oauth2 v0.0.0-20220223155221-ee480838109b // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/term v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/kube-openapi v0.0.0-20221012153701-172d655c2280 // indirect
	k8s.io/metrics v0.25.4
	k8s.io/utils v0.0.0-20221107191617-1a15be271d1d // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace k8s.io/api => k8s.io/api v0.24.2

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.24.2

replace k8s.io/apimachinery => k8s.io/apimachinery v0.24.3-rc.0

replace k8s.io/apiserver => k8s.io/apiserver v0.24.2

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.24.2

replace k8s.io/client-go => k8s.io/client-go v0.24.2

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.24.2

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.24.2

replace k8s.io/code-generator => k8s.io/code-generator v0.24.3-rc.0

replace k8s.io/component-base => k8s.io/component-base v0.24.2

replace k8s.io/component-helpers => k8s.io/component-helpers v0.24.2

replace k8s.io/controller-manager => k8s.io/controller-manager v0.24.2

replace k8s.io/cri-api => k8s.io/cri-api v0.25.0-alpha.0

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.24.2

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.24.2

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.24.2

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.24.2

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.24.2

replace k8s.io/kubectl => k8s.io/kubectl v0.24.2

replace k8s.io/kubelet => k8s.io/kubelet v0.24.2

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.24.2

replace k8s.io/metrics => k8s.io/metrics v0.24.2

replace k8s.io/mount-utils => k8s.io/mount-utils v0.24.3-rc.0

replace k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.24.2

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.24.2

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.24.2

replace k8s.io/sample-controller => k8s.io/sample-controller v0.24.2
