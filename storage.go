package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
)

type Storage interface {
	Init(string)
	Get(string) (string, error)
	New(string) (string, error)
	Cache() ([]urlEntity,error)
}

type PsqlStore struct {
	Db  *sql.DB
	cfg Config
}



func NewDb(cfg Config) (Storage, error) {
	if cfg.Host == "" || cfg.Port == "" || cfg.User == "" ||
		cfg.Password == "" || cfg.Database == "" {
		err := errors.New("All fields must be set.")
		return nil, err
	}
	
	var psq PsqlStore
	
	psq.cfg = cfg

	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port))
	if err != nil {
		fmt.Errorf("%s",err.Error())
		return nil, err
	}

	// Ping verifies if the connection to the database is alive or if a
	// new connection can be made.
	if err = db.Ping(); err != nil {
		fmt.Errorf("%s",err.Error())
		return nil, err
	}

	psq.Db = db

	return &psq, nil
}

func (p *PsqlStore) Init(name string) {
	/*create, err := ioutil.ReadFile("./DB_WEB_URL.sql")

	if _, err := p.Db.Exec(string(create)); err != nil {
		err = errors.New("Could not create database")
		return
	}*/ // this block is useless, an DB is needed in the config struct

	dat, err := ioutil.ReadFile("./" + name)
	fmt.Printf("Creating table with sql script %s \n %s",name,dat)
	
	if err != nil {
		fmt.Printf("Could not parse file %s.",name)
		return
	}
	if _, err := p.Db.Exec(string(dat)); err != nil {
		fmt.Printf("Could not create table from file %s ", name)
		return
	}
	return
}

func (p *PsqlStore) Get(shortUrl string) (string, error) {
	var url string = ""
	const query_ = `SELECT URL FROM WEBURL WHERE ID=$1` // limit not required because it is unique (?)
	var id = baseUDecode(shortUrl)
	err := p.Db.QueryRow(query_, id).Scan(&url)

	if err != nil {
		err = errors.New("Could not select URL from WEB_URL.")
		return "", err
	}
	
	return url, nil
}

// this returns the last ID used, this to give the respective id associated with the url
func (p *PsqlStore) New(fullUrl string) (string, error) {
	var lastId = 0
	const query = `INSERT INTO WEBURL (URL) VALUES ($1) RETURNING ID`
	err := p.Db.QueryRow(query,fullUrl).Scan(&lastId)
	
	if err != nil {
		return "", err
	}

	return baseUEncode(lastId), nil
}

func (p *PsqlStore) Cache() ([]urlEntity,error){
	
	
	return nil,nil
}