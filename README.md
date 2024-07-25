# Purpose

Create fake CSV file for testing.

# Requirements
- Go

# Instructions

To generate file with one million rows:
```
go run . -count 1_000_000 | pv > data.csv
```

That file can then be split into multiple smaller ones with the below command:
```
split -dl 1000 --additional-suffix=.csv data.csv dat
```