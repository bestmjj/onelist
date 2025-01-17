package thedb

type TheVideo struct {
	BackdropPath     string   `json:"backdrop_path"`
	FirstAirDate     string   `json:"first_air_date"`
	GenreIds         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	Title            string   `json:"title"`
	Name             string   `json:"name"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	Adult            bool     `json:"Adult"`
	OriginalTitle    string   `json:"original_title"`
	ReleaseDate      string   `json:"release_date"`
	Video            bool     `json:"video"`
}

type ThedbSearchRsp struct {
	Page         int        `json:"page"`
	Results      []TheVideo `json:"results"`
	TotalPages   int        `json:"total_pages"`
	TotalResults int        `json:"total_results"`
}

/*
https://api.themoviedb.org/3/search/movie?api_key=xxxxxxxx&language=zh&page=1&query=%E7%88%86%E6%AC%BE%E5%A5%BD%E4%BA%BA
"adult": false,
"backdrop_path": "/aVN85JU72T4np67TY3kTwtcvBAs.jpg",
"genre_ids": [18, 35],
"id": 1197437,
"original_language": "zh",
"original_title": "爆款好人",
"overview": "出租车司机张北京（葛优 饰）失去了在亲生儿子的婚宴上作为新郎父亲的发言权。备受打击的他与单亲母亲李小琴（李雪琴 饰）不打不相识，机缘巧合之下成为了“网红”，突如其来的爆红意外让他收获了想要的一切。在体验了一番荒诞的名利场沉浮后，张北京最终悟到了烟火人间的暖心真情。",
"popularity": 1.996,
"poster_path": "/tku8RKSNCF83KuHelmQdMSRej03.jpg",
"release_date": "2024-10-01",
"title": "爆款好人",
"video": false,
"vote_average": 5,
"vote_count": 8
-----------------
https://api.themoviedb.org/3/movie/1197437?api_key=%s&language=zh
{
  "adult": false,
  "backdrop_path": "/aVN85JU72T4np67TY3kTwtcvBAs.jpg",
  "belongs_to_collection": null,
  "budget": 0,
  "genres": [
    {
      "id": 18,
      "name": "剧情"
    },
    {
      "id": 35,
      "name": "喜剧"
    }
  ],
  "homepage": "",
  "id": 1197437,
  "imdb_id": "tt30497082",
  "origin_country": [
    "CN"
  ],
  "original_language": "zh",
  "original_title": "爆款好人",
  "overview": "出租车司机张北京（葛优 饰）失去了在亲生儿子的婚宴上作为新郎父亲的发言权。备受打击的他与单亲母亲李小琴（李雪琴 饰）不打不相识，机缘巧合之下成为了“网红”，突如其来的爆红意外让他收获了想要的一切。在体验了一番荒诞的名利场沉浮后，张北京最终悟到了烟火人间的暖心真情。",
  "popularity": 1.996,
  "poster_path": "/tku8RKSNCF83KuHelmQdMSRej03.jpg",
  "production_companies": [
    {
      "id": 98750,
      "logo_path": null,
      "name": "Huanxi Media Group",
      "origin_country": "CN"
    },
    {
      "id": 238498,
      "logo_path": null,
      "name": "海南如日方升影视文化传播有限公司",
      "origin_country": ""
    }
  ],
  "production_countries": [
    {
      "iso_3166_1": "CN",
      "name": "China"
    }
  ],
  "release_date": "2024-10-01",
  "revenue": 0,
  "runtime": 113,
  "spoken_languages": [
    {
      "english_name": "Mandarin",
      "iso_639_1": "zh",
      "name": "普通话"
    }
  ],
  "status": "Released",
  "tagline": "",
  "title": "爆款好人",
  "video": false,
  "vote_average": 5,
  "vote_count": 8
}
*/
