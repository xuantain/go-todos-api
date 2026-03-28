## NOTE: Model configs

1. The golang JSON Decoder expects for time values. It must be in the format of RFC3339 in order to be successfully decoded when using the JSON binder for ShouldBindJSON, or BindJSON.

EX: The "targetDate" in the requestion Json.body must be in format (YYYY-MM-DDTHH:mm:ssZ)
```
type Todo struct {
	TargetDate  time.Time `json:"targetDate" time_format:"2006-01-02T15:04:05Z07:00" gorm:"autoCreateTime:false"`
}
```

1. The json-type can be config different than model-type

EX: model-type => (ID  uint)  != json-type => `json:"id,string"`
```
type Todo struct {
	ID  uint  `json:"id,string" gorm:"primary_key"`
}
```