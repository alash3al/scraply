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

# aggregators enables you to call multiple macros in just one call!
aggregators {

    # /aggregators/exec/all
    all = ["nbe", "egbank"]

}