import scrapy
import datetime,time
from scrapy.spiders import CrawlSpider, Rule
from scrapy.linkextractors import LinkExtractor
from guolele.items import GuoleleItem
from scrapy.selector import Selector

class Guolele_spider(CrawlSpider):
    name = 'guolele'
    allowed_domains = ['tmall.guolele.com']
    start_urls = ['http://tmall.guolele.com/category/1/p1']

    rules = (
        Rule(LinkExtractor(allow=('.*/category/.*', ), deny=()), follow=True),
        Rule(LinkExtractor(allow=('.*/fruit/.*', )), callback='parse_item', follow=True),
    )

    def parse_item(self, response):
        sel = Selector(response)
        item = GuoleleItem()
        item['site_id'] = 22
        now = time.strftime("%Y-%m-%d %H:%M:%S")
        item['created_time'] = now
        item['updated_time'] = now
        item['name'] = sel.xpath('/html/body/div[1]/div[2]/ul/li/span[1]/text()').extract()[0]
        item['prod_id'] = sel.xpath('/html/body/div[1]/div[4]/ul/li[1]/text()').extract()[0]
        item['cates'] = sel.xpath('/html/body/div[1]/div[4]/ul/li[3]/text()').extract()[0]
        item['prod_desc'] = sel.xpath('/html/body/div[1]/div[6]').extract()[0]
        item['feature'] = sel.xpath('/html/body/div[1]/div[4]/ul/li[6]/text()').extract()[0]
        item['prop'] = sel.xpath('/html/body/div[1]/div[4]/ul/li[5]/text()').extract()[0]
        item['total_price'] = sel.xpath('/html/body/div[1]/div[2]/ul/li/span[2]/text()').extract()[0]
        item['price'] = sel.xpath('/html/body/div[1]/div[4]/ul/li[2]/text()').extract()[0]
        item['url'] = response.url
        #item['name'] = item['name'].encode('utf-8', 'ignore')
        # print item['name']
        text = item['name'].encode("utf-8", "ignore") + '&&' + item['prod_id'].encode("utf-8", 'ignore') + "\n"
        f = open("result.txt", "a+")
        f.write(text)
        f.close()
        return item

