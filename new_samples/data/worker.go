// Custom worker.go for Data samples - NOT generated
// This sample requires a custom DataConverter in worker options

package main

import (
	"fmt"

	"github.com/uber-go/tally"
	apiv1 "github.com/uber/cadence-idl/go/proto/api/v1"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/compatibility"
	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/peer"
	yarpchostport "go.uber.org/yarpc/peer/hostport"
	"go.uber.org/yarpc/transport/grpc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	HostPort       = "127.0.0.1:7833"
	Domain         = "cadence-samples"
	TaskListName   = "cadence-samples-worker"
	ClientName     = "cadence-samples-worker"
	CadenceService = "cadence-frontend"
)

// StartWorker creates and starts a Cadence worker with custom DataConverter.
func StartWorker() {
	logger, cadenceClient := BuildLogger(), BuildCadenceClient()

	// Create compressed JSON data converter
	dataConverter := NewCompressedJSONDataConverter()

	// Show compression statistics on startup
	printCompressionStats()

	workerOptions := worker.Options{
		Logger:        logger,
		MetricsScope:  tally.NewTestScope(TaskListName, nil),
		DataConverter: dataConverter, // Custom DataConverter for compression
	}

	w := worker.New(
		cadenceClient,
		Domain,
		TaskListName,
		workerOptions)

	// Register workflow and activity
	w.RegisterWorkflowWithOptions(LargeDataConverterWorkflow, workflow.RegisterOptions{Name: "cadence_samples.LargeDataConverterWorkflow"})
	w.RegisterActivityWithOptions(LargeDataConverterActivity, activity.RegisterOptions{Name: "cadence_samples.LargeDataConverterActivity"})

	err := w.Start()
	if err != nil {
		panic("Failed to start worker: " + err.Error())
	}
	logger.Info("Started Worker.", zap.String("worker", TaskListName))
}

// printCompressionStats displays compression statistics for the sample payload
func printCompressionStats() {
	largePayload := CreateLargePayload()
	originalSize, compressedSize, compressionPercentage, err := GetPayloadSizeInfo(largePayload, NewCompressedJSONDataConverter())
	if err != nil {
		fmt.Printf("Error calculating compression stats: %v\n", err)
		return
	}

	fmt.Printf("\n=== Compression Statistics ===\n")
	fmt.Printf("Original JSON size: %d bytes (%.2f KB)\n", originalSize, float64(originalSize)/1024.0)
	fmt.Printf("Compressed size: %d bytes (%.2f KB)\n", compressedSize, float64(compressedSize)/1024.0)
	fmt.Printf("Compression ratio: %.2f%% reduction\n", compressionPercentage)
	fmt.Printf("Space saved: %d bytes (%.2f KB)\n", originalSize-compressedSize, float64(originalSize-compressedSize)/1024.0)
	fmt.Printf("==============================\n\n")
}

func BuildCadenceClient(dialOptions ...grpc.DialOption) workflowserviceclient.Interface {
	grpcTransport := grpc.NewTransport()
	myChooser := peer.NewSingle(
		yarpchostport.Identify(HostPort),
		grpcTransport.NewDialer(dialOptions...),
	)
	outbound := grpcTransport.NewOutbound(myChooser)

	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: ClientName,
		Outbounds: yarpc.Outbounds{
			CadenceService: {Unary: outbound},
		},
	})
	if err := dispatcher.Start(); err != nil {
		panic("Failed to start dispatcher: " + err.Error())
	}

	clientConfig := dispatcher.ClientConfig(CadenceService)

	return compatibility.NewThrift2ProtoAdapter(
		apiv1.NewDomainAPIYARPCClient(clientConfig),
		apiv1.NewWorkflowAPIYARPCClient(clientConfig),
		apiv1.NewWorkerAPIYARPCClient(clientConfig),
		apiv1.NewVisibilityAPIYARPCClient(clientConfig),
	)
}

func BuildLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zapcore.InfoLevel)

	var err error
	logger, err := config.Build()
	if err != nil {
		panic("Failed to setup logger: " + err.Error())
	}

	return logger
}
