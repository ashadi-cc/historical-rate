# Historical-rate
API service to display latest EUR rates. rates raferences takes from Europe Bank Central. https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml

# Installation 
- make run - Running on development mode. make sure you have Golang installed on your sistem 

# Build and run app in container 
- make container - build docker container
- docker run -p 8001:8001 your_build_image_tag

# API Endpoint 
- GET localhost:8001/rates/latest