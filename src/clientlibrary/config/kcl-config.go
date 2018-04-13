package config

import (
	"clientlibrary/utils"
	"time"
)

// NewKinesisClientLibConfig to create a default KinesisClientLibConfiguration based on the required fields.
func NewKinesisClientLibConfig(applicationName, streamName, workerID string) *KinesisClientLibConfiguration {
	checkIsValueNotEmpty("ApplicationName", applicationName)
	checkIsValueNotEmpty("StreamName", streamName)
	checkIsValueNotEmpty("ApplicationName", applicationName)

	if empty(workerID) {
		workerID = utils.MustNewUUID()
	}

	// populate the KCL configuration with default values
	return &KinesisClientLibConfiguration{
		ApplicationName:                           applicationName,
		TableName:                                 applicationName,
		StreamName:                                streamName,
		WorkerID:                                  workerID,
		KinesisEndpoint:                           "",
		InitialPositionInStream:                   DEFAULT_INITIAL_POSITION_IN_STREAM,
		InitialPositionInStreamExtended:           *newInitialPosition(DEFAULT_INITIAL_POSITION_IN_STREAM),
		FailoverTimeMillis:                        DEFAULT_FAILOVER_TIME_MILLIS,
		MaxRecords:                                DEFAULT_MAX_RECORDS,
		IdleTimeBetweenReadsInMillis:              DEFAULT_IDLETIME_BETWEEN_READS_MILLIS,
		CallProcessRecordsEvenForEmptyRecordList:  DEFAULT_DONT_CALL_PROCESS_RECORDS_FOR_EMPTY_RECORD_LIST,
		ParentShardPollIntervalMillis:             DEFAULT_PARENT_SHARD_POLL_INTERVAL_MILLIS,
		ShardSyncIntervalMillis:                   DEFAULT_SHARD_SYNC_INTERVAL_MILLIS,
		CleanupTerminatedShardsBeforeExpiry:       DEFAULT_CLEANUP_LEASES_UPON_SHARDS_COMPLETION,
		TaskBackoffTimeMillis:                     DEFAULT_TASK_BACKOFF_TIME_MILLIS,
		MetricsBufferTimeMillis:                   DEFAULT_METRICS_BUFFER_TIME_MILLIS,
		MetricsMaxQueueSize:                       DEFAULT_METRICS_MAX_QUEUE_SIZE,
		ValidateSequenceNumberBeforeCheckpointing: DEFAULT_VALIDATE_SEQUENCE_NUMBER_BEFORE_CHECKPOINTING,
		RegionName:                                       "",
		ShutdownGraceMillis:                              DEFAULT_SHUTDOWN_GRACE_MILLIS,
		MaxLeasesForWorker:                               DEFAULT_MAX_LEASES_FOR_WORKER,
		MaxLeasesToStealAtOneTime:                        DEFAULT_MAX_LEASES_TO_STEAL_AT_ONE_TIME,
		InitialLeaseTableReadCapacity:                    DEFAULT_INITIAL_LEASE_TABLE_READ_CAPACITY,
		InitialLeaseTableWriteCapacity:                   DEFAULT_INITIAL_LEASE_TABLE_WRITE_CAPACITY,
		SkipShardSyncAtWorkerInitializationIfLeasesExist: DEFAULT_SKIP_SHARD_SYNC_AT_STARTUP_IF_LEASES_EXIST,
		WorkerThreadPoolSize:                             1,
	}
}

// WithTableName to provide alternative lease table in DynamoDB
func (c *KinesisClientLibConfiguration) WithTableName(tableName string) *KinesisClientLibConfiguration {
	c.TableName = tableName
	return c
}

func (c *KinesisClientLibConfiguration) WithKinesisEndpoint(kinesisEndpoint string) *KinesisClientLibConfiguration {
	c.KinesisEndpoint = kinesisEndpoint
	return c
}

func (c *KinesisClientLibConfiguration) WithInitialPositionInStream(initialPositionInStream InitialPositionInStream) *KinesisClientLibConfiguration {
	c.InitialPositionInStream = initialPositionInStream
	c.InitialPositionInStreamExtended = *newInitialPosition(initialPositionInStream)
	return c
}

func (c *KinesisClientLibConfiguration) WithTimestampAtInitialPositionInStream(timestamp *time.Time) *KinesisClientLibConfiguration {
	c.InitialPositionInStream = AT_TIMESTAMP
	c.InitialPositionInStreamExtended = *newInitialPositionAtTimestamp(timestamp)
	return c
}

