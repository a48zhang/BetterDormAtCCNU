# BetterDormAtCCNU
---



# structs

## User

```go
type User struct {
    Uid    string `json:"uid" bson:"uid"`
    CCNUid string `json:"ccnuid" bson:"ccnuid"`
    Name   string `json:"name" bson:"name"`
    Passwd string `json:"passwd" bson:"passwd"`
    Role   int    `json:"role" bson:"role"`
    Valid  int    `json:"valid" bson:"valid"`
}
```

### Role

0-student
1-teacher
2-admin\school

### valid

1 for valid, 0 for invalid (deleted user)

## Form

```go
type Form struct {
	Fid          string `json:"fid" bson:"fid"`
	StudentID    string `json:"student_id" bson:"student_id"`
	TeacherID    string `json:"teacher_id" bson:"teacher_id"`
	CreateAt     string `json:"create_at" bson:"create_at"`
	FromDorm     string `json:"from_dorm" bson:"from_dorm"`
	ToDorm       string `json:"to_dorm" bson:"to_dorm"`
	Context      string `json:"context" bson:"context"`
	Status       int    `json:"status" bson:"status"`
}
```