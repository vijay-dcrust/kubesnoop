/*


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

package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1alpha1 "github.com/vijay-dcrust/kubesnoop/api/v1alpha1"
)

// LifeCycleReconciler reconciles a LifeCycle object
type LifeCycleReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=core.kubesnoop.io,resources=lifecycles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core.kubesnoop.io,resources=lifecycles/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core.kubesnoop.io,resources=pods,verbs=get;list;watch;create;update

func (r *LifeCycleReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("lifecycle", req.NamespacedName)
	//Reconciler loop logic
	for {
		podList := &v1.PodList{}
		opts := []client.ListOption{
			client.InNamespace(req.NamespacedName.Namespace),
			//client.MatchingLabels{"k8s-app": "nexus-nginx"},
		}
		//err := r.client.Get(context.TODO(), req.NamespacedName, podList)
		err := r.List(ctx, podList, opts...)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println("PodName, StartTime, Status")
		for _, pod := range podList.Items {
			fmt.Println(pod.Name, pod.Status.StartTime, pod.Status.Phase)
		}
		//return ctrl.Result{}, nil
		time.Sleep(20 * time.Second)
	}
}

func (r *LifeCycleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1alpha1.LifeCycle{}).
		Complete(r)
}
