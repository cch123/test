本目录记录了一些在配置elasticsearch时需要修改的配置

要点有以下一些：

###1.host绑定

在es 2.0里，需要设置
```bash
在bin目录下的elasticsearch需要修改
ES_HEAP_SIZE=8g
JAVA_OPTS=" -Dcom.sun.management.jmxremote.port=9201 -Dcom.sun.management.jmxremote.ssl=false -Dcom.sun.management.jmxremote.authenticate=false"

在config目录下的elasticsearch.yml
network.host: 0.0.0.0
注意：不设置这个的话会无法在本地和远端同时访问
```

在1.7里也需要堆内存的限制，不过不太需要设置host

###2.插件安装

集群管理工具
>bin/plugin install mobz/elasticsearch-head

集群监控工具，可以通过它来查看es集群的各种状态，如：cpu、内存使用情况，索引数据、搜索情况，http连接数等。
>bin/plugin install lukas-vlcek/bigdesk

###3.jConsole调试
在上面的JAVA_OPTS选项里指定了jmxremote的端口，可以在本机用jConsole连接过去来查看应用程序的内存使用情况

刚开始部署的时候碰到了用jdbc导入数据时会频繁触发gc并stop-the-world的坑，后来发现是elasticsearch jdbc也需要设置一下使用时的运行内存，添加到jdbc运行时的指令里即可，参数可以动态调整

```
-Xms2g -Xmx2g -Djava.awt.headless=true -XX:+UseParNewGC -XX:+UseConcMarkSweepGC -XX:CMSInitiatingOccupancyFraction=75 -XX:+UseCMSInitiatingOccupancyOnly -XX:+HeapDumpOnOutOfMemoryError -XX:+DisableExplicitGC -Dfile.encoding=UTF-8 \
```

调整之后在导入百万数据时就不怎么会触发gc了~
