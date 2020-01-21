package model

type Property struct {
    Id       int    `gorm:"primary_key" json:"id"`
    FlgOpen  bool   `json:"flg_open"`
    PrefCd   int    `json:"pref_cd"`
    WalkTime string `json:"walk_time"`
    Area     string `json:"area"`
    Price    string `json:"price"`
}

type Properties []Property
