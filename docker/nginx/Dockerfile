FROM nginx

# Remove a configuração padrão do Nginx
#RUN rm /etc/nginx/conf.d/default.conf
COPY interface/index.html /usr/share/nginx/html/

# Copia a configuração personalizada para o contêiner
COPY nginx.conf /etc/nginx/conf.d/

