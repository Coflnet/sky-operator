package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	skyv1alpha1 "github.com/Coflnet/sky-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
)

// ProxyReconciler reconciles a Proxy object
type ProxyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

var (
	logger = ctrl.Log.WithName("proxy_controller")
)

//+kubebuilder:rbac:groups=sky.coflnet.com,resources=proxies,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=sky.coflnet.com,resources=proxies/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=sky.coflnet.com,resources=proxies/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Proxy object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *ProxyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// fetch the proxy
	proxy := &skyv1alpha1.Proxy{}
	err := r.Get(ctx, req.NamespacedName, proxy)
	if err != nil {
		logger.Error(err, "unable to fetch Proxy")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	logger.Info("reconciling proxy", "proxy", proxy)

	// fetch the desired amount of proxies
	logger.Info("Reconciling Proxy", "Proxy.Namespace", proxy.Namespace, "Proxy.Name", proxy.Name)

	// ensure proxy deployment exists and is correct
	logger.Info("Reconciling Proxy Deployment")

	// update the deployment replica count
	logger.Info("Updating Proxy Deployment replica count")

	// update the status
	proxy.Status.Replicas = 5
	r.Status().Update(ctx, proxy)
	logger.Info("Updated Proxy Deployment status")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProxyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&skyv1alpha1.Proxy{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}
