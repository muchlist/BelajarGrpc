package main

import (
	"context"
	"fmt"
	"log"

	"github.com/muchlist/BelajarGrpc/blog/blogpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client has Started")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	//Create blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Muchlis",
		Title:    "My First Blog",
		Content:  "Content of my first blog",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}
	fmt.Printf("Blog has been Created: %v", createBlogRes)
	blogID := createBlogRes.GetBlog().GetId()

	//Readblog with wrong id
	fmt.Println("Read the blog")
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "sadasdasd",
	})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v\n", err2)
	}

	//readblog
	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	res, err2 := c.ReadBlog(context.Background(), readBlogReq)
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v\n", err2)
	}

	fmt.Printf("Blog was succesfull read: %v\n", res)

}
