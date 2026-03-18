## Samples in this folder

This folder contains samples demonstrating custom data conversion patterns in Cadence.

### Data Converter Workflow

The `LargeDataConverterWorkflow` demonstrates how to use custom data converters with compression capabilities. Data converters control how workflow inputs, outputs, and activity parameters are serialized and deserialized.

#### What is a DataConverter?

A `DataConverter` is responsible for:
- Serializing workflow inputs and activity parameters before storage
- Deserializing workflow outputs and activity results when retrieved
- Enabling custom encoding formats (compression, encryption, etc.)

#### The Compressed JSON DataConverter

This sample implements `compressedJSONDataConverter` which:
- Serializes data to JSON format
- Compresses using gzip to reduce storage size
- Automatically decompresses when reading data back

#### Compression Benefits

For large payloads, compression typically provides:
- **60-80% size reduction** for JSON data
- **Lower storage costs** in Cadence history
- **Reduced bandwidth** for data transfer
- **Better scalability** with large payloads

#### Use Cases

Custom DataConverters are useful when you need to:
- Reduce storage costs with compression
- Add encryption/decryption for sensitive data
- Support legacy serialization formats
- Implement custom compression algorithms (LZ4, Snappy, etc.)
- Add data validation during serialization

#### How to Start the Workflow

```bash
cadence --domain cadence-samples \
  workflow start \
  --workflow_type cadence_samples.LargeDataConverterWorkflow \
  --tl cadence-samples-worker \
  --et 60
```

Note: The workflow generates its own large payload internally - no input is required. This design allows the workflow to be started from CLI without the client needing the custom DataConverter. The compression demonstration happens when data is passed between the workflow and activity. When the worker starts, it displays compression statistics showing the before/after sizes.

#### Key Implementation Details

The custom DataConverter is configured in `worker.go`:

```go
workerOptions := worker.Options{
    DataConverter: NewCompressedJSONDataConverter(),
    // ... other options
}
```

Both the worker AND any client triggering workflows must use the same DataConverter to properly encode/decode data.
