# Covid-19 Vaccine Appointment Availibility Monitor

A basic monitor that polls a vaccine appointment API to determine availability in a zip code within a certain radius. Supports noitification through Discord Webhook notifications

## Disclaimers
This is truly meant for people who really need the vaccine as soon as possible but have been having trouble getting appointments. If you don't need it right now, do not abuse this tool to get it before someone who needs it more than you.

I have not tested this at all on Mac but I did leave an OSX executable here that should work in theory.

This was written very quickly and is poorly error handled. A lot of the code is pretty bad right now but I just wanted to put something together as fast as a could. I may or may not push some more updates to it for improvements later.

A couple of notes about appointment monitoring:
- No Walgreens appointments are detected as they list theirs privately on their own website
- CVS lists availability as just one appointment per day, but following the link to their own site reveals a more comprehensive scheduling system
- I find this tool works best in detecting smaller vaccine sites that list all their appointments and brand availability on the public API I am polling, but obviously it misses anything that isn't listed there
- Brand availability (JJ/Pfizer/Moderna) will default to "unavailable" if they do not list, and a lot of sites don't list their brand availability through this API so double check if the brand you need is available

## Requirements
Go (I have 1.16 but it will likely work with older versions)

## Usage

### Building yourself
Clone the repository and compile into an executable by typing `go build`. 

### Using the precompiled binaries
Just clone this repository and run vaccinemonitor for your respective OS.

### Running the monitor
Once you have your executable, you can run it and follow the command line instructions to input your desired zip code, search radius, and Discord webhook URL. 

### Webhooks
If you don't know how to set up a Discord webhook, follow the instructions [here](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks) to get your webhook link.
Webhooks are not required, as the monitor will log any detected appointments to the console as well, but it is more convenient for receiving push notifications.
This is what a webhook will look like:
[webhook](./webhook.png)

