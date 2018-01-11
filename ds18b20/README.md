
# ds18b20
This activity will allow your Flogo application to read the temperature from a DS18b20 sensor through a Raspberry Pi.

## Installation

```bash
flogo install github.com/lhollyer/flogo-activities/ds18b20
```
Link for flogo web:
```
https://github.com/lhollyer/flogo-activities/ds18b20
```
------------
## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "delimiter",
      "type": "string"
    },
    {
      "name": "prefix",
      "type": "string"
    },
    {
      "name": "suffix",
      "type": "string"
    },
    {
      "name": "part1",
      "type": "string"
    },
    {
      "name": "part2",
      "type": "string"
    },
    {
      "name": "part3",
      "type": "string"
    },
    {
      "name": "part4",
      "type": "string"
    },
    {
      "name": "part5",
      "type": "string"
    },
    {
      "name": "part6",
      "type": "string"
    },
    {
      "name": "part7",
      "type": "string"
    },
    {
      "name": "part8",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}
```
