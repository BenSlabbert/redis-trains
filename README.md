# Redis Trains

Using redis streams and the graph module to track trains in a simulation.

## Hooks

The [hooks](./hooks) dir contains git hook scripts.

Make sure to run:
```shell
git config core.hooksPath hooks
```

So that `git` uses the [hooks](./hooks) dir for the got hooks.

## fmt

Go code uses `make fmt` for formatting, while SQL files uses `pg_format`, see [pre-commit](./hooks/pre-commit).
