package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ryuku-hisa/grpc-filestream/proto"
)

// DataStreamSender は，指定されたファイルを gRPC ストリームを使用してサーバーに送信します．
func DataStreamSender(client pb.DataStreamHandlerClient, filename string) error {
	// ファイルを開きます．
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// データストリームを開始します．
	stream, err := client.DataStream(context.Background())
	if err != nil {
		return err
	}

	// 1KB のバッファを作成します．
	buf := make([]byte, 1024)
	for {
		// ファイルからデータを読み取ります．
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// 読み取ったデータとファイル名をストリームに送信します．
		if err := stream.Send(&pb.DataStreamRequest{Data: buf[:n], FileName: filepath.Base(filename)}); err != nil {
			return err
		}
	}

	// ストリームを閉じてサーバーからの応答を受信します．
	_, err = stream.CloseAndRecv()
	return err
}

// main はアプリケーションのエントリポイントです．gRPC サーバーへの接続を設定し，
// DataStreamHandler クライアントを作成し，sendFile 関数を呼び出してファイルを送信します．
func main() {
	// コマンドライン引数をチェックします．
	if len(os.Args) != 2 {
		log.Println("Invalid argument")
		fmt.Println("Usage:", filepath.Base(os.Args[0]), "<file path>")
		os.Exit(1)
	}
	filename := os.Args[1]

	// gRPC サーバーへの接続を確立します．
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("could not connect:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 新しい DataStreamHandler クライアントを作成します．
	client := pb.NewDataStreamHandlerClient(conn)

	// sendFile 関数を呼び出してファイルを送信します．エラーが発生した場合はログに記録します．
	if err := DataStreamSender(client, filename); err != nil {
		log.Println("failed to send file:", err)
		os.Exit(1)
	}
}
