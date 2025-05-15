import React, { useState, useCallback } from 'react';
import Cropper, { Area } from 'react-easy-crop';
import '../styles/ImageCropperModal.scss';

interface ImageCropperModalProps {
  open: boolean;
  imageSrc: string | null;
  onClose: () => void;
  onCropComplete?: (croppedBlob: Blob) => void;
  aspect?: number;
  outputWidth?: number;
  outputHeight?: number;
  apiEndpoint?: string;
  formData?: Record<string, string>;
  onSuccess?: () => void;
  onError?: (err: any) => void;
  objectName?: string;
  personId?: number;
}

// Helper to get cropped image as a blob
async function getCroppedImg(
  imageSrc: string,
  crop: { x: number; y: number },
  zoom: number,
  croppedAreaPixels: Area,
  outputWidth: number,
  outputHeight: number
): Promise<Blob> {
  const image = new window.Image();
  image.src = imageSrc;
  await new Promise((resolve) => {
    image.onload = resolve;
  });
  const canvas = document.createElement('canvas');
  canvas.width = outputWidth;
  canvas.height = outputHeight;
  const ctx = canvas.getContext('2d');
  if (!ctx) throw new Error('No 2d context');
  ctx.drawImage(
    image,
    croppedAreaPixels.x,
    croppedAreaPixels.y,
    croppedAreaPixels.width,
    croppedAreaPixels.height,
    0,
    0,
    outputWidth,
    outputHeight
  );
  return new Promise((resolve) => {
    canvas.toBlob((blob) => {
      if (blob) resolve(blob);
    }, 'image/jpeg');
  });
}

const ImageCropperModal: React.FC<ImageCropperModalProps> = ({
  open,
  imageSrc,
  onClose,
  onCropComplete,
  aspect = 1,
  outputWidth = 512,
  outputHeight = 512,
  apiEndpoint,
  formData = {},
  onSuccess,
  onError,
  objectName,
  personId,
}) => {
  const [crop, setCrop] = useState<{ x: number; y: number }>({ x: 0, y: 0 });
  const [zoom, setZoom] = useState(1);
  const [croppedAreaPixels, setCroppedAreaPixels] = useState<Area | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const onCropCompleteHandler = useCallback(
    (_: Area, croppedAreaPixels: Area) => {
      setCroppedAreaPixels(croppedAreaPixels);
    },
    []
  );

  const handleDone = async () => {
    if (!imageSrc || !croppedAreaPixels) return;
    setLoading(true);
    setError(null);
    try {
      const croppedBlob = await getCroppedImg(
        imageSrc,
        crop,
        zoom,
        croppedAreaPixels,
        outputWidth,
        outputHeight
      );
      if (onCropComplete) onCropComplete(croppedBlob);
      let uploadSuccess = false;
      if (apiEndpoint) {
        const fd = new FormData();
        Object.entries(formData).forEach(([key, value]) => {
          fd.append(key, value);
        });
        fd.append('file', croppedBlob, objectName || 'cropped.jpg');
        const res = await fetch(apiEndpoint, {
          method: 'POST',
          body: fd,
        });
        if (!res.ok) {
          const errText = await res.text();
          setError(errText);
          if (onError) onError(errText);
          setLoading(false);
          return;
        }
        uploadSuccess = true;
      }
      // After successful upload, update the person's photo_url in the database
      if (uploadSuccess && personId && objectName) {
        const updateRes = await fetch('http://localhost:8080/persons/update', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            id: personId,
            updates: { photo_url: objectName },
          }),
        });
        if (!updateRes.ok) {
          const errText = await updateRes.text();
          setError('Image uploaded but failed to update database: ' + errText);
          if (onError) onError(errText);
          setLoading(false);
          return;
        }
      }
      if (onSuccess) onSuccess();
      setLoading(false);
      onClose();
    } catch (err: any) {
      setError(err.message || 'Unknown error');
      if (onError) onError(err);
      setLoading(false);
    }
  };

  if (!open || !imageSrc) return null;

  return (
    <div className="image-cropper-modal-overlay">
      <div className="image-cropper-modal-content">
        <div className="image-cropper-modal-cropper">
          <Cropper
            image={imageSrc}
            crop={crop}
            zoom={zoom}
            aspect={aspect}
            onCropChange={setCrop}
            onZoomChange={setZoom}
            onCropComplete={onCropCompleteHandler}
          />
        </div>
        <div className="image-cropper-modal-actions">
          <button onClick={handleDone} disabled={loading}>
            {loading ? 'Uploading...' : 'Done'}
          </button>
          <button onClick={onClose} disabled={loading}>
            Cancel
          </button>
        </div>
        {error && <div style={{ color: 'red', marginTop: 8 }}>{error}</div>}
      </div>
    </div>
  );
};

export default ImageCropperModal;
