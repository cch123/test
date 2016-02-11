# -*- coding: UTF-8 -*-
import smtplib
import datetime
from email.mime.text import MIMEText
#这个脚本要使用的话，必须有已经申请好的发送邮箱
#并且需要该邮箱打开了smtp发送服务，现在也就sina申请这样的邮箱比较方便

mailto_list=["cao1988228@163.com"]
mail_host="smtp.sina.com"  #设置服务器
mail_user="username"    #用户名
mail_pass="passwd"   #口令
mail_postfix="sina.com"  #发件箱的后缀

def send_mail(to_list,sub,content):
    me="<"+mail_user+"@"+mail_postfix+">"
    msg = MIMEText(content,_subtype='plain',_charset='gb2312')
    msg['Subject'] = sub
    msg['From'] = me
    msg['To'] = ";".join(to_list)
    try:
        server = smtplib.SMTP()
        server.connect(mail_host)
        server.login(mail_user,mail_pass)
        server.sendmail(me, to_list, msg.as_string())
        server.close()
        return True
    except Exception, e:
        print str(e)
        return False

if __name__ == '__main__':
    lastDate = datetime.date.today() - datetime.timedelta(days=1)
    dateStr = str(lastDate.year) + str(lastDate.month) + str(lastDate.day)
    title = dateStr

    content = ""
    #如果文件存在，那么需要改变content
    #if os.path.exists(report_file_path + dateStr):
    #python读取文件内容
    content = open(dateStr).read()
    #else :
    #content = "报表文件没有生成，请检查目录"

    if send_mail(mailto_list, title, content):
        print "发送成功"
    else:
        print "发送失败"
