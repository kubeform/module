/*
Copyright AppsCode Inc. and Contributors

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

package cmds

import (
	"fmt"
	"os"

	//+kubebuilder:scaffold:imports

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	"kmodules.xyz/client-go/discovery"
	"kmodules.xyz/client-go/meta"
	"kmodules.xyz/client-go/tools/cli"
	"kmodules.xyz/client-go/tools/clusterid"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	auditlib "go.bytebuilders.dev/audit/lib"
	licenseapi "go.bytebuilders.dev/license-verifier/apis/licenses/v1alpha1"
	license "go.bytebuilders.dev/license-verifier/kubernetes"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/klog/v2/klogr"
	"kmodules.xyz/client-go/tools/queue"
	tfv1alpha1 "kubeform.dev/module/api/v1alpha1"
	"kubeform.dev/module/pkg/controller"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	scheme   = clientgoscheme.Scheme
	setupLog = ctrl.Log.WithName("setup")
)

var (
	licenseFile          string
	metricsAddr          string
	enableLeaderElection bool
	probeAddr            string
	secretKey            string
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(tfv1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func NewCmdRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Launch Kubeform Module controller",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			klog.Infoln("Starting kf-module-controller...")

			ctrl.SetLogger(klogr.New())

			ctx := ctrl.SetupSignalHandler()

			mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
				Scheme:                 scheme,
				MetricsBindAddress:     metricsAddr,
				Port:                   9443,
				HealthProbeBindAddress: probeAddr,
				LeaderElection:         enableLeaderElection,
				LeaderElectionID:       "0e49653b.kubeform.com",
			})
			if err != nil {
				setupLog.Error(err, "unable to start manager")
				os.Exit(1)
			}
			cfg := mgr.GetConfig()

			restrictToNamespace := queue.NamespaceDemo
			if licenseFile != "" {
				info := license.NewLicenseEnforcer(cfg, licenseFile).LoadLicense()
				if info.Status != licenseapi.LicenseActive {
					klog.Infof("License status %s", info.Status)
					os.Exit(1)
				}
				if sets.NewString(info.Features...).Has("kubeform-enterprise") {
					restrictToNamespace = ""
				} else if !sets.NewString(info.Features...).Has("kubeform-community") {
					setupLog.Error(fmt.Errorf("not a valid license for this product"), "")
					os.Exit(1)
				}
			}

			// audit event publisher
			var auditor *auditlib.EventPublisher
			if licenseFile != "" && cli.EnableAnalytics {
				kc, err := kubernetes.NewForConfig(cfg)
				if err != nil {
					setupLog.Error(err, "unable to create Kubernetes client")
					os.Exit(1)
				}
				cmeta, err := clusterid.ClusterMetadata(kc.CoreV1().Namespaces())
				if err != nil {
					setupLog.Error(err, "failed to extract cluster metadata, reason: %v")
					os.Exit(1)
				}
				mapper := discovery.NewResourceMapper(mgr.GetRESTMapper())
				fn := auditlib.BillingEventCreator{
					Mapper:          mapper,
					ClusterMetadata: cmeta,
				}
				auditor = auditlib.NewResilientEventPublisher(func() (*auditlib.NatsConfig, error) {
					return auditlib.NewNatsConfig(cmeta.UID, licenseFile)
				}, mapper, fn.CreateEvent)
			}

			if err = (&controller.ModuleReconciler{
				Client: mgr.GetClient(),
				Log:    ctrl.Log.WithName("controllers").WithName("Module"),
				Gvk: schema.GroupVersionKind{
					Group:   "tf.kubeform.com",
					Version: "v1alpha1",
					Kind:    "Module",
				},
				SecretKey: secretKey,
				Scheme:    mgr.GetScheme(),
			}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
				setupLog.Error(err, "unable to create controller", "controller", "Module")
				os.Exit(1)
			}
			//+kubebuilder:scaffold:builder

			// Start periodic license verification
			//nolint:errcheck
			go license.VerifyLicensePeriodically(mgr.GetConfig(), licenseFile, ctx.Done())

			if auditor != nil {
				if err := auditor.SetupSiteInfoPublisherWithManager(mgr); err != nil {
					setupLog.Error(err, "unable to setup site info publisher")
					os.Exit(1)
				}
			}

			setupLog.Info("starting manager")
			if err := mgr.Start(ctx); err != nil {
				setupLog.Error(err, "problem running manager")
				os.Exit(1)
			}
		},
	}

	meta.AddLabelBlacklistFlag(cmd.Flags())
	cmd.Flags().StringVar(&licenseFile, "license-file", licenseFile, "Path to license file")
	cmd.Flags().StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	cmd.Flags().StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	cmd.Flags().BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	cmd.Flags().StringVar(&secretKey, "secret-key", "YXBwc2NvZGVrdWJlZm9ybXNlY3JldGtleWFhYWFhYQo=", "The base64 encoded secret key to use during encode and decode tfstate")

	return cmd
}
