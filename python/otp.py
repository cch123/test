import pyotp
totp = pyotp.TOTP("2SVVV5X3W4DJWBCX")
print("Current OTP:", totp.now())

