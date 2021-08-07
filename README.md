# RoadRunner Plugin Template

<p align="center">
 <a href="https://github.com/peterfox/roadrunner-plugin-template/releases"><img src="https://img.shields.io/github/v/release/peterfox/roadrunner-plugin-template.svg?maxAge=30"></a>
	<a href="https://pkg.go.dev/github.com/peterfox/roadrunner-plugin-template"><img src="https://godoc.org/github.com/peterfox/roadrunner-plugin-template?status.svg"></a>
	<a href="https://github.com/peterfox/roadrunner-plugin-template/actions"><img src="https://github.com/peterfox/roadrunner-plugin-template/workflows/tests/badge.svg"></a>
	<a href="https://goreportcard.com/report/github.com/peterfox/roadrunner-plugin-template"><img src="https://goreportcard.com/badge/github.com/peterfox/roadrunner-plugin-template"></a>
	<a href="https://lgtm.com/projects/g/peterfox/roadrunner-plugin-template/alerts/"><img alt="Total alerts" src="https://img.shields.io/lgtm/alerts/g/peterfox/roadrunner-plugin-template.svg?logo=lgtm&logoWidth=18"/></a>
    <img alt="All releases" src="https://img.shields.io/github/downloads/peterfox/roadrunner-plugin-template/total">
</p>

<todo>
This repo can be used to scaffold a RoadRunner Plugin. Follow these steps to get started:

1. Update the module name in go.mod

   ```
   module github.com/peterfox/roadrunner-plugin-template
   ```
2. Ideally you should rename the package in all the `.go` files from `plugin` to something else.

   ```go
   package plugin
   ```
3. Rename the plugin to whatever is appropriate, this should be kept to snake case:
   ```go
   const PluginName = "plugin"
   ```
4. Update the LICENSE.md file
5. If your plugin is not going to be public on GitHub then you should remove `.github/workflows/codeql-analysis.yml`. You may also need to [enable code scanning](https://docs.github.com/en/code-security/secure-coding/automatically-scanning-your-code-for-vulnerabilities-and-errors/about-code-scanning) for the repository
6. Update the contributors in the README.md.
7. Update the README.md references to the plugin, installation and usage.
8. Update the badges at the top of README.md to point to your own repository.
9. Remove the todo section from your README.md

</todo>

## Installation

To use this plugin with RoadRunner you will need to fork or clone your
own copy of the [RoadRunner binary](https://github.com/spiral/roadrunner-binary).

You can import the plugin via go modules:

```sh
go get github.com/peterfox/roadrunner-plugin-demo
```

From there you can edit the [plugins.go](https://github.com/spiral/roadrunner-binary/blob/stable/internal/container/plugins.go) file to
import the plugin.

```go
package container

import (
    // ...
    demoPlugin "github.com/peterfox/roadrunner-plugin-template"
    // ...
)

// Plugins returns active plugins for the endure container. Feel free to add or remove any plugins.
func Plugins() []interface{} {
	return []interface{}{
        // ...
        &demoPlugin.Plugin{},
        // ...
    }
}

```

By importing this plugin and registering it the plugin will be compiled into the final binary.

The plugin will require that the _.rr.yaml_ config has the key `plugin` for the plugin won't initialise with roadrunner.

```yaml
plugin:
  value: foobar
```

## Usage

To make use of this plugin via PHP you must install the [Spiral Goridge](https://github.com/spiral/goridge-php) library.

You can use the following code as an example in php:

```php
<?php

use Spiral\Goridge\RPC\RPC;
use Spiral\RoadRunner\Environment;

$rpc = RPC::create(Environment::fromGlobals()->getRPCAddress());

// returns ['message' => 'test']
$output = $rpc->call('plugin.Message', ['message' => 'test']);
```

## Testing

You may download the project and test the plugin using the following command.

```bash
go test
```

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information on what has changed recently.

## Contributing

Please see [CONTRIBUTING](.github/CONTRIBUTING.md) for details.

## Security Vulnerabilities

Please review [our security policy](../../security/policy) on how to report security vulnerabilities.

## Credits

- [Peter Fox](https://github.com/peterfox)
- [All Contributors](../../contributors)

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
