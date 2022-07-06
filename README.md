Simple Scraping Tool
======================
> Scraply, is a very simple html scraping tool, if you know css & jQuery then you can use it!

Overview
========
> you can use `scraply` within your stack via `cli` or `http`.  
```bash
# here is the CLI usage

# extracting the title and the description from scraply github repo page
$ scraply extract \
    -u "https://github.com/alash3al/scraply" \
    -x title="select('title').text()" \
    -x description="select('meta[name=description]').attr('content')"

# same thing but with custom user agent
$ scraply extract \
    -u "https://github.com/alash3al/scraply" \
    -ua "OptionalCustomUserAgent"\
    -x title="select('title').text()" \
    -x description="select('meta[name=description]').attr('content')"

# same thing but with asking scraply to return the response body for debuging purposes
$ scraply extract \
    --return-body \
    -u "https://github.com/alash3al/scraply" \
    -x title="select('title').text()" \
    -x description="select('meta[name=description]').attr('content')"
```

> for `http` usage, we will run the http server then using any http client to interact with it.  
```bash
# running the http server
# by default it listens on address ":8010" which equals to "0.0.0.0:8010"
# for more information execute `$ scraply help`
$ scraply serve

# then in another shell let's execute the following curl 
$ curl http://localhost:8010/extract \
    -H "Content-Type: application/json" \
    -s \
    -d '{"url": "https://github.com/alash3al/scraply", "extractors": {"title": "$(\"title\").text()"}, "return_body": false, "user_agent": "CustomeUserAgent"}'
```

Download ?
==========
> you can go to the [releases page](https://github.com/alash3al/scraply/releases) and pick the latest version.


Contribution ?
==============
> for sure you can contribute, how?

- clone the repo
- create your fix/feature branch
- create a pull request

nothing else, enjoy!

About
=====
> I'm [Mohamed Al Ashaal](https://alash3al.com), a software engineer :)