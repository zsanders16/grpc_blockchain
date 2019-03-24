package main

import (
	"context"
	"fmt"
	pb "grpc/grpc_blockchain/pb"
	"log"

	"google.golang.org/grpc"
)

const port = ":8080"

func main() {

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost"+port, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewBlockchainClient(conn)
	hash, err := client.AddBlock(context.Background(), &pb.AddBlockRequest{Data: "This is my first block."})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hash.Hash)

	protoBlocks, err := client.GetBlockchain(context.Background(), &pb.GetBlockchainRequest{})
	if err != nil {
		log.Fatal(err)
	}
	blocks := convertProtoBlocksToBlocks(protoBlocks)

	for _, b := range blocks {
		fmt.Println(b)
	}

}

func convertProtoBlocksToBlocks(res *pb.GetBlockchainResponse) []Block {

	blocks := []Block{}
	for _, b := range res.Blocks {
		block := Block{
			Hash:          b.Hash,
			PrevBlockHash: b.PrevBlockHash,
			Data:          b.Data,
		}
		blocks = append(blocks, block)
	}
	return blocks
}

type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}
