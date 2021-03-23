package client

import (
	"encoding/xml"
	"fmt"
	"github.com/gotomicro/gopay/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWechatDecode(t *testing.T) {
	data := `
<xml><return_code>SUCCESS</return_code><appid><![CDATA[wx99afd8defe4a1f74]]></appid><mch_id><![CDATA[1580519131]]></mch_id><nonce_str><![CDATA[fefeb22c3bf218cda55bf7bd4960824e]]></nonce_str><req_info><![CDATA[qEICGbF9wN5GeTvF3BKHBZOVaSgUWriefjL8GWvjo03T3YRXYnehaL+VsOIrewVWDEr3KOxwknpfI3VTRoTA4sN391Ar7agh1sKL1LQpzyiHh2lq4bXn/gxm9qjTTbFfrHzsRjVIngYahk8hB89BK3qGt9aKuwIdn77R7UyT8udYXRuXCrfpnBn6H4mRQgMblZEAcQRUrahBLhljcFRSkRI/4EJKWz8G+TpqWZsl8Wih7Xlnj6MJfBfMyNcmCH2Hb4GKwnhazjbBrHKchM+h30UvarO3ajTHtrNC/SjzRL2x7cyGh15DNrPVdO1OrHFToR5KBFBeJ1h+xF09MuNW/kAIwWJsIhjnFtTsXTz6LoKVYwLhpG2Q708xw8hOnIclBiscrcPqsKflaIIiwpR4Hu2RpJ17yuhXewmP8C1DpIvNekilQ9WRg5liE7rXo8K1lS3ty+0zBY+2CdlBZT3kQpYfyojUyW3YKF6UzrxwdbTiHsHVZEcGFwKjvHjDurGCPZUfX2UpQOQ1PQAldLW/aY+9+EIiZFNkLVQdCWrxBepCZ3sjaz/akytlf/EvZqEEFravH/IGQd5VCiNJ8wN28/yQjhQ7k9MgUzzqr63VvkLrbn0ycjrnvKxvmRArST1DVJhqGNnj7hNlzTnW2jnU80Swjz79LjYVF8RB29/u7WrY32L6bEBrCS0eptXjFko1CBGP5s2gPNKs96aP1mNqQDzCxHyv+qoaJT0vK+cTCOYcIOkK6qR8d1VvXpYpR0FFAKrRvmCa42PDxGwAocO1Af1pqykIscEIqCr0eTeRWKC0WljBdPjE2WPPrktsK1wJKmI3nCcIkPFUhmG8sgTOxj9pW2KgC+UpoG1oET3XsiprT31+s14ghZcWWVJEwU2aWp5H2hm7NAsH+bCW6V9kZCUcFndBEWH6B5ns1uwKX+SgUOdBZUbk1JMulJ5UOpDXGhI6EKWjQzzgBIrrsIIk1XAy5BtXODPxy3TbtblInWxCfwOLRrZBWqzH2ggS/WbeOFsAarOaNXFLdetMzOvUoJLX9lduKGHYbn+V5DIt0TfCtq21ciuA3g0r2bLWlUflwerfbVveWtY45tuXvWZHNA==]]></req_info></xml>
`
	wechatRes := &common.WechatRefundResultOriginalResp{}
	err := xml.Unmarshal([]byte(data), wechatRes)
	assert.Nil(t, err)

	decodeStr, err := WechatDecode(wechatRes.ReqInfo, "xxx")
	assert.Nil(t, err)
	fmt.Printf(string(decodeStr))

}
