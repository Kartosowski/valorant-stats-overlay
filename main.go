package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Config struct {
	Apikey   string `json:"apikey"`
	Username string `json:"username"`
	Tag      string `json:"tag"`
	Port     int    `json:"port"`
}

type MMRData struct {
	Name string `json:"name"`
	Data []struct {
		CurrentTierPatched  string `json:"currenttierpatched"`
		RankingInTier       int    `json:"ranking_in_tier"`
		MMRChangeToLastGame int    `json:"mmr_change_to_last_game"`

		Images struct {
			Small string `json:"small"`
		} `json:"images"`
	} `json:"data"`
}

func main() {
	config := loadConfig("config.json")

	mux := http.NewServeMux()

	mux.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w, r)

		client := &http.Client{Timeout: 10 * time.Second}
		url := fmt.Sprintf("https://api.henrikdev.xyz/valorant/v1/mmr-history/eu/%s/%s", config.Username, config.Tag)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", config.Apikey)

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Failed to fetch stats", http.StatusInternalServerError)
			log.Println("Fetch error:", err)
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		logFetch(string(body))
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	})

	mux.Handle("/script.js", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("public/script.js")
		if err != nil {
			http.Error(w, "JS file not found", 404)
			return
		}
		js := strings.ReplaceAll(string(content), "{PORT}", fmt.Sprintf("%d", config.Port))
		w.Header().Set("Content-Type", "application/javascript")
		w.Write([]byte(js))
	}))

	mux.Handle("/", http.FileServer(http.Dir("./public")))

	fmt.Printf("VALORANT STATS OVERLAY\nAutor: Kartos\nhttps://github.com/Kartosowski/valorant-stats-overlay\n\nUruchomione na porcie: %d\nLink: http://localhost:%d\n", config.Port, config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), mux))
}

func loadConfig(file string) Config {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Nie znaleziono pliku %s!\nWięcej informacji na https://github.com/Kartosowski/valorant-stats-overlay", file)
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Błędna konfiguracja JSON: %v\n\nWięcej informacji na https://github.com/Kartosowski/valorant-stats-overlay", err)
	}
	if config.Apikey == "" || config.Username == "" || config.Tag == "" || config.Port == 0 {
		log.Fatal("Niepełna konfiguracja w config.json!\nWięcej informacji na https://github.com/Kartosowski/valorant-stats-overlay")
	}
	return config
}

func logFetch(body string) {
	var data struct {
		Name string `json:"name"`
		Tag  string `json:"tag"`
		Data []struct {
			CurrentTierPatched  string `json:"currenttierpatched"`
			MMRChangeToLastGame int    `json:"mmr_change_to_last_game"`
		} `json:"data"`
	}

	if err := json.Unmarshal([]byte(body), &data); err != nil {
		fmt.Println("Błąd parsowania JSON do loga:", err)
		return
	}

	if len(data.Data) == 0 {
		fmt.Println("Brak danych o MMR")
		return
	}

	latest := data.Data[0]
	fmt.Printf(
		"\nZdobyte dane [%s]\n%s#%s - %s - %+d\n\n",
		time.Now().Format("15:04:05"),
		data.Name,
		data.Tag,
		latest.CurrentTierPatched,
		latest.MMRChangeToLastGame,
	)
}

func enableCORS(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
	}
}
