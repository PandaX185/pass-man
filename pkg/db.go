package pkg

import (
	"fmt"
	"os"
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
	db.DB, err = bolt.Open(os.Getenv("HOME")+"/.config/passman/.passwords.db", 0600, &bolt.Options{
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

func (db *BoltDB) GetPasswordsByDomain(domain string) (string, error) {
	var passwords []string
	err := db.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(domain))
		if bucket == nil {
			return fmt.Errorf("domain %s not found", domain)
		}

		return bucket.ForEach(func(k, v []byte) error {
			decPassword, err := decrypt(string(v))
			if err != nil {
				return fmt.Errorf("error decrypting password for %s: %v", string(k), err)
			}
			passwords = append(passwords, fmt.Sprintf("%s: %s", string(k), decPassword))
			return nil
		})
	})

	if err != nil {
		return "", err
	}

	if len(passwords) == 0 {
		return "", fmt.Errorf("no passwords found for domain %s", domain)
	}

	return strings.Join(passwords, "\n"), nil
}

func (db *BoltDB) GetAllPasswords() (map[string][]string, error) {
	allPasswords := make(map[string][]string)
	err := db.DB.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			return b.ForEach(func(k, v []byte) error {
				decPassword, err := decrypt(string(v))
				if err != nil {
					return fmt.Errorf("error decrypting password for %s: %v", string(k), err)
				}
				allPasswords[string(name)] = append(allPasswords[string(name)], fmt.Sprintf("%s: %s", string(k), decPassword))
				return nil
			})
		})
	})

	if err != nil {
		return nil, err
	}

	if len(allPasswords) == 0 {
		return nil, fmt.Errorf("no passwords found")
	}

	return allPasswords, nil
}
