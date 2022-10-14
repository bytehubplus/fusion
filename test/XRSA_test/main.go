// Use the XRSA library to encrypt DID documents. Before encryption,
// need to know the key of the receiver to complete the double encryption.
package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"os"
	"xrsa"

	did "github.com/bytehubplus/fusion/did"
)

var (
	//go:embed test/did1.json
	did1Json string
)

func main() {
	file, err := os.Open("private.pem")
	defer file.Close()
	if err != nil {
		log.Println(err)
		return
	}
	//获取文件信息
	info, err := file.Stat()
	if err != nil {
		log.Println(err)
		return
	}
	//创建切片，用于存储文件中读取到的公钥信息
	privateKey := make([]byte, info.Size())
	//读取公钥文件
	file.Read(privateKey)
	file.Close()

	file, err = os.Open("public.pem")
	defer file.Close()
	if err != nil {
		log.Println(err)
		return
	}
	//获取文件信息
	info, err = file.Stat()
	if err != nil {
		log.Println(err)
		return
	}
	//创建切片，用于存储文件中读取到的公钥信息
	publicKey := make([]byte, info.Size())
	//读取公钥文件
	file.Read(publicKey)
	file.Close()

	const publicKey_Bob = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlmnDvTCGosp52aXJXEOd
ZPb3rDmbk9oa+D7dtSKf9/aBqD4Ptowl12wZX+viqE/Kq+/zKzVsxWbOCecpJv5f
20m20ZtV3BhuE40WvCpOSVyivnXtId7VH+aw4enq24DZY4umEfY3JNxd1j25mLlB
c8yblvnjs/i4jG56urA5IL1yDWDs6TOu2I9AmHPjYu5a+eErGHOWFfDeW0K/OQaL
eE8Uqzi2WGajGUGU8jE58v3nyGMhH+x6kMuVl9NdEKzyGmrcEIiQF0H6uPY6e1xy
WI7j3cv/0nUtofRNW/0/5MnBrcTERWrsNQrn0j7uOPrtUuF8OdE+KImULq9oeUGX
ZQIDAQAB
-----END PUBLIC KEY-----
`
	const privateKey_Bob = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCWacO9MIaiynnZ
pclcQ51k9vesOZuT2hr4Pt21Ip/39oGoPg+2jCXXbBlf6+KoT8qr7/MrNWzFZs4J
5ykm/l/bSbbRm1XcGG4TjRa8Kk5JXKK+de0h3tUf5rDh6erbgNlji6YR9jck3F3W
PbmYuUFzzJuW+eOz+LiMbnq6sDkgvXINYOzpM67Yj0CYc+Ni7lr54SsYc5YV8N5b
Qr85Bot4TxSrOLZYZqMZQZTyMTny/efIYyEf7HqQy5WX010QrPIaatwQiJAXQfq4
9jp7XHJYjuPdy//SdS2h9E1b/T/kycGtxMRFauw1CufSPu44+u1S4Xw50T4oiZQu
r2h5QZdlAgMBAAECggEAUwP+x41f0btksy5gS276EL6KBeEpr9nB5t9zqER6+/Vu
rDfMnlkNja8Y9isPxwt69ZiSondzGCRcdXTC7sWYjERMBbXxFm/ZWSsWsDW9TZo2
LF6TyYzeHiRJ3fYn7IxZ7yolN2aoGs0RcWxR4ivlJw93lEVJWoxc9w0G6cDXVuy/
x1RVZvFLfHdBwlJhUjDMDtI1m+qlz2CBJRvlbZCkJDHAs3U8IzWRxeMt5AUQNHbq
gCJFpS0IO5CR5hazuj2l2sEho+qbO4aSYC/AB4SEm8ptBrY0TA+SBM7hQKNcOrZG
h7DHiinKMrXAtS73iHViF8vYTwjpYR+9sBVB8GMlgQKBgQDF+tihbwb6k+Wihqoz
zusl+6Yk4MdfX04v9tV16DBkaBkCmSVv1+se0UPbPDIEfM2zx6YNs5YrXdsavJU1
OYltEmql3kenER7DUUOZ/7J3xwbiUS0E29yiAPoJl1NEf49Crawf9ukEgV9b8V2A
sOn6AiGKXaWD4c8NanekJAO8xQKBgQDCfksj3qCsRbgMAAW4PKsY+Ar+Tc8tsH8/
YnHlqVGg7L1UTOAktCZ3dBzPtNRWwB5yKgXwsqWx2RIdrwLj1xnHxOgCzmjclK/N
8B9ka+t1d4rfrOn6/yy7wMqpPAzFRfiA+3jzYkeVgBNVup84THXNAsGZo9NIy6Zx
HE/Q61VaIQKBgQC2nA5usLsOK5aW14FsMgJBUaFIyK/87ypuIU146MaiBkZzWBDo
3Y64KSter2IvM3KEzbUDVE9CBtsPCTzTEBQLL/6AqcsLdUYbv/wLsobJ5iEaZOeS
YL1cDGyUpiieuM3KIejuvs7lYfM0Gig1iHj4KjkHqCL/xys046whETsFRQKBgEaM
QP0119olP+k6aNi3SOi5sGMmHixYhQ+cyeugSzewJ8zYRcWKgxl5SKoe0ZbWATUw
AQ+Z/XhRYLzbqpSyhWttBStfa7H3TP3KR9jJPJNTVVBlDEo7HSLNf7V7TraP9t/y
V6EQ2R5lYp4KH47wTRqcEVWFkzIojKsiteD2O/OBAoGACo1bnC4CJo9zpXVB/no1
+TUJSAXdCBThxJ/n6ADiBJ7e4nixO7EjOzdRLkFSY/XqB7Qob6sz0oIKsLdzIYzI
HU2/x53f1H0Kps+ViCEeV7v9dlPzpw9KbG4HaPnfOJIOYlD3mmzSSd245TRrqWpY
bjluVAchFOtRbBKv7QCsz88=
-----END PRIVATE KEY-----	
`

	//RsaGenKey(2048)
	var doc did.Document
	json.Unmarshal([]byte(did1Json), &doc)
	rawData, _ := json.Marshal(doc)
	//log.Println(rawData)
	d := &did.Document{}
	_ = json.Unmarshal(rawData, d)
	log.Println(d.ID)

	//fmt.Printf("%q\n", publicKey)
	//log.Println(string(privateKey))
	log.Println(publicKey)
	//log.Println(privateKey)
	if err != nil {
		log.Println(err)
		return
	}

	xrsa_Alice, err := xrsa.NewXRsa(publicKey, privateKey)
	if err != nil {
		log.Println(err)
		return
	}

	xrsa_Bob, err := xrsa.NewXRsa([]byte(publicKey_Bob), []byte(privateKey_Bob))
	if err != nil {
		log.Println(err)
		return
	}

	data := string(rawData)
	encrypted, _ := xrsa_Alice.PrivateEncrypt(data)
	//log.Println(encrypted)

	encrypted_Bob, _ := xrsa_Bob.PublicEncrypt(encrypted)
	//log.Println(encrypted_Bob)
	decrypted_Bob, _ := xrsa_Bob.PrivateDecrypt(encrypted_Bob)

	decrypted, _ := xrsa_Alice.PublicDecrypt(decrypted_Bob)
	//log.Println(reflect.TypeOf(decrypted))
	de_byte := []byte(decrypted)
	dd := &did.Document{}
	_ = json.Unmarshal(de_byte, dd)
	log.Println(dd.ID)

	sign, err := xrsa_Alice.Sign(data)
	err = xrsa_Alice.Verify(data, sign)

}
