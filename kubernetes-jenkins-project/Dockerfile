FROM centos:7
LABEL maintainer=alexrondon89@gmail.com
RUN yum install -y httpd zip unzip
COPY . /var/www/html/
WORKDIR /var/www/html/
RUN unzip photogenic.zip
RUN cp -rvf photogenic/* .
RUN rm -rf photogenic photogenic.zip
CMD ["/usr/sbin/httpd", "-D", "FOREGROUND"]
EXPOSE 80 443