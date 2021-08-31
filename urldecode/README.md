### INSTALL

#### Build 

`go build urldecode.go`

### Example

`echo "GET /query?db=statsd&epoch=s&q=SELECT+last%28%22value%22%29+FROM+%221minute%22.%22custom_fax_sent_faxSentOrderFailures%22+WHERE+time+%3E+now%28%29+-+15m+GROUP+BY+time%281m%29%2C+%22type%22+fill%28null%29 HTTP/1.1" | urldecode`

Should see this.

`GET /query?db=statsd&epoch=s&q=SELECT+last("value")+FROM+"1minute"."custom_fax_sent_faxSentOrderFailures"+WHERE+time+>+now()+-+15m+GROUP+BY+time(1m),+"type"+fill(null) HTTP/1.1`
