import socket
s = socket.socket()
s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEPORT, 1)
s2 = socket.socket()
s2.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEPORT, 1)
s.bind(('0.0.0.0', 12345))
s2.bind(('0.0.0.0', 12345))
s.connect(('220.181.57.217', 80))
s2.connect(('54.222.60.252', 80))
