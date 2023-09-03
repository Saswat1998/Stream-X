"use client";
import axios from "axios";
import { useState } from "react";

const Upload = () => {
  const [file, setFile] = useState(null);
  const [uploadProgress, setUploadProgress] = useState(0);
  const [showProgressBar, setShowProgressBar] = useState(false);
  const [s3ObjectUrl, setS3ObjectUrl] = useState("");
  const [cancelUpload, setCancelUpload] = useState(null);

  const uploadFile = async () => {
    setShowProgressBar(true);
    const source = axios.CancelToken.source("Upload cancelled by user");
    setCancelUpload(() => () => source.cancel("Upload Cancelled by user"));
    if (!file) {
      alert("Please select a file first.");
      setShowProgressBar(false);
      return;
    }

    const body = {
      videoName: file.name,
    };
    const res = await axios.post(
      "https://n30f91l0xh.execute-api.ap-south-1.amazonaws.com/prod/",
      body
    );
    const signedUrl = res.data.preSignedURL;

    const options = {
      headers: {
        "Content-Type": file.type,
      },
      onUploadProgress: (progressEvent) => {
        const percentCompleted = Math.round(
          (progressEvent.loaded * 100) / progressEvent.total
        );
        setUploadProgress(percentCompleted);
      },
      cancelToken: source.token,
    };
    try {
      const response = await axios.put(signedUrl, file, options);
      console.log("response after uploading: ", response);
      const s3Url = `https://x-stream-videos.s3.ap-south-1.amazonaws.com/uploads/${file.name}`;
      const sanitizedURL = s3Url.replace(/%20/g, "+");
      setS3ObjectUrl(sanitizedURL);

      alert("File Uploaded");
    } catch (error) {
        if (axios.isCancel(error)) {
          console.log("Upload canceled by user");
        } else {
          console.log("Upload failed:", error);
        }
    }

    setUploadProgress(0);
    setShowProgressBar(false);
    setCancelUpload(null);
  };

  const onFileChange = (e) => {
    setFile(e.target.files[0]);
    setS3ObjectUrl("");
  };

  return (
    <div>
      <h1>Upload Video</h1>
      <input type="file" accept="video/*" onChange={onFileChange} />
      <button onClick={uploadFile}>Upload</button>
      {showProgressBar && (
        <progress value={uploadProgress} max="100"></progress>
      )}
      {cancelUpload && <button onClick={cancelUpload}>Cancel</button>}
      {s3ObjectUrl != "" && (
        <div>
          <video controls width="250">
            <source src={s3ObjectUrl} type="video/mp4"></source>
          </video>
        </div>
      )}
    </div>
  );
};

export default Upload;
