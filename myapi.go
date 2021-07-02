package api

import(
    "github.com/imroc/req"
    "Defender-cli/utils"
    "time"
    "log"
)
type New struct{
    Token string
}

func (a New) Scan(file string) utils.ScannerR{

    var result map[string]string
    var returnResult utils.ScannerR
    url := "https://api.metadefender.com/v4/file/"
    authHeader := req.Header{
        "apikey" : a.Token,
    }
    headers := req.Header{
        "apikey": a.Token,
        "content-type": "application/json",
    }
    resp, err := req.Post(url, headers, req.File(file))
    
    if err != nil{
        log.Fatal(err)
    }

    resp.ToJSON(&result)
    time.Sleep(5 * time.Second)
    body, err := req.Get(url + result["data_id"], authHeader)

    if err != nil{
        log.Fatal(err)
    }
    body.ToJSON(&returnResult)

    return returnResult
}


