This is an *experimental* SDK for FamilySearch written in go (golang).
It's pretty minimal right now. 
I'll be adding to it over time, though I don't intend to make it as complete as the
[familysearch-javascript-sdk](https://github.com/rootsdev/familysearch-javascript-sdk)
without help from others.
This is my first large-ish go project, so it may not be idiomatic go.
If you notice something that could be improved, please post an issue or a pull request.
If you'd like to help design/develop, *please do so*!

## Example

Here's how you might use the SDK

      package main
      
      import (
         "github.com/rootsdev/familysearch"
         "log"
      )
      
      func main() {
         // the context can be shared among go-routines
         ctx := familysearch.NewContext("your client id goes here", "sandbox")
      
         // the FamilySearch object is very lightweight
         fs := &familysearch.FamilySearch {
            Context: ctx,
            AccessToken: "you must provide your own; the SDK doesn't currently have a way to generate one for you",
         }
      
         user, err := fs.GetCurrentUser()
         if err != nil {
            log.Panic(err)
         }
      
         log.Printf("id=%s personId=%s treeUserId=%s\n", user.Id, user.PersonId, user.TreeUserId)
      }