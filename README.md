# hyrumrand

Generate random integers using the fact that Golang map iteration order is, as
of 2025-12-13, guranteed to be random.

Snippet from [the runtime/maps.go source](https://go.googlesource.com/go/+/refs/heads/dev.go2go/src/runtime/map.go#831):

```golang
// decide where to start
r := uintptr(fastrand())
if h.B > 31-bucketCntBits {
    r += uintptr(fastrand()) << 31
}
```

## Why?

[Inspired by this Reddit comment](https://www.reddit.com/r/golang/comments/1plrkq0/comment/ntujtux/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button), which answered the question "Why does Go’s map always give random key/value order under the hood?" with:

> So that no one gets dumb enough to depend on them being in any particular order because the order can change for reasons under the hood and out of your control

I took that personally.

[Hyrum's law](https://www.hyrumslaw.com/) states:

> With a sufficient number of users of an API, it does not matter what you promise in the 
> contract: all observable behaviors of your system will be depended on by somebody.

So, now there is a Golang library that depends on the observable behaviour that the
iteration order of Go map types is random.

## Testing

The unit test ensures that at least 2 unique
values are returned for 10,000 attempts at getting a number between 0 and 10,000. Odds of
this failing randomly is 10,000 to the 10,000 power.

## Should I use this?

No.

## But what if...

No.

## But I really want to...

No. It's also almost a thousand times slower than `math/rand`; you can test it yourself using `go run cmd/main.go`.

```console
[mtu@archlap hyrumrand]$ go run cmd/main.go 
Elapsed time using hyrumrand: 61.5919ms
Sum using hyrumrand: 365195
Elapsed time using math/rand: 13.378µs
Sum using math/rand: 502368
```

## But it's random, so I can use it to...

Please don't.
