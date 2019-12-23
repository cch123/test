int shutdown(int sock, int howto);
  成功返回0，失败返回-1
  @sock  需要断开的套接字文件描述符。
  @howto 传递断开方式信息，可选值为`SHUT_RD/SHUT_WR/SHUT_RDWR`

服务器关闭输出流时向客户端传输EOF
