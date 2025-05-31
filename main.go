package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
)

// TODO: 録音できるようにする
// https://help.twilio.com/articles/4408190859931
// https://help.twilio.com/articles/4408182810523

// 初期化関数
func init() {
	// .envファイルから環境変数を読み込む
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or error loading")
	}
}

// 着信に応答するハンドラー
func handleIncomingCall(w http.ResponseWriter, r *http.Request) {
	log.Println("着信を受け取りました")

	// TwiMLレスポンスを作成
	twiml := `<?xml version="1.0" encoding="UTF-8"?>
<Response>
	<Say voice="woman" language="ja-JP">こんにちは。いい天気ですね。お電話ありがとうございます。</Say>
</Response>`

	// レスポンスをXMLとして送信
	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(twiml))

	log.Println("着信に応答しました")
}

// Twilioクライアントのインスタンスを取得
func getTwilioClient() *twilio.RestClient {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	// 認証情報が設定されているか確認
	if accountSid == "" || authToken == "" {
		log.Println("Warning: Twilio credentials not set in environment variables")
	}

	// Twilioクライアントの作成
	return twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}

func main() {
	// 環境変数からポート番号を取得、設定されていなければ8080を使用
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Twilioクライアントの初期化（認証情報の確認のみを目的として実行）
	_ = getTwilioClient() // クライアントを使用しない場合はアンダースコアで変数を無視
	log.Printf("Twilio client initialized successfully")

	// Webhookエンドポイントの登録
	http.HandleFunc("/voice", handleIncomingCall)

	// サーバーの起動
	log.Printf("サーバーを起動しています: http://localhost:%s", port)
	log.Printf("Twilioの設定でWebhook URLを http://あなたのドメイン/voice に設定してください")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}
