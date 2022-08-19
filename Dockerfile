# image: 853539536475.dkr.ecr.us-east-1.amazonaws.com/yidemo
# docker build . --network=host -t yidemo:v01
# docker run --network=host -it yidemo:v01 /bin/bash
# docker ps -a
# docker exec -it {ID} /bin/bash
# curl localhost
# curl localhost/cgi-bin/test.py
# docker stop {ID}
FROM python:3.9.13 AS build
RUN apt-get update && apt-get install -y git
RUN mkdir /var/www/ && cd /var/www/ && git clone https://github.com/yz124/demo1.git
FROM python:3.9.13
RUN apt-get update && apt-get install -y apache2 python3-pip vim
RUN a2enmod cgid && mkdir /var/run/apache2
RUN echo '#########     Adding capaility to run CGI-scripts #################' >> /etc/apache2/apache2.conf && \
	echo 'ServerName localhost' >> /etc/apache2/apache2.conf && \
	echo 'ScriptAlias /cgi-bin/ /var/www/cgi-bin/' >> /etc/apache2/apache2.conf && \
	echo 'Options +ExecCGI' >> /etc/apache2/apache2.conf && \
	echo 'AddHandler cgi-script .cgi .pl .py' >> /etc/apache2/apache2.conf
RUN echo '<IfModule mod_alias.c>' > /etc/apache2/conf-available/serve-cgi-bin.conf && \
	echo '<IfModule mod_cgi.c>' >> /etc/apache2/conf-available/serve-cgi-bin.conf && \
	echo '	Define ENABLE_USR_LIB_CGI_BIN' >> /etc/apache2/conf-available/serve-cgi-bin.conf && \
	echo '</IfModule>' >> /etc/apache2/conf-available/serve-cgi-bin.conf && \
	echo '<IfModule mod_cgid.c>' >> /etc/apache2/conf-available/serve-cgi-bin.conf && \
	echo '	Define ENABLE_USR_LIB_CGI_BIN' >> /etc/apache2/conf-available/serve-cgi-bin.conf && \
	echo '</IfModule>' >> /etc/apache2/conf-available/serve-cgi-bin.conf && \
	echo '<IfDefine ENABLE_USR_LIB_CGI_BIN>' >> /etc/apache2/conf-available/serve-cgi-bin.conf && \
	echo '	ScriptAlias /cgi-bin/ /var/www/cgi-bin/' >> /etc/apache2/conf-available/serve-cgi-bin.conf && \
	echo '	<Directory "/usr/lib/cgi-bin">' >> /etc/apache2/conf-available/serve-cgi-bin.conf && \
	echo '		AllowOverride None' >> /etc/apache2/conf-available/serve-cgi-bin.conf && \
	echo '		Options +ExecCGI' >> /etc/apache2/conf-available/serve-cgi-bin.conf && \
	echo '	</Directory>' >> /etc/apache2/conf-available/serve-cgi-bin.conf && \
	echo '</IfDefine>' >> /etc/apache2/conf-available/serve-cgi-bin.conf && \
	echo '</IfModule>' >> /etc/apache2/conf-available/serve-cgi-bin.conf
RUN mkdir /var/www/cgi-bin/
COPY --from=build /var/www/demo1/index.html /var/www/html/
COPY --from=build /var/www/demo1/test.py /var/www/cgi-bin/
RUN chown -R www-data:www-data /var/www/ && chmod -R +x /var/www/cgi-bin/
RUN echo '#!/bin/bash' > /start.sh && \
	echo 'source /etc/apache2/envvars && /usr/sbin/apache2 -D FOREGROUND' >> /start.sh && \
	chmod +x /start.sh
EXPOSE 80

CMD /start.sh
