package models

type Config struct {
	Title         string `json:"title"`
	DownLoadImage string `json:"download_image"`
	ImgUrl        string `json:"img_url"`
	KeyDb         string `json:"key_db"`
	FaviconicoUrl string `json:"faviconico_url"`
	VideoTypes    string `json:"video_types"`
	TGBotKey      string `json:"tg_bot_key"`
	TGChatID      string `json:"tg_chat_id"`
}
