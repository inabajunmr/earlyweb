package main

import (
	"net"
	"time"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// メインのFlexコンテナ (垂直方向)
	mainFlex := tview.NewFlex().
		SetDirection(tview.FlexRow)

	// 上部の入力エリア (水平方向)
	inputFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn)

	// 入力フィールドの作成
	destInput := tview.NewInputField().
		SetLabel("dest: ").
		SetFieldWidth(30)
	requestInput := tview.NewInputField().
		SetLabel("request: ").
		SetFieldWidth(30)

	titleArea := tview.NewTextView().
		SetDynamicColors(true).
		SetText("title")
	titleArea.SetBorder(true)

	textArea := tview.NewTextView().
		SetDynamicColors(true).
		SetText("body")

	// ボタンの作成
	button := tview.NewButton("送信").
		SetSelectedFunc(func() {
			// 入力テキストを取得
			dest := destInput.GetText()
			req := requestInput.GetText()
			textArea.SetText(sendHTTP09Request(dest, req))
		})

	// 入力エリアの配置
	inputFlex.AddItem(destInput, 0, 1, true).AddItem(requestInput, 0, 1, true).AddItem(button, 10, 0, false)

	// テキストエリアの枠設定
	textArea.SetBorder(true)

	// メインFlexに配置
	mainFlex.AddItem(inputFlex, 3, 0, true).
		AddItem(titleArea, 3, 0, true).
		AddItem(textArea, 0, 1, false)

	// アプリケーションの実行
	if err := app.SetRoot(mainFlex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func sendHTTP09Request(dest string, req string) string {
	conn, err := net.DialTimeout("tcp", dest, 5*time.Second)
	if err != nil {
		return err.Error()
	}
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	// リクエストを送信
	_, err = conn.Write([]byte("GET " + req + "\n"))
	if err != nil {
		return err.Error()
	}

	// レスポンスを受信
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		return err.Error()
	}

	return string(buffer[:n])
}
