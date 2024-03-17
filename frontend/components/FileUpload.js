import React, { useState } from 'react';
import axios from 'axios'; 

const FileUpload = () => {
    const [selectedFile, setSelectedFile] = useState(null);
    const [uploading, setUploading] = useState(false);
    const [uploadError, setUploadError] = useState(null);
    const [uploadSuccess, setUploadSuccess] = useState(false);

    const handleFileChange = (e) => {
        setSelectedFile(e.target.files[0]);
        setUploadError(null); 
    };

    const handleUpload = async () => {
        if (!selectedFile) {
            setUploadError('Please select a file.');
            return;
        }

        setUploading(true);

        try {
            const formData = new FormData();
            formData.append('file', selectedFile);

           
            const response = await axios.post('/api/upload', formData, {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
            });

            if (response.status === 200) {
                setUploadSuccess(true);
            } else {
                setUploadError('Failed to upload file. Please try again.');
            }
        } catch (error) {
            console.error('Upload error:', error);
            setUploadError('An error occurred while uploading the file. Please try again.');
        } finally {
            setUploading(false);
        }
    };

    return (
        <div>
            <h2>Upload GeoJSON/KML File</h2>
            <input type="file" onChange={handleFileChange} />
            <button onClick={handleUpload} disabled={uploading}>
                {uploading ? 'Uploading...' : 'Upload'}
            </button>
            {uploadError && <p style={{ color: 'red' }}>{uploadError}</p>}
            {uploadSuccess && <p style={{ color: 'green' }}>File uploaded successfully!</p>}
        </div>
    );
};

export default FileUpload;
