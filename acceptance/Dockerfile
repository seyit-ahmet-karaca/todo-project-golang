FROM node:lts

COPY . .
RUN apt update
RUN apt install libasound2 libnss3-dev libgdk-pixbuf2.0-dev libgtk-3-dev libxss-dev -y
RUN npm install

CMD ["yarn","start"]
#CMD ["sleep", "3600"]