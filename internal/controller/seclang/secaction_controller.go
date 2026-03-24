/*
Copyright 2025 Buzz-IT GmbH.

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

package seclang

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	seclangv1beta1 "github.com/buzz-it/kubewaf/api/seclang/v1beta1"
)

// SecActionReconciler reconciles a SecAction object
type SecActionReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=seclang.kubewaf.io,resources=secactions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=seclang.kubewaf.io,resources=secactions/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=seclang.kubewaf.io,resources=secactions/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SecAction object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.22.4/pkg/reconcile
func (r *SecActionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = logf.FromContext(ctx)

	// secAction := seclangv1beta1.SecAction{}
	// _, err := controller.InitHandler(ctx, req, &secAction, r.Client)
	// if err != nil {
	// 	return ctrl.Result{}, err
	// }

	// crslangSecAction, err := convert.ConvertSecRule(secAction.)
	// if err != nil {
	// 	return ctrl.Result{}, err
	// }

	// bla, err := yaml.Marshal(crslangSecAction)
	// if err != nil {
	// 	return ctrl.Result{}, err
	// }
	// fmt.Println(string(bla))

	// secAction.Status.SecRuleString = crslangSecAction.ToSeclang()

	// err = r.Client.Status().Update(ctx, &secAction)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SecActionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&seclangv1beta1.SecAction{}).
		Named("seclang-secaction").
		Complete(r)
}
