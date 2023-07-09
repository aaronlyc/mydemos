package config

func RecommendedDefaultSampleControllerConfiguration(obj *SamplecontrollerConfiguration) {
	if obj.ConcurrentSampleSyncs == 0 {
		obj.ConcurrentSampleSyncs = 2
	}

	if len(obj.ControllerName) == 0 {
		obj.ControllerName = "sample-controller"
	}
}
