
https://blog.lilydjwg.me/2015/8/19/tcp-fun.180084.html

人们对 TCP 的误解
因为我们的教育总是只教人「怎么做」，而根本不管「为什么这么做」，所以造成了很多误解。

今天（恰巧是今天）看到有人在 SegmentFault 上问「TCP server 为什么一个端口可以建立多个连接？」。提问者认为 client 端就不能使用相同的本地端口了。理论上来说，确定一条链路，只要五元组（源IP、源端口号、目标IP、目标端口号、协议）唯一就可以了，所以这不应该是技术限制。而实际上，Linux 3.9 之后确实可以让客户端使用相同的地址来连接不同的目标，只不过要提前跟内核说好而已。

当然，你不能使用同一个 socket，不然调用connect连接的时候会报错：

[Errno 106] (EISCONN) Transport endpoint is already connected
man 2 connect里说了：

Generally, connection-based protocol sockets may successfully connect() only once; connectionless protocol sockets may use connect() multiple times to change their association.
想也是，一个 socket 连接到多个目标，那发送的时候到底发给谁呢？TCP 又不像 UDP 那样无状态的，以前做过什么根本不管。

那用多个 socket 就可以了嘛。服务端其实也一直是用多个 socket 来处理多个连接的不是么，每次accept都生成个新的 socket。

>>> import socket
>>> s = socket.socket()
# since Linux 3.9, 见 man 7 socket
>>> s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEPORT, 1)
>>> s2 = socket.socket()
>>> s2.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEPORT, 1)
>>> s.bind(('127.0.0.1', 12345))
>>> s2.bind(('127.0.0.1', 12345))
# 都可以使用同一本地地址来连接哦
>>> s.connect(('127.0.0.1', 80))
>>> s2.connect(('127.0.0.1', 4321))
连上去之后 netstat 的输出（4568 进程是上边这个程序，另两个进程一个是 nginx，另一个是我的另一个 Python 程序）：

>>> netstat -npt | grep 12345
(Not all processes could be identified, non-owned process info
 will not be shown, you would have to be root to see it all.)
tcp        0      0 127.0.0.1:4321          127.0.0.1:12345         ESTABLISHED 18284/python3
tcp        0      0 127.0.0.1:12345         127.0.0.1:4321          ESTABLISHED 4568/python3
tcp        0      0 127.0.0.1:80            127.0.0.1:12345         ESTABLISHED -
tcp        0      0 127.0.0.1:12345         127.0.0.1:80            ESTABLISHED 4568/python3
当然你要是连接相同的地址会报错的：

OSError: [Errno 99] Cannot assign requested address
那个五元组已经被占用啦。

同时创建连接：恰巧你也在这里
有时候，我们不能一个劲地等待。主动出击也是可以的，即便对方并没有在等待。

这个在 TCP 里叫「simultaneous open」，用于 TCP 打洞。但是比起 UDP 打洞难多了，因为那个「simultaneous」字眼：必须同时调用connect，双方的 SYN 包要交叉，早了或者晚了都是会被拒绝的。

所以手工就办不到啦，在本地测试也不容易办到。我本地的系统时间是使用 NTP 同步的，再用一个时钟也和 NTP 同步的 VPS 就可以啦，我这里延迟 80ms 左右，足够那两个 SYN 「在空中会面」了。以下是代码：

#!/usr/bin/env python3

import time
import sys
import socket
import datetime

def wait_until(t):
  deadline = t.timestamp()
  to_wait = deadline - time.time()
  time.sleep(to_wait)

s = socket.socket()
s.bind(('', 1314))

if sys.argv[1] == 'local':
  ip = 'VPS 的地址'
else:
  ip = '我的地址'

t = datetime.datetime(2015, 8, 19, 22, 14, 30)
wait_until(t)
s.connect((ip, 1314))

s.send(b'I love you.')
print(s.recv(1024))
当然，我是公网 IP。在内网里包就不容易进来啦。

然后双方在约定的时间之前跑起来即可，结果是这样子的：

# 本地
>>> python3 t.py local
b'I love you.'

# VPS 上
>>> python3 t.py remote
b'I love you.'
一个人也可以建立 TCP 连接呢
如果你没有 VPS，或者没有公网 IP，也是有活动可以参与的哦。即使只有一个 socket，也可以自己连接到自己的：

>>> import socket                                                               
>>> s = socket.socket()
>>> s.bind(('127.0.0.1', 1314))
>>> s.connect(('127.0.0.1', 1314))
>>> s.send(b'I love you.')
11
>>> s.recv(1024)
b'I love you.'
netstat 输出：

>>> netstat -npt | grep 1314
tcp        0      0 127.0.0.1:1314          127.0.0.1:1314  
