package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

// Définition des structures Go pour les données JSON

// Thumbnail représente les différentes tailles de vignettes d'une photo
type Thumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Photo contient les informations d'une photo et ses vignettes
type Photo struct {
	ID         string              `json:"id"`
	Width      int                 `json:"width"`
	Height     int                 `json:"height"`
	URL        string              `json:"url"`
	Filename   string              `json:"filename"`
	Size       int                 `json:"size"`
	Type       string              `json:"type"`
	Thumbnails map[string]Thumbnail `json:"thumbnails"`
}

// Fields contient les champs de chaque enregistrement
type Fields struct {
	MagicSeaweedLink        string   `json:"Magic Seaweed Link"`
	Photos                  []Photo  `json:"Photos"`
	Geocode                 string   `json:"Geocode"`
	Influencers             []string `json:"Influencers"`
	SurfBreak               []string `json:"Surf Break"`
	PeakSurfSeasonBegins    string   `json:"Peak Surf Season Begins"`
	DestinationStateCountry string   `json:"Destination State/Country"`
	PeakSurfSeasonEnds      string   `json:"Peak Surf Season Ends"`
	DifficultyLevel         int      `json:"Difficulty Level"`
	Destination             string   `json:"Destination"`
	Address                 string   `json:"Address"`
}

// Event représente chaque enregistrement avec son ID, l'heure de création et les champs
type Event struct {
	ID          string `json:"id"`
	CreatedTime string `json:"createdTime"`
	Fields      Fields `json:"fields"`
}

type AllEvents struct {
	Records []Event `json:"records"`
}

