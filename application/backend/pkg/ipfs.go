package pkg

import (
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

var sh = shell.NewShell("127.0.0.1:6001") //ipfs api

func IpfsAdd(filename string) (string, error) {
	ipfsfile, _ := os.Open(fmt.Sprintf("./files/uploadfiles/%v", filename))
	defer ipfsfile.Close()
	cid, err := sh.Add(ipfsfile)
	if err != nil {
		return "", fmt.Errorf("ipfs add file faile,err:%v", err)
	}
	return cid, nil
}

func IpfsGet(cid string, filename string) error {
	// fmt.Printf("cid:%v,filename:%v", cid, filename)
	err := sh.Get(cid, fmt.Sprintf("./files/downloadfiles/%v", filename))
	if err != nil {
		fmt.Printf("err:%v", err)
		return fmt.Errorf("ipfs get file faile,err:%v", err)
	}
	return nil
}
