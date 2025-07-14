/*
Copyright 2025.

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

// Package v1 contains API Schema definitions for the demo v1 API group.
// +kubebuilder:object:generate=true
// +groupName=demo.jeff
package v1

import (
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects.
	// group.domain   kubebuilder init --domain jeff
	GroupVersion = schema.GroupVersion{Group: "demo.jeff", Version: "v1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	// 被app_types.go line
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
/*
scheme:
在 Kubebuilder 和 controller-runtime 项目中，schema（Scheme）是 Kubernetes 中用于资源类型注册和识别的核心机制。它是 API 类型和其对应的 GroupVersionKind（GVK） 之间的映射表
一个注册表，记录了哪些 Go 结构体（如 App）对应于哪些 apiVersion 和 kind（GVK）。

不把 App 注册到 Scheme 中，controller-runtime 就无法反序列化这个资源，也就无法监听它或进行 Reconcile

在 main.go 中：
utilruntime.Must(demov1.AddToScheme(scheme))
就完成注册了。
*/ 
