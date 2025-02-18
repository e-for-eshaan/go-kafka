# GO_KAFKA app
This is a simple app built using Golang, Kafka, and PostgreSQL. It exposes several APIs to create and fetch users and products. Utilizes Kafka for retrieval and publishing of data.

<img width="844" alt="image" src="https://github.com/user-attachments/assets/342e4699-9d1c-4883-9b36-3789b88bba0e" />

*Fig: Workflow of the APP*

## GO RESTful APIs
I built a server using `Golang` and `PostgreSQL`. It exposed many API endpoints to create users as well as products. It also exposes a special `GET` endpoint that can be used to listen to the `Kafka Queue`

![257329078-bcba6529-4334-4b3f-82d0-b87b69d1bfef](https://github.com/e-for-eshaan/go-kafka/assets/76566992/905a4de7-393d-46e6-bba4-af801eafbe85)

*Fig: Directory structure*

![257329347-86eee4ba-3256-49cc-8e4a-0cc63ac61581](https://github.com/e-for-eshaan/go-kafka/assets/76566992/ea3c992c-8629-4bbb-9d59-e051427b48f2)

*Fig: API endpoints for the app*

## Kafka server

![257328276-17bc84c2-abe8-4728-bcc4-97b408b8c132](https://github.com/e-for-eshaan/go-kafka/assets/76566992/14cd3f8f-69f6-411c-bce1-bc50c027ead5)

*Fig: Kafka Server is ran locally on the machine as looks like this*

### 1. Post a product
Create a post request for adding `Product` to the database

![257330140-e58fe8b1-fe5b-4477-a91f-83f17db7c99c](https://github.com/e-for-eshaan/go-kafka/assets/76566992/6126cc06-ab4f-4932-975c-bebb030de4fb)

*Fig: `POST` API for the creation of products. This houses a function to publish onto a Kafka queue.*

### 2. Publish to Kafka

The API function calls in itself, a function to publish onto the Kafka queue

![257328628-fd16e38e-44db-4483-a31f-f8b02f668ad4](https://github.com/e-for-eshaan/go-kafka/assets/76566992/d5c0f32d-761d-4119-acac-b77af2d3a043)

*Figure: Kafka consumer acknowledges that it can see a value.*

![image](https://github.com/e-for-eshaan/go-kafka/assets/76566992/6ecd6dd6-fe67-446a-b612-3f968016553f)

*Figure: Logger for successful publishing*


### 3. Consume Kafka

We try to get the data from the Kafka queue, using a function for consumer

![257329724-49a7fb8d-0d9b-4ab8-8f0c-898a9237e7dc](https://github.com/e-for-eshaan/go-kafka/assets/76566992/5b53cd21-04e8-4364-ac79-e13e79d4db5f)

*Figure: A `GET` API to call the Kafak consumer, fetch from DB and return value.*

### 4. Return value
We try to get the data from the kafka queue, using a function for consumer

![257329603-7c650357-8177-4f9d-8311-78d4f8cc39fe](https://github.com/e-for-eshaan/go-kafka/assets/76566992/db74b2ae-dc1f-45d7-bd48-bd767e9ab31f)

*Fig: Finally returns a value.*
