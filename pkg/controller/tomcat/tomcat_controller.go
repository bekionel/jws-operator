package tomcat

import (
	"context"

	"k8s.io/apimachinery/pkg/util/intstr"

	jwsv1alpha1 "github.com/jws-image-operator/pkg/apis/jws/v1alpha1"

	appsv1 "github.com/openshift/api/apps/v1"
	buildv1 "github.com/openshift/api/build/v1"
	imagev1 "github.com/openshift/api/image/v1"
	routev1 "github.com/openshift/api/route/v1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_tomcat")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Tomcat Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileTomcat{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("tomcat-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Tomcat
	err = c.Watch(&source.Kind{Type: &jwsv1alpha1.Tomcat{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Tomcat
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &jwsv1alpha1.Tomcat{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileTomcat{}

// ReconcileTomcat reconciles a Tomcat object
type ReconcileTomcat struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Tomcat object and makes changes based on the state read
// and what is in the Tomcat.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileTomcat) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Tomcat")

	// Fetch the Tomcat tomcat
	tomcat := &jwsv1alpha1.Tomcat{}
	err := r.client.Get(context.TODO(), request.NamespacedName, tomcat)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			reqLogger.Info("Tomcat resource not found. Ignoring since object must be deleted")
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		reqLogger.Error(err, "Failed to get Tomcat")
		return reconcile.Result{}, err
	}

	// Check if the Service already exists, if not create a new one
	list := &corev1.ServiceList{}
	opts := &client.ListOptions{}
	err = r.client.List(context.TODO(), opts, list)
	if (err != nil && errors.IsNotFound(err)) || len(list.Items) == 1 {
		// Define a new Service
		ser := r.serviceForTomcat(tomcat)
		reqLogger.Info("Creating a new Service.", "Service.Namespace", ser.Namespace, "Service.Name", ser.Name)
		err = r.client.Create(context.TODO(), ser)
		if err != nil {
			reqLogger.Error(err, "Failed to create new Service.", "Service.Namespace", ser.Namespace, "Service.Name", ser.Name)
			return reconcile.Result{}, err
		}
		// Service created successfully - return and requeue
		return reconcile.Result{Requeue: true}, nil
	} else if err != nil {
		reqLogger.Error(err, "Failed to get Service.")
		return reconcile.Result{}, err
	}

	// Check if the Route already exists, if not create a new one
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: tomcat.Spec.ApplicationName, Namespace: tomcat.Namespace}, &routev1.Route{})
	if err != nil && errors.IsNotFound(err) {
		// Define a new Route
		rou := r.routeForTomcat(tomcat)
		reqLogger.Info("Creating a new Route.", "Route.Namespace", rou.Namespace, "Route.Name", rou.Name)
		err = r.client.Create(context.TODO(), rou)
		if err != nil {
			reqLogger.Error(err, "Failed to create new Route.", "Route.Namespace", rou.Namespace, "Route.Name", rou.Name)
			return reconcile.Result{}, err
		}
		// Route created successfully - return and requeue
		return reconcile.Result{Requeue: true}, nil
	} else if err != nil {
		reqLogger.Error(err, "Failed to get Service.")
		return reconcile.Result{}, err
	}

	// Check if the ImageStream already exists, if not create a new one
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: tomcat.Spec.ApplicationName, Namespace: tomcat.Namespace}, &imagev1.ImageStream{})
	if err != nil && errors.IsNotFound(err) {
		// Define a new ImageStream
		img := r.imageStreamForTomcat(tomcat)
		reqLogger.Info("Creating a new ImageStream.", "ImageStream.Namespace", img.Namespace, "ImageStream.Name", img.Name)
		err = r.client.Create(context.TODO(), img)
		if err != nil {
			reqLogger.Error(err, "Failed to create new ImageStream.", "ImageStream.Namespace", img.Namespace, "ImageStream.Name", img.Name)
			return reconcile.Result{}, err
		}
		// ImageStream created successfully - return and requeue
		return reconcile.Result{Requeue: true}, nil
	} else if err != nil {
		reqLogger.Error(err, "Failed to get ImageStream.")
		return reconcile.Result{}, err
	}

	// Check if the BuildConfig already exists, if not create a new one
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: tomcat.Spec.ApplicationName, Namespace: tomcat.Namespace}, &buildv1.BuildConfig{})
	if err != nil && errors.IsNotFound(err) {
		// Define a new BuildConfig
		bui := r.buildConfigForTomcat(tomcat)
		reqLogger.Info("Creating a new BuildConfig.", "BuildConfig.Namespace", bui.Namespace, "BuildConfig.Name", bui.Name)
		err = r.client.Create(context.TODO(), bui)
		if err != nil {
			reqLogger.Error(err, "Failed to create new BuildConfig.", "BuildConfig.Namespace", bui.Namespace, "BuildConfig.Name", bui.Name)
			return reconcile.Result{}, err
		}
		// BuildConfig created successfully - return and requeue
		return reconcile.Result{Requeue: true}, nil
	} else if err != nil {
		reqLogger.Error(err, "Failed to get Service.")
		return reconcile.Result{}, err
	}

	// Check if the DeploymentConfig already exists, if not create a new one
	foundDeployment := &appsv1.DeploymentConfig{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: tomcat.Spec.ApplicationName, Namespace: tomcat.Namespace}, foundDeployment)
	if err != nil && errors.IsNotFound(err) {
		// Define a new DeploymentConfig
		dep := r.deploymentConfigForTomcat(tomcat)
		reqLogger.Info("Creating a new DeploymentConfig.", "DeploymentConfig.Namespace", dep.Namespace, "DeploymentConfig.Name", dep.Name)
		err = r.client.Create(context.TODO(), dep)
		if err != nil {
			reqLogger.Error(err, "Failed to create new DeploymentConfig.", "DeploymentConfig.Namespace", dep.Namespace, "DeploymentConfig.Name", dep.Name)
			return reconcile.Result{}, err
		}
		// DeploymentConfig created successfully - return and requeue
		return reconcile.Result{Requeue: true}, nil
	} else if err != nil {
		reqLogger.Error(err, "Failed to get Service.")
		return reconcile.Result{}, err
	}

	// Handle Scaling
	size := tomcat.Spec.Size
	if foundDeployment.Spec.Replicas != size {
		foundDeployment.Spec.Replicas = size
		err = r.client.Update(context.TODO(), foundDeployment)
		if err != nil {
			reqLogger.Error(err, "Failed to update Deployment.", "Deployment.Namespace", foundDeployment.Namespace, "Deployment.Name", foundDeployment.Name)
			return reconcile.Result{}, err
		}
		// Spec updated - return and requeue
		return reconcile.Result{Requeue: true}, nil
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileTomcat) serviceForTomcat(t *jwsv1alpha1.Tomcat) *corev1.Service {

	service := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      t.Spec.ApplicationName,
			Namespace: t.Namespace,
			Labels: map[string]string{
				"application": t.Spec.ApplicationName,
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{
				Name:       "ui",
				Port:       8080,
				TargetPort: intstr.FromInt(8080),
			}},
			Selector: map[string]string{
				"deploymentConfig": t.Spec.ApplicationName,
			},
		},
	}

	return service
}

