package recorder2

import "time"

type Record struct {
  GroupId string
  ID string
  Name string
  Email string
  AuthDomain string
  ReceivedAt time.Time
}
