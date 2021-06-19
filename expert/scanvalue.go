package main

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"database/sql/driver"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	_ "github.com/mattn/go-sqlite3"
)

var stmtInit string = `CREATE TABLE IF NOT EXISTS test (
                            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                            data TEXT NOT NULL
                      )`

var encryptionKey []byte = []byte("2234|;gZ:H3AHhDz*aZ9XM4*!zA4nJZd")

// EncryptedBytes is our type which will automatically encrypt itself as it
// is inserted into the database, and decrypt itself as it is retrieved.
type EncryptedBytes []byte

// Value implements the transparent encryption for the EncryptedBytes type
func (e EncryptedBytes) Value() (driver.Value, error) {
	if e == nil {
		return nil, nil
	}

	ct, err := encrypt(e, encryptionKey)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	wb64 := base64.NewEncoder(base64.StdEncoding, &buf)
	wgz := gzip.NewWriter(wb64)

	_, err = wgz.Write(ct)
	if err != nil {
		wgz.Close()
		wb64.Close()
		return nil, err
	}
	wgz.Close()
	wb64.Close()

	return buf.Bytes(), nil
}

// Scan implements the transparent decryption for the EncryptedBytes type
func (e *EncryptedBytes) Scan(value interface{}) error {
	bv, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source was not []byte")
	}

	buf := bytes.NewBuffer(bv)
	r, err := gzip.NewReader(base64.NewDecoder(base64.StdEncoding, buf))
	if err != nil {
		return err
	}
	defer r.Close()

	ct, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	*e, err = decrypt(ct, encryptionKey)
	return err
}

// https://astaxie.gitbooks.io/build-web-application-with-golang/en/09.6.html
func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func main() {
	db, err := sql.Open("sqlite3", "./scanvalue.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(stmtInit)
	if err != nil {
		panic(err)
	}

	eb := EncryptedBytes("Hello, world!")
	fmt.Println("Start value: " + string(eb))

	res, err := db.Exec("INSERT INTO test (data) VALUES (?)", eb)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	var rb EncryptedBytes
	rows, err := db.Query("SELECT data FROM test WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&rb)
		if err != nil {
			panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Retrieved value: " + string(rb))
	fmt.Println("Now run `sqlite3 scanvalue.db 'select * from test;'` to see the encrypted and base64'd value")
}
