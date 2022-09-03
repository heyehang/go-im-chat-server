package mongo

import (
	"testing"
)

//var configFile = flag.String("f", "../../../etc/chat.yaml", "the config file")

func TestGetMembersByChatID(t *testing.T) {
	//flag.Parse()
	//var c config.Config
	//conf.MustLoad(*configFile, &c)
	//Init(context.Background(), c)
	//var members []interface{}
	//for i := 0; i < 3; i++ {
	//	mem := new(model.ChatMember)
	//	mem.ID = primitive.NewObjectID()
	//	mem.ChatId = "1"
	//	mem.ChatMemberID = mem.ID.Hex()
	//	mem.CTime = time.Now().Unix()
	//	mem.Role = 1
	//	mem.UTime = mem.CTime
	//	mem.UserID = int64(i) + int64(1)
	//	members = append(members, mem)
	//}
	//err := InsertChatMember(context.Background(), members)
	//if err != nil {
	//	panic(err)
	//	return
	//}

	//uids, err := GetMembersByChatID(context.Background(), "1", 10, 0)
	//if err != nil {
	//	panic(err)
	//	return
	//}

	//fmt.Println("查询到的:", uids)

}
