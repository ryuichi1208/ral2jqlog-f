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
