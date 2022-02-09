// Package inkcaller is an API to call Ink.
//
// Each new call is independent.
// A call will force the ink state, (optional) set the knot, (optional) answer a previous choice.
// A call returns the current ink state (to inject in next call), the text, the choices.
//
// This does not allow all the features of ink (e.g. no callback or external function).
//
// Ink will need to interact with the game model. Set needed data inside ink state in the input. Use formatted text and a parser for the actions.
//
// Based on https://github.com/inkle/ink
package inkcaller
