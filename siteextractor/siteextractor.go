/**
 * Created with IntelliJ IDEA.
 * User: corvinus
 * Date: 04.06.13
 * Time: 6:55
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"net/http";
	"net/http/httputil";
	"strings";
	"errors";
)

func getSiteContent(siteString string) (string, error) {

	response, err := http.Get(siteString)
	if err != nil {
		return "", err
	}

	if strings.ToLower(response.Status) != "200 ok" {
		err = errors.New("Site is not responding " + response.Status)
		return "", err
	}

	_, err = httputil.DumpResponse(response, false);
	if err != nil {
		return "", err
	}

	contentType := response.Header["Content-Type"]

    if !checkContents(contentType, "utf-8") {
		err = errors.New("Cannot handle that content type");
		return "", err
	}

	var buf [1024]byte
	var siteCont string

	reader := response.Body
	for {
		n, err := reader.Read(buf[0:])
		if err != nil && err.Error() != "EOF" { return "", err }

		if n == 0 { break }
		siteCont = siteCont + string(buf[0:]);
	}

	return siteCont, nil
}

func checkContents (srcStrings []string, content string) bool {
	// each site should content like [text/html; charset=UTF-8]
	// we want the UTF-8 only
	for _, sS := range srcStrings {
		if strings.Index(strings.ToLower(sS), content) != -1 {
			return true
		}
	}
	return false
}


