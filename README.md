# GodRider

Welcome! This is a project I developed as an icebreaker for my first steps in Go programming language. All suggestions in terms of best practices and performance are welcome :wink:

## Main objective
This project is a fictional REST API which allow a consumer to see and take incoming orders from different rider platforms in an unified format. The premise is the following:

* Any consumer needs an account in order to interact with the API. All incoming requests need to be sent with an authentication token, which is valid for a 24h time window. The consumer needs to call to Login endpoint in order to get or refresh this token.
* The consumer can see all available order providers (the rider platforms which register) and subscribe/unsuscribe for their order feeds
* When requested, the API will return a collection of open orders with two conditions:
    1. They must come only from user's subscribed providers.
    2. All orders must have the same format. If required, an order will be able to show a provider's company logo, but the objective is return them in an unified format.

