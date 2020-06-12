# go-vk-news-loader

### Description

Loads news from popular vk news groups into PostgreSQL DB.

### Installation

- Install app:
  - ```git clone https://github.com/LeadNess/go-vk-news-loader.git```
- Set list of vk groups screen names in config/groups.json. Example of config/groups.json:
  -  ```["meduzaproject", "ria", "kommersant_ru", "tj", "rbc"]```
- Run 'deploy_container' script for setting custom app config (such as vk token and PostgreSQL connection information):
  - ```./deploy/deploy_container```
- Build docker image:
  - ```docker build -t news-service .```
- Run app as docker container (running PostgreSQL required):
  - ```docker run --name news-service-app news-service ```

