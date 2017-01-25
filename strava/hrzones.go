package main

import (
   "fmt"
   "encoding/json"
   "time"
   "io/ioutil"
   "net/http"
   "os"
)

type DistributionBucket struct {
  Max int
  Min int
  Time int
  Percent float64
  Tag string
  Label_short string
  Label_long string
}

type HeartRateZones struct {
   Score int
   Distribution_buckets []DistributionBucket
   Type string
   Resource_state int
   Sensor_based bool
   Points int
   Custom_zones bool
   Calculating bool
}

func Bar (n float64) string {
   barChar := "|"
   result := ""
   for i := 0; i < int(n/2); i++ {
      result+=barChar
   }
   return result
}

func main() {

  if len(os.Args[1:]) != 1 {
     fmt.Println("Usage: hrzones2.go <activity id>")
     os.Exit(1)
  }
  client := &http.Client{}
  req, _ := http.NewRequest("GET", "https://www.strava.com/activities/"+os.Args[1]+"/heartrate_zones", nil)
  req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0")
  req.Header.Del("Accept-Encoding")
  req.Header.Set("Accept", "*/*")
  res, _ := client.Do(req)
  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)
  fmt.Println(string(body))
  zones := HeartRateZones{}
  json.Unmarshal(body, &zones)

  fmt.Println("Your Zones")
  for _, bucket:= range zones.Distribution_buckets {
    fmt.Println(bucket.Label_long,":", time.Duration(bucket.Time) * time.Second)
  }

 fmt.Println("--------------------------------------------------------")
  for _, bucket:= range zones.Distribution_buckets {
    fmt.Println(bucket.Label_short,":", Bar(bucket.Percent*100) )
    fmt.Println("--------------------------------------------------------")
  }


}
