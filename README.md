
# RoboVac Rooms

Reads Room IDs from [Valetudo](https://github.com/Hypfer/Valetudo) in a format that can be directly appended to `customize.yaml` for use in home automation. 

Intended to be run as part of a init script for [Home Assistant](https://www.home-assistant.io/). Compiles as a static binary to avoid needing to introduce additional dependencies. 

## Why

I have a Dreame branded robot vacuum and use Home Assistant to control it. I have various automations to handle the schedule, such as "Don't vacuum the office if I'm in a meeting" and "Vacuum Living room overnight". These automations rely on HA knowing which room is which and unfortunately Dreame vacuums have a hardware issue that causes the virtual rooms to drift out of sync with the real world over time, requiring remapping occasionally. When new maps are created, the room IDs can change. This was created because I got tired of needing to update IDs in HA.

## Usage

```
./robovac-rooms VALETUDO_IP_OR_HOSTNAME >> /config/customize.yaml;
```

## Output

```yaml
input_boolean.vacuum_office:
  room_id: 1 
input_boolean.vacuum_bedroom:
  room_id: 2 
input_boolean.vacuum_bathroom:
  room_id: 4 
input_boolean.vacuum_hallway:
  room_id: 5 
input_boolean.vacuum_livingroom:
  room_id: 6 
input_boolean.vacuum_kitchen:
  room_id: 7 

```
## License

[MIT](LICENSE.txt)

