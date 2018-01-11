
# ds18b20
This activity will allow your Flogo application to read the temperature from a DS18b20 sensor through a Raspberry Pi.
Using [yryz's Golang library](https://github.com/yryz/ds18b20).

## Installation

```bash
flogo install github.com/lhollyer/flogo-activities/ds18b20
```
Link for flogo web:
```
https://github.com/lhollyer/flogo-activities/ds18b20
```


## Schema
Inputs and Outputs:

```json
{
  "inputs":[],
  "outputs": [
    {
      "name": "result",
      "type": "integer"
    }
  ]
}
```

## Outputs
| Output   | Description    |
|:----------|:---------------|
| result    | The temperature integer output |
