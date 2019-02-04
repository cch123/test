package main

func onlineGetUserID(userID int) int64 {
	return int64(userID)
}

func onlinePushMSG(userID int64) error {
	println("push to user(online)", userID)
	return nil
}

func onlineGetLimited(driverID int64) bool {
	return true
}

func onlineGetWhiteList(date string) map[int]int64 {
	return nil
}
