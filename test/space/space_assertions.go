package space

import (
	"context"

	toolchainv1alpha1 "github.com/codeready-toolchain/api/api/v1alpha1"
	"github.com/codeready-toolchain/toolchain-common/pkg/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Assertion struct {
	space          *toolchainv1alpha1.Space
	client         client.Client
	namespacedName types.NamespacedName
	t              test.T
}

func (a *Assertion) loadResource() error {
	tier := &toolchainv1alpha1.Space{}
	err := a.client.Get(context.TODO(), a.namespacedName, tier)
	a.space = tier
	return err
}

// AssertThatSpace helper func to begin with the assertions on a Space
func AssertThatSpace(t test.T, namespace, name string, client client.Client) *Assertion {
	return &Assertion{
		client:         client,
		namespacedName: test.NamespacedName(namespace, name),
		t:              t,
	}
}

func (a *Assertion) Exists() *Assertion {
	err := a.loadResource()
	require.NoError(a.t, err)
	return a
}

func (a *Assertion) HasFinalizer(name string) *Assertion {
	err := a.loadResource()
	require.NoError(a.t, err)
	assert.Contains(a.t, a.space.Finalizers, name)
	return a
}

func (a *Assertion) HasNoFinalizers() *Assertion {
	err := a.loadResource()
	require.NoError(a.t, err)
	assert.Empty(a.t, a.space.Finalizers)
	return a
}

func (a *Assertion) HasNoStatusTargetCluster() *Assertion {
	err := a.loadResource()
	require.NoError(a.t, err)
	assert.Empty(a.t, a.space.Status.TargetCluster)
	return a
}

func (a *Assertion) HasStatusTargetCluster(targetCluster string) *Assertion {
	err := a.loadResource()
	require.NoError(a.t, err)
	assert.Equal(a.t, targetCluster, a.space.Status.TargetCluster)
	return a
}

func (a *Assertion) HasConditions(expected ...toolchainv1alpha1.Condition) *Assertion {
	err := a.loadResource()
	require.NoError(a.t, err)
	test.AssertConditionsMatch(a.t, a.space.Status.Conditions, expected...)
	return a
}

func (a *Assertion) HasNoConditions() *Assertion {
	err := a.loadResource()
	require.NoError(a.t, err)
	assert.Empty(a.t, a.space.Status.Conditions)
	return a
}

func Provisioning() toolchainv1alpha1.Condition {
	return toolchainv1alpha1.Condition{
		Type:   toolchainv1alpha1.ConditionReady,
		Status: corev1.ConditionFalse,
		Reason: toolchainv1alpha1.SpaceProvisioningReason,
	}
}

func ProvisioningPending(msg string) toolchainv1alpha1.Condition {
	return toolchainv1alpha1.Condition{
		Type:    toolchainv1alpha1.ConditionReady,
		Status:  corev1.ConditionFalse,
		Reason:  toolchainv1alpha1.SpaceProvisioningPendingReason,
		Message: msg,
	}
}

func ProvisioningFailed(msg string) toolchainv1alpha1.Condition {
	return toolchainv1alpha1.Condition{
		Type:    toolchainv1alpha1.ConditionReady,
		Status:  corev1.ConditionFalse,
		Reason:  toolchainv1alpha1.SpaceProvisioningFailedReason,
		Message: msg,
	}
}

func UnableToCreateNSTemplateSet(msg string) toolchainv1alpha1.Condition {
	return toolchainv1alpha1.Condition{
		Type:    toolchainv1alpha1.ConditionReady,
		Status:  corev1.ConditionFalse,
		Reason:  toolchainv1alpha1.SpaceUnableToCreateNSTemplateSetReason,
		Message: msg,
	}
}

func Ready() toolchainv1alpha1.Condition {
	return toolchainv1alpha1.Condition{
		Type:   toolchainv1alpha1.ConditionReady,
		Status: corev1.ConditionTrue,
		Reason: toolchainv1alpha1.SpaceProvisionedReason,
	}
}

func Terminating() toolchainv1alpha1.Condition {
	return toolchainv1alpha1.Condition{
		Type:   toolchainv1alpha1.ConditionReady,
		Status: corev1.ConditionFalse,
		Reason: toolchainv1alpha1.SpaceTerminatingReason,
	}
}

func TerminatingFailed(msg string) toolchainv1alpha1.Condition {
	return toolchainv1alpha1.Condition{
		Type:    toolchainv1alpha1.ConditionReady,
		Status:  corev1.ConditionFalse,
		Reason:  toolchainv1alpha1.SpaceTerminatingFailedReason,
		Message: msg,
	}
}