package main

func testGetUserID(userID int) int64 {
	return 1
}

func testPushMSG(userID int64) error {
	println("push to user(online)", userID)
	return nil
}

func testGetLimited(driverID int64) bool {
	return false
}

func testGetWhiteList(date string) map[int]int64 {
	return nil
}
