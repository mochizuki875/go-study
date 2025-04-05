module go-study/k8s/path

go 1.23.0

toolchain go1.23.5

replace (
	k8s.io/api => k8s.io/api v0.32.1
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.32.1
	k8s.io/apiserver => k8s.io/apiserver v0.32.1
	k8s.io/component-base => k8s.io/component-base v0.32.1
	k8s.io/component-helpers => k8s.io/component-helpers v0.32.1
	k8s.io/controller-manager => k8s.io/controller-manager v0.32.1
)

require k8s.io/utils v0.0.0-20241104100929-3ea5e8cea738
