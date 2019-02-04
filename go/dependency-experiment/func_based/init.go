package main

func onlineInitDependency() {
	exactDependency = dependencyList{
		GetUserID:    onlineGetUserID,
		PushMSG:      onlinePushMSG,
		GetLimited:   onlineGetLimited,
		GetWhiteList: onlineGetWhiteList,
	}
}

func testInitDependency() {
	exactDependency = dependencyList{
		GetUserID:    testGetUserID,
		PushMSG:      testPushMSG,
		GetLimited:   testGetLimited,
		GetWhiteList: testGetWhiteList,
	}
}
