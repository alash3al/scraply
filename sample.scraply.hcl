macro scraply {
    url = "https://github.com/alash3al/scraply"
    ttl = 120
    exec = <<JS
        exports = {
            title: $('title').Text(),
            description: $('meta[name=description]').AttrOr('content', '')
        }
    JS
}

macro sqler {
    url = "https://github.com/alash3al/sqler"
    ttl = 120
    exec = <<JS
        exports = {
            title: $('title').Text(),
            description: $('meta[name=description]').AttrOr('content', '')
        }
    JS
}

macro redix {
    url = "https://github.com/alash3al/redix"
    ttl = 120
    exec = <<JS
        exports = {
            title: $('title').Text(),
            description: $('meta[name=description]').AttrOr('content', '')
        }
    JS
}

macro all {
    exec = <<JS
        // console.log(scraply.params)
        exports = {
            redis: scraply.macro("redix", scraply.params),
            sqler: scraply.macro("sqler", scraply.params)
        }
    JS
}
