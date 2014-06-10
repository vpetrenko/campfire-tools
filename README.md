Campfire tools perform two tasks:

1. Import Campfire history into MySQL database for further message integrity checks.
1. Check that Campfire and Slack histories are matching.

## Config file format

config.json contains following settings:

```js
{
    "DbName": "cftools",
    "ApiKey": "<XXXxxxxXXXXXXXXXXXxxxxxXXXXXXXxxxXXXXXxx>",
    "Domain": "<cfdomain>",
    "SlackData": "data"
}
```

Where:
* DbName - MySQL Database name for Campfire history crawling
* ApiKey - Your Campfire API key
* Domain - Your Campfire domain
* SlackData - Path to Slack's 'Export Data' messages archive (unzip it to some folder)


## CLI

```bash
    # Imports Campfire history to MySQL database
    campfire-tools import

    # Perform histories match check
    campfire-tools -cfroom CfRoomName -slroom SlackRoomName [-from YYYY-MM-DD] [-to YYYY-MM-DD] check
```

## License

Copyright (c) 2014, Victor Petrenko.
Distributed under the MIT license.
