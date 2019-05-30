# /macros/exec/nbe
macro nbe {
    url = "https://www.nbe.com.eg/ExchangeRate.aspx"
    ttl = 120
    exec = <<JS
        exports = {
            usd: {
                buy: parseFloat($('.td1 #txtBanKNote_BuyPrice').Attr('value')),
                sell: parseFloat($('.td1 #txtBanKNote_SellPrice').Attr('value')),
            }
        }
    JS
}

# /macros/exec/egbank
macro egbank {
    url = "https://www.eg-bank.com/Ar/ExchangeRate"
    ttl = 120
    exec = <<JS
        exports = {
            usd: {
                buy: parseFloat($('.content_of_currency').First().Find('.currency').First().Text().split(' ')[1].trim())
                sell: parseFloat($('.content_of_currency').First().Find('.currency').Last().Text().split(' ')[1].trim()),
            }
        }
    JS
}

macro scraply {
    url = "https://github.com/alash3al/scraply"
    ttl = 120
    exec = <<JS
        exports = {
            title: $('title').Text(),
            description: $('meta[name="description"]').AttrOr('content', '')
        }
    JS
}

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
    projects = [
        "scraply",
        "sqler",
        "redix"
    ]
}