var events = AllEvents{
	Records: []Event{
		{
			ID:          "recTY7UmjDJl0VCvI",
			CreatedTime: "2018-05-31T00:16:16.000Z",
			Fields: Fields{
				MagicSeaweedLink: "https://magicseaweed.com/Pipeline-Backdoor-Surf-Report/616/",
				Photos: []Photo{
					{
						ID:       "attf6yu03NAtCuv5L",
						Width:    2233,
						Height:   1536,
						URL:      "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/wX1jRSJPeSuql1StB4zTLQ/_QEAcliZ1VmxvL4_kGbu4hfHEJhJu3fZBUi7OQ2FHTmURbLfJQmWqP33kuZU9J6CoflJZ7QY-1KVDPZ-MrYbQoNvmgjCwwUw_bFl_iJnsv8crJy0nXFIPa0gSDE57paU/u78tYA5D0MsHiYkreazUHtNnM6iH4jY6CtCIsEDNYPg",
						Filename: "thomas-ashlock-64485-unsplash.jpg",
						Size:     688397,
						Type:     "image/jpeg",
						Thumbnails: map[string]Thumbnail{
							"small": {
								URL:    "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/6Mvid7iBhitC3mnPR55T9A/Nl4M9mOBeEDP4wnnyrgTag7PtaIOfX14yMVSlhxA8DipkttMZs4qVWrBkQIlUOShovDypE6fnZ85xi856HOP_g8bBjhMWQrG0EyhoKFDZDjc5p_EqsC-cdTcYBqFCQVV-cXfGSoo0LulN1UB6Rrcmg/7McZvlKXCa_l2tOh9igLIk9-KXKrtocgbk8JyZPHPCM",
								Width:  52,
								Height: 36,
							},
							"large": {
								URL:    "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/PWd4HHrlDXs02WZJnUJrvg/3jB0G7KwFtdQPdDTlrmToouqWP8LXQXWy9TM4OpiKuU2fggGgLMq9dQi38wX4jk82GaZoqDKI68HxJ3X11oPWrqinlYzyk1TOsBKpUqVCfbmTmXrke5in6ut3CTcPGsqoIb98w9cbumg1lBFCKfexw/EQowCeHCBaHU4cu6VqTtcniQc7kwmm4_DdcbV8au_Po",
								Width:  744,
								Height: 512,
							},
							"full": {
								URL:    "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/RiaM5hlHlVxkAqid1qJ-WQ/BstFxa0p9hrC9Jt78XOSQOKPRq5Boj15CjMKnAT9nq4vuNK36XB5qCsulI0VTVscYn_3hQmKhNh_9ViTZx7goW9ZpXpwhZFt9_k53IWR1KJdGA3qwukWSdDMbAddfS2H/usms1qTLNWpRj4leEQjxbEQkS_-APLyVnfp1BrMHobo",
								Width:  2233,
								Height: 1536,
							},
						},
					},
				},
				Geocode:                 "eyJpIjoiUGlwZWxpbmUsIE9haHUsIEhhd2FpaSIsIm8iOnsic3RhdHVzIjoiT0siLCJmb3JtYXR0ZWRBZGRyZXNzIjoiRWh1a2FpIEJlYWNoIFBhcmssIEhhbGVpd2EsIEhJIDk2NzEyLCBVbml0ZWQgU3RhdGVzIiwibGF0IjoyMS42NjUwNTYyLCJsbmciOi0xNTguMDUxMjA0Njk5OTk5OTd9LCJlIjoxNTM1MzA3MDE5OTE1fQ==",
				Influencers:             []string{"recrP1aupHoWQxMe0", "recPdVWkPoHCQawnl"},
				SurfBreak:               []string{"Reef Break"},
				PeakSurfSeasonBegins:    "2024-07-22",
				DestinationStateCountry: "Oahu, Hawaii",
				PeakSurfSeasonEnds:      "2024-08-31",
				DifficultyLevel:         4,
				Destination:             "Pipeline",
				Address:                 "Pipeline, Oahu, Hawaii",
			},
		},
		{
			ID:          "recvQG8QHwdybkGPr",
			CreatedTime: "2018-05-31T00:16:16.000Z",
			Fields: Fields{
				MagicSeaweedLink: "https://magicseaweed.com/Surfers-Paradise-Gold-Coast-Surf-Report/1012/",
				Photos: []Photo{
					{
						ID:       "attmtbEOAQteRjz2p",
						Width:    4000,
						Height:   3000,
						URL:      "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/shMRDGex6-AmMnqlU6auzA/G6zFnkSlDl7YUSk1a5aEcC5D4He9o_-mkBqPlimff72fdOhNA6Inc7IT0jyDRMwMW0PWvIY0sUyvbUzLUpGQfvtgR2NXQj3Y2qkmTHvbFsXxciBuNm7v7SujxOs4aAj4/7rQZaIKntZOn6Q6ESUyo5ZuRoSTT2Tq4a6KkWobfibg",
						Filename: "jeremy-bishop-80371-unsplash.jpg",
						Size:     1524876,
						Type:	  "image/jpeg",
						Thumbnails: map[string]Thumbnail{
							"small": {
								URL:    "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/BrCzerLd2GQ39mAwbScn9w/T8NQbtKxtxLgaz15t-LgJcZu0g4vPVFxfELXfYzeQLcbjdJ9cfSabNYFeRNOr8sjlEoMEPwZdPaJ2628c8nLs9wTomMVqZ9aqumBSFsrP6VCU-bzX-Fm6x_PgjlbYPue/obOzU_nWR7kIbjuH76EJSDUMxgqlkKbO3kQCkJBeYSg",
								Width:  48,
								Height: 36,
							},
							"large": {
								URL:    "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/8suHtFSemJOT5enIsPcmwQ/tWQyh5QP_46EW09Nxb6utw0pzFPbDlNvcC5GghVy9RfdAOPTJbDIQ0qEqmsaCkzo79Z5UN7LGko280hjbDnNZ5ztRJf5V_qNTlp2TGhFf3fZEaqyznOoILVB2oZK0eNu/HGaageoGwi9I5EMBNpZSM1BDwKAkjw-HInv3Pf9M4sM",
								Width:  683,
								Height: 512,
							},
							"full": {
								URL:    "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/lKJrn_ZCk53_vm6jB5QxfA/SQ_2JEHRm0gm9TdW31rcXuQX4A1Mxa2KpgkFvLZccAsKAsZfZHeE77cvk7dXulpdiUXK0xW1e_VZwVr2F1omK3cR7XBUkKC93haf2SdhhJbK2BwgXW_M_dbQ2HZyp4Fe/yf1NmbYd2a1rzdxbtXRDweqz8dGcwY3LCUv-fK-dt2M",
								Width:  3000,
								Height: 2250,
							},
						},
					},
				},
				Geocode:                 "eyJpIjoiU3VwZXJiYW5rLCBHb2xkIENvYXN0LCBBdXN0cmFsaWEiLCJvIjp7InN0YXR1cyI6Ik9LIiwiZm9ybWF0dGVkQWRkcmVzcyI6IlNuYXBwZXIgUm9ja3MgUmQsIENvb2xhbmdhdHRhIFFMRCA0MjI1LCBBdXN0cmFsaWEiLCJsYXQiOi0yOC4xNjI1NTcsImxuZyI6MTUzLjU1MDAyODgwMDAwMDA2fSwiZSI6MTUzNTMwNzAyMjAxNX0=",
				Influencers:             []string{"recG8bPaumZkCk66b"},
				SurfBreak:               []string{"Point Break"},
				PeakSurfSeasonBegins:    "2022-11-28",
				DestinationStateCountry: "Gold Coast, Australia",
				PeakSurfSeasonEnds:      "2024-02-01",
				DifficultyLevel:         4,
				Destination:             "Superbank",
				Address:                 "Superbank, Gold Coast, Australia",
			},
		},
		{
			ID:          "recXD6YjnydQ68I1H",
			CreatedTime: "2018-05-31T00:51:11.000Z",
			Fields: Fields{
				MagicSeaweedLink: "https://magicseaweed.com/Jeffreys-Bay-J-Bay-Surf-Report/88/",
				Photos: []Photo{
					{
						ID:       "attw1P2npyb2BdUwL",
						Width:    3648,
						Height:   2432,
						URL:      "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/C7WRcIzcumHS6pQon4o3qw/BqeLi0bLvEGtWSW5lXVrET4AsJF4tRhx3-3pH1CLq_AmWb_RqhL56sQf3cLin9lFxmzl4u7mx8BZbUT5dYnC2GmXPR4jZdHexNfmka1I8Ty2HOi6S1WP3ZiYEMylB4cE/xxblOW_rrotw_a4ISXKCjETcb0a-yCrhTVYpG4qY_lM",
						Filename: "cesar-couto-477018-unsplash (1).jpg",
						Size:     3066389,
						Type:	  "image/jpeg",
						Thumbnails: map[string]Thumbnail{
							"small": {
								URL:    "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/xAicNLwOjlAFx6htll7Kzg/oFdlORhGz4URfJ9gH2Qam_DtvOtve_PKTLxBX9VSe-wKvIB0eW99blaizJo_MFwhTKZTKFNFPBtxsenxFzzNwiV-jilPkmk0liw0_ztybO-acHeHC-F4KxvP3XaSRmdpk8RR6kyF36VpdADLRp5Ccw/15KFN2s29dB9PPUTWpd-_M-i28GzH72umy5zdqQ_u2Y",
								Width:  54,
								Height: 36,
							},
							"large": {
								URL:    "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/ROldw1PNSPZQB6Js__wtPw/huAfXjn3jroovp9SiC-IppDvWCeQLJfeUqA0DG-QMMRmlIKu3SSIAuJS1MGvQ6mIXmYU-j_bUCLVsPMCACWUONTdrHU6v6yqibT2JbHI5cVII6x82aKwTjgr5_EF0CP_Ax5EqIgikbEkujS5kr3-zw/8ye1FLfb44Ilf3k3_Ey2y7AYGso63UasEz_-aCHJxK8",
								Width:  768,
								Height: 512,
							},
							"full": {
								URL:    "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/BSFJ-PE3YYVQ5Y9j8Ze9Vw/3Yntz0A5ozs2D_1GphA9ftSWnln1bJRdyV3ZQ98GSjUuN6pq6Hh4FHmTM6tr2NLWXjXsspfFKzDlAwNVp40O13o_x--Kuz0aD4gAsfB0W4gBQAfkwMvgpaiJHV_gii3L_AxyzgZ9iWeoGMUQJ0KlxA/xRcbMtSbsHJBnJhwxuaW70fCZoU2GKHw1TsPZfiJkxQ",
								Width:  3000,
								Height: 2000,
							},
						},
					},
				},
				Geocode:                 "eyJpIjoiU3VwZXJ0dWJlcywgSmVmZnJleXMgQmF5LCBTb3V0aCBBZnJpY2EiLCJvIjp7InN0YXR1cyI6Ik9LIiwiZm9ybWF0dGVkQWRkcmVzcyI6IjEyIFBlcHBlciBTdCwgRmVycmVpcmEgVG93biwgSmVmZnJleXMgQmF5LCA2MzMwLCBTb3V0aCBBZnJpY2EiLCJsYXQiOi0zNC4wMzE3ODMsImxuZyI6MjQuOTMxNTk0MDAwMDAwMDJ9LCJlIjoxNTM1MzA3MDI3OTc1fQ==",
				Influencers:             []string{"recG8bPaumZkCk66b", "recPdVWkPoHCQawnl", "recrP1aupHoWQxMe0"},
				SurfBreak:               []string{"Point Break"},
				PeakSurfSeasonBegins:    "2024-08-01",
				DestinationStateCountry: "Jeffreys Bay, South Africa",
				PeakSurfSeasonEnds:      "2024-10-09",
				DifficultyLevel:         5,
				Destination:             "Supertubes",
				Address:                 "Supertubes, Jeffreys Bay, South Africa",
			},
		},
	},
}


//EVENTS
func getEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

//ONE EVENT
func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events.Records {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
			return
		}
	}
}

type SimpleSpots struct {
	SurfBreak []string `json:"surfBreak"`
	Address string `json:"address"`
}

//SUMMARY SPOTS LIST
func getSimpleSpotsList(w http.ResponseWriter, r *http.Request) {
	var simpleSpotsList []SimpleSpots

	for _, singleEvent := range events.Records {
		spotsListInfos := SimpleSpots{
			SurfBreak: singleEvent.Fields.SurfBreak,
			Address: singleEvent.Fields.Address,
		}
		simpleSpotsList = append(simpleSpotsList, spotsListInfos)
	}
	json.NewEncoder(w).Encode(simpleSpotsList)
}

func getSimpleSpot(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events.Records {
		if singleEvent.ID == eventID {
			oneSpotInfos := SimpleSpots{
				SurfBreak: singleEvent.Fields.SurfBreak,
				Address: singleEvent.Fields.Address,
			}
			json.NewEncoder(w).Encode(oneSpotInfos)
			return
		}
	}

}

