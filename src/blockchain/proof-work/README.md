Code your own Proof of Work algorithm!

Read our blog post first to see a walkthrough of the code.
Ask us anything!

Join our Telegram chat and follow us on Twitter!
Deployment steps:

    clone this repo
    navigate to this directory and rename the example file mv example.env .env
    go run main.go
    open a web browser and visit http://localhost:8080/
    to write new blocks, send a POST request (I like to use Postman) to http://localhost:8080/ with a JSON payload with BPM as the key and an integer as the value. For example:

{"BPM":50}

    watch your Terminal to see Proof of Work in action!
