/**
 * Date: 04.06.13
 * Time: 4:42
 * To change this template use File | Settings | File Templates.
 */
package main

import ("fmt"
)


func parseSiteContent (siteContent , startMask, endMask string) ([]string) {

	var urls []string;
	var url string;
//	var cleanurls []string;

	//searching on mask href="http://<sitename>/"  onmousedown="rc(this, '//yandex.ru/clck/redir
	// then extract <sitename> and put it to map or set
	for i:=0; i < len(siteContent) - len(endMask); i++ {
		if siteContent[i:i+len(startMask)] == startMask {
			//found startmask, fixing 'i-counter', and seeking endmask
			for j:=i; j < len(siteContent) - len(endMask); j++ {
				if siteContent[j] == ' '{
					break
				}
				if siteContent[j:j+len(endMask)] == endMask {
					//it seems we found url between i and j
					url = siteContent[i+len(startMask):j];
					url = cleanURL(url);
					if !hasAlwaysURL(urls, url) {
						urls = append(urls, url); //memory expensive part for gc - optimizations are welcome
					}
				}
			}
		}
	}

	//here we going to clean found urls from subdirectories

	for i:=0; i<len(urls); i++ {
	    fmt.Println(urls[i]);
	}
	return urls
}


func cleanURL(url string) string {
	for i := 0; i<len(url); i++ {
		if url[i] == '/' {
			url = url[0:i];
			break
		}
	}
	return url
}

func hasAlwaysURL (urls []string, url string) bool {
	for i := 0; i < len(urls); i++ {
		if urls[i] == url {
			return true
		}
	}
	return false
}
