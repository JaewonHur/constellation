
package main

import (
	"context"
	"flag"
	"os"
	"strings"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	updatev1alpha1 "github.com/edgelesssys/constellation/operators/constellation-node-operator/api/v1alpha1"
	"github.com/edgelesssys/constellation/operators/constellation-node-operator/controllers"
	nodemaintenancev1beta1 "github.com/medik8s/node-maintenance-operator/api/v1beta1"
	//+kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(nodemaintenancev1beta1.AddToScheme(scheme))
	utilruntime.Must(updatev1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	var csp string
	var metricsAddr string
	var enableLeaderElection bool
	var probeAddr string
	flag.StringVar(&csp, "csp", "", "Cloud Service Provider the image is running on")
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	var cspClient cspAPI
	switch strings.ToLower(csp) {
	case "azure":
		panic("Azure is not supported yet")
	case "gcp":
		panic("GCP is not supported yet")
	default:
		setupLog.Info("Unknown CSP", "csp", csp)
		os.Exit(1)
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "38cc1645.edgeless.systems",
	})
	if err != nil {
		setupLog.Error(err, "Unable to start manager")
		os.Exit(1)
	}

	if err = controllers.NewNodeImageReconciler(
		cspClient, mgr.GetClient(), mgr.GetScheme(),
	).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "Unable to create controller", "controller", "NodeImage")
		os.Exit(1)
	}
	if err = (&controllers.AutoscalingStrategyReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "Unable to create controller", "controller", "AutoscalingStrategy")
		os.Exit(1)
	}
	if err = controllers.NewScalingGroupReconciler(
		cspClient, mgr.GetClient(), mgr.GetScheme(),
	).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "Unable to create controller", "controller", "ScalingGroup")
		os.Exit(1)
	}
	if err = controllers.NewPendingNodeReconciler(
		cspClient, mgr.GetClient(), mgr.GetScheme(),
	).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "Unable to create controller", "controller", "PendingNode")
		os.Exit(1)
	}
	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "Unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "Unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("Starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "Problem running manager")
		os.Exit(1)
	}
}

type cspAPI interface {
	// GetNodeImage retrieves the image currently used by a node.
	GetNodeImage(ctx context.Context, providerID string) (string, error)
	// GetScalingGroupID retrieves the scaling group that a node is part of.
	GetScalingGroupID(ctx context.Context, providerID string) (string, error)
	// CreateNode creates a new node inside a specified scaling group at the CSP and returns its future name and provider id.
	CreateNode(ctx context.Context, scalingGroupID string) (nodeName, providerID string, err error)
	// DeleteNode starts the termination of the node at the CSP.
	DeleteNode(ctx context.Context, providerID string) error
	// GetNodeState retrieves the state of a pending node from a CSP.
	GetNodeState(ctx context.Context, providerID string) (updatev1alpha1.CSPNodeState, error)
	// GetScalingGroupImage retrieves the image currently used by a scaling group.
	GetScalingGroupImage(ctx context.Context, scalingGroupID string) (string, error)
	// SetScalingGroupImage sets the image to be used by newly created nodes in a scaling group.
	SetScalingGroupImage(ctx context.Context, scalingGroupID, imageURI string) error
}
