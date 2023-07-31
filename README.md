# GO_KAFKA app
This is a simple app built using Golang, Kafka, and PostgreSQL. It exposes several APIs to create and fetch users and products. Utilizes Kafka for retrieval and publishing of data.

![image](https://github.com/e-for-eshaan/go-kafka/assets/76566992/22337373-582a-4acd-be3a-121108ef19fd)

*Fig: Workflow of the APP*

## GO RESTful APIs
I built a server using `Golang` and `PostgreSQL`. It exposed many API endpoints to create users as well as products. It also exposes a special `GET` endpoint that can be used to listen to the `Kafka Queue`

![app structure](https://github.com/e-for-eshaan/go-kafka/assets/76566992/bcba6529-4334-4b3f-82d0-b87b69d1bfef)

*Fig: Directory structure*

![API end-points](https://github.com/e-for-eshaan/go-kafka/assets/76566992/86eee4ba-3256-49cc-8e4a-0cc63ac61581)

*Fig: API endpoints for the app*

## Kafka server
![image](https://github.com/e-for-eshaan/go-kafka/assets/76566992/17bc84c2-abe8-4728-bcc4-97b408b8c132)

*Fig: Kafka Server is ran locally on the machine as looks like this*

### 1. Post a product
Create a post request for adding `Product` to the database

![image](https://github.com/e-for-eshaan/go-kafka/assets/76566992/e58fe8b1-fe5b-4477-a91f-83f17db7c99c)

*Fig: `POST` API for the creation of products. This houses a function to publish onto a Kafka queue.*

### 2. Publish to Kafka

The API function calls in itself, a function to publish onto the Kafka queue

![image](https://github.com/e-for-eshaan/go-kafka/assets/76566992/fd16e38e-44db-4483-a31f-f8b02f668ad4)

*Figure: Kafka consumer acknowledges that it can see a value.*

![image](https://github.com/e-for-eshaan/go-kafka/assets/76566992/83e9beef-9381-4b60-9266-a3e96593ff55)

*Figure: Logger for successful publishing*


### 3. Consume Kafka

We try to get the data from the Kafka queue, using a function for consumer

![image](https://github.com/e-for-eshaan/go-kafka/assets/76566992/49a7fb8d-0d9b-4ab8-8f0c-898a9237e7dc)

*Figure: A `GET` API to call the Kafak consumer, fetch from DB and return value.*

### 4. Return value
We try to get the data from the kafka queue, using a function for consumer

![image](https://github.com/e-for-eshaan/go-kafka/assets/76566992/7c650357-8177-4f9d-8311-78d4f8cc39fe)

*Fig: Finally returns a value.*
