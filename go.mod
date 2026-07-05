module github.com/vincentkerdraon/inkcaller

replace github.com/vincentkerdraon/inkcaller/inkcallerlib => ./inkcallerlib

go 1.26

require (
	github.com/vincentkerdraon/inkcaller/inkcallerlib v0.0.0-20250628143758-23299900f503
	golang.org/x/sync v0.21.0
	rogchap.com/v8go v0.9.0
)
