# lr_event_driven_go

Example project: building event driven microservices in Go.

![Event Queue](diagrams/queue2.svg)

See [project tutorial](https://www.literank.com/project/18/intro) here.

This project provides a comprehensive guide on building event-driven microservices using Apache Kafka and the Gin web framework in Golang.

The event-driven microservices tutorial with Kafka and Gin provides a comprehensive guide to building a modern, scalable application architecture. Starting with preparation and objectives, the tutorial progresses through the creation of a web service using Gin, including designing a webpage layout, implementing data models, and integrating event-producing functionality with Kafka.

Subsequently, it explores the development of consumer services such as the Trend Service and Recommendation Service, each with their own design considerations, Gin API servers, and event consumers to process Kafka events generated by the web service.

The tutorial concludes with deployment instructions, utilizing Docker and Docker Compose to containerize and orchestrate the microservices and Kafka infrastructure.

Overall, this tutorial equips readers with the knowledge and practical experience necessary to architect, develop, and deploy event-driven microservices using Kafka and Gin, paving the way for the creation of robust and scalable applications.

## Build

```bash
make build
```

## Run in Docker Compose

Create `compose/.env` file:

```bash
REDIS_PASSWORD=your_pass
MYSQL_PASSWORD=your_pass
MYSQL_ROOT_PASSWORD=your_root_pass
```

Run it:

```bash
cd compose
docker compose up
```

See [project tutorial](https://www.literank.com/project/18/intro) here.
