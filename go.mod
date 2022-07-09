module github.com/vincentkerdraon/inkcaller

replace github.com/vincentkerdraon/inkcaller/inkcallerlib => ./inkcallerlib

go 1.18

require (
	github.com/vincentkerdraon/inkcaller/inkcallerlib v0.0.0-20220324184416-1a0971bbb504
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f
	rogchap.com/v8go v0.7.0
)
