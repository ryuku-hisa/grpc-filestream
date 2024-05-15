#!/bin/bash

# プロトコルバッファーをGoに変換
protoc --go_out=. \
       --go_opt=paths=source_relative \
       --go-grpc_out=. \
       --go-grpc_opt=paths=source_relative \
        file_stream.proto

# --go_out=. \                           # Goのコードを生成
# --go_opt=paths=source_relative \       # ファイルのパスを相対パスで生成
# --go-grpc_out=. \                      # GoのgRPCコードを生成
#  --go-grpc_opt=paths=source_relative \  # ファイルのパスを相対パスで生成
# ${filename}.proto                      # プロトコルバッファーのファイル名

