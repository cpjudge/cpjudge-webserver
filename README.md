# Welcome to Buffalo!

Thank you for choosing Buffalo for your web development needs.

## Database Setup

> This section is only if you're **not** using docker to run the app. If you are using docker, follow the [Docker setup](#docker-setup) section.

It looks like you chose to set up your application using a mysql database! Fantastic!

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

You will also need to make sure that **you** start/install the database of your choice. Buffalo **won't** install and start mysql for you.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started mysql, now Buffalo can create the databases in that file for you:

	$ buffalo db create -a

### Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## Docker setup

To run the application using docker, you need to have the following installed:
+ [Docker](https://docs.docker.com/install/)
+ [Docker Compose](https://docs.docker.com/compose/install/)

We have used two services in our [docker-compose file](./docker-compose.yml) - `db` and `app`.

The `db` service is for the database, which in our case is `mysql v5.7`. It has the `/var/lib/mysql` directory mounted as a [docker volume](https://docs.docker.com/storage/volumes/) to have persistent database between containers.

The `app` service is for the Buffalo web application. It has current directory mounted as a volume.

### Starting the Containers

To start the containers, run the following command from the root directory of this project:

	$ docker-compose up -d

> `-d` flag is used to run the containers in detached mode (background)

After that, to connect to the container and open a terminal, run:

	$ docker-compose exec app bash

You are now inside the container with access to its terminal.

### Starting the Application

Make sure all dependencies are installed by using "dep":

	$ dep ensure

Create the databases by running the following command:

	$ buffalo db create -a

Then, run the app by using the buffalo command line utility:

	$ buffalo dev run

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

Good luck!

[Powered by Buffalo](http://gobuffalo.io)
