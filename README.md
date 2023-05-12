# ral2jqlog-f

ral2jqlog-f is a tool that retrieves the RDS Audit logs in the specified bucket and uploads them in HIVE format for easy analysis.  HIVE format has the advantage of easy partitioning.

## Install

```
$ go install github.com/ryuichi1208/ral2jqlog-f@v0.0.2
```

## Usage

```
Usage:
  ral2jqlog-f [OPTIONS]

Application Options:
  -d, --dst-bucket= audit log file
  -s, --src-bucket= File Content Type
      --date=       date

Help Options:
  -h, --help        Show this help message

Usage:
  ral2jqlog-f [OPTIONS]

Application Options:
  -d, --dst-bucket= audit log file
  -s, --src-bucket= File Content Type
      --date=       date

Help Options:
  -h, --help        Show this help message
```

## Athena

```
# CREATE TABLE
CREATE EXTERNAL TABLE IF NOT EXISTS `rds_log`.`hive_query_log` (
  `timestamp` string,
  `user` string,
  `client` string,
  `host` string,
  `command` string,
  `query` string
)
PARTITIONED BY (`dt` string)
ROW FORMAT SERDE 'org.openx.data.jsonserde.JsonSerDe'
WITH SERDEPROPERTIES (
  'ignore.malformed.json' = 'FALSE',
  'dots.in.keys' = 'FALSE',
  'case.insensitive' = 'TRUE',
  'mapping' = 'TRUE'
)
STORED AS INPUTFORMAT 'org.apache.hadoop.mapred.TextInputFormat' OUTPUTFORMAT 'org.apache.hadoop.hive.ql.io.HiveIgnoreKeyTextOutputFormat'
LOCATION 's3://rds-querylog-hive/'
TBLPROPERTIES ('classification' = 'json');

# Run Load
MSCK REPAIR TABLE `hive_query_log`

```
