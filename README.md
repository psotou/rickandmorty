# Rick and Morty Challange

### Generals 

Personal approach to solve the Rick and Morty challenge. It provides the code and an executable for thouse who aren't really into Go (although you totally should, the language is awesome). So if you aren't a gopher, go clone this repo and then run the executable:

```bash
git clone https://github.com/psotou/rickandmortyapi.git
```

And then run

```bash
bin/rickandmorty
```

This should produce a JSON file (results.json) in the same directory where you run the executable.

Also, in case you are unfamiliar with how to structure projects with Go, the `cmd/` directory is where you add the main application entry point files for the project.

I also setup a pipeline (github action where one can see the test code coverage).

### Project Structure

```bash
cmd/
├── fixtures/
├── char_counter.go
├── char_counter_test.go
├── character.go
├── character_test.go
├── episode.go
├── episode_locations.go
├── episode_locations_test.go
├── episode_test.go
├── final_answer.go
├── http.go
├── info.go
├── location.go
├── location_test.go
└── main.go
```

