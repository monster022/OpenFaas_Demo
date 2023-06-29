package function

import (
	"bytes"
	"fmt"
	"net/http"
)

// Handle a serverless request
func Handle(req []byte) string {
	jsonStr := []byte(fmt.Sprintf(`{
		"mob": %s,
  		"msg": "尊敬的 李华 您已成功预订中国联航KN5737上海虹桥T2\t-缅甸丹老机场，05月29日07:25起飞11:25到达,行李额0KG,请您提前120分钟到达机场,祝您一路平安！"
	}`, string(req)))
	xx, _ := http.NewRequest("POST", "http://message.dev/Msg/SendSms", bytes.NewBuffer(jsonStr))
	xx.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(xx)
	return fmt.Sprintf("Hello, Go. You said: %s", resp.Status)
}
