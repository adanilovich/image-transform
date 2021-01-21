# Image-transform
This one is my test task as part of the interview process.

# Task description:
## Objective: Create a simple REST API in Go.
1. Application should allow users to provide a single file (as url).
2. Application should validate if the file is a valid png or jpeg image (don't rely on Content-Type).
3. Application should be able to perform different transformations on the provided image (additional credit for not using any external packages):
   a) crop - crop the image by the provided bounding rectangle,
   b) rotate - rotate the image by the multiple of 90 degrees,
   c) remove exif - remove exif metadata from the image.
4. Application should accept the transformation parameters and file url from a POST request with JSON body (e.g. {"file": "...", "transformations": [...]}) on '/accept' endpoint.
5. User should be able to define multiple transformations in one request (e.g. rotate by 90 degrees, crop, rotate by 180 degrees, remove exif).
6. Transformations should be run concurrently and each should produce a separate file (named respectively to the transformation parameters).
7. There should be a limit of maximum number of concurrently running transformations adjustable in a JSON configuration file.
8. All the results should be compressed into a single zip archive and streamed back to the user in response.
9. The streaming should begin as soon as the first transformation is complete.
Bonus: Application should work in second mode - instead of streaming the zip file back in response body, it should stream it to a cloud storage (e.g. S3) and return a JSON payload containing a temporary url to the resource.
Bonus: Add resize transformation using the standard library only.
Please deliver a public Github repository with README explaining how the application can be run locally.
What we will be looking at (in the order of importance):
1. Understanding of requirements.
2. The correctness of requirements implementation.
3. Code organization.
4. Code quality.
5. Presence of tests.
