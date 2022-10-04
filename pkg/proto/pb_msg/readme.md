```
message MessageBody {
//  string srv_msg_id = 1;
  string cli_msg_id = 2;
  string sender_id = 3;
  string receiver_id = 4;
  int32 sender_platform = 5;
  string sender_nickname = 6;
  string sender_avatar_key = 7;
  string chat_id = 8;
  chat_type chat_type = 9;
//  int32 seq_id = 10;
  int32 msg_from = 11;
  int32 content_type = 12;
  bytes content = 13;
//  int32 status = 14;
  int64 sent_ts = 15;
//  int64 received_ts = 16;
}
```