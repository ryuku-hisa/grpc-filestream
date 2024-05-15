// Package main は，ストリーミングを使用してファイルをサーバーにアップロードおよび保存する
// gRPC サーバーの実装を提供します．
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/ryuku-hisa/grpc-filestream/proto"
)

const (
	// port は gRPC サーバーがリッスンするポート番号です．
	port = ":8080"
)

// server は，DataStreamHandler サービスを実装する gRPC サーバーです．
type server struct {
	pb.UnimplementedDataStreamHandlerServer
}

// DataStream は，クライアントからのデータストリームを受信してファイルに保存します．
func (s *server) DataStream(stream pb.DataStreamHandler_DataStreamServer) error {
	fmt.Println("Receiving...")

	// 最初のメッセージを受け取り，ファイル名を取得します．
	firstMessage, err := stream.Recv()
	if err != nil {
		return err
	}
	filename := firstMessage.FileName

	// "receivedData" ディレクトリを作成します．
	if err := os.MkdirAll("receivedData", 0777); err != nil {
		return err
	}

	// "receivedData" ディレクトリに受け取ったファイル名でファイルを作成します．
	filePath := filepath.Join("receivedData", filename)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 最初のデータチャンクをファイルに書き込みます．
	if _, err := file.Write(firstMessage.Data); err != nil {
		return err
	}

	// 残りのデータを受信し，ファイルに書き込みます．
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break // データの受信が完了したらループを終了します．
		}
		if err != nil {
			return err
		}
		if _, err := file.Write(resp.Data); err != nil {
			return err
		}
	}

	// クライアントに対してストリームの終了とステータスを送信します．
	if err := stream.SendAndClose(&pb.DataStreamResponse{
		DataStreamStatus: "OK",
	}); err != nil {
		return err
	}

	fmt.Println("DONE")
	return nil
}

// main はアプリケーションのエントリポイントです．gRPCサーバーを設定し，起動します．
func main() {
	fmt.Println("Preparing...")

	// 指定したポートでTCPリスナーを作成します．
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("failed to listen:", err)
		os.Exit(1)
	}
	defer lis.Close()

	// gRPC サーバーを作成します．
	gserver := grpc.NewServer()

	// DataStreamHandler サーバーを登録し，gRPC サーバーで使用できるようにします．
	pb.RegisterDataStreamHandlerServer(gserver, &server{})

	// サーバーリフレクションを登録します（クライアントがサーバーのメタデータを問い合わせることができるようにするため）．
	reflection.Register(gserver)

	// gRPC サーバーを起動し，リスナーで指定されたアドレスで接続を受け付けます．
	fmt.Println("--- Receiving ---")
	if err := gserver.Serve(lis); err != nil {
		log.Println("server ended:", err)
		os.Exit(1)
	}
}
