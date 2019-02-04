package main

// normal init
func normalInit() {
	exactDependency = dependencyList{
		priceServiceInstance: onlinePriceService{},
		pushServiceInstance: onlinePushService{
			userServiceInstance: onlineUserService{
				// may be some many other service
				// may be some many other service
				// may be some many other service
				// may be some many other service
			},
		},
	}
}

// test init
func testInit() {
	exactDependency = dependencyList{
		priceServiceInstance: testPriceService{},
		pushServiceInstance: testPushService{
			userServiceInstance: testUserService{},
		},
	}
}
