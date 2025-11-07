# Secret Santa Generator

Secret Santa Generator is a simple CLI script that generates the Secret Santa's couples for you.
Once you've added all the participants, their email, and optionally a propic and some ideas, it will
send and email to each one of them, with the extracted gift-receiver.

<img src="example.jpg">

You can run it with `go run . <email_file> <config_file>`.

An `email_file` has the following structure:

```jsonc
{
    "address": "", // the sender's email; required
    "host": "",    // the mail server's host; required
    "port": 0,     // the mail server's port; required
    "login": "",   // used when logging into the mail server; required
    "password": "" // used when logging into the mail server; required
}
```

A `config_file` has the following structure:
```jsonc
{
    "subject": "", // the emails subject; default "SecretSantaGenerator"

    "lang": "", // one of the ones available; default: en

    "players": [ // the list of players; required
        {
            "name": "",     // the player's name; required
            "email": "",    // the player's email; required
            "pic_path": "", // the path of the player's picture; default: pics/_missing.png
            "ideas": [      // the list of options; default: none
                {
                    "description": "", // required
                    "link": "",        // default: none
                }
                // ...
            ]
        }
        // ...
    ]
}
```

The code is released with the [Unlicense](LICENSE). Feel free to contribute.
