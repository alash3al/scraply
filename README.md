Scraply
========
> Scraply, is a very simple html scraping tool, if you know css & jQuery then you can use it!, scraply should be simple and tiny as well it could be used as a component in a large system something like [this use-case](https://github.com/alash3al/scraply/issues/1#issuecomment-570082314)

Overview
========
> you can use `scraply` within your stack via `cli` or `http`.  
```bash
# here is the CLI usage

# extracting the title and the description from scraply github repo page
$ scraply extract \
    -u "https://github.com/alash3al/scraply" \
    -x title='$("title").text()' \
    -x description='$("meta[name=description]").attr("content")'

# same thing but with custom user agent
$ scraply extract \
    -u "https://github.com/alash3al/scraply" \
    -ua "OptionalCustomUserAgent"\
    -x title='$("title").text()' \
    -x description='$("meta[name=description]").attr("content")'

# same thing but with asking scraply to return the response body for debugging purposes
$ scraply extract \
    --return-body \
    -u "https://github.com/alash3al/scraply" \
    -x title='$("title").text()' \
    -x description='$("meta[name=description]").attr("content")'
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

> for debugging, there is `shell`
```bash
$ scraply shell -u https://github.com/alash3al/scraply
➜ (scraply) > $("title").text()
GitHub - alash3al/scraply: Scraply a simple dom scraper to fetch information from any html based website and convert that info to JSON APIs

➜ (scraply) > request.url
https://github.com/alash3al/scraply

➜ (scraply) > response.status_code
200

➜ (scraply) > response.url
https://github.com/alash3al/scraply

➜ (scraply) > response.body
<html>.....
```

Download ?
==========
> you can go to the [releases page](https://github.com/alash3al/scraply/releases) and pick the latest version.
> or you can `$ docker run --rm -it ghcr.io/alash3al/scraply scraply help`


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