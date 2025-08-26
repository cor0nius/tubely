# Tubely - A Video Upload and Streaming Backend with AWS S3

Tubely is a backend service for a video-sharing application, built in Go. It provides a JSON API for user authentication and video metadata management. The key feature of this project is its integration with AWS S3 for scalable, cloud-based video storage and AWS CloudFront for content delivery.

This project was completed as part of the ["Learn File Servers and CDNs with S3 and CloudFront"](https://www.boot.dev/courses/learn-file-servers-s3-cloudfront-golang) course on boot.dev, starting from a pre-built backend and enhancing it with cloud storage capabilities.

## Features Implemented

*   **Direct-to-S3 Video Uploads:** Replaced the local file storage system with direct uploads to an AWS S3 bucket.
*   **Video Processing:** The application processes videos locally to ensure they are optimized for web streaming ("fast start") before uploading them to the cloud.
*   **CDN Integration:** Serves video files through a CDN (AWS CloudFront) for fast, low-latency streaming across the globe.
*   **Local Thumbnail Storage:** The starter project's functionality for handling thumbnail uploads to the local filesystem was retained.

## Key Cloud & Backend Skills Learned

*   **Cloud Storage with AWS S3:**
    *   Configured and integrated the AWS SDK for Go (v2) into an existing application.
    *   Wrote the logic to securely upload large video files from the server to an S3 bucket.
    *   Managed S3 object keys to organize files within the bucket.
*   **Content Delivery Networks (CDN):**
    *   Understood the role of a CDN in a modern file-serving architecture.
    *   Constructed file URLs that point to a CloudFront distribution for efficient, cached content delivery.
*   **Backend API Enhancement:**
    *   Modified existing API handlers to incorporate the new cloud upload functionality.
    *   Updated the database logic to store the S3/CDN URL of the video file instead of a local path.
*   **Configuration Management:**
    *   Managed environment variables for all AWS-related settings, including bucket name, region, and CloudFront distribution domain.
*   **Handling Multipart File Uploads:**
    *   Worked with multipart form data from HTTP requests to receive and process video files on the backend before uploading to the cloud.
