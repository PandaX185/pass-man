package pkg

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

type BoltDB struct {
	DB *bolt.DB
}

func (db *BoltDB) OpenBoltDB() error {
	if db.DB != nil {
		return nil
	}
	var err error
	db.DB, err = bolt.Open("passwords.db", 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})
	if err != nil {
		return err
	}
	return nil
}

func (db *BoltDB) AddPassword(email, password string) error {
	valid, _ := regexp.Match(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, []byte(email))
	if !valid {
		return fmt.Errorf("invalid email format")
	}

	domain := strings.Split(strings.Split(email, "@")[1], ".")[0]
	db.DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(domain))
		if err != nil {
			return fmt.Errorf("error creating bucket: %v", err)
		}

		encPassword, err := encrypt(password)
		if err != nil {
			return fmt.Errorf("error encrypting password: %v", err)
		}

		err = bucket.Put([]byte(email), []byte(encPassword))
		if err != nil {
			return fmt.Errorf("error storing password: %v", err)
		}
		fmt.Printf("Password added successfully for %s\n", email)
		return nil
	})
	return nil
}

func (db *BoltDB) GetPasswordByEmail(email string) (string, error) {
	valid, _ := regexp.Match(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, []byte(email))
	if !valid {
		return "", fmt.Errorf("invalid email format")
	}

	domain := strings.Split(strings.Split(email, "@")[1], ".")[0]

	var password string
	err := db.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(domain))
		if bucket == nil {
			return fmt.Errorf("domain %s not found", domain)
		}
		encPassword := bucket.Get([]byte(email))
		if encPassword == nil {
			return fmt.Errorf("email %s not found in domain %s", email, domain)
		}
		decPassword, err := decrypt(string(encPassword))
		if err != nil {
			return fmt.Errorf("error decrypting password: %v", err)
		}

		password = decPassword
		return nil
	})

	if err != nil {
		return "", err
	}
	return password, nil
}
