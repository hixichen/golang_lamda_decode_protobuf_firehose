# golang_lamda_decode_protobuf_for_firehose


When backup data to s3 via firehose, firehose will trigger this lambda function to decode data from protobuf to json format.
    
How to use it:

0.set GOPATH to {pwd}  
1.get import packages via 'go get' 
2.execute ./zip.sh to build and package 
3.upload main.zip to aws lambda function via console 
