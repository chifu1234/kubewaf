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
	"github.com/buzz-it/kubewaf/api/seclang/v1beta1/convert"
	"github.com/buzz-it/kubewaf/internal/controller"
	types "github.com/coreruleset/crslang/types"
)

// SecRuleReconciler reconciles a SecRule object
type SecRuleReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=seclan.kubewaf.io,resources=secrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=seclan.kubewaf.io,resources=secrules/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=seclan.kubewaf.io,resources=secrules/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SecRule object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.22.4/pkg/reconcile
func (r *SecRuleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := logf.FromContext(ctx)
	var (
		secRule = &seclangv1beta1.SecRule{}
		updated bool
	)

	updated, err := controller.InitHandler(ctx, req, secRule, r.Client)
	if err != nil {
		return ctrl.Result{}, err
	}

	crslangSecRule, err := convert.ConvertSecRule(*secRule)
	if err != nil {
		return ctrl.Result{}, err
	}
	currentSecRuleString := secRule.Status.SecRuleString
	secRuleString := convertToSecLangString(crslangSecRule)

	secRule.Status.SecRuleString = secRuleString
	if updated || secRule.Status.SecRuleString != currentSecRuleString {
		l.Info("Updated SecRule status with generated SecLang string")
		if err := r.Client.Status().Update(ctx, secRule); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SecRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&seclangv1beta1.SecRule{}).
		Named("secrule").
		Complete(r)
}

// func secRuleSpecToSecRule(secRuleSpec seclangv1beta1.SecRuleSpec) (crslang_types.SecRule, error) {
// 	data, err := json.Marshal(secRuleSpec)
// 	if err != nil {
// 		return crslang_types.SecRule{}, err
// 	}
// 	fmt.Println(fmt.Println(string(data)))
// 	var rule crslang_types.SecRule
// 	if err := json.Unmarshal(data, &rule); err != nil {
// 		return crslang_types.SecRule{}, fmt.Errorf("Could not Unmarshal Object into crs Object=%s", err.Error())
// 	}
// 	ruleData, err := json.Marshal(rule)
// 	if err != nil {
// 		return crslang_types.SecRule{}, err
// 	}
// 	fmt.Println(string(ruleData))

// 	return rule, nil
// }

//	func copySecRuleSpecToCRS(secRule *seclangv1beta1.SecRule) crslang_types.SecRule {
//		return crslang_types.SecRule{}
//	}
func convertToSecLangString(rules []types.SeclangDirective) string {
	configList := types.ConfigurationList{
		DirectiveList: []types.DirectiveList{{
			Directives: rules,
		}},
	}

	unformatted := types.FromCRSLangToUnformattedDirectives(configList)
	return unformatted.DirectiveList[0].ToSeclang()
}
