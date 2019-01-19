package tests

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"log"
	"testing"
)

func TestGzip(t *testing.T) {
	var b bytes.Buffer
	gz, _ := gzip.NewWriterLevel(&b, gzip.HuffmanOnly)
	if _, err := gz.Write([]byte(`eyJhbGciOiJSUzM4NCIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIxIiwiVXNlck5hbWUiOiIiLCJFbWFpbCI6Ijc3NDU3MzIwOUBxcS5jb20iLCJQaG9uZSI6IjE2Njc4NjA1NTM2IiwiSWQiOjEsIkF2YXRhciI6IiJ9.cVXWQtUXHkCWHmG2LN4Pgfzx-94fpbpEo_PPt9a6mHi94yBrscwPO3IHg7RXheaFMs3z99XSOG0ZCiYlttkfMU4u8xKyT07mWrAt9kqayAQlAqfbFZ8vnoU0_eIY9_bTTFE6IUhSI5EZNR_2QSyfzavVyFqanUrEW7hobmm9St_dlcrvjwyaVtyKTsS7pmLr19m2RDOdJVAs-dD4_tx1ZfWG1ASAtSQpUD9gNSMaLru_DA09Bq7d_bwZ1YNiMXbDAqaFULkEivwpXDJD-WLEF1AO6QQB6y6-_BUPSaTp_5piOIFGS_GzuwGrjYgUeQt0iylZ-KXdIduni-7D3os5-A`)); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	str := base64.StdEncoding.EncodeToString(b.Bytes())
	fmt.Println(str)
	//data, _ := base64.StdEncoding.DecodeString(str)
	//fmt.Println(data)
	//rdata := bytes.NewReader(data)
	//r, _ := gzip.NewReader(rdata)
	//s, _ := ioutil.ReadAll(r)
	//fmt.Println(string(s))

	// 一个缓存区压缩的内容
	buf := bytes.NewBuffer(nil)

	// 创建一个flate.Writer
	flateWrite, err := flate.NewWriter(buf, flate.BestCompression)
	if err != nil {
		log.Fatalln(err)
	}
	defer flateWrite.Close()
	// 写入待压缩内容
	flateWrite.Write([]byte("eyJhbGciOiJSUzM4NCIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIxIiwiVXNlck5hbWUiOiIiLCJFbWFpbCI6Ijc3NDU3MzIwOUBxcS5jb20iLCJQaG9uZSI6IjE2Njc4NjA1NTM2IiwiSWQiOjEsIkF2YXRhciI6IiJ9.cVXWQtUXHkCWHmG2LN4Pgfzx-94fpbpEo_PPt9a6mHi94yBrscwPO3IHg7RXheaFMs3z99XSOG0ZCiYlttkfMU4u8xKyT07mWrAt9kqayAQlAqfbFZ8vnoU0_eIY9_bTTFE6IUhSI5EZNR_2QSyfzavVyFqanUrEW7hobmm9St_dlcrvjwyaVtyKTsS7pmLr19m2RDOdJVAs-dD4_tx1ZfWG1ASAtSQpUD9gNSMaLru_DA09Bq7d_bwZ1YNiMXbDAqaFULkEivwpXDJD-WLEF1AO6QQB6y6-_BUPSaTp_5piOIFGS_GzuwGrjYgUeQt0iylZ-KXdIduni-7D3os5-A"))
	flateWrite.Flush()
	fmt.Printf("压缩后的内容：%s\n", string(buf.Bytes()))

}