package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	str := `{"result":true,"code":0,"message":"success","data":[{"id":{"value":"fcaa4892-29dc-4981-b8ba-9e5ebf11c9bf-O30939"},"framework_id":{"value":"6a5f2b34-6c71-40c5-9488-8fd3eaca4dc2-0000"},"agent_id":{"value":"03e27556-9568-47b3-8b83-486b8d9dcc63-S1"},"hostname":"bcs-test-mesos-slave-2","url":{"scheme":"http","address":{"hostname":"bcs-test-mesos-slave-2","ip":"10.241.97.39","port":8080},"path":"/slave(1)"},"resources":[{"name":"cpus","type":0,"scalar":{"value":6.99},"role":"*"},{"name":"mem","type":0,"scalar":{"value":14570},"role":"*"},{"name":"disk","type":0,"scalar":{"value":891261},"role":"*"},{"name":"ports","type":1,"ranges":{"range":[{"begin":31000,"end":32000}]},"role":"*"}],"attributes":[{"name":"InnerIP","type":3,"text":{"value":"10.241.97.39"}},{"name":"ip-resources","type":0,"scalar":{"value":1}}],"executor_ids":[{"value":"1.deploy-nginx1.jiayuan.20000.1534423031569104026"}]},{"id":{"value":"fcaa4892-29dc-4981-b8ba-9e5ebf11c9bf-O30940"},"framework_id":{"value":"6a5f2b34-6c71-40c5-9488-8fd3eaca4dc2-0000"},"agent_id":{"value":"03e27556-9568-47b3-8b83-486b8d9dcc63-S0"},"hostname":"bcs-test-mesos-slave-1","url":{"scheme":"http","address":{"hostname":"bcs-test-mesos-slave-1","ip":"10.241.97.28","port":8080},"path":"/slave(1)"},"resources":[{"name":"cpus","type":0,"scalar":{"value":4.98},"role":"*"},{"name":"mem","type":0,"scalar":{"value":14498},"role":"*"},{"name":"disk","type":0,"scalar":{"value":891197},"role":"*"},{"name":"ports","type":1,"ranges":{"range":[{"begin":31000,"end":32000}]},"role":"*"}],"attributes":[{"name":"InnerIP","type":3,"text":{"value":"10.241.97.28"}},{"name":"ip-resources","type":0,"scalar":{"value":1}}],"executor_ids":[{"value":"0.deploy-nginx1.jiayuan.20000.1534423031430495401"},{"value":"0.test123-deployment-1.bellke-test2.20000.1533220199742403112"}]}]}`

	str1 := `{"resources": [{"name":"cpus"},{"name":"mem"}]}`

	type res struct {
		Name   string
		Scalar map[string]interface{} `json:"scalar"`
	}

	type atts struct {
		Name   string
		Scalar map[string]interface{} `json:"scalar"`
	}

	type data struct {
		Resources  []res
		Attributes []atts
	}
	type respBody struct {
		Result  bool   `json:"result"`
		Code    uint   `json:"code"`
		Message string `json:"message"`
		Data    []data `json:"data"`
	}
	respContent := respBody{}
	err := json.Unmarshal([]byte(str), &respContent)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(respContent)

	type Test struct {
		Resources []res
	}
	var test1 Test
	err = json.Unmarshal([]byte(str1), &test1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test1)
}
