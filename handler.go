package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Hello string `json:"hello"`
	}

	resp := response {
		Hello: "world",
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&resp)
}


func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html; charset=utf8")
	fmt.Fprint(w, "ok")
}

func createMerchantHandler(w http.ResponseWriter, r *http.Request)  {
	// parse
	merchantReq := CreateMerchantRequest{}
	if err := json.NewDecoder(r.Body).Decode(&merchantReq); err != nil {
		fmt.Errorf("failed to decode %w", err)
		return
	}
	// validate

	// save(DB)
	var db *sqlx.DB
	// ファイル名 この処理を実行後にファイルがカレントディレクトリで作成される
	db, err := sqlx.Connect("sqlite3", "sqlite.db")
	if err != nil {
		fmt.Fprintf(w, fmt.Errorf("sqlite: could not open database: %w", err).Error())
		return
	}
	if err := db.Ping(); err != nil {
		fmt.Fprintf(w, fmt.Errorf("sqlite: could not ping database: %w", err).Error())
		return
	}
	ctx := context.Background()
	result := db.MustExecContext(ctx,
		"INSERT INTO merchants(email, name, photo_url)\nVALUES (?,?,?)",
		merchantReq.Email, merchantReq.Name, merchantReq.PhotoURL)

	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	var merchant = Merchant{
		ID:       id,
		Email:    merchantReq.Email,
		Name:     merchantReq.Name,
		PhotoURL: merchantReq.PhotoURL,
	}

	// response
	merchantResp := CreateMerchantResponse{}
	merchantResp.Merchant = merchant
	merchantResp.Success = true

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&merchantResp)
}

func listMerchantHandler(w http.ResponseWriter, r *http.Request)  {
	// read(DB)
	var db *sqlx.DB
	// ファイル名 この処理を実行後にファイルがカレントディレクトリで作成される
	db, err := sqlx.Connect("sqlite3", "sqlite.db")
	if err != nil {
		fmt.Fprintf(w, fmt.Errorf("sqlite: could not open database: %w", err).Error())
		return
	}
	if err := db.Ping(); err != nil {
		fmt.Fprintf(w, fmt.Errorf("sqlite: could not ping database: %w", err).Error())
		return
	}
	//ctx := context.Background()
	merchants := make([]*Merchant, 0)
	if err := db.Select(&merchants, "Select id, name, email, photo_url FROM merchants"); err != nil {
		return
	}
	// response
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&merchants)
}

func dbInitHandler(w http.ResponseWriter, r *http.Request) {
	// ここからDB接続処理開始
	var db *sqlx.DB
	// ファイル名 この処理を実行後にファイルがカレントディレクトリで作成される
	db, err := sqlx.Connect("sqlite3", "sqlite.db")
	if err != nil {
		fmt.Fprintf(w, fmt.Errorf("sqlite: could not open database: %w", err).Error())
		return
	}
	if err := db.Ping(); err != nil {
		fmt.Fprintf(w, fmt.Errorf("sqlite: could not ping database: %w", err).Error())
		return
	}
	// DB接続処理完了
	// スキーマ（テーブル）作成
	if err := prepareSchema(db); err != nil {
		fmt.Errorf("sqlite: could not prepare schema: %w", err)
		return
	}
	w.Header().Set("content-Type", "text/html; charset=utf8")
	fmt.Fprint(w, "succeed")
}

func dbSeedHandler(w http.ResponseWriter, r *http.Request) {
	// ここからDB接続処理開始
	var db *sqlx.DB
	// ファイル名 この処理を実行後にファイルがカレントディレクトリで作成される
	db, err := sqlx.Connect("sqlite3", "sqlite.db")
	if err != nil {
		fmt.Fprintf(w, fmt.Errorf("sqlite: could not open database: %w", err).Error())
		return
	}
	if err := db.Ping(); err != nil {
		fmt.Fprintf(w, fmt.Errorf("sqlite: could not ping database: %w", err).Error())
		return
	}
	// DB接続処理完了
	// データ作成
	if err := seed(db); err != nil {
		fmt.Fprintf(w, fmt.Errorf("sqlite: could not seed: %w", err).Error())
		return
	}
	w.Header().Set("content-Type", "text/html; charset=utf8")
	fmt.Fprint(w, "succeed")
}
