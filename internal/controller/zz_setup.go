/*
Copyright 2021 The Crossplane Authors.

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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	config "github.com/crossplane-contrib/provider-jet-xo/internal/controller/cloud/config"
	providerconfig "github.com/crossplane-contrib/provider-jet-xo/internal/controller/providerconfig"
	set "github.com/crossplane-contrib/provider-jet-xo/internal/controller/resource/set"
	acl "github.com/crossplane-contrib/provider-jet-xo/internal/controller/xenorchestra/acl"
	vm "github.com/crossplane-contrib/provider-jet-xo/internal/controller/xenorchestra/vm"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		config.Setup,
		providerconfig.Setup,
		set.Setup,
		acl.Setup,
		vm.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
