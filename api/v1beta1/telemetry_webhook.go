/*
Copyright 2022.
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

//
// Generated by:
//
// operator-sdk create webhook --group telemetry --version v1beta1 --kind Telemetry --programmatic-validation --defaulting
//

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// TelemetryDefaults -
type TelemetryDefaults struct {
	CentralContainerImageURL      string
	CentralInitContainerImageURL  string
	ComputeContainerImageURL      string
	ComputeInitContainerImageURL  string
	NotificationContainerImageURL string
	SgCoreContainerImageURL       string
	NodeExporterContainerImageURL string
	IpmiContainerImageURL         string
}

var telemetryDefaults TelemetryDefaults

// log is for logging in this package.
var telemetrylog = logf.Log.WithName("telemetry-resource")

// SetupTelemetryDefaults - initialize Telemetry spec defaults for use with either internal or external webhooks
func SetupTelemetryDefaults(defaults TelemetryDefaults) {
	telemetryDefaults = defaults
	telemetrylog.Info("Telemetry defaults initialized", "defaults", defaults)
}

// SetupWebhookWithManager sets up the webhook with the Manager
func (r *Telemetry) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-telemetry-openstack-org-v1beta1-telemetry,mutating=true,failurePolicy=fail,sideEffects=None,groups=telemetry.openstack.org,resources=telemetries,verbs=create;update,versions=v1beta1,name=mtelemetry.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Telemetry{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Telemetry) Default() {
	telemetrylog.Info("default", "name", r.Name)

	r.Spec.Default()
}

// Default - set defaults for this Telemetry spec
func (spec *TelemetrySpec) Default() {
	if spec.CeilometerCentral.CentralImage == "" {
		spec.CeilometerCentral.CentralImage = telemetryDefaults.CentralContainerImageURL
	}
	if spec.CeilometerCentral.InitImage == "" {
		spec.CeilometerCentral.InitImage = telemetryDefaults.CentralInitContainerImageURL
	}
	if spec.CeilometerCompute.ComputeImage == "" {
		spec.CeilometerCompute.ComputeImage = telemetryDefaults.ComputeContainerImageURL
	}
	if spec.CeilometerCompute.IpmiImage == "" {
		spec.CeilometerCompute.IpmiImage = telemetryDefaults.IpmiContainerImageURL
	}
	if spec.CeilometerCompute.InitImage == "" {
		spec.CeilometerCompute.InitImage = telemetryDefaults.ComputeInitContainerImageURL
	}
	if spec.CeilometerCentral.NotificationImage == "" {
		spec.CeilometerCentral.NotificationImage = telemetryDefaults.NotificationContainerImageURL
	}
	if spec.CeilometerCentral.SgCoreImage == "" {
		spec.CeilometerCentral.SgCoreImage = telemetryDefaults.SgCoreContainerImageURL
	}
	if spec.InfraCompute.NodeExporterImage == "" {
		spec.InfraCompute.NodeExporterImage = telemetryDefaults.NodeExporterContainerImageURL
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-telemetry-openstack-org-v1beta1-telemetry,mutating=false,failurePolicy=fail,sideEffects=None,groups=telemetry.openstack.org,resources=telemetries,verbs=create;update,versions=v1beta1,name=vtelemetry.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Telemetry{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Telemetry) ValidateCreate() error {
	telemetrylog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Telemetry) ValidateUpdate(old runtime.Object) error {
	telemetrylog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Telemetry) ValidateDelete() error {
	telemetrylog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
