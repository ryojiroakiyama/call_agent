package main

import (
	"log"
	"net/http"
	"os"

	"github.com/twilio/twilio-go/twiml"
)

// 着信に応答するハンドラー
func handleIncomingCall(w http.ResponseWriter, r *http.Request) {
	// Twilio TwiMLレスポンスを作成
	response := twiml.NewVoiceResponse()
	
	// 特定の言葉を返す
	response.Say(twiml.Say{
		Text:     "こんにちは、こちらは自動応答システムです。お電話ありがとうございます。",
		Language: "ja-JP", // 日本語で応答
		Voice:    "woman", // 女性の声を使用
	})
	
	// レスポンスをXMLとして送信
	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(response.String()))
	
	log.Println("着信に応答しました")
}

func main() {
	// 環境変数からポート番号を取得、設定されていなければ8080を使用
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	// Webhookエンドポイントの登録
	http.HandleFunc("/voice", handleIncomingCall)
	
	// サーバーの起動
	log.Printf("サーバーを起動しています: http://localhost:%s", port)
	log.Printf("Twilioの設定でWebhook URLを http://あなたのドメイン/voice に設定してください")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}