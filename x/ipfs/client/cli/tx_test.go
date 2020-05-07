package cli

import (
	"bufio"
	"fmt"
	chunker "github.com/ipfs/go-ipfs-chunker"
	"io"
	"os"
	"testing"
)

func TestGetCmdPut(t *testing.T) {
	handler, err := os.Open("/home/angelo/Musica/test.mp3")
	if err != nil {
		t.Fatalf("%s", err)
	}

	chnk, err := chunker.FromString(bufio.NewReader(handler), "size-262144")
	if err != nil {
		t.Fatalf("failed to chunk block: %s", err)
	}

	for {
		chunk, err := chnk.NextBytes()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("%s", err.Error())
		}
		fmt.Println(len(chunk))
	}

	// ipfs dag get QmV1jwrSwA5XbcDK42ME1jS29jFmsEcCAhjjpgTdANeTTk
	// {"data":"CAIYtP4/IICAECCAgBAggIAQILT+Dw==","links":[{"Name":"","Size":262158,"Cid":{"/":"QmdQhBSL2ar6fhLYdCcf2spPjMBj5mp1qeJecGB3dnC3xx"}},{"Name":"","Size":262158,"Cid":{"/":"QmPtbdeZ9yGNMiogwGrUpXic63dsszqPhFj91UvhVyB7ZH"}},{"Name":"","Size":262158,"Cid":{"/":"QmfEDPxvUeRxx88qg5K3nnMutuqjY6wAqkixamyujERR79"}},{"Name":"","Size":261954,"Cid":{"/":"QmNZDhGQWx6SfVRrcPVYsSebVYqcDxC6JV9tY7DzwqLk9J"}}]}
}