func (c *KinesisClientLibConfiguration) WithFailoverTimeMillis(failoverTimeMillis int) *KinesisClientLibConfiguration {
	checkIsValuePositive("FailoverTimeMillis", failoverTimeMillis)
	c.FailoverTimeMillis = failoverTimeMillis
	return c
}

func (c *KinesisClientLibConfiguration) WithShardSyncIntervalMillis(shardSyncIntervalMillis int) *KinesisClientLibConfiguration {
	checkIsValuePositive("ShardSyncIntervalMillis", shardSyncIntervalMillis)
	c.ShardSyncIntervalMillis = shardSyncIntervalMillis
	return c
}

func (c *KinesisClientLibConfiguration) WithMaxRecords(maxRecords int) *KinesisClientLibConfiguration {
	checkIsValuePositive("MaxRecords", maxRecords)
	c.MaxRecords = maxRecords
	return c
}

/**
 * Controls how long the KCL will sleep if no records are returned from Kinesis
 *
 * <p>
 * This value is only used when no records are returned; if records are returned, the {@link com.amazonaws.services.kinesis.clientlibrary.lib.worker.ProcessTask} will
 * immediately retrieve the next set of records after the call to
 * {@link com.amazonaws.services.kinesis.clientlibrary.interfaces.v2.IRecordProcessor#processRecords(ProcessRecordsInput)}
 * has returned. Setting this value to high may result in the KCL being unable to catch up. If you are changing this
 * value it's recommended that you enable {@link #withCallProcessRecordsEvenForEmptyRecordList(boolean)}, and
 * monitor how far behind the records retrieved are by inspecting
 * {@link com.amazonaws.services.kinesis.clientlibrary.types.ProcessRecordsInput#getMillisBehindLatest()}, and the
 * <a href=
 * "http://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html#kinesis-metrics-stream">CloudWatch
 * Metric: GetRecords.MillisBehindLatest</a>
 * </p>
 *
 * @param IdleTimeBetweenReadsInMillis
 *            how long to sleep between GetRecords calls when no records are returned.
 * @return KinesisClientLibConfiguration
 */
func (c *KinesisClientLibConfiguration) WithIdleTimeBetweenReadsInMillis(idleTimeBetweenReadsInMillis int) *KinesisClientLibConfiguration {
	checkIsValuePositive("IdleTimeBetweenReadsInMillis", idleTimeBetweenReadsInMillis)
	c.IdleTimeBetweenReadsInMillis = idleTimeBetweenReadsInMillis
	return c
}

func (c *KinesisClientLibConfiguration) WithCallProcessRecordsEvenForEmptyRecordList(callProcessRecordsEvenForEmptyRecordList bool) *KinesisClientLibConfiguration {
	c.CallProcessRecordsEvenForEmptyRecordList = callProcessRecordsEvenForEmptyRecordList
	return c
}

func (c *KinesisClientLibConfiguration) WithTaskBackoffTimeMillis(taskBackoffTimeMillis int) *KinesisClientLibConfiguration {
	checkIsValuePositive("TaskBackoffTimeMillis", taskBackoffTimeMillis)
	c.TaskBackoffTimeMillis = taskBackoffTimeMillis
	return c
}

// WithMetricsBufferTimeMillis configures Metrics are buffered for at most this long before publishing to CloudWatch
func (c *KinesisClientLibConfiguration) WithMetricsBufferTimeMillis(metricsBufferTimeMillis int) *KinesisClientLibConfiguration {
	checkIsValuePositive("MetricsBufferTimeMillis", metricsBufferTimeMillis)
	c.MetricsBufferTimeMillis = metricsBufferTimeMillis
	return c
}

// WithMetricsMaxQueueSize configures Max number of metrics to buffer before publishing to CloudWatch
func (c *KinesisClientLibConfiguration) WithMetricsMaxQueueSize(metricsMaxQueueSize int) *KinesisClientLibConfiguration {
	checkIsValuePositive("MetricsMaxQueueSize", metricsMaxQueueSize)
	c.MetricsMaxQueueSize = metricsMaxQueueSize
	return c
}

// WithRegionName configures region for the stream
func (c *KinesisClientLibConfiguration) WithRegionName(regionName string) *KinesisClientLibConfiguration {
	checkIsValueNotEmpty("RegionName", regionName)
	c.RegionName = regionName
	return c
}

// WithWorkerThreadPoolSize configures worker thread pool size
func (c *KinesisClientLibConfiguration) WithWorkerThreadPoolSize(n int) *KinesisClientLibConfiguration {
	checkIsValuePositive("WorkerThreadPoolSize", n)
	c.WorkerThreadPoolSize = n
	return c
}
