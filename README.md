# adjective-animal

Prints a random adjective and a random animal.

### Build:
```
go build
```

### Use:
```
$ ./adjective-animal
shocking mole
```

Alternatively, run `./adjective-animal -server` to host it as an http service.
```
$ curl localhost:8080
careful elephant
```
I used nginx to bind this to a specific URL pattern
```
location ~* /adjective.?animal {
    proxy_pass http://localhost:8080;
}
```

### FAQ
> Where'd you get the words?

Why do you ask? Are you the words police? Words are public domain aren't they? They're not secret words or patented incantations. They're in alphabetical order so it's not like--

I scraped them from Wikipedia, ok?
