package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

var (
	user    string
	host    string
	port    string
	keyfile string
)

func init() {
	flag.StringVar(&user, "u", "root", "SSH user")
	flag.StringVar(&host, "h", "example.tld", "Host")
	flag.StringVar(&port, "p", "22", "SSH port")
	flag.StringVar(&keyfile, "pk", "", "Public key file, e.g.: \"~/.ssh/id_rsa\"")
}

func main() {
	flag.Parse()

	if host == "example.tld" {
		fmt.Println("Usage: go run ssh-client.go -h <host> -p <port> -pk <path_to_private_key>")
		flag.PrintDefaults()
		return
	}

	var client *ssh.Client
	var err error

	if keyfile != "" {
		client, err = connectToHostWithPublicKey(user, fmt.Sprintf("%v:%v", host, port), keyfile)
	} else {
		client, err = connectToHost(user, fmt.Sprintf("%v:%v", host, port))
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect: %v\n", err)
		return
	}
	defer client.Close()

	runRemoteCommands(client)
}

func connectToHost(user, host string) (*ssh.Client, error) {
	var password string
	fmt.Print("SSH Password: ")
	fmt.Scanf("%s\n", &password)

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func connectToHostWithPublicKey(user, host, publicKeyFile string) (*ssh.Client, error) {
	key, err := ioutil.ReadFile(publicKeyFile)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func runRemoteCommands(client *ssh.Client) {
	commands := []string{
		"ls -al",
		"df -h",
		"uptime",
		"whoami",
	}

	for _, cmd := range commands {
		output, err := executeRemoteCommand(client, cmd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to run command '%s': %v\n", cmd, err)
			continue
		}
		fmt.Printf("Output of '%s':\n%s\n", cmd, output)
	}
}

func executeRemoteCommand(client *ssh.Client, command string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
