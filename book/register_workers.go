package book

import (
	"../crawlsave"
)

func RegisterWorkers() {

	crawlInfo := crawlsave.CrawlNode{}

	crawlInfo.COLL_NAME = "books"
	crawlInfo.REDIS_SET_NAME = "saved_books"
	crawlInfo.QUEUE_NAME = "resque:queue:booksavedb"
	crawlInfo.QUEUE_NEXT = "resque:queue:bookinfocrawl"
	crawlInfo.CLASS_NAME = "Booksavedb"
	crawlInfo.CLASS_NEXT = "Bookinfocrawl"
	crawlInfo.CLASS_PARSE = "Bookcrawl"

	crawlInfo.Scrape = Scrape
	crawlInfo.GetMongoObj = getMongoObj

	crawlsave.RegisterWorkers(&crawlInfo)

}
