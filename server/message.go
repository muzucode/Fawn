package server

import "time"

type Message struct {
	Id             [16]byte       `json:"id" binding:"required"`
	MetadataHeader MetadataHeader `json:"metadata-header" binding:"required"`
	DataPayload    DataPayload    `json:"data-payload" binding:"required"`
	OptionsHeader  OptionsHeader  `json:"options-header" binding:"required"`
}

type MetadataHeader struct {
	Time       time.Time `json:"time" binding:"required"`
	OwnerId    [16]byte  `json:"owner-id" binding:"required"`
	DestAddr   string    `json:"dest-addr" binding:"required"`
	DestPort   string    `json:"dest-port" binding:"required"`
	SourceAddr string    `json:"source-addr" binding:"required"`
	SourcePort string    `json:"source-port" binding:"required"`
}

type OptionsHeader struct {
}

type DataPayload struct {
	Data [128]byte
}
