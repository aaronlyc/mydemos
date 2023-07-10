package app

import (
	"context"
	"time"

	"mydemos/sample-controller-manager/pkg/controller/sample"
	sampleclientset "mydemos/sample-controller-manager/pkg/generated/clientset/versioned"
	sampleinformers "mydemos/sample-controller-manager/pkg/generated/informers/externalversions"

	"k8s.io/controller-manager/controller"
)

func startSampleController(ctx context.Context, controllerContext ControllerContext) (controller.Interface, bool, error) {

	restConfig := controllerContext.ClientBuilder.ConfigOrDie(controllerContext.ComponentConfig.Samplecontroller.ControllerName)

	exampleClient := sampleclientset.NewForConfigOrDie(restConfig)
	exampleInformerFactory := sampleinformers.NewSharedInformerFactory(exampleClient, time.Second*30)

	samplecontroller := sample.NewController(
		ctx,
		controllerContext.ClientBuilder.ClientOrDie(controllerContext.ComponentConfig.Samplecontroller.ControllerName),
		exampleClient,
		controllerContext.InformerFactory.Apps().V1().Deployments(),
		exampleInformerFactory.Samplecontroller().V1alpha1().Foos(),
	)

	exampleInformerFactory.Start(ctx.Done())

	go samplecontroller.Run(ctx, int(controllerContext.ComponentConfig.Samplecontroller.ConcurrentSampleSyncs))
	return nil, true, nil
}
