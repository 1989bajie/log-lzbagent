package po
type MongoRecord struct {
	Target string `bson:"target"`
	Target_id string `bson:"target_id"`
	Action string `bson:"action"`
	Description string `bson:"description"`
	Operater_type string `bson:"operater_type"`
	Operater_id string `bson:"operater_id"`
	Operater_name string `bson:"operater_name"`
	Source string `bson:"source"`
	New_data string `bson:"new_data"`
	Old_data string `bson:"old_data"`
	Ip string `bson:"ip"`
	Add_time int64 `bson:"add_time"`
}
func (MongoRecord)GetTable() string{
	return "record"
}