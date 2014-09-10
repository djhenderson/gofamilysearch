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
         "github.com/rootsdev/gofamilysearch"
         "log"
         "net/http"
      )
      
      func main() {
         // pass in client to allow running on appengine
         client := &http.Client{};
         env := gofamilysearch.NewEnvironment("your client id goes here", "sandbox", client)
         c := gofamilysearch.NewContext(env, "access token for the requesting user goes here", client)
      
         user, err := c.GetCurrentUser()
         if err != nil {
            log.Panic(err)
         }
      
         log.Printf("id=%s personId=%s treeUserId=%s\n", user.Id, user.PersonId, user.TreeUserId)
      }
      