package main

import (
	"context"
	pb "grpc/grpc_blockchain/pb"
	bc "grpc/grpc_blockchain/server/blockchain"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = ":8080"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	pb.RegisterBlockchainServer(server, &blockchainService{Blockchain: bc.NewBlockchain()})
	log.Println("Starting server on port ", port)

	server.Serve(lis)
}

type blockchainService struct {
	Blockchain *bc.Blockchain
}

func (b *blockchainService) AddBlock(ctx context.Context, req *pb.AddBlockRequest) (*pb.AddBlockResponse, error) {

	block := b.Blockchain.AddBlock(req.Data)
	return &pb.AddBlockResponse{Hash: block.Hash}, nil
}

func (b *blockchainService) GetBlockchain(ctx context.Context, req *pb.GetBlockchainRequest) (*pb.GetBlockchainResponse, error) {

	resp := new(pb.GetBlockchainResponse)
	for _, b := range b.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, convertBlockToProtoBlock(b))
	}

	return &pb.GetBlockchainResponse{Blocks: resp.Blocks}, nil
}

func convertBlockToProtoBlock(block *bc.Block) *pb.Block {
	return &pb.Block{
		Hash:          block.Hash,
		PrevBlockHash: block.PrevBlockHash,
		Data:          block.Data,
	}
}
