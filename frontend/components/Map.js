import React, { useState } from 'react';
import { MapContainer, TileLayer, GeoJSON } from 'react-leaflet';

const Map = () => {
    const [geojsonData, setGeojsonData] = useState(null);
    const [drawnShapes, setDrawnShapes] = useState([]);

    const handleGeoJSONChange = (e) => {
       
        setGeojsonData(e.target.toGeoJSON());
    };

    const handleShapeCreation = (e) => {
      
        const { layer } = e;
        const shapeData = layer.toGeoJSON();
        setDrawnShapes([...drawnShapes, shapeData]);
    };

    return (
        <div>
            <MapContainer center={[51.505, -0.09]} zoom={13} style={{ height: '400px', width: '100%' }}>
                <TileLayer
                    url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                    attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                />
                <GeoJSON data={geojsonData} onEachFeature={handleGeoJSONChange} />
                <GeoJSON data={drawnShapes} />
            </MapContainer>
            <button onClick={() => setGeojsonData(null)}>Clear Shapes</button>
        </div>
    );
};

export default Map;
