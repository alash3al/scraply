Scraply
========
> Scraply a simple dom scraper to fetch information from any html based website and convert that info to JSON APIs

Overview
=======
```hcl
# this is the configurations file [sample]

# /macros/exec/scraply
macro scraply {
    // the url to scrap
    // we will scrap scraply github page and get information from it
    url = "https://github.com/alash3al/scraply"

    // cache [time to live] in seconds
    ttl = 120

    // code to be executed
    //
    // this is a javascript code
    // you must set your returns in the exports variable
    exec = <<JS
        exports = {
            // fetching the title
            // similar to jQuery, right?
            title: $("title").Text(),
            description: $('meta[name=description]').AttrOr('content', '')
        }
    JS

    // schedule this macro to run at the specified cron style spec
    schedule = "* * * * *"

    // notify an endpoint with the result
    // the payload is a json object just like: {"error": "an error if any", "result": "the result will be here"}
    webhook = "http://some.endpoint.com"

    // whether you don't want to expose this macro to the api or not
    private = true

    // our $(..).Method() is just like jQuery's $(..).method()
    // our $(..).Method() is an alias for document.Find(..).Method()
    // 
    // here is a table shows you jQuery methods and scraply Methods:
    //
    //  jQuery              :   Scraply
    //  -------------           ---------------
    //  $(..).first()       :   $(..).First()
    //  $(..).html()        :   $(..).Html()
    //  $(..).text()        :   $(..).Text()
    //  $(..).last()        :   $(..).Last()
    //  $(..).find()        :   $(..).Find()
    //  $(..).attr()        :   $(..).Attr() | $(..).AttrOr(needle, defaultValue)
    //  $(..).children()    :   $(..).Children()
    //  $(..).prev()        :   $(..).Prev()
    //  $(..).next()        :   $(..).Next()
    //  $(..).has()         :   $(..).Has()
}

# /macros/exec/sqler
macro sqler {
    url = "https://github.com/alash3al/sqler"
    ttl = 120
    exec = <<JS
        exports = {
            title: $('title').Text(),
            description: $('meta[name="description"]').AttrOr('content', '')
        }
    JS
}

# /macros/exec/redix
macro redix {
    url = "https://github.com/alash3al/redix"
    ttl = 120
    exec = <<JS
        exports = {
            title: $('title').Text(),
            description: $('meta[name="description"]').AttrOr('content', '')
        }
    JS
}

# aggregators enables you to call multiple macros in just one call!
aggregators {
    # /aggregators/exec/projects
    projects = [
        "scraply",
        "sqler",
        "redix"
    ]
}
```

Why?
====
> I wanted a simple tool that fetches the required information in a simple way from web pages, I'm using it in the following cases:

- Scraping data from currency rates websites
- Scraping product pricing data from e-commerce sites
- Scraping news from news websites
- Scraping search data
- there are more use cases ...

How?
====
- Download the binary that fits your OS from [here](https://github.com/alash3al/scraply/releases)
- Create a configuration file i.e `scraply.hcl`
- Run scrapply `./path/to/downloaded/scrapply --config=./scraply.hcl --listen=:9080`
