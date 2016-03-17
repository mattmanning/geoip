# Geoip

JSON API for the MaxMind GeoLite2 City database.

## Run locally

    $ convox start

## Deploy

    $ convox deploy

## HTTP requests

    http://<docker host>/city/<ip address>

## Example

    $ curl http://convox.local/city/50.180.47.55
    {"City":{"GeoNameID":4180439,"Names":{"de":"Atlanta","en":"Atlanta","es":"Atlanta","fr":"Atlanta","ja":"アトランタ","pt-BR":"Atlanta","ru":"Атланта","zh-CN":"亚特兰大"}},"Continent":{"Code":"NA","GeoNameID":6255149,"Names":{"de":"Nordamerika","en":"North America","es":"Norteamérica","fr":"Amérique du Nord","ja":"北アメリカ","pt-BR":"América do Norte","ru":"Северная Америка","zh-CN":"北美洲"}},"Country":{"GeoNameID":6252001,"IsoCode":"US","Names":{"de":"USA","en":"United States","es":"Estados Unidos","fr":"États-Unis","ja":"アメリカ合衆国","pt-BR":"Estados Unidos","ru":"США","zh-CN":"美国"}},"Location":{"Latitude":33.7884,"Longitude":-84.3491,"MetroCode":524,"TimeZone":"America/New_York"},"Postal":{"Code":"30306"},"RegisteredCountry":{"GeoNameID":6252001,"IsoCode":"US","Names":{"de":"USA","en":"United States","es":"Estados Unidos","fr":"États-Unis","ja":"アメリカ合衆国","pt-BR":"Estados Unidos","ru":"США","zh-CN":"美国"}},"RepresentedCountry":{"GeoNameID":0,"IsoCode":"","Names":null,"Type":""},"Subdivisions":[{"GeoNameID":4197000,"IsoCode":"GA","Names":{"en":"Georgia","es":"Georgia","fr":"Géorgie","ja":"ジョージア州","pt-BR":"Geórgia","ru":"Джорджия"}}],"Traits":{"IsAnonymousProxy":false,"IsSatelliteProvider":false}}

## Attribution

This product includes GeoLite2 data created by MaxMind, available from
<a href="http://www.maxmind.com">http://www.maxmind.com</a>.
