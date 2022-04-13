package addon

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	operatorsv1 "github.com/operator-framework/api/pkg/operators/v1"
	"go.opentelemetry.io/otel"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	addonsv1alpha1 "github.com/openshift/addon-operator/apis/addons/v1alpha1"
	"github.com/openshift/addon-operator/internal/controllers"
)

// Ensures the presense or absense of an OperatorGroup depending on the Addon install type.
func (r *AddonReconciler) ensureOperatorGroup(
	ctx context.Context, log logr.Logger, addon *addonsv1alpha1.Addon) (requeueResult, error) {

	_, span := otel.Tracer(addonTracer).Start(ctx, "ensureOperatorGroup")
	defer span.End()

	targetNamespace, _, stop := r.parseAddonInstallConfig(log, addon)
	if stop {
		return resultStop, nil
	}

	desiredOperatorGroup := &operatorsv1.OperatorGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      controllers.DefaultOperatorGroupName,
			Namespace: targetNamespace,
			Labels:    map[string]string{},
		},
	}
	if addon.Spec.Install.Type == addonsv1alpha1.OLMOwnNamespace {
		desiredOperatorGroup.Spec.TargetNamespaces = []string{targetNamespace}
	}

	controllers.AddCommonLabels(desiredOperatorGroup, addon)
	if err := controllerutil.SetControllerReference(addon, desiredOperatorGroup, r.Scheme); err != nil {
		return resultNil, fmt.Errorf("setting controller reference: %w", err)
	}
	return resultNil, r.reconcileOperatorGroup(ctx, desiredOperatorGroup, addon.Spec.ResourceAdoptionStrategy)
}

// Reconciles the Spec of the given OperatorGroup if needed by updating or creating the OperatorGroup.
// The given OperatorGroup is updated to reflect the latest state from the kube-apiserver.
func (r *AddonReconciler) reconcileOperatorGroup(
	ctx context.Context, operatorGroup *operatorsv1.OperatorGroup, strategy addonsv1alpha1.ResourceAdoptionStrategyType) error {
	currentOperatorGroup := &operatorsv1.OperatorGroup{}

	err := r.Get(ctx, client.ObjectKeyFromObject(operatorGroup), currentOperatorGroup)
	if errors.IsNotFound(err) {
		return r.Create(ctx, operatorGroup)
	}
	if err != nil {
		return fmt.Errorf("getting OperatorGroup: %w", err)
	}

	currentLabels := labels.Set(currentOperatorGroup.Labels)
	newLabels := labels.Merge(currentLabels, labels.Set(operatorGroup.Labels))
	ownedByAddon := controllers.HasEqualControllerReference(currentOperatorGroup, operatorGroup)
	specChanged := !equality.Semantic.DeepEqual(currentOperatorGroup.Spec, operatorGroup.Spec)
	if specChanged || !ownedByAddon || !labels.Equals(currentLabels, newLabels) {
		// TODO: remove this condition once resourceAdoptionStrategy is discontinued
		if strategy != addonsv1alpha1.ResourceAdoptionAdoptAll && !ownedByAddon {
			return controllers.ErrNotOwnedByUs
		}
		currentOperatorGroup.Spec = operatorGroup.Spec
		currentOperatorGroup.OwnerReferences = operatorGroup.OwnerReferences
		currentOperatorGroup.Labels = newLabels
		return r.Update(ctx, currentOperatorGroup)
	}
	return nil
}
