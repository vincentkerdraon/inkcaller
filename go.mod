module github.com/vincentkerdraon/inkcaller

replace github.com/vincentkerdraon/inkcaller/inkcallerlib => ./inkcallerlib

go 1.17

require (
	github.com/vincentkerdraon/inkcaller/inkcallerlib v0.0.0-20220303022905-855faab75d25
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	rogchap.com/v8go v0.7.0
)
