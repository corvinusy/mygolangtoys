/**
 * Date: 04.06.13
 * Time: 1:56
 * To change this template use File | Settings | File Templates.
 */
package main

import ( "fmt";
	"os"
)

const YANDEX_SITE_STRING = "http://yandex.ru/yandsearch?lr=11091&numdoc=50&text=vaska";
const YANDEX_MASK_START="href=\"http://";
const YANDEX_MASK_END="\" onmousedown=\"rc(this, '//yandex.ru/clck/redir";

const WHOIS_SITE_STRING="http://api.robowhois.com/v1/whois/"
const WHOIS_START_MASK="Creation Date: "
const WHOIS_END_MASK=" "

//const FILE_NAME = "yandex_buffer.html"

func main() {

	var urls []string;
//	var strDate []string;

	siteContent, err := getSiteContent(YANDEX_SITE_STRING);
	if err != nil {
		fmt.Println(err.Error());
		os.Exit(1)
	}

	urls = parseSiteContent(siteContent, YANDEX_MASK_START, YANDEX_MASK_END);

	for i := 0; i<len(urls); i++ {
		siteContent, err = getSiteContent(WHOIS_SITE_STRING+urls[i]);
		fmt.Println(siteContent);
		strDate := parseSiteContent(siteContent, WHOIS_START_MASK, WHOIS_END_MASK);

		if strDate != nil {
			fmt.Println(urls[i] + "  Date: " + strDate[0]);
		} else {
			fmt.Println("Date of " + urls[i] + " not found");
		}

	}

	os.Exit(0)
}

