This is an *in-progress* SDK for FamilySearch written in go (golang).
It includes functions for most of the read calls but none of the update calls.
I plan to eventually add functions for most of the update calls, 
though I don't intend to include every function like the 
[familysearch-javascript-sdk](https://github.com/rootsdev/familysearch-javascript-sdk) does - just the ones I use.

This is my first large-ish go project, so it may not be idiomatic go.
If you notice something that could be improved, please post an issue or a pull request.

If you'd like to help develop, *please do so*!

## Documentation

[GoDoc.org](http://godoc.org/github.com/rootsdev/gofamilysearch)

## Example

Here's how you might use the SDK

      package main
      
      import (
         "github.com/rootsdev/gofamilysearch"
         "log"
         "net/http"
      )
      
      func main() {
          // Context can be shared among go-routines
         ctx := &gofamilysearch.Context{
            Environment: "sandbox",
         }
         // Client is specific to a user
         c := &gofamilysearch.Client{
            Context: ctx,
            AccessToken: "access token for the requesting user goes here",
            Transport: DefaultTransport, // pass in transport to allow running normally or on appengine
         }
      
         user, err := c.GetCurrentUser()
         if err != nil {
            log.Panic(err)
         }
         log.Printf("ID=%s personID=%s treeUserID=%s\n", user.ID, user.PersonID, user.TreeUserID)
      }      
