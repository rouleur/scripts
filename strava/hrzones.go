package main

import "encoding/json"
import "fmt"
import "time"
import "io/ioutil"
import "os"

type Activity struct{
   Time            []int 
   Distance        []int 
   Altitude        []int
   Heartrate       []int
   Cadence         []int
   Watts           []int
   Watts_calc      []int
   Temp            []int
   Grade_adjusted_distance []int
   Grade_smooth    []int
   Moving          []bool
   Velocity_smooth []int
}

func Bar (n int) string {
	  barChar := "|" 
    result := ""
    for i:=0; i<int(n/2); i ++ {
       result+=barChar
    }     
    return result
} 

func ZonePercentage (zoneTime int, activity Activity) int {
	return int(float32(zoneTime)/float32(len(activity.Time))*100)
}

func main(){
  
    if len(os.Args[1:]) != 1 {
      fmt.Println("Usage: hrzones.go <activity json file>")
      os.Exit(1)
    }

	  str,_ := ioutil.ReadFile(os.Args[1]);
    activity := Activity{}
    json.Unmarshal([]byte(str), &activity)
    fmt.Println("HR sample count:", len(activity.Heartrate))

    z2 := 109
    z3 := 145
    z4 := 162
    z5 := 181 

    z1Seconds := 0
    z2Seconds := 0
    z3Seconds := 0
    z4Seconds := 0
    z5Seconds := 0

    for j := range activity.Heartrate  {
    	if j == 0 { continue }
    	var entry = activity.Heartrate[j]
    	switch {
    	   case entry < z2:
    	      z1Seconds++
    	   case entry >= z2 && entry < z3:
    	      z2Seconds++
    	   case entry >= z3 && entry < z4:
              z3Seconds++
           case entry >= z4 && entry < z5:
              z4Seconds++
           case entry >= z4:
              z5Seconds++         
    	   }
    } 

    fmt.Println("Total: ", time.Duration(len(activity.Time)-1)*time.Second)
    fmt.Println("Your Zones")
    fmt.Println("Zone 1 (Endurance) :", time.Duration(z1Seconds) * time.Second)
    fmt.Println("Zone 2 (Moderate)  :", time.Duration(z2Seconds) * time.Second)
    fmt.Println("Zone 3 (Tempo)     :", time.Duration(z3Seconds) * time.Second)
    fmt.Println("Zone 4 (Threshold) :", time.Duration(z4Seconds) * time.Second)
    fmt.Println("Zone 5 (Anaerobic) :", time.Duration(z5Seconds) * time.Second)

    fmt.Println("--------------------------------------------------------")
    fmt.Println("Zone 1", Bar(ZonePercentage(z1Seconds, activity )))
    fmt.Println("--------------------------------------------------------")
    fmt.Println("Zone 2", Bar(ZonePercentage(z2Seconds, activity )))
    fmt.Println("--------------------------------------------------------")
    fmt.Println("Zone 3", Bar(ZonePercentage(z3Seconds, activity )))
    fmt.Println("--------------------------------------------------------")
    fmt.Println("Zone 4", Bar(ZonePercentage(z4Seconds, activity )))
    fmt.Println("--------------------------------------------------------")
    fmt.Println("Zone 5", Bar(ZonePercentage(z5Seconds, activity )))
    fmt.Println("--------------------------------------------------------")

      
}
