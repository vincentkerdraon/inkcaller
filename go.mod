module github.com/vincentkerdraon/inkcaller

replace github.com/vincentkerdraon/inkcaller/inkcallerlib => ./inkcallerlib

go 1.23

require (
	github.com/vincentkerdraon/inkcaller/inkcallerlib v0.0.0-20241004162243-b437b2929cad
	golang.org/x/sync v0.8.0
	rogchap.com/v8go v0.9.0
)