func (r *ReconcileTomcat) deploymentConfigForTomcat(t *jwsv1alpha1.Tomcat) *appsv1.DeploymentConfig {

	terminationGracePeriodSeconds := int64(60)

	deploymentConfig := &appsv1.DeploymentConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps.openshift.io/v1",
			Kind:       "DeploymentConfig",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      t.Spec.ApplicationName,
			Namespace: t.Namespace,
			Labels: map[string]string{
				"application": t.Spec.ApplicationName,
			},
		},
		Spec: appsv1.DeploymentConfigSpec{
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.DeploymentStrategyTypeRecreate,
			},
			Triggers: []appsv1.DeploymentTriggerPolicy{{
				Type: appsv1.DeploymentTriggerOnImageChange,
				ImageChangeParams: &appsv1.DeploymentTriggerImageChangeParams{
					Automatic:      true,
					ContainerNames: []string{t.Spec.ApplicationName},
					From: corev1.ObjectReference{
						Kind: "ImageStreamTag",
						Name: t.Spec.ApplicationName + ":latest",
					},
				},
			},
				{
					Type: appsv1.DeploymentTriggerOnConfigChange,
				}},
			Replicas: 1,
			Selector: map[string]string{
				"deploymentConfig": t.Spec.ApplicationName,
			},
			Template: &corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: t.Spec.ApplicationName,
					Labels: map[string]string{
						"application":      t.Spec.ApplicationName,
						"deploymentConfig": t.Spec.ApplicationName,
					},
				},
				Spec: corev1.PodSpec{
					TerminationGracePeriodSeconds: &terminationGracePeriodSeconds,
					Containers: []corev1.Container{{
						Name:            t.Spec.ApplicationName,
						Image:           t.Spec.ApplicationName,
						ImagePullPolicy: "Always",
						ReadinessProbe: &corev1.Probe{
							Handler: corev1.Handler{
								Exec: &corev1.ExecAction{
									Command: []string{
										"/bin/bash",
										"-c",
										"curl --noproxy '*' -s -u ${JWS_ADMIN_USERNAME}:${JWS_ADMIN_PASSWORD} 'http://localhost:8080/manager/jmxproxy/?get=Catalina%3Atype%3DServer&att=stateName' |grep -iq 'stateName *= *STARTED'",
									},
								},
							},
						},
						Ports: []corev1.ContainerPort{{
							Name:          "jolokia",
							ContainerPort: 8778,
							Protocol:      corev1.ProtocolTCP,
						}, {
							Name:          "http",
							ContainerPort: 8080,
							Protocol:      corev1.ProtocolTCP,
						}},
						Env: []corev1.EnvVar{{
							Name:  "JWS_ADMIN_USERNAME",
							Value: t.Spec.JwsAdminUsername,
						}, {
							Name:  "JWS_ADMIN_PASSWORD",
							Value: t.Spec.JwsAdminPassword,
						}},
					}},
				},
			},
		},
	}

	return deploymentConfig
}

