import pika
import sys

credentials = pika.PlainCredentials("guest", "guest")
conn_params = pika.ConnectionParameters("localhost", credentials=credentials)
conn_broker = pika.BlockingConnection(conn_params)
channel = conn_broker.channel()
# 注意下面这一句
channel.confirm_delivery()

channel.exchange_declare(exchange="hello-exchange",
                         type="direct",
                         passive=False,
                         durable=True,
                         auto_delete=False)


msg = sys.argv[1]

msg_props = pika.BasicProperties()

msg_props.content_type = "text/plain"

if channel.basic_publish(body=msg,
                         exchange="hello-exchange",
                         properties=msg_props,
                         routing_key="hola"):
    print "confirm received"
else:
    "message lost"

channel.close()
