deploy procedure:
- create and push tag locally (git push origin v1.3.0)
- this triggers pipeline that creates draft release on github
- open that release and publish it
- this triggers another pipeline that publishes image to the docker
