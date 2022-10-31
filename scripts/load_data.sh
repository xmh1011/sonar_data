#!/bin/bash

START_TIME=${START_TIME:-946656000}
FIANL_TIME=${FINAL_TIME:-1640966400}
STEP_TIME=${STEP_TIME:-6000}
DATABASE_PORT=${DATABASE_PORT:-8086}
while true
do
#generate data
EXE_FILE_NAME_1=${EXE_FILE_NAME_1:-$(which sonar_data)}
# shellcheck disable=SC1072
$EXE_FILE_NAME_1

EXE_FILE_NAME_2=${EXE_FILE_NAME_2:-$(which load_cnosdb)}


# Load parameters
BATCH_SIZE=${BATCH_SIZE:-10000}
NUM_WORKERS=${NUM_WORKERS:-4}
BACKOFF_SECS=${BACKOFF_SECS:-1s}
REPORTING_PERIOD=${REPORTING_PERIOD:-10s}
DO_ABORT_ON_EXIST=${DO_ABORT_ON_EXIST:-false}
DO_CREATE_DB=${DO_CREATE_DB:-false}

$EXE_FILE_NAME_1  generatedata5 --startTime=${START_TIME} --endTime=$((START_TIME+STEP_TIME)) | $EXE_FILE_NAME_2 \
                                                                                            --do-create-db=${DO_CREATE_DB} \
                                                                                            --do-abort-on-exist=${DO_ABORT_ON_EXIST} \
                                                                                            --backoff=${BACKOFF_SECS} \
                                                                                            --workers=${NUM_WORKERS} \
                                                                                            --batch-size=${BATCH_SIZE} \
                                                                                            --reporting-period=${REPORTING_PERIOD} \
START_TIME=$((START_TIME+STEP_TIME))
sleep 5
done
