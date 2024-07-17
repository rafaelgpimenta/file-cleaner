#!/bin/bash
awslocal s3 mb s3://bucket
awslocal s3 cp /home/localstack/wikipedia.png s3://bucket
