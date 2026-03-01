# Wargame bot database

## Wargame data

| Mode |  |
|:-|:-|
|__id__| int |
| name | string |
| serever_name | string |
| map_pool | list of map sessions |
| auto_rotate | bool |
| enable vote | bool |
| players | int |
| req_players | int |
| gameMode | int |

| Map_session |  |
|:-|:-|
| __id__ | int |
| name | string |
| income_rate | int |
| init_money | int |
| score_limit | int |
| time_limit | int |

| Map |  |
|:-|:-|
| __id__ | string |
| name | string |
| image | string |
| size | string |

| Nation | |
|:-|:-|
| __id__ | int |
| name | string |
| code | string |
| emote_id | string |

| Specializasion | |
|:-|:-|
| __id__ | int |
| name | string |
| code | string |
| emote_id | string |

| Era | |
|:-|:-|
| __id__ | int |
| name | string |
| emote_id | string |

| Wargame_Player |  |
|:-|:-:|
| __id__ | int |
| name | string |
| sessions | table |
| discord_player | bool? |

| Discord_Player |  |
|:- |:-:|
| __id__ | string |
| name | string |
| sessions | table |
| wargame_id | int? |

