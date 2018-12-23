package main

import "github.com/110y/gae-go-grpc/internal/api/cmd"

func main() {
	cmd.CheckError(cmd.Execute())
}
