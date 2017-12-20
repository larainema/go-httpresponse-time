package httpresponse

import (        
  "testing"
)    

func Test_GetTime(t *testing.T) {
  GetTime()
  t.Log("Testing GetTime")
}

func Test_CronJob(t *testing.T) {
  CronJob()
  t.Log("Testing CronJob")
}
