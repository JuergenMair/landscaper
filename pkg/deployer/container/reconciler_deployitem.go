// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors.
//
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/gardener/landscaper/apis/deployer/container"
	deployerlib "github.com/gardener/landscaper/pkg/deployer/lib"
	"github.com/gardener/landscaper/pkg/kubernetes"
	componentsregistry "github.com/gardener/landscaper/pkg/landscaper/registry/components"

	lsv1alpha1 "github.com/gardener/landscaper/apis/core/v1alpha1"
	lsv1alpha1helper "github.com/gardener/landscaper/apis/core/v1alpha1/helper"
	containerv1alpha1 "github.com/gardener/landscaper/apis/deployer/container/v1alpha1"
	kutil "github.com/gardener/landscaper/pkg/utils/kubernetes"
)

const (
	cacheIdentifier = "container-deployer-controller"
)

func NewDeployItemReconciler(log logr.Logger, lsClient, hostClient, directHostClient client.Client, config *containerv1alpha1.Configuration) (*DeployItemReconciler, error) {
	componentRegistryMgr, err := componentsregistry.SetupManagerFromConfig(log, config.OCI, cacheIdentifier)
	if err != nil {
		return nil, err
	}
	return &DeployItemReconciler{
		log:                   log,
		config:                config,
		lsClient:              lsClient,
		hostClient:            hostClient,
		directHostClient:      directHostClient,
		scheme:                kubernetes.LandscaperScheme,
		componentsRegistryMgr: componentRegistryMgr,
	}, nil
}

type DeployItemReconciler struct {
	log                   logr.Logger
	lsClient              client.Client
	hostClient            client.Client
	directHostClient      client.Client
	scheme                *runtime.Scheme
	config                *containerv1alpha1.Configuration
	componentsRegistryMgr *componentsregistry.Manager
}

func (r *DeployItemReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	logger := r.log.WithValues("resource", req.NamespacedName)
	deployItem, err := GetAndCheckReconcile(r.log, r.lsClient, r.config)(ctx, req)
	if err != nil {
		return reconcile.Result{}, err
	}
	if deployItem == nil {
		return reconcile.Result{}, nil
	}

	if deployItem.Status.ObservedGeneration == deployItem.Generation && !lsv1alpha1helper.HasOperation(deployItem.ObjectMeta, lsv1alpha1.ReconcileOperation) {
		logger.V(5).Info("Version already reconciled")
		return reconcile.Result{}, nil
	}

	logger.Info("Reconcile container deploy item")
	errHdl := deployerlib.HandleErrorFunc(logger, r.lsClient, deployItem)
	if err := errHdl(ctx, r.reconcile(ctx, deployItem)); err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func (r *DeployItemReconciler) reconcile(ctx context.Context, deployItem *lsv1alpha1.DeployItem) (err error) {
	containerOp, err := New(r.log, r.lsClient, r.hostClient, r.directHostClient, r.config, deployItem, r.componentsRegistryMgr)
	if err != nil {
		return err
	}

	if !kutil.HasFinalizer(deployItem, lsv1alpha1.LandscaperFinalizer) {
		controllerutil.AddFinalizer(deployItem, lsv1alpha1.LandscaperFinalizer)
		if err := r.lsClient.Update(ctx, deployItem); err != nil {
			return err
		}
		return nil
	}

	if !deployItem.DeletionTimestamp.IsZero() {
		return containerOp.Delete(ctx)
	}

	return containerOp.Reconcile(ctx, container.OperationReconcile)
}
