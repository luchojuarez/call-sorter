# phone-company-invoice

## Run
`make web`

## Get invoice
curl --location --request GET 'localhost:8080/v1/invoice/2021/1' --header 'phone_number: +5491167980953'


## task pending
[-]Add wrapper in generate proccess (storage generated invoices for future gets)

# For challenge reviewers 👋
## Some disclaimers
in challenge says: "generate", in order of this, will be use a POST method in this senario, for the meneang of POST verb.

choose use `phone_numner` header, to get key to identify the current user in order to emulate some user token.

This code probably taste like an overkill solution, but it was funny to wrote.
I hope you enjoy reading and testing
