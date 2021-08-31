// Description: Decode URL's
// e.g. echo "GET /query?db=statsd&epoch=s&q=SELECT+last%28%22value%22%29+FROM+%221minute%22.%22custom_fax_sent_faxSentOrderFailures%22+WHERE+time+%3E+now%28%29+-+15m+GROUP+BY+time%281m%29%2C+%22type%22+fill%28null%29 HTTP/1.1" | urldecode

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"net/url"
)

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	str := string(bytes)
	path := (str)
	unescapedPath, err := url.PathUnescape(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unescapedPath)
}
