# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# http://doc.scrapy.org/en/latest/topics/items.html

import scrapy


class GuoleleItem(scrapy.Item):
    # define the fields for your item here like:
    # name = scrapy.Field()
    name = scrapy.Field()
    site_id = scrapy.Field()
    prod_id = scrapy.Field()
    cates = scrapy.Field()
    prod_desc = scrapy.Field()
    feature = scrapy.Field()
    origin = scrapy.Field()
    price = scrapy.Field()
    prop = scrapy.Field()
    size = scrapy.Field()
    total_price = scrapy.Field()
    url = scrapy.Field()
    created_time = scrapy.Field()
    updated_time = scrapy.Field()
    pass
