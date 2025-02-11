/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and cf-service-operator contributors
SPDX-License-Identifier: Apache-2.0
*/

package v1alpha1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var servicebindinglog = logf.Log.WithName("servicebinding-resource")

func (r *ServiceBinding) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:path=/mutate-cf-cs-sap-com-v1alpha1-servicebinding,mutating=true,failurePolicy=fail,sideEffects=None,groups=cf.cs.sap.com,resources=servicebindings,verbs=create;update,versions=v1alpha1,name=mservicebinding.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &ServiceBinding{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *ServiceBinding) Default() {
	servicebindinglog.Info("default", "name", r.Name)

	if r.Labels == nil {
		r.Labels = make(map[string]string)
	}
	r.Labels[LabelKeyServiceInstance] = r.Spec.ServiceInstanceName

	if r.Spec.Name == "" {
		r.Spec.Name = r.Name
	}
	if r.Spec.SecretName == "" {
		r.Spec.SecretName = r.Name
	}
}

// +kubebuilder:webhook:path=/validate-cf-cs-sap-com-v1alpha1-servicebinding,mutating=false,failurePolicy=fail,sideEffects=None,groups=cf.cs.sap.com,resources=servicebindings,verbs=create;update,versions=v1alpha1,name=vservicebinding.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &ServiceBinding{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ServiceBinding) ValidateCreate() error {
	servicebindinglog.Info("validate create", "name", r.Name)

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ServiceBinding) ValidateUpdate(old runtime.Object) error {
	servicebindinglog.Info("validate update", "name", r.Name)
	s := old.(*ServiceBinding)
	// Call the defaulting webhook logic for the old object (because defaulting through the webhook might be incomplete in case of generateName usage)
	s.Name = r.Name
	s.Default()

	// TODO: why not to allow name updates ?
	if r.Spec.Name != s.Spec.Name {
		return fmt.Errorf("spec.name is immutable")
	}

	if r.Spec.ServiceInstanceName != s.Spec.ServiceInstanceName {
		return fmt.Errorf("spec.serviceInstanceName is immutable")
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ServiceBinding) ValidateDelete() error {
	servicebindinglog.Info("validate delete", "name", r.Name)

	return nil
}
