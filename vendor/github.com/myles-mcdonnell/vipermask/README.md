## Vipermask

A common requirement is the ability to mask one set of configuration with another. 

For example there may be a default configuration set for an application that should be masked with environment specific values.

Consider the following TOML config files:


```
//Base Config
MaxWorkers = 10
LogLevel = Debug
DatabaseHost = localhost
```

```
//Prod config
LogLevel = Info
DatabaseHost = a.prod.db.server.com
```
By masking the Base with the Prod config the result is:

```
//Resulting config
MaxWorkers = 10
LogLevel = Info
DatabaseHost = a.prod.db.server.com
```

Vipermask simply composes multiple [Vipers](https://github.com/spf13/viper) to achieve this.  Use either ```NewFromFiles``` or ```NewFromVipers``` constructors to create a masked Config set with the base config set passed as the last argument.

See the [tests](configreader_test.go) for examples.
