# later-execution

# Development

## Table: Execution

|AttrName|Type|Schema|Description |UserIndex|
|--------|----|------|------------|---------|
|Id      |S   |PK    |            |         |
|UserId  |S   |      |            |PK       |
|Created |N   |      |            |SK       |
|At      |N   |      |Timestamp   |         |
|Expired |N   |      |TTL         |         |
|RuleName|S   |      |EventsBridge|         |

## Table: AccessKey

|AttrName|Type|Schema|Description|UserIndex|
|--------|----|------|-----------|---------|
|Id      |S   |PK    |           |         |
|Key     |S   |      |           |         |
|UserId  |S   |      |           |PK       |
|Created |N   |      |           |SK       |
|Expired |N   |      |TTL        |         |
