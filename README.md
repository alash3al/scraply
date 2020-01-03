Scraply
========
> Scraply a simple dom scraper to fetch information from any html based website using `jQuery` like syntax and convert that info to JSON APIs

How it works?
==============
> it works by simple define some `macros`/`endpoints` in `HCL` format, and let the magic begins, here is an example:
```hcl
# /scraply
macro scraply {
    // the url to scrap
    // we will scrap scraply github page and get information from it
    url = "https://github.com/alash3al/scraply"

    // cache [time to live] in seconds
    // set it to any value < 1 to disable it.
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
    // it extends the cronjob with an additional field in the first
    // to supports seconds.
    schedule = "* * * * * *"

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

    // also you have the following functions in js context
    // println()/console.log()
    // time() the current timestamp
    // sleep(ms) sleep the execution for x of milliseconds
    // macro(macro_name, paramsObject) executes the specified macro name and return its result
    // scraply.macro(...) an alias for macro(...)
    // scraply.params is an object containing the current request GET query params
}

# /sqler
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

# /redix
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

# aggregate ?
macro all {
    exec = <<JS
        exports = {
            redis: macro("redix"),
            sqler: macro("sqler")
        }
    JS
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

Features
========
- Tiny & Portable Engine.
- You can scale & distribute it easily.
- Private/Public Macros.
- Cron like scheduler.
- Webhook Support.
- jQuery like API.
- Customize everythin in javascript.

How?
====
- Download the binary that fits your OS from [here](https://github.com/alash3al/scraply/releases)
- Create a configuration file i.e `scraply.hcl`
- Run scrapply `./path/to/downloaded/scrapply --config=./scraply.hcl --listen=:9080`

Usage
=====
- Download binaries from [here](https://github.com/alash3al/scraply/releases)

```bash
# let's assume that you downloaded
# goto the the downloads directory
# then extract it
# then create your macros file i.e `example.scraply.hcl`
# then execute this command
$ ./scraply_linux_amd64
# you will see something like this, saying that everything is running under http://<yourhost>:9080/
⇨ http server started on [::]:9080
```