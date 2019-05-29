Scraply
========
> Scraply a simple dom scraper to fetch information from any html based website and convert that info to JSON APIs

Why?
====
> I wanted a simple tool that fetches the required information in a simple and neat way from web pages, I'm using it in the following cases:

- Scraping data from currency rates websites
- Scraping pricing data from e-commerce sites
- Scraping news from news websites
- Scraping search data
- there are more use cases ...

How?
====
> just download the binary that fits your OS from [here](https://github.com/alash3al/scraply), then try to understand the following

```hcl
# /macros/exec/m1
macro m1 {
    // the url to scrap
    url = "http://webpage.url/here"

    // cache [time to live] in seconds
    ttl = 120

    // code to execute
    // this is a javascript code
    // you must set your returns in the exports variable
    // there is two global variables available there `document` and $
    // `document` is the DOM object you use to work with the DOM
    // `$` is like jQuery, it will help you to select anything from the document
    exec = <<JS
        exports = {
            usd: {
                buy: parseFloat($('.td1 #txtBanKNote_BuyPrice').Attr('value')),
                sell: parseFloat($('.td1 #txtBanKNote_SellPrice').Attr('value')),
            }
        }
    JS
}

# aggregators enables you to call multiple macros in just one call!
aggregators {
    # /aggregators/exec/all
    all = ["m1"]
}
```

DOM Methods
============
all methods from `document` or `$(..)` are similar to `jQuery` but with one exception, it is using `Golang` convention, for example

```js
//  jQuery              :   Scraply
//  -------------           ---------------
//  $(..).first()       :   $(..).First()
//  $(..).html()        :   $(..).Html()
//  $(..).text()        :   $(..).Text()
//  $(..).last()        :   $(..).Last()
//  $(..).find()        :   $(..).Find()
//  $(..).attr()        :   $(..).Attr()
//  $(..).children()    :   $(..).Children()
//  $(..).prev()        :   $(..).Prev()
//  $(..).next()        :   $(..).Next()
//  $(..).has()         :   $(..).Has()
```