func (r *ReconcileTomcat) routeForTomcat(t *jwsv1alpha1.Tomcat) *routev1.Route {

	route := &routev1.Route{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "route.openshift.io/v1",
			Kind:       "Route",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      t.Spec.ApplicationName,
			Namespace: t.Namespace,
			Labels: map[string]string{
				"application": t.Spec.ApplicationName,
			},
			Annotations: map[string]string{
				"description": "Route for application's http service.",
			},
		},
		Spec: routev1.RouteSpec{
			Host: t.Spec.HostnameHttp,
			To: routev1.RouteTargetReference{
				Name: t.Spec.ApplicationName,
			},
		},
	}

	return route
}

func (r *ReconcileTomcat) imageStreamForTomcat(t *jwsv1alpha1.Tomcat) *imagev1.ImageStream {

	imageStream := &imagev1.ImageStream{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "image.openshift.io/v1",
			Kind:       "ImageStream",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      t.Spec.ApplicationName,
			Namespace: t.Namespace,
			Labels: map[string]string{
				"application": t.Spec.ApplicationName,
			},
		},
	}

	return imageStream
}

func (r *ReconcileTomcat) buildConfigForTomcat(t *jwsv1alpha1.Tomcat) *buildv1.BuildConfig {

	buildConfig := &buildv1.BuildConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "build.openshift.io/v1",
			Kind:       "BuildConfig",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      t.Spec.ApplicationName,
			Namespace: t.Namespace,
			Labels: map[string]string{
				"application": t.Spec.ApplicationName,
			},
		},
		Spec: buildv1.BuildConfigSpec{
			CommonSpec: buildv1.CommonSpec{
				Source: buildv1.BuildSource{
					Type: "Git",
					Git: &buildv1.GitBuildSource{
						URI: t.Spec.SourceRepositoryUrl,
						Ref: t.Spec.SourceRepositoryRef,
					},
					ContextDir: t.Spec.ContextDir,
				},
				Strategy: buildv1.BuildStrategy{
					Type: "Source",
					SourceStrategy: &buildv1.SourceBuildStrategy{
						Env: []corev1.EnvVar{{
							Name:  "MAVEN_MIRROR_URL",
							Value: t.Spec.MavenMirrorUrl,
						}, {
							Name:  "ARTIFACT_DIR",
							Value: t.Spec.ArtifactDir,
						}},
						ForcePull: true,
						From: corev1.ObjectReference{
							Kind:      "ImageStreamTag",
							Namespace: t.Spec.ImageStreamNamespace,
							Name:      "jboss-webserver53-tomcat9-openshift:latest",
						},
					},
				},
				Output: buildv1.BuildOutput{
					To: &corev1.ObjectReference{
						Kind: "ImageStreamTag",
						Name: t.Spec.ApplicationName + ":latest",
					},
				},
			},
			Triggers: []buildv1.BuildTriggerPolicy{{
				Type: "Github",
				GitHubWebHook: &buildv1.WebHookTrigger{
					Secret: t.Spec.GithubWebhookSecret,
				},
			}, {
				Type: "Generic",
				GenericWebHook: &buildv1.WebHookTrigger{
					Secret: t.Spec.GenericWebhookSecret,
				},
			}, {
				Type:        "ImageChange",
				ImageChange: &buildv1.ImageChangeTrigger{},
			}, {
				Type: "ConfigChange",
			}},
		},
	}

	return buildConfig
}
