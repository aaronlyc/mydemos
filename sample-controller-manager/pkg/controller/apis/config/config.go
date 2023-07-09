package config

import (
	sampleconfig "mydemos/sample-controller-manager/pkg/controller/sample/config"

	cmconfig "k8s.io/controller-manager/config"
)

type ControllerManagerConfiguration struct {
	Generic cmconfig.GenericControllerManagerConfiguration

	//TODO: 下面是注册的所有自实现的控制器的配置文件
	Samplecontroller sampleconfig.SamplecontrollerConfiguration
}
