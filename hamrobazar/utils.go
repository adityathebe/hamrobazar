package hamrobazar

import (
	"fmt"
	"strconv"
	"strings"
)

func formatMessage(title string, price float64, location string, creator Creator, summary string, url string) string {
	msg := fmt.Sprintf("<b>%s</b>\n\n", title)
	msg += fmt.Sprintf("Rs %s\n", numberWithCommas(price))
	msg += fmt.Sprintf("<b>Seller</b>: %s (%d ads)\n", creator.CreatedByName, creator.CreatorAdCount)
	msg += fmt.Sprintf("<b>Location</b>: %s\n\n", location)

	summaryLenLimit := 1024 - len(msg) - len(url) - len("<pre></pre>\n\n")
	if summaryLenLimit > len(summary) {
		summaryLenLimit = len(summary)
	}

	msg += fmt.Sprintf("<pre>%s</pre>\n\n", summary[:summaryLenLimit])
	msg += url
	return msg
}

func formatImgURL(imgURL string) string {
	suffix := "?x-image-process=image/resize,m_lfit,h_500,w_500"
	if imgURL[0] == '/' {
		return "https://hamrobazaar.com/layout_images/noimage_100.gif" + suffix
	}

	return strings.Replace(imgURL, "_small.jpg", ".jpg", 1) + suffix
}

func getItemURL(id string) string {
	return fmt.Sprintf("https://hamrobazaar.com/for-rent-house/does-not-matter/%s", id)
}

func numberWithCommas(x float64) string {
	return strconv.FormatInt(int64(x), 10)
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false
}
