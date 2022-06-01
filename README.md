# deep-search

This repository was created to compare golang and javascript implementations of deep search with a json.

Both go and js implementations are tested with the same data.

The results are very similar probably because I don't know go at all (github copilot did all of the work here), I expected go to be way faster.

## What it does

- Given an array of objects, it must return all the objects that have the insensitive string in any value inside de object.
- Must return the object even if the insensitive string is inside a deeply nested object.

## Results

Both implementations are very similar and using the same dataset with 19547 records

| language   | sensitive case | insensitive case |
| ---------- | -------------- | ---------------- |
| golang     | 6ms            | 32ms             |
| javascript | 23ms           | 33ms             |

## Conclusions

You get the same results in go or js in about the same time. When it's not an insensitive implementation you get different results, go was taking around 6~7ms and js was taking around 23~24ms, in this case go was 3~4 times faster than js.

I learned a little bit of go, so it's a win for me regardless.

## Next steps

- Double check both implementations, if you search any genre you get different results, for example `Science Fiction` returns 1473 results in js and 25 in go, it's probably a js quirk where `typeof [] === 'object'`.
- Get memory usage and cpu usage of both implementations.
- Create a frontend to query both implementations.
- Test with fuzzy search.
- Test with different data.
- Try a more manual implementation instead of using the language features.
- Get a bigger average time, maybe with some scripts.

## Running locally

You need both node and go installed to run this example

#### js

- cd into `apps/js/js-search`
- run `yarn` or `npm install`
- run `yarn dev`

#### go

- cd into `apps/go/go-search`
- run `go run main.go`
