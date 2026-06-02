# StartTech System Architecture

## Architecture Overview

StartTech uses a cloud-native architecture hosted on AWS with Infrastructure as Code managed through Terraform and deployment automation using GitHub Actions.

The application consists of:

* React Frontend hosted on Amazon S3
* CloudFront CDN for content delivery
* Golang Backend API running on EC2 instances
* Application Load Balancer (ALB)
* Auto Scaling Group (ASG)
* Amazon ElastiCache Redis Cluster
* MongoDB Atlas Database
* Amazon CloudWatch Monitoring and Logging

---

## High-Level Architecture

Users access the React frontend through CloudFront. Static assets are served from Amazon S3.

API requests are routed through an Application Load Balancer to EC2 instances running the Golang application. The EC2 instances are managed by an Auto Scaling Group for high availability and scalability.

Redis is used for caching and session management through Amazon ElastiCache.

Persistent application data is stored in MongoDB Atlas.

CloudWatch collects logs, metrics, and alarms for monitoring and observability.

---

## Networking Architecture

### VPC

CIDR Block:

10.0.0.0/16

### Public Subnets

* 10.0.1.0/24
* 10.0.2.0/24

Resources:

* Application Load Balancer
* NAT Gateway

### Private Subnets

* 10.0.11.0/24
* 10.0.12.0/24

Resources:

* EC2 Instances
* Redis Cluster

---

## Security Architecture

### Application Load Balancer Security Group

Inbound:

* HTTP (80)
* HTTPS (443)

Outbound:

* All traffic

### EC2 Security Group

Inbound:

* Port 8080 from ALB
* Port 22 from Administrator IP

Outbound:

* All traffic

### Redis Security Group

Inbound:

* Port 6379 from EC2 Security Group

Outbound:

* All traffic

---

## Monitoring Architecture

CloudWatch is used for:

* Application Logs
* EC2 Metrics
* Auto Scaling Metrics
* Alarm Notifications
* Dashboard Visualization

---

## CI/CD Architecture

### Infrastructure Pipeline

GitHub Actions

* Terraform Format
* Terraform Validate
* Terraform Plan
* Terraform Apply

### Frontend Pipeline

GitHub Actions

* npm install
* npm test
* npm audit
* npm run build
* S3 Deployment
* CloudFront Cache Invalidation

### Backend Pipeline

GitHub Actions

* Go Tests
* Linting
* Vulnerability Scanning
* Docker Build
* ECR Push
* Rolling Deployment
* Smoke Tests

---

## AWS Services Used

* Amazon VPC
* Amazon EC2
* Auto Scaling Group
* Application Load Balancer
* Amazon S3
* Amazon CloudFront
* Amazon ElastiCache Redis
* Amazon CloudWatch
* AWS IAM
* Amazon ECR

---

## Database

MongoDB Atlas is used as the primary database for persistent storage.

Connection strings and credentials are stored securely using GitHub Secrets and environment variables.
