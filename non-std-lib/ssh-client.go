package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"

	"golang.org/x/crypto/ssh"
)

var user string
var host string
var port string
var keyfile string

func main() {
	flag.StringVar(&user, "u", "root", "ssh user")
	flag.StringVar(&host, "h", "example.tld", "host")
	flag.StringVar(&port, "p", "22", "ssh port")
	flag.StringVar(&keyfile, "pk", "", "public key file, eg.: \"~/.ssh/id_rsa\"")
	flag.Parse()

	var client *ssh.Client
	var session *ssh.Session
	var err error

	if host == "example.tld" {
		fmt.Println("go run ssh-client.go -h example.tld -p 22 -pk /Users/johndoe/.ssh/id_rsa")
		flag.PrintDefaults()
		return
	}

	if keyfile != "" {
		client, session, err = connectToHostWithPublickey(user, fmt.Sprintf("%v:%v", host, port), keyfile)
	} else {
		client, session, err = connectToHost(user, fmt.Sprintf("%v:%v", host, port))
	}

	if err != nil {
		panic(err)
	}
	out, err := session.CombinedOutput("ls -al")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	client.Close()
}

func connectToHost(user, host string) (*ssh.Client, *ssh.Session, error) {
	var pass string
	fmt.Print("SSH-Password: ")
	fmt.Scanf("%s\n", &pass)

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}

func connectToHostWithPublickey(user, host, publickeyfile string) (*ssh.Client, *ssh.Session, error) {
	key, err := ioutil.ReadFile(publickeyfile)
	if err != nil {
		return nil, nil, err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, nil, err
	}
	client, err := ssh.Dial("tcp", host, &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.HostKeyCallback(func(string, net.Addr, ssh.PublicKey) error { return nil }),
	})
	if client == nil || err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}
