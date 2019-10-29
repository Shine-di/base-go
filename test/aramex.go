//@author D-S
//Created on 2019/10/24 11:50 上午
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const  (
	ARAMEX_URL = "https://ws.dev.aramex.net/shippingapi/tracking/service_1_0.svc"
     USER_NAME = "testingapi@aramex.com"
    PASSWORD = "R123456789$r"
     VERSION =   "v1"
     ACCOUNT_NUMBER =  "20016"
     ACCOUNT_PIN =  "331421"
     ACCOOUNT_ENTITY =   "AMM"
     ACCOUNT_COUNTRY_CODE =  "JO"
)

type (
	ClientInfo struct {
		UserName string `xml:"UserName"`
		Password  string  `xml:"Password"`
		Version string  `xml:"Version"`
		AccountNumber  string  `xml:"AccountNumber"`
		AccountPin  string  `xml:"AccountPin"`
		AccountEntity  string  `xml:"AccountEntity"`
		AccountCountryCode  string  `xml:"AccountCountryCode"`
	}
	Transaction struct {
		Reference1  string `xml:"Reference1"`
		Reference2  string `xml:"Reference2"`
		Reference3  string `xml:"Reference3"`
		Reference4  string `xml:"Reference4"`
		Reference5  string `xml:"Reference5"`
		Reference6  string `xml:"Reference6"`
	}

	TrackIngRequest struct {
		ClientInfo ClientInfo  `xml:"ClientInfo"`
		Transaction Transaction    `xml:"Transaction"`
		Shipments   []string  `xml:"Shipments"`
		GetLastTrackingUpdateOnly bool  `xml:"GetLastTrackingUpdateOnly"`
	}


	Notifications struct {
		Code string  `xml:"Code"`
		Message  string `xml:"Message"`
	}
	TrackingResults struct {
		WayBillNumber string `xml:"WayBillNumber"`
		UpdateCode  string  `xml:"UpdateCode"`
		UpdateDescription string  `xml:"UpdateDescription"`
		UpdateDateTime time.Time  `xml:"UpdateDateTime"`
		UpdateLocation  string  `xml:"UpdateLocation"`
		Comments   string   `xml:"Comments"`
		ProblemCode string   `xml:"ProblemCode"`
	}

	TrackIngResponse struct {
		Transaction Transaction     `xml:"Transaction"`
		Notifications Notifications   `xml:"Notification"`
		HasErrors  bool `xml:"HasErrors"`
		TrackingResults TrackingResults  `xml:"TrackingResults"`
	}
)

func DoARAMEX() error {
	request := TrackIngRequest{
		//Shipments:                 nil,
		GetLastTrackingUpdateOnly: true,
		Transaction:Transaction{
			Reference1: "123",
			Reference2: "123",
			Reference3: "13",
			Reference4: "1312",
			Reference5: "13",
			Reference6: "1312",
		},
		Shipments: []string{
			"12312",
			"123123",
		},
	}
	request.ClientInfo.UserName = USER_NAME
	request.ClientInfo.Password = PASSWORD
	request.ClientInfo.AccountCountryCode = ACCOUNT_COUNTRY_CODE
	request.ClientInfo.AccountNumber  = ACCOUNT_NUMBER
	request.ClientInfo.AccountEntity = ACCOOUNT_ENTITY
	request.ClientInfo.AccountPin = ACCOUNT_PIN
	request.ClientInfo.Version = VERSION

	response,err := Post(ARAMEX_URL,request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		return err
	}
	fmt.Println(response)
	return nil
}


func Post(url string, body interface{}) ( *TrackIngResponse,error) {
	bodyByte, _ := xml.Marshal(body)
	requestBody := bytes.NewBuffer(bodyByte)
	fmt.Println(string(bodyByte))
	rsp, err := http.Post(url, "text/xml; charset=utf-8", requestBody)
	if err != nil {
		return nil,err
	}
	fmt.Println(rsp)
	rb, _ := ioutil.ReadAll(rsp.Body)
	var resBody TrackIngResponse
	if err := xml.Unmarshal(rb, &resBody); err != nil {
		return  nil,err
	}
	return  &resBody,nil
}

func main()  {
	DoARAMEX()
}