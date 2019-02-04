package main

// normal init
func normalInit() {
	exactDependency = dependencyList{
		pushServiceInstance: pushService{
			PushMSG: onlinePushMSG,
		},
	}
}

// test init
func testInit() {
	exactDependency = dependencyList{
		pushServiceInstance: pushService{
			PushMSG: testPushMSG,
		},
	}
}
