# tado-temp-guard

Do you use a [Tado smart thermostat](https://www.tado.com/gb/)? Does someone in the insist on turning the temperature up way too high? Well no more!

## Use

Add your Tado credentials (email/password) to the config file, set a max temperature, run it, and sit back while the temperature never goes too high again :-)

## Output
```
./tado-temp-guard -v

tado-temp-guard (tts) v0.0.1 - woanware

Home:
{
  "homes": [
    {
      "id": 123,
      "name": "Mark and Co's Home"
    }
  ]
}

Zones:
[
  {
    "id": 1,
    "name": "Heating",
    "type": "HEATING"
  }
]

Temperature has been set TOOOOOO HIGH: 25
Temperature set to: 21
```
