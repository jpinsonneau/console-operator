// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	configv1alpha1 "github.com/openshift/api/config/v1alpha1"
	internal "github.com/openshift/client-go/config/applyconfigurations/internal"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// InsightsDataGatherApplyConfiguration represents a declarative configuration of the InsightsDataGather type for use
// with apply.
type InsightsDataGatherApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *InsightsDataGatherSpecApplyConfiguration `json:"spec,omitempty"`
	Status                           *configv1alpha1.InsightsDataGatherStatus  `json:"status,omitempty"`
}

// InsightsDataGather constructs a declarative configuration of the InsightsDataGather type for use with
// apply.
func InsightsDataGather(name string) *InsightsDataGatherApplyConfiguration {
	b := &InsightsDataGatherApplyConfiguration{}
	b.WithName(name)
	b.WithKind("InsightsDataGather")
	b.WithAPIVersion("config.openshift.io/v1alpha1")
	return b
}

// ExtractInsightsDataGather extracts the applied configuration owned by fieldManager from
// insightsDataGather. If no managedFields are found in insightsDataGather for fieldManager, a
// InsightsDataGatherApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// insightsDataGather must be a unmodified InsightsDataGather API object that was retrieved from the Kubernetes API.
// ExtractInsightsDataGather provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractInsightsDataGather(insightsDataGather *configv1alpha1.InsightsDataGather, fieldManager string) (*InsightsDataGatherApplyConfiguration, error) {
	return extractInsightsDataGather(insightsDataGather, fieldManager, "")
}

// ExtractInsightsDataGatherStatus is the same as ExtractInsightsDataGather except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractInsightsDataGatherStatus(insightsDataGather *configv1alpha1.InsightsDataGather, fieldManager string) (*InsightsDataGatherApplyConfiguration, error) {
	return extractInsightsDataGather(insightsDataGather, fieldManager, "status")
}

func extractInsightsDataGather(insightsDataGather *configv1alpha1.InsightsDataGather, fieldManager string, subresource string) (*InsightsDataGatherApplyConfiguration, error) {
	b := &InsightsDataGatherApplyConfiguration{}
	err := managedfields.ExtractInto(insightsDataGather, internal.Parser().Type("com.github.openshift.api.config.v1alpha1.InsightsDataGather"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(insightsDataGather.Name)

	b.WithKind("InsightsDataGather")
	b.WithAPIVersion("config.openshift.io/v1alpha1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *InsightsDataGatherApplyConfiguration) WithKind(value string) *InsightsDataGatherApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *InsightsDataGatherApplyConfiguration) WithAPIVersion(value string) *InsightsDataGatherApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *InsightsDataGatherApplyConfiguration) WithName(value string) *InsightsDataGatherApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *InsightsDataGatherApplyConfiguration) WithGenerateName(value string) *InsightsDataGatherApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *InsightsDataGatherApplyConfiguration) WithNamespace(value string) *InsightsDataGatherApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *InsightsDataGatherApplyConfiguration) WithUID(value types.UID) *InsightsDataGatherApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *InsightsDataGatherApplyConfiguration) WithResourceVersion(value string) *InsightsDataGatherApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *InsightsDataGatherApplyConfiguration) WithGeneration(value int64) *InsightsDataGatherApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *InsightsDataGatherApplyConfiguration) WithCreationTimestamp(value metav1.Time) *InsightsDataGatherApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *InsightsDataGatherApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *InsightsDataGatherApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *InsightsDataGatherApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *InsightsDataGatherApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *InsightsDataGatherApplyConfiguration) WithLabels(entries map[string]string) *InsightsDataGatherApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *InsightsDataGatherApplyConfiguration) WithAnnotations(entries map[string]string) *InsightsDataGatherApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *InsightsDataGatherApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *InsightsDataGatherApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *InsightsDataGatherApplyConfiguration) WithFinalizers(values ...string) *InsightsDataGatherApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *InsightsDataGatherApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *InsightsDataGatherApplyConfiguration) WithSpec(value *InsightsDataGatherSpecApplyConfiguration) *InsightsDataGatherApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *InsightsDataGatherApplyConfiguration) WithStatus(value configv1alpha1.InsightsDataGatherStatus) *InsightsDataGatherApplyConfiguration {
	b.Status = &value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *InsightsDataGatherApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.Name
}
