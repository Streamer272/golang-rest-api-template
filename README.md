# Rest API

### Setup DB
`sudo -S docker-compose -f db.yml up -d`
### Build container
`sudo -S docker build -t rest-api-image .`
### Run container from image
`sudo -S docker run --net host -it --name rest-api rest-api-image`
<br>
If you want to run it in detached mode, add `-d` flag.
<br>
If container already exists, run it with `sudo docker start -a -i rest-api`
<br>
If you don't want to start it in attached mode, remove `-a` flag.
